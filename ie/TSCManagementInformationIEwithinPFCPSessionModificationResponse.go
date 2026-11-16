package ie

import (
	"github.com/reogac/pfcp/tlv"
)

/* 'TSC Management Information IE within PFCP Session Modification Response' data structure */
type TSCManagementInformationIEwithinPFCPSessionModificationResponse struct {
	PortManagementInformationContainer          []byte   //Port Management Information Container
	UserPlaneNodeManagementInformationContainer []byte   //Bridge Management Information Container
	NWTTPortNumber                              []byte   //NW-TT Port Number
	Unknown                                     []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
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
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	var seen [3]bool
	for _, de := range elements {
		switch de.Type {
		case PORT_MANAGEMENT_INFORMATION_CONTAINER_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.PortManagementInformationContainer = de.Value
		case BRIDGE_MANAGEMENT_INFORMATION_CONTAINER_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			t.UserPlaneNodeManagementInformationContainer = de.Value
		case NW_TT_PORT_NUMBER_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			t.NWTTPortNumber = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	return
}
