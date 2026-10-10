package ie

import (
	"github.com/reogac/pfcp/tlv"
)

func NewGBR(ul, dl uint64) *GBR {
	return &GBR{
		Ul: ul,
		Dl: dl,
	}
}

// GBR is the uplink and downlink bitrate in kbps, each five octets on the
// wire (TS 29.244 §8.2.9).
type GBR struct {
	Ul, Dl uint64
}

func (ie *GBR) Bytes() []byte {
	return appendBitrate(appendBitrate(make([]byte, 0, 2*bitrateLen), ie.Ul), ie.Dl)
}
func (ie *GBR) UnmarshalBinary(wire []byte) error {
	if len(wire) < 2*bitrateLen {
		return tlv.ErrIncomplete
	} else if len(wire) > 2*bitrateLen {
		return tlv.ErrTail
	}
	ie.Ul, ie.Dl = bitrate(wire[:bitrateLen]), bitrate(wire[bitrateLen:])
	return nil
}
