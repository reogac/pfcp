package ie

import (
	"github.com/reogac/pfcp"

	"github.com/reogac/pfcp/tlv"
)

/* 'Remove BAR' data structure */
type RemoveBAR struct {
	BARID uint8 //BAR ID
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
		return
	}
	for _, de := range elements {
		switch de.Type {
		case BAR_ID_TYPE:
			err = pfcp.UnmarshalUint8(de.Value, &t.BARID)
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
