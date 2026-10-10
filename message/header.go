package message

import (
	"encoding/binary"
	"errors"
)

// Version is the only PFCP version there is (TS 29.244 §7.2.2).
const Version = 1

const (
	//the four octets every header starts with: flags, type, length
	fixedLen = 4
	//a node message's header: the fixed part, the sequence number and a spare
	//octet
	nodeHeaderLen = 8
	//a session message's header adds the SEID
	sessionHeaderLen = 16

	flagS  = 0x01
	flagMP = 0x02
	flagFO = 0x04

	maxSeq = 1<<24 - 1
)

var (
	// ErrVersion is a message of a version other than 1, which a receiver
	// answers with a Version Not Supported Response.
	ErrVersion = errors.New("PFCP version not supported")
	// ErrSFlag is a header whose S flag disagrees with its message type: a
	// session message without a SEID, or a node message with one.
	ErrSFlag = errors.New("PFCP S flag does not match the message type")
	// ErrUnknownType is a message type this library has no body for.
	ErrUnknownType = errors.New("unknown PFCP message type")
	// ErrShort is input shorter than the header, or than the length it states.
	ErrShort = errors.New("PFCP message shorter than its header or length")
	// ErrTail is input left after a message whose FO flag says nothing follows.
	ErrTail = errors.New("junk after the PFCP message")
	// ErrTooLong is a message longer than its length field holds.
	ErrTooLong = errors.New("PFCP message longer than 65535 octets after its first four")
)

// Header is a PFCP header (TS 29.244 §7.2.2) without its length, which is the
// codec's to compute and check.
type Header struct {
	Type Type
	// Seid is the receiver's SEID: 0 in a Session Establishment Request, whose
	// sender's SEID travels in its CP F-SEID. Only a session message has one.
	Seid uint64
	// Seq is the sequence number, 24 bits; a response carries its request's.
	Seq uint32
	// Priority is the message priority, 0 (highest) to 15, sent only when
	// HasPriority is set.
	Priority    uint8
	HasPriority bool
	// FollowOn is set on a message another follows in the same datagram.
	FollowOn bool
}

// Len is how many octets the header takes on the wire.
func (h *Header) Len() int {
	if h.Type.IsSession() {
		return sessionHeaderLen
	}
	return nodeHeaderLen
}

// appendHeader writes h with a zero length and returns where the length goes.
func appendHeader(b []byte, h *Header) []byte {
	flags := byte(Version << 5)
	if h.FollowOn {
		flags |= flagFO
	}
	if h.HasPriority {
		flags |= flagMP
	}
	session := h.Type.IsSession()
	if session {
		flags |= flagS
	}
	b = append(b, flags, byte(h.Type), 0, 0)
	if session {
		b = binary.BigEndian.AppendUint64(b, h.Seid)
	}
	var last byte
	if h.HasPriority {
		//bits 8-5 of the octet after the sequence number
		last = h.Priority << 4
	}
	return append(b, byte(h.Seq>>16), byte(h.Seq>>8), byte(h.Seq), last)
}

// DecodeHeader reads the header at the start of b and returns the body it
// frames and what follows it, which is another message only when FollowOn is
// set. It allocates nothing.
//
// The checks run before anything is read past the fixed part: the version,
// then the stated length against what arrived, then the S flag against the
// type. A type this library does not know is returned with ErrUnknownType and
// a header that is otherwise complete, so a receiver can still answer it.
func DecodeHeader(b []byte) (h Header, body, rest []byte, err error) {
	if len(b) < fixedLen {
		return h, nil, nil, ErrShort
	}
	flags := b[0]
	if flags>>5 != Version {
		h.Type = Type(b[1])
		return h, nil, nil, ErrVersion
	}
	h.Type = Type(b[1])
	end := fixedLen + int(binary.BigEndian.Uint16(b[2:4]))
	if len(b) < end {
		return h, nil, nil, ErrShort
	}
	h.FollowOn = flags&flagFO != 0
	h.HasPriority = flags&flagMP != 0
	hasSeid := flags&flagS != 0

	hlen := nodeHeaderLen
	if hasSeid {
		hlen = sessionHeaderLen
	}
	if end < hlen {
		return h, nil, nil, ErrShort
	}
	off := fixedLen
	if hasSeid {
		h.Seid = binary.BigEndian.Uint64(b[off:])
		off += 8
	}
	h.Seq = uint32(b[off])<<16 | uint32(b[off+1])<<8 | uint32(b[off+2])
	if h.HasPriority {
		h.Priority = b[off+3] >> 4
	}
	body, rest = b[hlen:end], b[end:]

	if !h.Type.Known() {
		return h, body, rest, ErrUnknownType
	}
	if hasSeid != h.Type.IsSession() {
		return h, body, rest, ErrSFlag
	}
	if !h.FollowOn && len(rest) > 0 {
		return h, body, rest, ErrTail
	}
	return h, body, rest, nil
}

// putLength fills in the length of the message that starts at start and ends
// at the end of b.
func putLength(b []byte, start int) error {
	n := len(b) - start - fixedLen
	if n > 0xffff {
		return ErrTooLong
	}
	binary.BigEndian.PutUint16(b[start+2:], uint16(n))
	return nil
}
