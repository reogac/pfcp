package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
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
	Unknown                                    []uint32                                    `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid                                    []uint32                                    `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
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
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [9]bool
	for _, de := range elements {
		opt := false
		switch de.Type {
		case MAR_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.MARID = de.Value
		case STEERING_FUNCTIONALITY_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			t.SteeringFunctionality = de.Value
		case STEERING_MODE_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			t.SteeringMode = de.Value
		case UPDATE__ACCESS_FORWARDING_ACTION_INFORMATION_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			var v UpdateAccessForwardingActionInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.UpdateAccessForwardingActionInformation = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case UPDATE_NON__ACCESS_FORWARDING_ACTION_INFORMATION_TYPE:
			if seen[4] {
				continue
			}
			seen[4] = true
			var v UpdateNonAccessForwardingActionInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.UpdateNonAccessForwardingActionInformation = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case _ACCESS_FORWARDING_ACTION_INFORMATION_TYPE:
			if seen[5] {
				continue
			}
			seen[5] = true
			var v AccessForwardingActionInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.AccessForwardingActionInformation = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case NON__ACCESS_FORWARDING_ACTION_INFORMATION_TYPE:
			if seen[6] {
				continue
			}
			seen[6] = true
			var v NonAccessForwardingActionInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.NonAccessForwardingActionInformation = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case THRESHOLDS_TYPE:
			if seen[7] {
				continue
			}
			seen[7] = true
			t.Thresholdvalues = de.Value
		case STEERING_MODE_INDICATOR_TYPE:
			if seen[8] {
				continue
			}
			seen[8] = true
			t.SteeringModeIndicator = de.Value
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
		return &pfcp.MissingIE{Type: MAR_ID_TYPE}
	}
	return
}
