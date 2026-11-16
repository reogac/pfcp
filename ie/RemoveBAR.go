package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Remove BAR' data structure */
type RemoveBAR struct {
	BARID   uint8    //BAR ID
	Unknown []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *RemoveBAR) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(BAR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint8(t.BARID))))
	return tlv.TLV(REMOVE_BAR_TYPE, fields...)
}

func (t *RemoveBAR) UnmarshalBinary(buf []byte) (err error) {
	*t = RemoveBAR{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [1]bool
	for _, de := range elements {
		switch de.Type {
		case BAR_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &t.BARID)
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return pfcp.Faulty(de.Type, err)
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: BAR_ID_TYPE}
	}
	return
}
