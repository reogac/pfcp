package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Create MAR' data structure */
type CreateMAR struct {
	MARID                                []byte                                //MAR ID
	SteeringFunctionality                []byte                                //Steering Functionality
	SteeringMode                         []byte                                //Steering Mode
	AccessForwardingActionInformation    *AccessForwardingActionInformation    //3GPP Access Forwarding Action Information
	NonAccessForwardingActionInformation *NonAccessForwardingActionInformation //Non-3GPP Access Forwarding Action Information
	ThresholdValues                      []byte                                //Thresholds
	SteeringModeIndicator                []byte                                //Steering Mode Indicator
}

func (t *CreateMAR) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(MAR_ID_TYPE, tlv.Bytes(t.MARID)))
	fields = append(fields, tlv.TLVFrom(STEERING_FUNCTIONALITY_TYPE, tlv.Bytes(t.SteeringFunctionality)))
	fields = append(fields, tlv.TLVFrom(STEERING_MODE_TYPE, tlv.Bytes(t.SteeringMode)))
	if t.AccessForwardingActionInformation != nil {
		fields = append(fields, t.AccessForwardingActionInformation.Field())
	}
	if t.NonAccessForwardingActionInformation != nil {
		fields = append(fields, t.NonAccessForwardingActionInformation.Field())
	}
	if len(t.ThresholdValues) > 0 {
		fields = append(fields, tlv.TLVFrom(THRESHOLDS_TYPE, tlv.Bytes(t.ThresholdValues)))
	}
	if len(t.SteeringModeIndicator) > 0 {
		fields = append(fields, tlv.TLVFrom(STEERING_MODE_INDICATOR_TYPE, tlv.Bytes(t.SteeringModeIndicator)))
	}
	return tlv.TLV(CREATE_MAR_TYPE, fields...)
}

func (t *CreateMAR) UnmarshalBinary(buf []byte) (err error) {
	*t = CreateMAR{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case MAR_ID_TYPE:
			t.MARID = de.Value
		case STEERING_FUNCTIONALITY_TYPE:
			t.SteeringFunctionality = de.Value
		case STEERING_MODE_TYPE:
			t.SteeringMode = de.Value
		case _ACCESS_FORWARDING_ACTION_INFORMATION_TYPE:
			var v AccessForwardingActionInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.AccessForwardingActionInformation = &v
			}
		case NON__ACCESS_FORWARDING_ACTION_INFORMATION_TYPE:
			var v NonAccessForwardingActionInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.NonAccessForwardingActionInformation = &v
			}
		case THRESHOLDS_TYPE:
			t.ThresholdValues = de.Value
		case STEERING_MODE_INDICATOR_TYPE:
			t.SteeringModeIndicator = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
