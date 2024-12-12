package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'ATSSS-LL Parameters' data structure */
type ATSSSLLParameters struct {
	ATSSSLLInformation []byte //ATSSS-LL Information
}

func (t *ATSSSLLParameters) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(ATSSS_LL_INFORMATION_TYPE, tlv.Bytes(t.ATSSSLLInformation)))
	return tlv.TLV(ATSSS_LL_PARAMETERS_TYPE, fields...)
}

func (t *ATSSSLLParameters) UnmarshalBinary(buf []byte) (err error) {
	*t = ATSSSLLParameters{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case ATSSS_LL_INFORMATION_TYPE:
			t.ATSSSLLInformation = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
