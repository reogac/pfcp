package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Remove MAR' data structure */
type RemoveMAR struct {
	MARID   []byte   //MAR ID
	Unknown []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
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
		return pfcp.Faulty(0, err)
	}
	var seen [1]bool
	for _, de := range elements {
		switch de.Type {
		case MAR_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.MARID = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return pfcp.Faulty(de.Type, err)
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: MAR_ID_TYPE}
	}
	return
}
