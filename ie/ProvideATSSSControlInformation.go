package ie

import (
	"github.com/reogac/pfcp/tlv"
)

/* 'Provide ATSSS Control Information' data structure */
type ProvideATSSSControlInformation struct {
	MPTCPControlInformation   []byte   //MPTCP Control Information
	ATSSSLLControlInformation []byte   //ATSSS-LL Control Information
	PMFControlInformation     []byte   //PMF Control Information
	Unknown                   []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *ProvideATSSSControlInformation) Field() tlv.Field {
	fields := []tlv.Field{}
	if len(t.MPTCPControlInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(MPTCP_CONTROL_INFORMATION_TYPE, tlv.Bytes(t.MPTCPControlInformation)))
	}
	if len(t.ATSSSLLControlInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(ATSSS_LL_CONTROL_INFORMATION_TYPE, tlv.Bytes(t.ATSSSLLControlInformation)))
	}
	if len(t.PMFControlInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(PMF_CONTROL_INFORMATION_TYPE, tlv.Bytes(t.PMFControlInformation)))
	}
	return tlv.TLV(PROVIDE_ATSSS_CONTROL_INFORMATION_TYPE, fields...)
}

func (t *ProvideATSSSControlInformation) UnmarshalBinary(buf []byte) (err error) {
	*t = ProvideATSSSControlInformation{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	var seen [3]bool
	for _, de := range elements {
		switch de.Type {
		case MPTCP_CONTROL_INFORMATION_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.MPTCPControlInformation = de.Value
		case ATSSS_LL_CONTROL_INFORMATION_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			t.ATSSSLLControlInformation = de.Value
		case PMF_CONTROL_INFORMATION_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			t.PMFControlInformation = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	return
}
