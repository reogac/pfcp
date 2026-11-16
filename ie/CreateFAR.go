package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Create FAR' data structure */
type CreateFAR struct {
	FARID                                     uint32                                     //FAR ID
	ApplyAction                               ApplyAction                                //Apply Action
	ForwardingParameters                      *ForwardingParameters                      //Forwarding Parameters
	DuplicatingParameters                     []DuplicatingParameters                    //Duplicating Parameters
	BARID                                     *uint8                                     //BAR ID
	RedundantTransmissionForwardingParameters *RedundantTransmissionForwardingParameters //Redundant Transmission Forwarding Parameters
	MBSMulticastParameters                    *MBSMulticastParameters                    //MBS Multicast Parameters
	AddMBSUnicastParameters                   []AddMBSUnicastParameters                  //Add MBS Unicast Parameters
	Unknown                                   []uint32                                   `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid                                   []uint32                                   `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
}

func (t *CreateFAR) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(FAR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(t.FARID))))
	fields = append(fields, tlv.TLVFrom(APPLY_ACTION_TYPE, tlv.Bytes(t.ApplyAction.Bytes())))
	if t.ForwardingParameters != nil {
		fields = append(fields, t.ForwardingParameters.Field())
	}
	for _, ie := range t.DuplicatingParameters {
		fields = append(fields, ie.Field())
	}
	if t.BARID != nil {
		fields = append(fields, tlv.TLVFrom(BAR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint8(*t.BARID))))
	}
	if t.RedundantTransmissionForwardingParameters != nil {
		fields = append(fields, t.RedundantTransmissionForwardingParameters.Field())
	}
	if t.MBSMulticastParameters != nil {
		fields = append(fields, t.MBSMulticastParameters.Field())
	}
	for _, ie := range t.AddMBSUnicastParameters {
		fields = append(fields, ie.Field())
	}
	return tlv.TLV(CREATE_FAR_TYPE, fields...)
}

func (t *CreateFAR) UnmarshalBinary(buf []byte) (err error) {
	*t = CreateFAR{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [6]bool
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
			if err = de.UnmarshalValue(&t.ApplyAction); err == nil {
			}
		case FORWARDING_PARAMETERS_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			var v ForwardingParameters
			if err = de.UnmarshalValue(&v); err == nil {
				t.ForwardingParameters = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case DUPLICATING_PARAMETERS_TYPE:
			var v DuplicatingParameters
			if err = de.UnmarshalValue(&v); err == nil {
				t.DuplicatingParameters = append(t.DuplicatingParameters, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case BAR_ID_TYPE:
			opt = true
			if seen[3] {
				continue
			}
			seen[3] = true
			var v uint8
			if err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &v); err == nil {
				t.BARID = &v
			}
		case REDUNDANT_TRANSMISSION_FORWARDING_PARAMETERS_TYPE:
			if seen[4] {
				continue
			}
			seen[4] = true
			var v RedundantTransmissionForwardingParameters
			if err = de.UnmarshalValue(&v); err == nil {
				t.RedundantTransmissionForwardingParameters = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case MBS_MULTICAST_PARAMETERS_TYPE:
			if seen[5] {
				continue
			}
			seen[5] = true
			var v MBSMulticastParameters
			if err = de.UnmarshalValue(&v); err == nil {
				t.MBSMulticastParameters = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ADD_MBS_UNICAST_PARAMETERS_TYPE:
			var v AddMBSUnicastParameters
			if err = de.UnmarshalValue(&v); err == nil {
				t.AddMBSUnicastParameters = append(t.AddMBSUnicastParameters, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
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
	if !seen[1] {
		return &pfcp.MissingIE{Type: APPLY_ACTION_TYPE}
	}
	return
}
