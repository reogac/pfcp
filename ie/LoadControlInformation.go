package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Load Control Information' data structure */
type LoadControlInformation struct {
	LoadControlSequenceNumber []byte   //Sequence Number
	LoadMetric                []byte   //Metric
	Unknown                   []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid                   []uint32 `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
}

func (t *LoadControlInformation) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(SEQUENCE_NUMBER_TYPE, tlv.Bytes(t.LoadControlSequenceNumber)))
	fields = append(fields, tlv.TLVFrom(METRIC_TYPE, tlv.Bytes(t.LoadMetric)))
	return tlv.TLV(LOAD_CONTROL_INFORMATION_TYPE, fields...)
}

func (t *LoadControlInformation) UnmarshalBinary(buf []byte) (err error) {
	*t = LoadControlInformation{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [2]bool
	for _, de := range elements {
		opt := false
		switch de.Type {
		case SEQUENCE_NUMBER_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.LoadControlSequenceNumber = de.Value
		case METRIC_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			t.LoadMetric = de.Value
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
	return
}
