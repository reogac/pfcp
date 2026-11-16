package ie

import (
	"github.com/reogac/pfcp"

	"github.com/reogac/pfcp/tlv"
)

/* 'Remove URR' data structure */
type RemoveURR struct {
	URRID   uint32   //URR ID
	Unknown []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *RemoveURR) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(URR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(t.URRID))))
	return tlv.TLV(REMOVE_URR_TYPE, fields...)
}

func (t *RemoveURR) UnmarshalBinary(buf []byte) (err error) {
	*t = RemoveURR{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	var seen [1]bool
	for _, de := range elements {
		switch de.Type {
		case URR_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &t.URRID)
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: URR_ID_TYPE}
	}
	return
}
