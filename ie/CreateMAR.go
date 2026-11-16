package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
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
	Unknown                              []uint32                              `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid                              []uint32                              `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
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
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [7]bool
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
		case _ACCESS_FORWARDING_ACTION_INFORMATION_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			var v AccessForwardingActionInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.AccessForwardingActionInformation = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case NON__ACCESS_FORWARDING_ACTION_INFORMATION_TYPE:
			if seen[4] {
				continue
			}
			seen[4] = true
			var v NonAccessForwardingActionInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.NonAccessForwardingActionInformation = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case THRESHOLDS_TYPE:
			if seen[5] {
				continue
			}
			seen[5] = true
			t.ThresholdValues = de.Value
		case STEERING_MODE_INDICATOR_TYPE:
			if seen[6] {
				continue
			}
			seen[6] = true
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
	if !seen[1] {
		return &pfcp.MissingIE{Type: STEERING_FUNCTIONALITY_TYPE}
	}
	if !seen[2] {
		return &pfcp.MissingIE{Type: STEERING_MODE_TYPE}
	}
	return
}
