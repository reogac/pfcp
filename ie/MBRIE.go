package ie

import (
	"github.com/reogac/pfcp"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

func NewMBR(ul, dl uint32) *MBR {
	return &MBR{
		Ul: ul,
		Dl: dl,
	}
}

type MBR struct {
	Ul, Dl uint32
}

func (ie *MBR) Bytes() []byte {
	return append(pfcp.MarshalUint32(ie.Ul), pfcp.MarshalUint32(ie.Dl)...)
}
func (ie *MBR) UnmarshalBinary(wire []byte) error {
	if len(wire) < 8 {
		return tlv.ErrIncomplete
	} else if len(wire) > 8 {
		return tlv.ErrTail
	}
	pfcp.UnmarshalUint32(wire[0:4], &ie.Ul)
	pfcp.UnmarshalUint32(wire[4:], &ie.Dl)
	return nil
}
