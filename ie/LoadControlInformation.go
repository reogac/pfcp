package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Load Control Information' data structure */
type LoadControlInformation struct {
	LoadControlSequenceNumber []byte //Sequence Number
	LoadMetric                []byte //Metric
}

func (t *LoadControlInformation) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(SEQUENCE_NUMBER_TYPE, tlv.Bytes(t.LoadControlSequenceNumber)))
	fields = append(fields, tlv.TLVFrom(METRIC_TYPE, tlv.Bytes(t.LoadMetric)))
	return tlv.TLV(LOAD_CONTROL_INFORMATION_TYPE, fields...)
}

func (t *LoadControlInformation) UnmarshalBinary(buf []byte) (err error) {
	*t = LoadControlInformation{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case SEQUENCE_NUMBER_TYPE:
			t.LoadControlSequenceNumber = de.Value
		case METRIC_TYPE:
			t.LoadMetric = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
