package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Overload Control Information' data structure */
type OverloadControlInformation struct {
	OverloadControlSequenceNumber   []byte   //Sequence Number
	OverloadReductionMetric         []byte   //Metric
	PeriodofValidity                []byte   //Timer
	OverloadControlInformationFlags []byte   //OCI Flags
	Unknown                         []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid                         []uint32 `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
}

func (t *OverloadControlInformation) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(SEQUENCE_NUMBER_TYPE, tlv.Bytes(t.OverloadControlSequenceNumber)))
	fields = append(fields, tlv.TLVFrom(METRIC_TYPE, tlv.Bytes(t.OverloadReductionMetric)))
	fields = append(fields, tlv.TLVFrom(TIMER_TYPE, tlv.Bytes(t.PeriodofValidity)))
	if len(t.OverloadControlInformationFlags) > 0 {
		fields = append(fields, tlv.TLVFrom(OCI_FLAGS_TYPE, tlv.Bytes(t.OverloadControlInformationFlags)))
	}
	return tlv.TLV(OVERLOAD_CONTROL_INFORMATION_TYPE, fields...)
}

func (t *OverloadControlInformation) UnmarshalBinary(buf []byte) (err error) {
	*t = OverloadControlInformation{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [4]bool
	for _, de := range elements {
		opt := false
		switch de.Type {
		case SEQUENCE_NUMBER_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.OverloadControlSequenceNumber = de.Value
		case METRIC_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			t.OverloadReductionMetric = de.Value
		case TIMER_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			t.PeriodofValidity = de.Value
		case OCI_FLAGS_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			t.OverloadControlInformationFlags = de.Value
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
		return &pfcp.MissingIE{Type: SEQUENCE_NUMBER_TYPE}
	}
	if !seen[1] {
		return &pfcp.MissingIE{Type: METRIC_TYPE}
	}
	if !seen[2] {
		return &pfcp.MissingIE{Type: TIMER_TYPE}
	}
	return
}
