package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Application ID's PFDs' data structure */
type ApplicationIDsPFDs struct {
	ApplicationID ApplicationID //Application ID
	PFDcontext    *PFDcontext   //PFD context
}

func (t *ApplicationIDsPFDs) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(APPLICATION_ID_TYPE, tlv.Bytes(t.ApplicationID.Bytes())))
	if t.PFDcontext != nil {
		fields = append(fields, t.PFDcontext.Field())
	}
	return tlv.TLV(APPLICATION_ID_S_PFDS_TYPE, fields...)
}

func (t *ApplicationIDsPFDs) UnmarshalBinary(buf []byte) (err error) {
	*t = ApplicationIDsPFDs{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case APPLICATION_ID_TYPE:
			err = de.UnmarshalValue(&t.ApplicationID)
		case PFD_CONTEXT_TYPE:
			var v PFDcontext
			if err = de.UnmarshalValue(&v); err == nil {
				t.PFDcontext = &v
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
