package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'ATSSS Control Parameters' data structure */
type ATSSSControlParameters struct {
	MPTCPParameters   *MPTCPParameters   //MPTCP Parameters
	ATSSSLLParameters *ATSSSLLParameters //ATSSS-LL Parameters
	PMFParameters     *PMFParameters     //PMF Parameters
}

func (t *ATSSSControlParameters) Field() tlv.Field {
	fields := []tlv.Field{}
	if t.MPTCPParameters != nil {
		fields = append(fields, t.MPTCPParameters.Field())
	}
	if t.ATSSSLLParameters != nil {
		fields = append(fields, t.ATSSSLLParameters.Field())
	}
	if t.PMFParameters != nil {
		fields = append(fields, t.PMFParameters.Field())
	}
	return tlv.TLV(ATSSS_CONTROL_PARAMETERS_TYPE, fields...)
}

func (t *ATSSSControlParameters) UnmarshalBinary(buf []byte) (err error) {
	*t = ATSSSControlParameters{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case MPTCP_PARAMETERS_TYPE:
			var v MPTCPParameters
			if err = de.UnmarshalValue(&v); err == nil {
				t.MPTCPParameters = &v
			}
		case ATSSS_LL_PARAMETERS_TYPE:
			var v ATSSSLLParameters
			if err = de.UnmarshalValue(&v); err == nil {
				t.ATSSSLLParameters = &v
			}
		case PMF_PARAMETERS_TYPE:
			var v PMFParameters
			if err = de.UnmarshalValue(&v); err == nil {
				t.PMFParameters = &v
			}
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
