package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Remove SRR' data structure */
type RemoveSRR struct {
	SRRID []byte //SRR ID
}

func (t *RemoveSRR) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(SRR_ID_TYPE, tlv.Bytes(t.SRRID)))
	return tlv.TLV(REMOVE_SRR_TYPE, fields...)
}

func (t *RemoveSRR) UnmarshalBinary(buf []byte) (err error) {
	*t = RemoveSRR{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case SRR_ID_TYPE:
			t.SRRID = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
