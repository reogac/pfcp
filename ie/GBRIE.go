package ie

import (
	"github.com/reogac/pfcp"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

func NewGBR(ul, dl uint32) *GBR {
	return &GBR{
		Ul: ul,
		Dl: dl,
	}
}

type GBR struct {
	Ul, Dl uint32
}

func (ie *GBR) Bytes() []byte {
	return append(pfcp.MarshalUint32(ie.Ul), pfcp.MarshalUint32(ie.Dl)...)
}
func (ie *GBR) UnmarshalBinary(wire []byte) error {
	if len(wire) < 8 {
		return tlv.ErrIncomplete
	} else if len(wire) > 8 {
		return tlv.ErrTail
	}
	pfcp.UnmarshalUint32(wire[0:4], &ie.Ul)
	pfcp.UnmarshalUint32(wire[4:], &ie.Dl)
	return nil
}
