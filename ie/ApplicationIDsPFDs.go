package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Application ID's PFDs' data structure */
type ApplicationIDsPFDs struct {
	ApplicationID ApplicationID //Application ID
	PFDcontext    []PFDcontext  //PFD context
	Unknown       []uint32      `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *ApplicationIDsPFDs) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(APPLICATION_ID_TYPE, tlv.Bytes(t.ApplicationID.Bytes())))
	for _, ie := range t.PFDcontext {
		fields = append(fields, ie.Field())
	}
	return tlv.TLV(APPLICATION_ID_S_PFDS_TYPE, fields...)
}

func (t *ApplicationIDsPFDs) UnmarshalBinary(buf []byte) (err error) {
	*t = ApplicationIDsPFDs{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [1]bool
	for _, de := range elements {
		switch de.Type {
		case APPLICATION_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			if err = de.UnmarshalValue(&t.ApplicationID); err == nil {
			}
		case PFD_CONTEXT_TYPE:
			var v PFDcontext
			if err = de.UnmarshalValue(&v); err == nil {
				t.PFDcontext = append(t.PFDcontext, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
			}
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return pfcp.Faulty(de.Type, err)
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: APPLICATION_ID_TYPE}
	}
	return
}
