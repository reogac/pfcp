package ie

import (
	"github.com/reogac/pfcp"

	"github.com/reogac/pfcp/tlv"
)

/* 'PFD context' data structure */
type PFDcontext struct {
	PFDContents [][]byte //PFD contents
	Unknown     []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *PFDcontext) Field() tlv.Field {
	fields := []tlv.Field{}
	for _, ie := range t.PFDContents {
		fields = append(fields, tlv.TLVFrom(PFD_CONTENTS_TYPE, tlv.Bytes(ie)))
	}
	return tlv.TLV(PFD_CONTEXT_TYPE, fields...)
}

func (t *PFDcontext) UnmarshalBinary(buf []byte) (err error) {
	*t = PFDcontext{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	for _, de := range elements {
		switch de.Type {
		case PFD_CONTENTS_TYPE:
			t.PFDContents = append(t.PFDContents, de.Value)
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	if len(t.PFDContents) == 0 {
		return &pfcp.MissingIE{Type: PFD_CONTENTS_TYPE}
	}
	return
}
