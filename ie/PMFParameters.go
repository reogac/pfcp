package ie

import (
	"github.com/reogac/pfcp"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'PMF Parameters' data structure */
type PMFParameters struct {
	PMFAddressInformation []byte //PMF Address Information
	QoSflowidentifier     *uint8 //QFI
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
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case PMF_ADDRESS_INFORMATION_TYPE:
			t.PMFAddressInformation = de.Value
		case QFI_TYPE:
			var v uint8
			if err = pfcp.UnmarshalUint8(de.Value, &v); err == nil {
				t.QoSflowidentifier = &v
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
