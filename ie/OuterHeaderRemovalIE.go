package ie

import (
	"github.com/reogac/pfcp/tlv"
)

const (
	OuterHeaderRemovalGtpUUdpIpv4 uint8 = iota
	OuterHeaderRemovalGtpUUdpIpv6
	OuterHeaderRemovalUdpIpv4
	OuterHeaderRemovalUdpIpv6
)

func NewOuterHeaderRemoval(desc uint8) *OuterHeaderRemoval {
	return &OuterHeaderRemoval{
		Desc: desc,
	}
}

type OuterHeaderRemoval struct {
	Desc uint8
	// ExtHeaderDeletion is the optional second octet, GTP-U Extension Header
	// Deletion (TS 29.244 §8.2.64); nil when it is absent.
	ExtHeaderDeletion *uint8
}

func (ie *OuterHeaderRemoval) Bytes() []byte {
	if ie.ExtHeaderDeletion != nil {
		return []byte{ie.Desc, *ie.ExtHeaderDeletion}
	}
	return []byte{ie.Desc}
}
func (ie *OuterHeaderRemoval) UnmarshalBinary(wire []byte) (err error) {
	if len(wire) > 2 {
		return tlv.ErrTail
	} else if len(wire) < 1 {
		return tlv.ErrIncomplete
	}
	ie.Desc = wire[0]
	ie.ExtHeaderDeletion = nil
	if len(wire) == 2 {
		ext := wire[1]
		ie.ExtHeaderDeletion = &ext
	}
	return nil
}
