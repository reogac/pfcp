package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Clock Drift Report' data structure */
type ClockDriftReport struct {
	TimeDomainNumber               []byte   //Time Domain Number
	TimeOffsetMeasurement          []byte   //Time Offset Measurement
	CumulativerateRatioMeasurement []byte   //Cumulative rateRatio Measurement
	TimeStamp                      []byte   //Time Stamp
	NetworkInstance                []byte   //Network Instance
	APNDNN                         []byte   //APN/DNN
	SNSSAI                         []byte   //S-NSSAI
	Unknown                        []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
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
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [7]bool
	for _, de := range elements {
		switch de.Type {
		case TIME_DOMAIN_NUMBER_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.TimeDomainNumber = de.Value
		case TIME_OFFSET_MEASUREMENT_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			t.TimeOffsetMeasurement = de.Value
		case CUMULATIVE_RATERATIO_MEASUREMENT_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			t.CumulativerateRatioMeasurement = de.Value
		case TIME_STAMP_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			t.TimeStamp = de.Value
		case NETWORK_INSTANCE_TYPE:
			if seen[4] {
				continue
			}
			seen[4] = true
			t.NetworkInstance = de.Value
		case APN_DNN_TYPE:
			if seen[5] {
				continue
			}
			seen[5] = true
			t.APNDNN = de.Value
		case S_NSSAI_TYPE:
			if seen[6] {
				continue
			}
			seen[6] = true
			t.SNSSAI = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return pfcp.Faulty(de.Type, err)
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: TIME_DOMAIN_NUMBER_TYPE}
	}
	return
}
