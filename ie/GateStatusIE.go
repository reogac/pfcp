package ie

import "github.com/reogac/pfcp/tlv"

// A gate's value as TS 29.244 §8.2.27 defines it: 0 is OPEN and 1 CLOSED.
const (
	GateStatusOpen   uint8 = 0
	GateStatusClosed uint8 = 1
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

// Bytes packs both gates into the one octet of TS 29.244 §8.2.27: the DL gate
// in bits 1-2 and the UL gate in bits 3-4.
func (ie *GateStatus) Bytes() []byte {
	return []byte{(ie.Ul&0x03)<<2 | ie.Dl&0x03}
}
func (ie *GateStatus) UnmarshalBinary(wire []byte) (err error) {
	if len(wire) < 1 {
		return tlv.ErrIncomplete
	} else if len(wire) > 1 {
		return tlv.ErrTail
	}
	ie.Ul = wire[0] >> 2 & 0x03
	ie.Dl = wire[0] & 0x03
	return nil
}
