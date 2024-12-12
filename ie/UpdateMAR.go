package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Update MAR' data structure */
type UpdateMAR struct {
	MARID                                      []byte                                      //MAR ID
	SteeringFunctionality                      []byte                                      //Steering Functionality
	SteeringMode                               []byte                                      //Steering Mode
	UpdateAccessForwardingActionInformation    *UpdateAccessForwardingActionInformation    //Update 3GPP Access Forwarding Action Information
	UpdateNonAccessForwardingActionInformation *UpdateNonAccessForwardingActionInformation //Update Non-3GPP Access Forwarding Action Information
	AccessForwardingActionInformation          *AccessForwardingActionInformation          //3GPP Access Forwarding Action Information
	NonAccessForwardingActionInformation       *NonAccessForwardingActionInformation       //Non-3GPP Access Forwarding Action Information
	Thresholdvalues                            []byte                                      //Thresholds
	SteeringModeIndicator                      []byte                                      //Steering Mode Indicator
}

func (t *UpdateMAR) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(MAR_ID_TYPE, tlv.Bytes(t.MARID)))
	if len(t.SteeringFunctionality) > 0 {
		fields = append(fields, tlv.TLVFrom(STEERING_FUNCTIONALITY_TYPE, tlv.Bytes(t.SteeringFunctionality)))
	}
	if len(t.SteeringMode) > 0 {
		fields = append(fields, tlv.TLVFrom(STEERING_MODE_TYPE, tlv.Bytes(t.SteeringMode)))
	}
	if t.UpdateAccessForwardingActionInformation != nil {
		fields = append(fields, t.UpdateAccessForwardingActionInformation.Field())
	}
	if t.UpdateNonAccessForwardingActionInformation != nil {
		fields = append(fields, t.UpdateNonAccessForwardingActionInformation.Field())
	}
	if t.AccessForwardingActionInformation != nil {
		fields = append(fields, t.AccessForwardingActionInformation.Field())
	}
	if t.NonAccessForwardingActionInformation != nil {
		fields = append(fields, t.NonAccessForwardingActionInformation.Field())
	}
	if len(t.Thresholdvalues) > 0 {
		fields = append(fields, tlv.TLVFrom(THRESHOLDS_TYPE, tlv.Bytes(t.Thresholdvalues)))
	}
	if len(t.SteeringModeIndicator) > 0 {
		fields = append(fields, tlv.TLVFrom(STEERING_MODE_INDICATOR_TYPE, tlv.Bytes(t.SteeringModeIndicator)))
	}
	return tlv.TLV(UPDATE_MAR_TYPE, fields...)
}

func (t *UpdateMAR) UnmarshalBinary(buf []byte) (err error) {
	*t = UpdateMAR{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case MAR_ID_TYPE:
			t.MARID = de.Value
		case STEERING_FUNCTIONALITY_TYPE:
			t.SteeringFunctionality = de.Value
		case STEERING_MODE_TYPE:
			t.SteeringMode = de.Value
		case UPDATE__ACCESS_FORWARDING_ACTION_INFORMATION_TYPE:
			var v UpdateAccessForwardingActionInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.UpdateAccessForwardingActionInformation = &v
			}
		case UPDATE_NON__ACCESS_FORWARDING_ACTION_INFORMATION_TYPE:
			var v UpdateNonAccessForwardingActionInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.UpdateNonAccessForwardingActionInformation = &v
			}
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
			t.Thresholdvalues = de.Value
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
