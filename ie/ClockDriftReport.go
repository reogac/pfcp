package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Clock Drift Report' data structure */
type ClockDriftReport struct {
	TimeDomainNumber               []byte //Time Domain Number
	TimeOffsetMeasurement          []byte //Time Offset Measurement
	CumulativerateRatioMeasurement []byte //Cumulative rateRatio Measurement
	TimeStamp                      []byte //Time Stamp
	NetworkInstance                []byte //Network Instance
	APNDNN                         []byte //APN/DNN
	SNSSAI                         []byte //S-NSSAI
}

func (t *ClockDriftReport) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(TIME_DOMAIN_NUMBER_TYPE, tlv.Bytes(t.TimeDomainNumber)))
	if len(t.TimeOffsetMeasurement) > 0 {
		fields = append(fields, tlv.TLVFrom(TIME_OFFSET_MEASUREMENT_TYPE, tlv.Bytes(t.TimeOffsetMeasurement)))
	}
	if len(t.CumulativerateRatioMeasurement) > 0 {
		fields = append(fields, tlv.TLVFrom(CUMULATIVE_RATERATIO_MEASUREMENT_TYPE, tlv.Bytes(t.CumulativerateRatioMeasurement)))
	}
	if len(t.TimeStamp) > 0 {
		fields = append(fields, tlv.TLVFrom(TIME_STAMP_TYPE, tlv.Bytes(t.TimeStamp)))
	}
	if len(t.NetworkInstance) > 0 {
		fields = append(fields, tlv.TLVFrom(NETWORK_INSTANCE_TYPE, tlv.Bytes(t.NetworkInstance)))
	}
	if len(t.APNDNN) > 0 {
		fields = append(fields, tlv.TLVFrom(APN_DNN_TYPE, tlv.Bytes(t.APNDNN)))
	}
	if len(t.SNSSAI) > 0 {
		fields = append(fields, tlv.TLVFrom(S_NSSAI_TYPE, tlv.Bytes(t.SNSSAI)))
	}
	return tlv.TLV(CLOCK_DRIFT_REPORT_TYPE, fields...)
}

func (t *ClockDriftReport) UnmarshalBinary(buf []byte) (err error) {
	*t = ClockDriftReport{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case TIME_DOMAIN_NUMBER_TYPE:
			t.TimeDomainNumber = de.Value
		case TIME_OFFSET_MEASUREMENT_TYPE:
			t.TimeOffsetMeasurement = de.Value
		case CUMULATIVE_RATERATIO_MEASUREMENT_TYPE:
			t.CumulativerateRatioMeasurement = de.Value
		case TIME_STAMP_TYPE:
			t.TimeStamp = de.Value
		case NETWORK_INSTANCE_TYPE:
			t.NetworkInstance = de.Value
		case APN_DNN_TYPE:
			t.APNDNN = de.Value
		case S_NSSAI_TYPE:
			t.SNSSAI = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
