package ie

import (
	"github.com/reogac/pfcp"

	"github.com/reogac/pfcp/tlv"
)

/* 'Remove SRR' data structure */
type RemoveSRR struct {
	SRRID   []byte   //SRR ID
	Unknown []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *RemoveSRR) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(SRR_ID_TYPE, tlv.Bytes(t.SRRID)))
	return tlv.TLV(REMOVE_SRR_TYPE, fields...)
}

func (t *RemoveSRR) UnmarshalBinary(buf []byte) (err error) {
	*t = RemoveSRR{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	var seen [1]bool
	for _, de := range elements {
		switch de.Type {
		case SRR_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.SRRID = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: SRR_ID_TYPE}
	}
	return
}
