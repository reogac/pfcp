package ie

import (
	"github.com/reogac/pfcp/tlv"
)

/* 'Provide ATSSS Control Information' data structure */
type ProvideATSSSControlInformation struct {
	MPTCPControlInformation   []byte //MPTCP Control Information
	ATSSSLLControlInformation []byte //ATSSS-LL Control Information
	PMFControlInformation     []byte //PMF Control Information
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
	for _, de := range elements {
		switch de.Type {
		case MPTCP_CONTROL_INFORMATION_TYPE:
			t.MPTCPControlInformation = de.Value
		case ATSSS_LL_CONTROL_INFORMATION_TYPE:
			t.ATSSSLLControlInformation = de.Value
		case PMF_CONTROL_INFORMATION_TYPE:
			t.PMFControlInformation = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
