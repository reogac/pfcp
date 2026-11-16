package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'DSCP to PPI Control Information' data structure */
type DSCPtoPPIControlInformation struct {
	DSCPtoPPIMappingInformation [][]byte //DSCP to PPI Mapping Information
	QFI                         []uint8  //QFI
	Unknown                     []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid                     []uint32 `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
}

func (t *DSCPtoPPIControlInformation) Field() tlv.Field {
	fields := []tlv.Field{}
	for _, ie := range t.DSCPtoPPIMappingInformation {
		fields = append(fields, tlv.TLVFrom(DSCP_TO_PPI_MAPPING_INFORMATION_TYPE, tlv.Bytes(ie)))
	}
	for _, ie := range t.QFI {
		fields = append(fields, tlv.TLVFrom(QFI_TYPE, tlv.Bytes(pfcp.MarshalUint8(ie))))
	}
	return tlv.TLV(DSCP_TO_PPI_CONTROL_INFORMATION_TYPE, fields...)
}

func (t *DSCPtoPPIControlInformation) UnmarshalBinary(buf []byte) (err error) {
	*t = DSCPtoPPIControlInformation{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	for _, de := range elements {
		opt := false
		switch de.Type {
		case DSCP_TO_PPI_MAPPING_INFORMATION_TYPE:
			t.DSCPtoPPIMappingInformation = append(t.DSCPtoPPIMappingInformation, de.Value)
		case QFI_TYPE:
			opt = true
			var v uint8
			if err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &v); err == nil {
				t.QFI = append(t.QFI, v)
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
	if len(t.DSCPtoPPIMappingInformation) == 0 {
		return &pfcp.MissingIE{Type: DSCP_TO_PPI_MAPPING_INFORMATION_TYPE}
	}
	return
}
