package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Redundant Transmission Forwarding Parameters' data structure */
type RedundantTransmissionForwardingParameters struct {
	OuterHeaderCreation                     OuterHeaderCreation //Outer Header Creation
	NetworkInstanceforRedundantTransmission []byte              //Network Instance
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
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case OUTER_HEADER_CREATION_TYPE:
			err = de.UnmarshalValue(&t.OuterHeaderCreation)
		case NETWORK_INSTANCE_TYPE:
			t.NetworkInstanceforRedundantTransmission = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
