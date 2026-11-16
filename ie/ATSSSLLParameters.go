package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'ATSSS-LL Parameters' data structure */
type ATSSSLLParameters struct {
	ATSSSLLInformation []byte   //ATSSS-LL Information
	Unknown            []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid            []uint32 `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
}

func (t *ATSSSLLParameters) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(ATSSS_LL_INFORMATION_TYPE, tlv.Bytes(t.ATSSSLLInformation)))
	return tlv.TLV(ATSSS_LL_PARAMETERS_TYPE, fields...)
}

func (t *ATSSSLLParameters) UnmarshalBinary(buf []byte) (err error) {
	*t = ATSSSLLParameters{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [1]bool
	for _, de := range elements {
		opt := false
		switch de.Type {
		case ATSSS_LL_INFORMATION_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.ATSSSLLInformation = de.Value
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
	if !seen[0] {
		return &pfcp.MissingIE{Type: ATSSS_LL_INFORMATION_TYPE}
	}
	return
}
