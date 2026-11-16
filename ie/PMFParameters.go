package ie

import (
	"github.com/reogac/pfcp"

	"github.com/reogac/pfcp/tlv"
)

/* 'PMF Parameters' data structure */
type PMFParameters struct {
	PMFAddressInformation []byte   //PMF Address Information
	QoSflowidentifier     *uint8   //QFI
	Unknown               []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *PMFParameters) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(PMF_ADDRESS_INFORMATION_TYPE, tlv.Bytes(t.PMFAddressInformation)))
	if t.QoSflowidentifier != nil {
		fields = append(fields, tlv.TLVFrom(QFI_TYPE, tlv.Bytes(pfcp.MarshalUint8(*t.QoSflowidentifier))))
	}
	return tlv.TLV(PMF_PARAMETERS_TYPE, fields...)
}

func (t *PMFParameters) UnmarshalBinary(buf []byte) (err error) {
	*t = PMFParameters{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	var seen [2]bool
	for _, de := range elements {
		switch de.Type {
		case PMF_ADDRESS_INFORMATION_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.PMFAddressInformation = de.Value
		case QFI_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			var v uint8
			if err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &v); err == nil {
				t.QoSflowidentifier = &v
			}
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: PMF_ADDRESS_INFORMATION_TYPE}
	}
	return
}
