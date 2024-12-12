package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Clock Drift Control Information' data structure */
type ClockDriftControlInformation struct {
	RequestedClockDriftInformation []byte //Requested Clock Drift Information
	TimeDomainNumber               []byte //Time Domain Number
	ConfiguredTimeDomain           []byte //Configured Time Domain
	TimeOffsetThreshold            []byte //Time Offset Threshold
	CumulativerateRatioThreshold   []byte //Cumulative rateRatio Threshold
}

func (t *ClockDriftControlInformation) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(REQUESTED_CLOCK_DRIFT_INFORMATION_TYPE, tlv.Bytes(t.RequestedClockDriftInformation)))
	if len(t.TimeDomainNumber) > 0 {
		fields = append(fields, tlv.TLVFrom(TIME_DOMAIN_NUMBER_TYPE, tlv.Bytes(t.TimeDomainNumber)))
	}
	if len(t.ConfiguredTimeDomain) > 0 {
		fields = append(fields, tlv.TLVFrom(CONFIGURED_TIME_DOMAIN_TYPE, tlv.Bytes(t.ConfiguredTimeDomain)))
	}
	if len(t.TimeOffsetThreshold) > 0 {
		fields = append(fields, tlv.TLVFrom(TIME_OFFSET_THRESHOLD_TYPE, tlv.Bytes(t.TimeOffsetThreshold)))
	}
	if len(t.CumulativerateRatioThreshold) > 0 {
		fields = append(fields, tlv.TLVFrom(CUMULATIVE_RATERATIO_THRESHOLD_TYPE, tlv.Bytes(t.CumulativerateRatioThreshold)))
	}
	return tlv.TLV(CLOCK_DRIFT_CONTROL_INFORMATION_TYPE, fields...)
}

func (t *ClockDriftControlInformation) UnmarshalBinary(buf []byte) (err error) {
	*t = ClockDriftControlInformation{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case REQUESTED_CLOCK_DRIFT_INFORMATION_TYPE:
			t.RequestedClockDriftInformation = de.Value
		case TIME_DOMAIN_NUMBER_TYPE:
			t.TimeDomainNumber = de.Value
		case CONFIGURED_TIME_DOMAIN_TYPE:
			t.ConfiguredTimeDomain = de.Value
		case TIME_OFFSET_THRESHOLD_TYPE:
			t.TimeOffsetThreshold = de.Value
		case CUMULATIVE_RATERATIO_THRESHOLD_TYPE:
			t.CumulativerateRatioThreshold = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
