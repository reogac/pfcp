package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Update FAR' data structure */
type UpdateFAR struct {
	FARID                                     uint32                                     //FAR ID
	ApplyAction                               *ApplyAction                               //Apply Action
	UpdateForwardingparameters                *UpdateForwardingParameters                //Update Forwarding Parameters
	UpdateDuplicatingParameters               []UpdateDuplicatingParameters              //Update Duplicating Parameters
	RedundantTransmissionForwardingParameters *RedundantTransmissionForwardingParameters //Redundant Transmission Forwarding Parameters
	BARID                                     *uint8                                     //BAR ID
	AddMBSUnicastParameters                   []AddMBSUnicastParameters                  //Add MBS Unicast Parameters
	RemoveMBSUnicastParameters                [][]byte                                   //Remove MBS Unicast Parameters
	Unknown                                   []uint32                                   `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid                                   []uint32                                   `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
}

func (t *UpdateFAR) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(FAR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(t.FARID))))
	if t.ApplyAction != nil {
		fields = append(fields, tlv.TLVFrom(APPLY_ACTION_TYPE, tlv.Bytes(t.ApplyAction.Bytes())))
	}
	if t.UpdateForwardingparameters != nil {
		fields = append(fields, t.UpdateForwardingparameters.Field())
	}
	for _, ie := range t.UpdateDuplicatingParameters {
		fields = append(fields, ie.Field())
	}
	if t.RedundantTransmissionForwardingParameters != nil {
		fields = append(fields, t.RedundantTransmissionForwardingParameters.Field())
	}
	if t.BARID != nil {
		fields = append(fields, tlv.TLVFrom(BAR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint8(*t.BARID))))
	}
	for _, ie := range t.AddMBSUnicastParameters {
		fields = append(fields, ie.Field())
	}
	for _, ie := range t.RemoveMBSUnicastParameters {
		fields = append(fields, tlv.TLVFrom(REMOVE_MBS_UNICAST_PARAMETERS_TYPE, tlv.Bytes(ie)))
	}
	return tlv.TLV(UPDATE_FAR_TYPE, fields...)
}

func (t *UpdateFAR) UnmarshalBinary(buf []byte) (err error) {
	*t = UpdateFAR{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [5]bool
	for _, de := range elements {
		opt := false
		switch de.Type {
		case FAR_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &t.FARID)
		case APPLY_ACTION_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			var v ApplyAction
			if err = de.UnmarshalValue(&v); err == nil {
				t.ApplyAction = &v
			}
		case UPDATE_FORWARDING_PARAMETERS_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			var v UpdateForwardingParameters
			if err = de.UnmarshalValue(&v); err == nil {
				t.UpdateForwardingparameters = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case UPDATE_DUPLICATING_PARAMETERS_TYPE:
			var v UpdateDuplicatingParameters
			if err = de.UnmarshalValue(&v); err == nil {
				t.UpdateDuplicatingParameters = append(t.UpdateDuplicatingParameters, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case REDUNDANT_TRANSMISSION_FORWARDING_PARAMETERS_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			var v RedundantTransmissionForwardingParameters
			if err = de.UnmarshalValue(&v); err == nil {
				t.RedundantTransmissionForwardingParameters = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case BAR_ID_TYPE:
			if seen[4] {
				continue
			}
			seen[4] = true
			var v uint8
			if err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &v); err == nil {
				t.BARID = &v
			}
		case ADD_MBS_UNICAST_PARAMETERS_TYPE:
			var v AddMBSUnicastParameters
			if err = de.UnmarshalValue(&v); err == nil {
				t.AddMBSUnicastParameters = append(t.AddMBSUnicastParameters, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case REMOVE_MBS_UNICAST_PARAMETERS_TYPE:
			t.RemoveMBSUnicastParameters = append(t.RemoveMBSUnicastParameters, de.Value)
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
		return &pfcp.MissingIE{Type: FAR_ID_TYPE}
	}
	return
}
