package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Redundant Transmission Parameters' data structure */
type RedundantTransmissionParameters struct {
	NetworkInstanceforRedundantTransmission []byte //Network Instance
}

func (t *RedundantTransmissionParameters) Field() tlv.Field {
	fields := []tlv.Field{}
	if len(t.NetworkInstanceforRedundantTransmission) > 0 {
		fields = append(fields, tlv.TLVFrom(NETWORK_INSTANCE_TYPE, tlv.Bytes(t.NetworkInstanceforRedundantTransmission)))
	}
	return tlv.TLV(REDUNDANT_TRANSMISSION_PARAMETERS_TYPE, fields...)
}

func (t *RedundantTransmissionParameters) UnmarshalBinary(buf []byte) (err error) {
	*t = RedundantTransmissionParameters{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
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
