package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Overload Control Information' data structure */
type OverloadControlInformation struct {
	OverloadControlSequenceNumber   []byte //Sequence Number
	OverloadReductionMetric         []byte //Metric
	PeriodofValidity                []byte //Timer
	OverloadControlInformationFlags []byte //OCI Flags
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
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case SEQUENCE_NUMBER_TYPE:
			t.OverloadControlSequenceNumber = de.Value
		case METRIC_TYPE:
			t.OverloadReductionMetric = de.Value
		case TIMER_TYPE:
			t.PeriodofValidity = de.Value
		case OCI_FLAGS_TYPE:
			t.OverloadControlInformationFlags = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
