package ie

import (
	"github.com/reogac/pfcp/tlv"
)

/* 'Provide RDS Configuration Information' data structure */
type ProvideRDSConfigurationInformation struct {
	RDSConfigurationInformation []byte   //RDS Configuration Information
	Unknown                     []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
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
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	var seen [1]bool
	for _, de := range elements {
		switch de.Type {
		case RDS_CONFIGURATION_INFORMATION_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.RDSConfigurationInformation = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	return
}
