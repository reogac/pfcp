package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Provide RDS Configuration Information' data structure */
type ProvideRDSConfigurationInformation struct {
	RDSConfigurationInformation []byte //RDS Configuration Information
}

func (t *ProvideRDSConfigurationInformation) Field() tlv.Field {
	fields := []tlv.Field{}
	if len(t.RDSConfigurationInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(RDS_CONFIGURATION_INFORMATION_TYPE, tlv.Bytes(t.RDSConfigurationInformation)))
	}
	return tlv.TLV(PROVIDE_RDS_CONFIGURATION_INFORMATION_TYPE, fields...)
}

func (t *ProvideRDSConfigurationInformation) UnmarshalBinary(buf []byte) (err error) {
	*t = ProvideRDSConfigurationInformation{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case RDS_CONFIGURATION_INFORMATION_TYPE:
			t.RDSConfigurationInformation = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
