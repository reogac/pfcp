package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Remove MAR' data structure */
type RemoveMAR struct {
	MARID []byte //MAR ID
}

func (t *RemoveMAR) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(MAR_ID_TYPE, tlv.Bytes(t.MARID)))
	return tlv.TLV(REMOVE_MAR_TYPE, fields...)
}

func (t *RemoveMAR) UnmarshalBinary(buf []byte) (err error) {
	*t = RemoveMAR{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case MAR_ID_TYPE:
			t.MARID = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
