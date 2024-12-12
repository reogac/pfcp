package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Created Bridge Info for TSC' data structure */
type CreatedBridgeInfoforTSC struct {
	DSTTPortNumber      []byte //DS-TT Port Number
	fivegsUserPlaneNode []byte //5GS User Plane Node
}

func (t *CreatedBridgeInfoforTSC) Field() tlv.Field {
	fields := []tlv.Field{}
	if len(t.DSTTPortNumber) > 0 {
		fields = append(fields, tlv.TLVFrom(DS_TT_PORT_NUMBER_TYPE, tlv.Bytes(t.DSTTPortNumber)))
	}
	if len(t.fivegsUserPlaneNode) > 0 {
		fields = append(fields, tlv.TLVFrom(FIVEGS_USER_PLANE_NODE_TYPE, tlv.Bytes(t.fivegsUserPlaneNode)))
	}
	return tlv.TLV(CREATED_BRIDGE_INFO_FOR_TSC_TYPE, fields...)
}

func (t *CreatedBridgeInfoforTSC) UnmarshalBinary(buf []byte) (err error) {
	*t = CreatedBridgeInfoforTSC{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case DS_TT_PORT_NUMBER_TYPE:
			t.DSTTPortNumber = de.Value
		case FIVEGS_USER_PLANE_NODE_TYPE:
			t.fivegsUserPlaneNode = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
