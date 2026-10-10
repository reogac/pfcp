package ie

import (
	"github.com/reogac/pfcp/tlv"
)

/* 'Access Availability Report' data structure */
type AccessAvailabilityReport struct {
	AccessAvailabilityInformation []byte //Access Availability Information
}

func (t *AccessAvailabilityReport) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(ACCESS_AVAILABILITY_INFORMATION_TYPE, tlv.Bytes(t.AccessAvailabilityInformation)))
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
			t.AccessAvailabilityInformation = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
