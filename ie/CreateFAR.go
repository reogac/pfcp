package ie

import (
	"etrib5gc/pfcp"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Create FAR' data structure */
type CreateFAR struct {
	FARID                                     uint32                                     //FAR ID
	ApplyAction                               ApplyAction                                //Apply Action
	ForwardingParameters                      *ForwardingParameters                      //Forwarding Parameters
	DuplicatingParameters                     *DuplicatingParameters                     //Duplicating Parameters
	BARID                                     *uint8                                     //BAR ID
	RedundantTransmissionForwardingParameters *RedundantTransmissionForwardingParameters //Redundant Transmission Forwarding Parameters
	MBSMulticastParameters                    *MBSMulticastParameters                    //MBS Multicast Parameters
	AddMBSUnicastParameters                   *AddMBSUnicastParameters                   //Add MBS Unicast Parameters
}

func (t *CreateFAR) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(FAR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(t.FARID))))
	fields = append(fields, tlv.TLVFrom(APPLY_ACTION_TYPE, tlv.Bytes(t.ApplyAction.Bytes())))
	if t.ForwardingParameters != nil {
		fields = append(fields, t.ForwardingParameters.Field())
	}
	if t.DuplicatingParameters != nil {
		fields = append(fields, t.DuplicatingParameters.Field())
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
	if t.AddMBSUnicastParameters != nil {
		fields = append(fields, t.AddMBSUnicastParameters.Field())
	}
	return tlv.TLV(CREATE_FAR_TYPE, fields...)
}

func (t *CreateFAR) UnmarshalBinary(buf []byte) (err error) {
	*t = CreateFAR{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case FAR_ID_TYPE:
			err = pfcp.UnmarshalUint32(de.Value, &t.FARID)
		case APPLY_ACTION_TYPE:
			err = de.UnmarshalValue(&t.ApplyAction)
		case FORWARDING_PARAMETERS_TYPE:
			var v ForwardingParameters
			if err = de.UnmarshalValue(&v); err == nil {
				t.ForwardingParameters = &v
			}
		case DUPLICATING_PARAMETERS_TYPE:
			var v DuplicatingParameters
			if err = de.UnmarshalValue(&v); err == nil {
				t.DuplicatingParameters = &v
			}
		case BAR_ID_TYPE:
			var v uint8
			if err = pfcp.UnmarshalUint8(de.Value, &v); err == nil {
				t.BARID = &v
			}
		case REDUNDANT_TRANSMISSION_FORWARDING_PARAMETERS_TYPE:
			var v RedundantTransmissionForwardingParameters
			if err = de.UnmarshalValue(&v); err == nil {
				t.RedundantTransmissionForwardingParameters = &v
			}
		case MBS_MULTICAST_PARAMETERS_TYPE:
			var v MBSMulticastParameters
			if err = de.UnmarshalValue(&v); err == nil {
				t.MBSMulticastParameters = &v
			}
		case ADD_MBS_UNICAST_PARAMETERS_TYPE:
			var v AddMBSUnicastParameters
			if err = de.UnmarshalValue(&v); err == nil {
				t.AddMBSUnicastParameters = &v
			}
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
