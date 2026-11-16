package ie

import (
	"github.com/reogac/pfcp"

	"github.com/reogac/pfcp/tlv"
)

/* 'Remove QER' data structure */
type RemoveQER struct {
	QERID   uint32   //QER ID
	Unknown []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *RemoveQER) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(QER_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(t.QERID))))
	return tlv.TLV(REMOVE_QER_TYPE, fields...)
}

func (t *RemoveQER) UnmarshalBinary(buf []byte) (err error) {
	*t = RemoveQER{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	var seen [1]bool
	for _, de := range elements {
		switch de.Type {
		case QER_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &t.QERID)
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: QER_ID_TYPE}
	}
	return
}
