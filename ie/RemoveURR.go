package ie

import (
	"github.com/reogac/pfcp"

	"github.com/reogac/pfcp/tlv"
)

/* 'Remove URR' data structure */
type RemoveURR struct {
	URRID uint32 //URR ID
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
	for _, de := range elements {
		switch de.Type {
		case URR_ID_TYPE:
			err = pfcp.UnmarshalUint32(de.Value, &t.URRID)
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
