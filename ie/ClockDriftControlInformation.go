package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Clock Drift Control Information' data structure */
type ClockDriftControlInformation struct {
	RequestedClockDriftInformation []byte   //Requested Clock Drift Information
	TimeDomainNumber               [][]byte //Time Domain Number
	ConfiguredTimeDomain           []byte   //Configured Time Domain
	TimeOffsetThreshold            []byte   //Time Offset Threshold
	CumulativerateRatioThreshold   []byte   //Cumulative rateRatio Threshold
	Unknown                        []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid                        []uint32 `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
}

func (t *ClockDriftControlInformation) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(REQUESTED_CLOCK_DRIFT_INFORMATION_TYPE, tlv.Bytes(t.RequestedClockDriftInformation)))
	for _, ie := range t.TimeDomainNumber {
		fields = append(fields, tlv.TLVFrom(TIME_DOMAIN_NUMBER_TYPE, tlv.Bytes(ie)))
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
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [4]bool
	for _, de := range elements {
		opt := false
		switch de.Type {
		case REQUESTED_CLOCK_DRIFT_INFORMATION_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.RequestedClockDriftInformation = de.Value
		case TIME_DOMAIN_NUMBER_TYPE:
			t.TimeDomainNumber = append(t.TimeDomainNumber, de.Value)
		case CONFIGURED_TIME_DOMAIN_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			t.ConfiguredTimeDomain = de.Value
		case TIME_OFFSET_THRESHOLD_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			t.TimeOffsetThreshold = de.Value
		case CUMULATIVE_RATERATIO_THRESHOLD_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			t.CumulativerateRatioThreshold = de.Value
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
		return &pfcp.MissingIE{Type: REQUESTED_CLOCK_DRIFT_INFORMATION_TYPE}
	}
	return
}
