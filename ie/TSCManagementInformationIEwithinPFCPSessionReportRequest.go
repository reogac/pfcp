package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'TSC Management Information IE within PFCP Session Report Request' data structure */
type TSCManagementInformationIEwithinPFCPSessionReportRequest struct {
	PortManagementInformationContainer          []byte   //Port Management Information Container
	UserPlaneNodeManagementInformationContainer []byte   //Bridge Management Information Container
	NWTTPortNumber                              []byte   //NW-TT Port Number
	Unknown                                     []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid                                     []uint32 `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
}

func (t *TSCManagementInformationIEwithinPFCPSessionReportRequest) Field() tlv.Field {
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
	return tlv.TLV(TSC_MANAGEMENT_INFORMATION_IE_WITHIN_PFCP_SESSION_REPORT_REQUEST_TYPE, fields...)
}

func (t *TSCManagementInformationIEwithinPFCPSessionReportRequest) UnmarshalBinary(buf []byte) (err error) {
	*t = TSCManagementInformationIEwithinPFCPSessionReportRequest{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [3]bool
	for _, de := range elements {
		opt := false
		switch de.Type {
		case PORT_MANAGEMENT_INFORMATION_CONTAINER_TYPE:
			opt = true
			if seen[0] {
				continue
			}
			seen[0] = true
			t.PortManagementInformationContainer = de.Value
		case BRIDGE_MANAGEMENT_INFORMATION_CONTAINER_TYPE:
			opt = true
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
			if err = pfcp.Faulty(de.Type, err); opt && pfcp.Discardable(err) {
				t.Invalid = append(t.Invalid, de.Type)
				err = nil
				continue
			}
			return err
		}
	}
	return
}
