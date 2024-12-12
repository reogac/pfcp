package message

import (
	"etrib5gc/pfcp/ie"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'PFCP Version Not Supported Response' data structure */
type PFCPVersionNotSupportedResponse struct {
}

func (t *PFCPVersionNotSupportedResponse) Field() tlv.Field {
	fields := []tlv.Field{}
	return tlv.TLV(ie.MSG_PFCP_VERSION_NOT_SUPPORTED_RESPONSE_TYPE, fields...)
}

func (t *PFCPVersionNotSupportedResponse) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPVersionNotSupportedResponse{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}

func (t *PFCPVersionNotSupportedResponse) UnmarshalTLV(typ uint32, buf []byte) (err error) {
	if typ != ie.MSG_PFCP_VERSION_NOT_SUPPORTED_RESPONSE_TYPE {
		err = tlv.ErrCritical
		return
	}

	err = t.UnmarshalBinary(buf)
	return
}

func (t *PFCPVersionNotSupportedResponse) MarshalJSON() (buf []byte, err error) {
	buf, err = jsonMarshal(t)
	return
}

func (t *PFCPVersionNotSupportedResponse) UnmarshalJSON(buf []byte) (err error) {
	*t = PFCPVersionNotSupportedResponse{}
	err = jsonUnmarshal(buf, t)
	return
}
