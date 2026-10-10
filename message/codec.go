package message

import "github.com/reogac/pfcp/tlv"

// The codec works at two levels, and the body level stands alone: an SMF that
// hands its N4 to a proxy over SBI never sees a header, while the proxy and the
// UPF frame every body with one.

// AppendBody appends m's body, its IEs, to b.
func AppendBody(b []byte, m Message) ([]byte, error) {
	return tlv.AppendValue(b, m)
}

// EncodeBody is AppendBody into a new buffer.
func EncodeBody(m Message) ([]byte, error) {
	return AppendBody(nil, m)
}

// DecodeBody decodes a body of type t. The message aliases body.
func DecodeBody(t Type, body []byte) (Message, error) {
	m := New(t)
	if m == nil {
		return nil, ErrUnknownType
	}
	if err := m.UnmarshalBinary(body); err != nil {
		return nil, err
	}
	return m, nil
}

// AppendMessage appends m behind header h, in one pass: the header is written
// with a zero length, the body appended after it, and the length filled in
// last. h.Type is taken from m, so the two cannot disagree.
func AppendMessage(b []byte, h Header, m Message) ([]byte, error) {
	h.Type = m.MsgType()
	start := len(b)
	b = appendHeader(b, &h)
	b, err := AppendBody(b, m)
	if err != nil {
		return b[:start], err
	}
	if err = putLength(b, start); err != nil {
		return b[:start], err
	}
	return b, nil
}

// AppendRaw appends a body already encoded, for one that arrived encoded and is
// only being framed - as a proxy relays what an SMF sent it over SBI.
func AppendRaw(b []byte, h Header, body []byte) ([]byte, error) {
	if !h.Type.Known() {
		return b, ErrUnknownType
	}
	start := len(b)
	b = appendHeader(b, &h)
	b = append(b, body...)
	if err := putLength(b, start); err != nil {
		return b[:start], err
	}
	return b, nil
}

// Decode decodes one message from the start of b: its header, its body, and
// what follows it in the datagram.
func Decode(b []byte) (h Header, m Message, rest []byte, err error) {
	var body []byte
	if h, body, rest, err = DecodeHeader(b); err != nil {
		return
	}
	m, err = DecodeBody(h.Type, body)
	return
}
