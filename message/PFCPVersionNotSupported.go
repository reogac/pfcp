package message

import (
	"github.com/reogac/pfcp/v2/ie"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'PFCP Version Not Supported Response' data structure */
type PFCPVersionNotSupportedResponse struct {
	Unknown []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *PFCPVersionNotSupportedResponse) Field() tlv.Field {
	fields := []tlv.Field{}
	return tlv.TLV(ie.MSG_PFCP_VERSION_NOT_SUPPORTED_RESPONSE_TYPE, fields...)
}

func (t *PFCPVersionNotSupportedResponse) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPVersionNotSupportedResponse{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	for _, de := range elements {
		switch de.Type {
		default:
			t.Unknown = append(t.Unknown, de.Type)
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
