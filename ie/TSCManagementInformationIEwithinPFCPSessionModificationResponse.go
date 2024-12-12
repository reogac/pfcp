package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'TSC Management Information IE within PFCP Session Modification Response' data structure */
type TSCManagementInformationIEwithinPFCPSessionModificationResponse struct {
	PortManagementInformationContainer          []byte //Port Management Information Container
	UserPlaneNodeManagementInformationContainer []byte //Bridge Management Information Container
	NWTTPortNumber                              []byte //NW-TT Port Number
}

func (t *TSCManagementInformationIEwithinPFCPSessionModificationResponse) Field() tlv.Field {
	fields := []tlv.Field{}
	if len(t.PortManagementInformationContainer) > 0 {
		fields = append(fields, tlv.TLVFrom(PORT_MANAGEMENT_INFORMATION_CONTAINER_TYPE, tlv.Bytes(t.PortManagementInformationContainer)))
	}
	if len(t.UserPlaneNodeManagementInformationContainer) > 0 {
		fields = append(fields, tlv.TLVFrom(BRIDGE_MANAGEMENT_INFORMATION_CONTAINER_TYPE, tlv.Bytes(t.UserPlaneNodeManagementInformationContainer)))
	}
	if len(t.NWTTPortNumber) > 0 {
		fields = append(fields, tlv.TLVFrom(NW_TT_PORT_NUMBER_TYPE, tlv.Bytes(t.NWTTPortNumber)))
	}
	return tlv.TLV(TSC_MANAGEMENT_INFORMATION_IE_WITHIN_PFCP_SESSION_MODIFICATION_RESPONSE_TYPE, fields...)
}

func (t *TSCManagementInformationIEwithinPFCPSessionModificationResponse) UnmarshalBinary(buf []byte) (err error) {
	*t = TSCManagementInformationIEwithinPFCPSessionModificationResponse{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case PORT_MANAGEMENT_INFORMATION_CONTAINER_TYPE:
			t.PortManagementInformationContainer = de.Value
		case BRIDGE_MANAGEMENT_INFORMATION_CONTAINER_TYPE:
			t.UserPlaneNodeManagementInformationContainer = de.Value
		case NW_TT_PORT_NUMBER_TYPE:
			t.NWTTPortNumber = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
