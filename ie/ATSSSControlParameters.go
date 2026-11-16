package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'ATSSS Control Parameters' data structure */
type ATSSSControlParameters struct {
	MPTCPParameters   *MPTCPParameters   //MPTCP Parameters
	ATSSSLLParameters *ATSSSLLParameters //ATSSS-LL Parameters
	PMFParameters     []PMFParameters    //PMF Parameters
	Unknown           []uint32           `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid           []uint32           `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
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
		return pfcp.Faulty(0, err)
	}
	var seen [2]bool
	for _, de := range elements {
		opt := false
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
				t.Invalid = append(t.Invalid, v.Invalid...)
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
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case PMF_PARAMETERS_TYPE:
			var v PMFParameters
			if err = de.UnmarshalValue(&v); err == nil {
				t.PMFParameters = append(t.PMFParameters, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
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
