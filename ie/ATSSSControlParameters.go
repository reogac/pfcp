package ie

import (
	"github.com/reogac/pfcp/v2/tlv"
)

/* 'ATSSS Control Parameters' data structure */
type ATSSSControlParameters struct {
	MPTCPParameters   *MPTCPParameters   //MPTCP Parameters
	ATSSSLLParameters *ATSSSLLParameters //ATSSS-LL Parameters
	PMFParameters     []PMFParameters    //PMF Parameters
	Unknown           []uint32           `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *ATSSSControlParameters) Field() tlv.Field {
	fields := []tlv.Field{}
	if t.MPTCPParameters != nil {
		fields = append(fields, t.MPTCPParameters.Field())
	}
	if t.ATSSSLLParameters != nil {
		fields = append(fields, t.ATSSSLLParameters.Field())
	}
	for _, ie := range t.PMFParameters {
		fields = append(fields, ie.Field())
	}
	return tlv.TLV(ATSSS_CONTROL_PARAMETERS_TYPE, fields...)
}

func (t *ATSSSControlParameters) UnmarshalBinary(buf []byte) (err error) {
	*t = ATSSSControlParameters{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	var seen [2]bool
	for _, de := range elements {
		switch de.Type {
		case MPTCP_PARAMETERS_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			var v MPTCPParameters
			if err = de.UnmarshalValue(&v); err == nil {
				t.MPTCPParameters = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
			}
		case ATSSS_LL_PARAMETERS_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			var v ATSSSLLParameters
			if err = de.UnmarshalValue(&v); err == nil {
				t.ATSSSLLParameters = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
			}
		case PMF_PARAMETERS_TYPE:
			var v PMFParameters
			if err = de.UnmarshalValue(&v); err == nil {
				t.PMFParameters = append(t.PMFParameters, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
			}
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	return
}
