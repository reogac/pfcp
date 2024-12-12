package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
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
}

func (ie *OuterHeaderRemoval) Bytes() []byte {
	return []byte{ie.Desc}
}
func (ie *OuterHeaderRemoval) UnmarshalBinary(wire []byte) (err error) {
	if len(wire) > 1 {
		return tlv.ErrTail
	} else if len(wire) < 1 {
		return tlv.ErrIncomplete
	}
	ie.Desc = wire[0]
	return nil
}
