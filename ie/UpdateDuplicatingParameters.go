package ie

import (
	"etrib5gc/pfcp"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Update Duplicating Parameters' data structure */
type UpdateDuplicatingParameters struct {
	DestinationInterface  *uint8               //Destination Interface
	OuterHeaderCreation   *OuterHeaderCreation //Outer Header Creation
	TransportLevelMarking []byte               //Transport Level Marking
	ForwardingPolicy      []byte               //Forwarding Policy
}

func (t *UpdateDuplicatingParameters) Field() tlv.Field {
	fields := []tlv.Field{}
	if t.DestinationInterface != nil {
		fields = append(fields, tlv.TLVFrom(DESTINATION_INTERFACE_TYPE, tlv.Bytes(pfcp.MarshalUint8(*t.DestinationInterface))))
	}
	if t.OuterHeaderCreation != nil {
		fields = append(fields, tlv.TLVFrom(OUTER_HEADER_CREATION_TYPE, tlv.Bytes(t.OuterHeaderCreation.Bytes())))
	}
	if len(t.TransportLevelMarking) > 0 {
		fields = append(fields, tlv.TLVFrom(TRANSPORT_LEVEL_MARKING_TYPE, tlv.Bytes(t.TransportLevelMarking)))
	}
	if len(t.ForwardingPolicy) > 0 {
		fields = append(fields, tlv.TLVFrom(FORWARDING_POLICY_TYPE, tlv.Bytes(t.ForwardingPolicy)))
	}
	return tlv.TLV(UPDATE_DUPLICATING_PARAMETERS_TYPE, fields...)
}

func (t *UpdateDuplicatingParameters) UnmarshalBinary(buf []byte) (err error) {
	*t = UpdateDuplicatingParameters{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case DESTINATION_INTERFACE_TYPE:
			var v uint8
			if err = pfcp.UnmarshalUint8(de.Value, &v); err == nil {
				t.DestinationInterface = &v
			}
		case OUTER_HEADER_CREATION_TYPE:
			var v OuterHeaderCreation
			if err = de.UnmarshalValue(&v); err == nil {
				t.OuterHeaderCreation = &v
			}
		case TRANSPORT_LEVEL_MARKING_TYPE:
			t.TransportLevelMarking = de.Value
		case FORWARDING_POLICY_TYPE:
			t.ForwardingPolicy = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
