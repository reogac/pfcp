package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'PFD context' data structure */
type PFDcontext struct {
	PFDContents []byte //PFD contents
}

func (t *PFDcontext) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(PFD_CONTENTS_TYPE, tlv.Bytes(t.PFDContents)))
	return tlv.TLV(PFD_CONTEXT_TYPE, fields...)
}

func (t *PFDcontext) UnmarshalBinary(buf []byte) (err error) {
	*t = PFDcontext{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case PFD_CONTENTS_TYPE:
			t.PFDContents = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
