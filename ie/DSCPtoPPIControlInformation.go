package ie

import (
	"etrib5gc/pfcp"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'DSCP to PPI Control Information' data structure */
type DSCPtoPPIControlInformation struct {
	DSCPtoPPIMappingInformation []byte //DSCP to PPI Mapping Information
	QFI                         *uint8 //QFI
}

func (t *DSCPtoPPIControlInformation) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(DSCP_TO_PPI_MAPPING_INFORMATION_TYPE, tlv.Bytes(t.DSCPtoPPIMappingInformation)))
	if t.QFI != nil {
		fields = append(fields, tlv.TLVFrom(QFI_TYPE, tlv.Bytes(pfcp.MarshalUint8(*t.QFI))))
	}
	return tlv.TLV(DSCP_TO_PPI_CONTROL_INFORMATION_TYPE, fields...)
}

func (t *DSCPtoPPIControlInformation) UnmarshalBinary(buf []byte) (err error) {
	*t = DSCPtoPPIControlInformation{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case DSCP_TO_PPI_MAPPING_INFORMATION_TYPE:
			t.DSCPtoPPIMappingInformation = de.Value
		case QFI_TYPE:
			var v uint8
			if err = pfcp.UnmarshalUint8(de.Value, &v); err == nil {
				t.QFI = &v
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
