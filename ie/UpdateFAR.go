package ie

import (
	"etrib5gc/pfcp"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Update FAR' data structure */
type UpdateFAR struct {
	FARID                                     uint32                                     //FAR ID
	ApplyAction                               *ApplyAction                               //Apply Action
	UpdateForwardingparameters                *UpdateForwardingParameters                //Update Forwarding Parameters
	UpdateDuplicatingParameters               *UpdateDuplicatingParameters               //Update Duplicating Parameters
	RedundantTransmissionForwardingParameters *RedundantTransmissionForwardingParameters //Redundant Transmission Forwarding Parameters
	BARID                                     *uint8                                     //BAR ID
	AddMBSUnicastParameters                   *AddMBSUnicastParameters                   //Add MBS Unicast Parameters
	RemoveMBSUnicastParameters                []byte                                     //Remove MBS Unicast Parameters
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
	if t.UpdateDuplicatingParameters != nil {
		fields = append(fields, t.UpdateDuplicatingParameters.Field())
	}
	if t.RedundantTransmissionForwardingParameters != nil {
		fields = append(fields, t.RedundantTransmissionForwardingParameters.Field())
	}
	if t.BARID != nil {
		fields = append(fields, tlv.TLVFrom(BAR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint8(*t.BARID))))
	}
	if t.AddMBSUnicastParameters != nil {
		fields = append(fields, t.AddMBSUnicastParameters.Field())
	}
	if len(t.RemoveMBSUnicastParameters) > 0 {
		fields = append(fields, tlv.TLVFrom(REMOVE_MBS_UNICAST_PARAMETERS_TYPE, tlv.Bytes(t.RemoveMBSUnicastParameters)))
	}
	return tlv.TLV(UPDATE_FAR_TYPE, fields...)
}

func (t *UpdateFAR) UnmarshalBinary(buf []byte) (err error) {
	*t = UpdateFAR{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case FAR_ID_TYPE:
			err = pfcp.UnmarshalUint32(de.Value, &t.FARID)
		case APPLY_ACTION_TYPE:
			var v ApplyAction
			if err = de.UnmarshalValue(&v); err == nil {
				t.ApplyAction = &v
			}
		case UPDATE_FORWARDING_PARAMETERS_TYPE:
			var v UpdateForwardingParameters
			if err = de.UnmarshalValue(&v); err == nil {
				t.UpdateForwardingparameters = &v
			}
		case UPDATE_DUPLICATING_PARAMETERS_TYPE:
			var v UpdateDuplicatingParameters
			if err = de.UnmarshalValue(&v); err == nil {
				t.UpdateDuplicatingParameters = &v
			}
		case REDUNDANT_TRANSMISSION_FORWARDING_PARAMETERS_TYPE:
			var v RedundantTransmissionForwardingParameters
			if err = de.UnmarshalValue(&v); err == nil {
				t.RedundantTransmissionForwardingParameters = &v
			}
		case BAR_ID_TYPE:
			var v uint8
			if err = pfcp.UnmarshalUint8(de.Value, &v); err == nil {
				t.BARID = &v
			}
		case ADD_MBS_UNICAST_PARAMETERS_TYPE:
			var v AddMBSUnicastParameters
			if err = de.UnmarshalValue(&v); err == nil {
				t.AddMBSUnicastParameters = &v
			}
		case REMOVE_MBS_UNICAST_PARAMETERS_TYPE:
			t.RemoveMBSUnicastParameters = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
