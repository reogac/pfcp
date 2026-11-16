package ie

import (
	"github.com/reogac/pfcp"

	"github.com/reogac/pfcp/tlv"
)

/* 'Error Indication Report' data structure */
type ErrorIndicationReport struct {
	RemoteFTEID []FTEID  //F-TEID
	Unknown     []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *ErrorIndicationReport) Field() tlv.Field {
	fields := []tlv.Field{}
	for _, ie := range t.RemoteFTEID {
		fields = append(fields, tlv.TLVFrom(F_TEID_TYPE, tlv.Bytes(ie.Bytes())))
	}
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
			var v FTEID
			if err = de.UnmarshalValue(&v); err == nil {
				t.RemoteFTEID = append(t.RemoteFTEID, v)
			}
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	if len(t.RemoteFTEID) == 0 {
		return &pfcp.MissingIE{Type: F_TEID_TYPE}
	}
	return
}
