package ie

import (
	"github.com/reogac/pfcp/tlv"
)

/* 'Error Indication Report' data structure */
type ErrorIndicationReport struct {
	RemoteFTEID FTEID //F-TEID
}

func (t *ErrorIndicationReport) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(F_TEID_TYPE, tlv.Bytes(t.RemoteFTEID.Bytes())))
	return tlv.TLV(ERROR_INDICATION_REPORT_TYPE, fields...)
}

func (t *ErrorIndicationReport) UnmarshalBinary(buf []byte) (err error) {
	*t = ErrorIndicationReport{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	for _, de := range elements {
		switch de.Type {
		case F_TEID_TYPE:
			err = de.UnmarshalValue(&t.RemoteFTEID)
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
