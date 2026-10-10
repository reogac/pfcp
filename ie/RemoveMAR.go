package ie

import (
	"github.com/reogac/pfcp/tlv"
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
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	for _, de := range elements {
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
