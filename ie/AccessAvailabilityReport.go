package ie

import (
	"github.com/reogac/pfcp"

	"github.com/reogac/pfcp/tlv"
)

/* 'Access Availability Report' data structure */
type AccessAvailabilityReport struct {
	AccessAvailabilityInformation [][]byte //Access Availability Information
	Unknown                       []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *AccessAvailabilityReport) Field() tlv.Field {
	fields := []tlv.Field{}
	for _, ie := range t.AccessAvailabilityInformation {
		fields = append(fields, tlv.TLVFrom(ACCESS_AVAILABILITY_INFORMATION_TYPE, tlv.Bytes(ie)))
	}
	return tlv.TLV(ACCESS_AVAILABILITY_REPORT_TYPE, fields...)
}

func (t *AccessAvailabilityReport) UnmarshalBinary(buf []byte) (err error) {
	*t = AccessAvailabilityReport{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	for _, de := range elements {
		switch de.Type {
		case ACCESS_AVAILABILITY_INFORMATION_TYPE:
			t.AccessAvailabilityInformation = append(t.AccessAvailabilityInformation, de.Value)
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	if len(t.AccessAvailabilityInformation) == 0 {
		return &pfcp.MissingIE{Type: ACCESS_AVAILABILITY_INFORMATION_TYPE}
	}
	return
}
