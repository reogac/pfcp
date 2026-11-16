package ie

import (
	"github.com/reogac/pfcp"

	"github.com/reogac/pfcp/tlv"
)

/* 'Remove PDR' data structure */
type RemovePDR struct {
	PDRID   uint16   //PDR ID
	Unknown []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *RemovePDR) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(PDR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint16(t.PDRID))))
	return tlv.TLV(REMOVE_PDR_TYPE, fields...)
}

func (t *RemovePDR) UnmarshalBinary(buf []byte) (err error) {
	*t = RemovePDR{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	var seen [1]bool
	for _, de := range elements {
		switch de.Type {
		case PDR_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			err = pfcp.UnmarshalExtendable(de.Value, 2, 2, &t.PDRID)
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: PDR_ID_TYPE}
	}
	return
}
