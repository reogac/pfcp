package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Update Duplicating Parameters' data structure */
type UpdateDuplicatingParameters struct {
	DestinationInterface  *uint8               //Destination Interface
	OuterHeaderCreation   *OuterHeaderCreation //Outer Header Creation
	TransportLevelMarking []byte               //Transport Level Marking
	ForwardingPolicy      []byte               //Forwarding Policy
	Unknown               []uint32             `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
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
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	var seen [4]bool
	for _, de := range elements {
		switch de.Type {
		case DESTINATION_INTERFACE_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			var v uint8
			if err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &v); err == nil {
				t.DestinationInterface = &v
			}
		case OUTER_HEADER_CREATION_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			var v OuterHeaderCreation
			if err = de.UnmarshalValue(&v); err == nil {
				t.OuterHeaderCreation = &v
			}
		case TRANSPORT_LEVEL_MARKING_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			t.TransportLevelMarking = de.Value
		case FORWARDING_POLICY_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			t.ForwardingPolicy = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	return
}
