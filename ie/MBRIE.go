package ie

import (
	"github.com/reogac/pfcp/tlv"
)

func NewMBR(ul, dl uint64) *MBR {
	return &MBR{
		Ul: ul,
		Dl: dl,
	}
}

// MBR is the uplink and downlink bitrate in kbps, each five octets on the
// wire (TS 29.244 §8.2.8).
type MBR struct {
	Ul, Dl uint64
}

func (ie *MBR) Bytes() []byte {
	return appendBitrate(appendBitrate(make([]byte, 0, 2*bitrateLen), ie.Ul), ie.Dl)
}
func (ie *MBR) UnmarshalBinary(wire []byte) error {
	if len(wire) < 2*bitrateLen {
		return tlv.ErrIncomplete
	} else if len(wire) > 2*bitrateLen {
		return tlv.ErrTail
	}
	ie.Ul, ie.Dl = bitrate(wire[:bitrateLen]), bitrate(wire[bitrateLen:])
	return nil
}
