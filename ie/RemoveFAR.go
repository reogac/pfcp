package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Remove FAR' data structure */
type RemoveFAR struct {
	FARID   uint32   //FAR ID
	Unknown []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid []uint32 `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
}

func (t *RemoveFAR) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(FAR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(t.FARID))))
	return tlv.TLV(REMOVE_FAR_TYPE, fields...)
}

func (t *RemoveFAR) UnmarshalBinary(buf []byte) (err error) {
	*t = RemoveFAR{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [1]bool
	for _, de := range elements {
		opt := false
		switch de.Type {
		case FAR_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &t.FARID)
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
		return &pfcp.MissingIE{Type: FAR_ID_TYPE}
	}
	return
}
