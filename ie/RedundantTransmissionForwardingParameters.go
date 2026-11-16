package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Redundant Transmission Forwarding Parameters' data structure */
type RedundantTransmissionForwardingParameters struct {
	OuterHeaderCreation                     OuterHeaderCreation //Outer Header Creation
	NetworkInstanceforRedundantTransmission []byte              //Network Instance
	Unknown                                 []uint32            `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *RedundantTransmissionForwardingParameters) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(OUTER_HEADER_CREATION_TYPE, tlv.Bytes(t.OuterHeaderCreation.Bytes())))
	if len(t.NetworkInstanceforRedundantTransmission) > 0 {
		fields = append(fields, tlv.TLVFrom(NETWORK_INSTANCE_TYPE, tlv.Bytes(t.NetworkInstanceforRedundantTransmission)))
	}
	return tlv.TLV(REDUNDANT_TRANSMISSION_FORWARDING_PARAMETERS_TYPE, fields...)
}

func (t *RedundantTransmissionForwardingParameters) UnmarshalBinary(buf []byte) (err error) {
	*t = RedundantTransmissionForwardingParameters{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	var seen [2]bool
	for _, de := range elements {
		switch de.Type {
		case OUTER_HEADER_CREATION_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			if err = de.UnmarshalValue(&t.OuterHeaderCreation); err == nil {
			}
		case NETWORK_INSTANCE_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			t.NetworkInstanceforRedundantTransmission = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: OUTER_HEADER_CREATION_TYPE}
	}
	return
}
