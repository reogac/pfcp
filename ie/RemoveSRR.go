package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Remove SRR' data structure */
type RemoveSRR struct {
	SRRID   []byte   //SRR ID
	Unknown []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid []uint32 `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
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
		return pfcp.Faulty(0, err)
	}
	var seen [1]bool
	for _, de := range elements {
		opt := false
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
			if err = pfcp.Faulty(de.Type, err); opt && pfcp.Discardable(err) {
				t.Invalid = append(t.Invalid, de.Type)
				err = nil
				continue
			}
			return err
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: SRR_ID_TYPE}
	}
	return
}
