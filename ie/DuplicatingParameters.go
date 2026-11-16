package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Duplicating Parameters' data structure */
type DuplicatingParameters struct {
	DestinationInterface  uint8                //Destination Interface
	OuterHeaderCreation   *OuterHeaderCreation //Outer Header Creation
	TransportLevelmarking []byte               //Transport Level Marking
	ForwardingPolicy      []byte               //Forwarding Policy
	Unknown               []uint32             `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid               []uint32             `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
}

func (t *DuplicatingParameters) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(DESTINATION_INTERFACE_TYPE, tlv.Bytes(pfcp.MarshalUint8(t.DestinationInterface))))
	if t.OuterHeaderCreation != nil {
		fields = append(fields, tlv.TLVFrom(OUTER_HEADER_CREATION_TYPE, tlv.Bytes(t.OuterHeaderCreation.Bytes())))
	}
	if len(t.TransportLevelmarking) > 0 {
		fields = append(fields, tlv.TLVFrom(TRANSPORT_LEVEL_MARKING_TYPE, tlv.Bytes(t.TransportLevelmarking)))
	}
	if len(t.ForwardingPolicy) > 0 {
		fields = append(fields, tlv.TLVFrom(FORWARDING_POLICY_TYPE, tlv.Bytes(t.ForwardingPolicy)))
	}
	return tlv.TLV(DUPLICATING_PARAMETERS_TYPE, fields...)
}

func (t *DuplicatingParameters) UnmarshalBinary(buf []byte) (err error) {
	*t = DuplicatingParameters{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [4]bool
	for _, de := range elements {
		opt := false
		switch de.Type {
		case DESTINATION_INTERFACE_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &t.DestinationInterface)
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
			t.TransportLevelmarking = de.Value
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
			if err = pfcp.Faulty(de.Type, err); opt && pfcp.Discardable(err) {
				t.Invalid = append(t.Invalid, de.Type)
				err = nil
				continue
			}
			return err
		}
	}
	return
}
