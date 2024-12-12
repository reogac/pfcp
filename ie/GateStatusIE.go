package ie

import "github.com/usnistgov/ndn-dpdk/ndn/tlv"

const (
	GateStatusOpen   uint8 = 1
	GateStatusClosed uint8 = 0
)

func NewGateStatus(ul, dl uint8) *GateStatus {
	return &GateStatus{
		Ul: ul,
		Dl: dl,
	}
}

type GateStatus struct {
	Ul uint8
	Dl uint8
}

func (ie *GateStatus) Bytes() []byte {
	return []byte{ie.Ul, ie.Dl}
}
func (ie *GateStatus) UnmarshalBinary(wire []byte) (err error) {
	if len(wire) < 2 {
		return tlv.ErrIncomplete
	} else if len(wire) > 2 {
		return tlv.ErrTail
	}
	ie.Ul = wire[0]
	ie.Dl = wire[1]
	return nil
}
