package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Redundant Transmission Parameters' data structure */
type RedundantTransmissionParameters struct {
	LocalFTEIDforRedundantTransmission      FTEID    //F-TEID
	NetworkInstanceforRedundantTransmission []byte   //Network Instance
	Unknown                                 []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *RedundantTransmissionParameters) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(F_TEID_TYPE, tlv.Bytes(t.LocalFTEIDforRedundantTransmission.Bytes())))
	if len(t.NetworkInstanceforRedundantTransmission) > 0 {
		fields = append(fields, tlv.TLVFrom(NETWORK_INSTANCE_TYPE, tlv.Bytes(t.NetworkInstanceforRedundantTransmission)))
	}
	return tlv.TLV(REDUNDANT_TRANSMISSION_PARAMETERS_TYPE, fields...)
}

func (t *RedundantTransmissionParameters) UnmarshalBinary(buf []byte) (err error) {
	*t = RedundantTransmissionParameters{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [2]bool
	for _, de := range elements {
		switch de.Type {
		case F_TEID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			if err = de.UnmarshalValue(&t.LocalFTEIDforRedundantTransmission); err == nil {
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
			return pfcp.Faulty(de.Type, err)
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: F_TEID_TYPE}
	}
	return
}
