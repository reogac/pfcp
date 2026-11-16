package message

import (
	"github.com/reogac/pfcp/v2"
	"github.com/reogac/pfcp/v2/ie"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'PFCP Version Not Supported Response' data structure */
type PFCPVersionNotSupportedResponse struct {
	Unknown []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid []uint32 `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
}

func (t *PFCPVersionNotSupportedResponse) Field() tlv.Field {
	fields := []tlv.Field{}
	return tlv.TLV(ie.MSG_PFCP_VERSION_NOT_SUPPORTED_RESPONSE_TYPE, fields...)
}

// Skipped is what decoding left out: IEs unknown or not defined here
// (TS 29.244 7.6.9), and optional ones discarded for a faulty value (7.6.8).
func (t *PFCPVersionNotSupportedResponse) Skipped() (unknown, invalid []uint32) {
	return t.Unknown, t.Invalid
}

func (t *PFCPVersionNotSupportedResponse) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPVersionNotSupportedResponse{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	for _, de := range elements {
		opt := false
		switch de.Type {
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
