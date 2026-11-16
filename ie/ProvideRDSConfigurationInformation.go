package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Provide RDS Configuration Information' data structure */
type ProvideRDSConfigurationInformation struct {
	RDSConfigurationInformation []byte   //RDS Configuration Information
	Unknown                     []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid                     []uint32 `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
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
		return pfcp.Faulty(0, err)
	}
	var seen [1]bool
	for _, de := range elements {
		opt := false
		switch de.Type {
		case RDS_CONFIGURATION_INFORMATION_TYPE:
			opt = true
			if seen[0] {
				continue
			}
			seen[0] = true
			t.RDSConfigurationInformation = de.Value
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
