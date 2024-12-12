package ie

import (
	"etrib5gc/pfcp"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Application Detection Information' data structure */
type ApplicationDetectionInformation struct {
	ApplicationID         ApplicationID //Application ID
	ApplicationInstanceID []byte        //Application Instance ID
	FlowInformation       []byte        //Flow Information
	PDRID                 *uint16       //PDR ID
}

func (t *ApplicationDetectionInformation) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(APPLICATION_ID_TYPE, tlv.Bytes(t.ApplicationID.Bytes())))
	if len(t.ApplicationInstanceID) > 0 {
		fields = append(fields, tlv.TLVFrom(APPLICATION_INSTANCE_ID_TYPE, tlv.Bytes(t.ApplicationInstanceID)))
	}
	if len(t.FlowInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(FLOW_INFORMATION_TYPE, tlv.Bytes(t.FlowInformation)))
	}
	if t.PDRID != nil {
		fields = append(fields, tlv.TLVFrom(PDR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint16(*t.PDRID))))
	}
	return tlv.TLV(APPLICATION_DETECTION_INFORMATION_TYPE, fields...)
}

func (t *ApplicationDetectionInformation) UnmarshalBinary(buf []byte) (err error) {
	*t = ApplicationDetectionInformation{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case APPLICATION_ID_TYPE:
			err = de.UnmarshalValue(&t.ApplicationID)
		case APPLICATION_INSTANCE_ID_TYPE:
			t.ApplicationInstanceID = de.Value
		case FLOW_INFORMATION_TYPE:
			t.FlowInformation = de.Value
		case PDR_ID_TYPE:
			var v uint16
			if err = pfcp.UnmarshalUint16(de.Value, &v); err == nil {
				t.PDRID = &v
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
