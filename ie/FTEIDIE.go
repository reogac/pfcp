package ie

import (
	"encoding/binary"
	"net"

	"github.com/reogac/pfcp/tlv"
)

type FTEID struct {
	teid   uint32
	v4, v6 net.IP
	ch     bool
	//chV4 and chV6 are the address families a CHOOSE asks for, which TS 29.244
	//§8.2.3 carries in the V4 and V6 flags with no address behind them
	chV4, chV6 bool
	chId       *uint8
}

// NewChFTEID asks the UP function to choose an IPv4 F-TEID.
func NewChFTEID(chId *uint8) *FTEID {
	return NewChooseFTEID(true, false, chId)
}

// NewChooseFTEID asks the UP function to choose an F-TEID in the families
// named; at least one must be.
func NewChooseFTEID(v4, v6 bool, chId *uint8) *FTEID {
	return &FTEID{
		ch:   true,
		chV4: v4,
		chV6: v6,
		chId: chId,
	}
}
func NewFTEID(v4, v6 *net.IP, teid uint32) (ie *FTEID) {
	ie = &FTEID{
		teid: teid,
	}
	if v4 != nil {
		ie.v4 = v4.To4()
	}
	if v6 != nil {
		ie.v6 = v6.To16()
	}
	return
}

func (ie *FTEID) V4() net.IP {
	return ie.v4
}

func (ie *FTEID) V6() net.IP {
	return ie.v6
}

func (ie *FTEID) Teid() uint32 {
	return ie.teid
}

func (ie *FTEID) IsCh() bool {
	return ie.ch
}

func (ie *FTEID) ChooseId() *uint8 {
	return ie.chId
}

// ChooseV4 and ChooseV6 are the families a CHOOSE asks for.
func (ie *FTEID) ChooseV4() bool {
	return ie.chV4
}

func (ie *FTEID) ChooseV6() bool {
	return ie.chV6
}

func (ie *FTEID) Bytes() (wire []byte) {
	hasV4 := len(ie.v4) > 0 || ie.ch && ie.chV4
	hasV6 := len(ie.v6) > 0 || ie.ch && ie.chV6
	var flags uint8

	if hasV4 {
		flags = 1 // bit 0
	}
	if hasV6 {
		flags |= 1 << 1 //bit 1
	}
	if ie.ch {
		flags |= 1 << 2 //bit 2
	}
	if ie.chId != nil {
		flags |= 1 << 3 //bit 3
	}
	if !ie.ch {
		wireLen := 5
		if hasV4 {
			wireLen += net.IPv4len
		}
		if hasV6 {
			wireLen += net.IPv6len
		}
		wire = make([]byte, wireLen)
		wire[0] = flags
		binary.BigEndian.PutUint32(wire[1:5], ie.teid)
		offset := 5
		if hasV4 {
			copy(wire[offset:offset+net.IPv4len], ie.v4)
			offset += net.IPv4len
		}
		if hasV6 {
			copy(wire[offset:offset+net.IPv6len], ie.v6)
		}
	} else {
		if ie.chId != nil {
			wire = []byte{flags, *ie.chId}
		} else {
			wire = []byte{flags}
		}
	}
	return
}

func (ie *FTEID) UnmarshalBinary(wire []byte) error {
	wireLen := len(wire)
	if wireLen < 1 {
		return tlv.ErrIncomplete
	}
	hasV4 := (wire[0] & 1) == 1
	hasV6 := ((wire[0] & (1 << 1)) >> 1) == 1
	ie.ch = ((wire[0] & (1 << 2)) >> 2) == 1
	chId := ((wire[0] & (1 << 3)) >> 3) == 1
	offset := 1
	if !ie.ch {
		if offset+4 > wireLen {
			return tlv.ErrIncomplete
		}
		ie.teid = binary.BigEndian.Uint32(wire[offset : offset+4])
		offset += 4
		if hasV4 {
			if offset+net.IPv4len > wireLen {
				return tlv.ErrIncomplete
			}
			ie.v4 = net.IP(make([]byte, net.IPv4len))
			copy(ie.v4[:], wire[offset:offset+net.IPv4len])
			offset += net.IPv4len
		}
		if hasV6 {
			if offset+net.IPv6len > wireLen {
				return tlv.ErrIncomplete
			}
			ie.v6 = net.IP(make([]byte, net.IPv6len))
			copy(ie.v6[:], wire[offset:offset+net.IPv6len])
			offset += net.IPv6len

		}

	} else {
		ie.chV4, ie.chV6 = hasV4, hasV6
		if chId {
			if offset+1 > wireLen {
				return tlv.ErrIncomplete
			} else {
				ie.chId = new(uint8)
				*ie.chId = wire[1]
				offset += 1
			}
		}
	}

	//check there is no junk
	if offset < wireLen {
		return tlv.ErrTail
	}
	return nil
}
