package message

import (
	"github.com/reogac/pfcp/v2"
	"github.com/reogac/pfcp/v2/ie"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'PFCP Session Modification Request' data structure */
type PFCPSessionModificationRequest struct {
	CPFSEID                        *ie.FSEID                                                           //F-SEID
	RemovePDR                      []ie.RemovePDR                                                      //Remove PDR
	RemoveFAR                      []ie.RemoveFAR                                                      //Remove FAR
	RemoveURR                      []ie.RemoveURR                                                      //Remove URR
	RemoveQER                      []ie.RemoveQER                                                      //Remove QER
	RemoveBAR                      *ie.RemoveBAR                                                       //Remove BAR
	RemoveTrafficEndpoint          []ie.RemoveTrafficEndpoint                                          //Remove Traffic Endpoint
	CreatePDR                      []ie.CreatePDR                                                      //Create PDR
	CreateFAR                      []ie.CreateFAR                                                      //Create FAR
	CreateURR                      []ie.CreateURR                                                      //Create URR
	CreateQER                      []ie.CreateQER                                                      //Create QER
	CreateBAR                      *ie.CreateBAR                                                       //Create BAR
	CreateTrafficEndpoint          []ie.CreateTrafficEndpoint                                          //Create Traffic Endpoint
	UpdatePDR                      []ie.UpdatePDR                                                      //Update PDR
	UpdateFAR                      []ie.UpdateFAR                                                      //Update FAR
	UpdateURR                      []ie.UpdateURR                                                      //Update URR
	UpdateQER                      []ie.UpdateQER                                                      //Update QER
	UpdateBAR                      *ie.UpdateBARSessionModificationRequest                             //Update BAR Session Modification Request
	UpdateTrafficEndpoint          [][]byte                                                            //Update Traffic Endpoint
	PFCPSMReqFlags                 *uint8                                                              //PFCPSMReq-Flags
	QueryURR                       []ie.QueryURR                                                       //Query URR
	UserPlaneInactivityTimer       []byte                                                              //User Plane Inactivity Timer
	QueryURRReference              []byte                                                              //Query URR Reference
	TraceInformation               []byte                                                              //Trace Information
	RemoveMAR                      []ie.RemoveMAR                                                      //Remove MAR
	UpdateMAR                      []ie.UpdateMAR                                                      //Update MAR
	CreateMAR                      []ie.CreateMAR                                                      //Create MAR
	NodeID                         *ie.NodeID                                                          //Node ID
	TSCManagementInformation       []ie.TSCManagementInformationIEwithinPFCPSessionModificationRequest //TSC Management Information IE within PFCP Session Modification Request
	RemoveSRR                      []ie.RemoveSRR                                                      //Remove SRR
	CreateSRR                      []ie.CreateSRR                                                      //Create SRR
	UpdateSRR                      []ie.UpdateSRR                                                      //Update SRR
	ProvideATSSSControlInformation *ie.ProvideATSSSControlInformation                                  //Provide ATSSS Control Information
	EthernetContextInformation     *ie.EthernetContextInformation                                      //Ethernet Context Information
	AccessAvailabilityInformation  [][]byte                                                            //Access Availability Information
	QueryPacketRateStatus          []ie.QueryPacketRateStatusIEwithinPFCPSessionModificationRequest    //Query Packet Rate Status IE within PFCP Session Modification Request
	SNSSAI                         []byte                                                              //S-NSSAI
	RATType                        []byte                                                              //RAT Type
	GroupId                        []byte                                                              //Group ID
	MBSSessionN4ControlInformation []ie.MBSSessionN4ControlInformation                                 //MBS Session N4 Control Information
	DSCPtoPPIControlInformation    *ie.DSCPtoPPIControlInformation                                     //DSCP to PPI Control Information
	FQCSIDList                     []ie.FQCSID                                                         //FQ-CSID
	Unknown                        []uint32                                                            `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid                        []uint32                                                            `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
}

func (t *PFCPSessionModificationRequest) Field() tlv.Field {
	fields := []tlv.Field{}
	if t.CPFSEID != nil {
		fields = append(fields, tlv.TLVFrom(ie.F_SEID_TYPE, tlv.Bytes(t.CPFSEID.Bytes())))
	}
	for _, i := range t.RemovePDR {
		fields = append(fields, i.Field())
	}
	for _, i := range t.RemoveFAR {
		fields = append(fields, i.Field())
	}
	for _, i := range t.RemoveURR {
		fields = append(fields, i.Field())
	}
	for _, i := range t.RemoveQER {
		fields = append(fields, i.Field())
	}
	if t.RemoveBAR != nil {
		fields = append(fields, t.RemoveBAR.Field())
	}
	for _, i := range t.RemoveTrafficEndpoint {
		fields = append(fields, i.Field())
	}
	for _, i := range t.CreatePDR {
		fields = append(fields, i.Field())
	}
	for _, i := range t.CreateFAR {
		fields = append(fields, i.Field())
	}
	for _, i := range t.CreateURR {
		fields = append(fields, i.Field())
	}
	for _, i := range t.CreateQER {
		fields = append(fields, i.Field())
	}
	if t.CreateBAR != nil {
		fields = append(fields, t.CreateBAR.Field())
	}
	for _, i := range t.CreateTrafficEndpoint {
		fields = append(fields, i.Field())
	}
	for _, i := range t.UpdatePDR {
		fields = append(fields, i.Field())
	}
	for _, i := range t.UpdateFAR {
		fields = append(fields, i.Field())
	}
	for _, i := range t.UpdateURR {
		fields = append(fields, i.Field())
	}
	for _, i := range t.UpdateQER {
		fields = append(fields, i.Field())
	}
	if t.UpdateBAR != nil {
		fields = append(fields, t.UpdateBAR.Field())
	}
	for _, i := range t.UpdateTrafficEndpoint {
		fields = append(fields, tlv.TLVFrom(ie.UPDATE_TRAFFIC_ENDPOINT_TYPE, tlv.Bytes(i)))
	}
	if t.PFCPSMReqFlags != nil {
		fields = append(fields, tlv.TLVFrom(ie.PFCPSMREQ_FLAGS_TYPE, tlv.Bytes(pfcp.MarshalUint8(*t.PFCPSMReqFlags))))
	}
	for _, i := range t.QueryURR {
		fields = append(fields, i.Field())
	}
	if len(t.UserPlaneInactivityTimer) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.USER_PLANE_INACTIVITY_TIMER_TYPE, tlv.Bytes(t.UserPlaneInactivityTimer)))
	}
	if len(t.QueryURRReference) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.QUERY_URR_REFERENCE_TYPE, tlv.Bytes(t.QueryURRReference)))
	}
	if len(t.TraceInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.TRACE_INFORMATION_TYPE, tlv.Bytes(t.TraceInformation)))
	}
	for _, i := range t.RemoveMAR {
		fields = append(fields, i.Field())
	}
	for _, i := range t.UpdateMAR {
		fields = append(fields, i.Field())
	}
	for _, i := range t.CreateMAR {
		fields = append(fields, i.Field())
	}
	if t.NodeID != nil {
		fields = append(fields, tlv.TLVFrom(ie.NODE_ID_TYPE, tlv.Bytes(t.NodeID.Bytes())))
	}
	for _, i := range t.TSCManagementInformation {
		fields = append(fields, i.Field())
	}
	for _, i := range t.RemoveSRR {
		fields = append(fields, i.Field())
	}
	for _, i := range t.CreateSRR {
		fields = append(fields, i.Field())
	}
	for _, i := range t.UpdateSRR {
		fields = append(fields, i.Field())
	}
	if t.ProvideATSSSControlInformation != nil {
		fields = append(fields, t.ProvideATSSSControlInformation.Field())
	}
	if t.EthernetContextInformation != nil {
		fields = append(fields, t.EthernetContextInformation.Field())
	}
	for _, i := range t.AccessAvailabilityInformation {
		fields = append(fields, tlv.TLVFrom(ie.ACCESS_AVAILABILITY_INFORMATION_TYPE, tlv.Bytes(i)))
	}
	for _, i := range t.QueryPacketRateStatus {
		fields = append(fields, i.Field())
	}
	if len(t.SNSSAI) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.S_NSSAI_TYPE, tlv.Bytes(t.SNSSAI)))
	}
	if len(t.RATType) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.RAT_TYPE_TYPE, tlv.Bytes(t.RATType)))
	}
	if len(t.GroupId) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.GROUP_ID_TYPE, tlv.Bytes(t.GroupId)))
	}
	for _, i := range t.MBSSessionN4ControlInformation {
		fields = append(fields, i.Field())
	}
	if t.DSCPtoPPIControlInformation != nil {
		fields = append(fields, t.DSCPtoPPIControlInformation.Field())
	}
	for _, i := range t.FQCSIDList {
		fields = append(fields, tlv.TLVFrom(ie.FQ_CSID_TYPE, tlv.Bytes(i.Bytes())))
	}
	return tlv.TLV(ie.MSG_PFCP_SESSION_MODIFICATION_REQUEST_TYPE, fields...)
}

// Skipped is what decoding left out: IEs unknown or not defined here
// (TS 29.244 7.6.9), and optional ones discarded for a faulty value (7.6.8).
func (t *PFCPSessionModificationRequest) Skipped() (unknown, invalid []uint32) {
	return t.Unknown, t.Invalid
}

func (t *PFCPSessionModificationRequest) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPSessionModificationRequest{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [15]bool
	for _, de := range elements {
		opt := false
		switch de.Type {
		case ie.F_SEID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			var v ie.FSEID
			if err = de.UnmarshalValue(&v); err == nil {
				t.CPFSEID = &v
			}
		case ie.REMOVE_PDR_TYPE:
			var v ie.RemovePDR
			if err = de.UnmarshalValue(&v); err == nil {
				t.RemovePDR = append(t.RemovePDR, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.REMOVE_FAR_TYPE:
			var v ie.RemoveFAR
			if err = de.UnmarshalValue(&v); err == nil {
				t.RemoveFAR = append(t.RemoveFAR, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.REMOVE_URR_TYPE:
			var v ie.RemoveURR
			if err = de.UnmarshalValue(&v); err == nil {
				t.RemoveURR = append(t.RemoveURR, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.REMOVE_QER_TYPE:
			var v ie.RemoveQER
			if err = de.UnmarshalValue(&v); err == nil {
				t.RemoveQER = append(t.RemoveQER, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.REMOVE_BAR_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			var v ie.RemoveBAR
			if err = de.UnmarshalValue(&v); err == nil {
				t.RemoveBAR = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.REMOVE_TRAFFIC_ENDPOINT_TYPE:
			var v ie.RemoveTrafficEndpoint
			if err = de.UnmarshalValue(&v); err == nil {
				t.RemoveTrafficEndpoint = append(t.RemoveTrafficEndpoint, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.CREATE_PDR_TYPE:
			var v ie.CreatePDR
			if err = de.UnmarshalValue(&v); err == nil {
				t.CreatePDR = append(t.CreatePDR, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.CREATE_FAR_TYPE:
			var v ie.CreateFAR
			if err = de.UnmarshalValue(&v); err == nil {
				t.CreateFAR = append(t.CreateFAR, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.CREATE_URR_TYPE:
			var v ie.CreateURR
			if err = de.UnmarshalValue(&v); err == nil {
				t.CreateURR = append(t.CreateURR, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.CREATE_QER_TYPE:
			var v ie.CreateQER
			if err = de.UnmarshalValue(&v); err == nil {
				t.CreateQER = append(t.CreateQER, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.CREATE_BAR_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			var v ie.CreateBAR
			if err = de.UnmarshalValue(&v); err == nil {
				t.CreateBAR = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.CREATE_TRAFFIC_ENDPOINT_TYPE:
			var v ie.CreateTrafficEndpoint
			if err = de.UnmarshalValue(&v); err == nil {
				t.CreateTrafficEndpoint = append(t.CreateTrafficEndpoint, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.UPDATE_PDR_TYPE:
			var v ie.UpdatePDR
			if err = de.UnmarshalValue(&v); err == nil {
				t.UpdatePDR = append(t.UpdatePDR, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.UPDATE_FAR_TYPE:
			var v ie.UpdateFAR
			if err = de.UnmarshalValue(&v); err == nil {
				t.UpdateFAR = append(t.UpdateFAR, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.UPDATE_URR_TYPE:
			var v ie.UpdateURR
			if err = de.UnmarshalValue(&v); err == nil {
				t.UpdateURR = append(t.UpdateURR, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.UPDATE_QER_TYPE:
			var v ie.UpdateQER
			if err = de.UnmarshalValue(&v); err == nil {
				t.UpdateQER = append(t.UpdateQER, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.UPDATE_BAR_SESSION_MODIFICATION_REQUEST_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			var v ie.UpdateBARSessionModificationRequest
			if err = de.UnmarshalValue(&v); err == nil {
				t.UpdateBAR = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.UPDATE_TRAFFIC_ENDPOINT_TYPE:
			t.UpdateTrafficEndpoint = append(t.UpdateTrafficEndpoint, de.Value)
		case ie.PFCPSMREQ_FLAGS_TYPE:
			if seen[4] {
				continue
			}
			seen[4] = true
			var v uint8
			if err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &v); err == nil {
				t.PFCPSMReqFlags = &v
			}
		case ie.QUERY_URR_TYPE:
			var v ie.QueryURR
			if err = de.UnmarshalValue(&v); err == nil {
				t.QueryURR = append(t.QueryURR, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.USER_PLANE_INACTIVITY_TIMER_TYPE:
			if seen[5] {
				continue
			}
			seen[5] = true
			t.UserPlaneInactivityTimer = de.Value
		case ie.QUERY_URR_REFERENCE_TYPE:
			opt = true
			if seen[6] {
				continue
			}
			seen[6] = true
			t.QueryURRReference = de.Value
		case ie.TRACE_INFORMATION_TYPE:
			opt = true
			if seen[7] {
				continue
			}
			seen[7] = true
			t.TraceInformation = de.Value
		case ie.REMOVE_MAR_TYPE:
			var v ie.RemoveMAR
			if err = de.UnmarshalValue(&v); err == nil {
				t.RemoveMAR = append(t.RemoveMAR, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.UPDATE_MAR_TYPE:
			var v ie.UpdateMAR
			if err = de.UnmarshalValue(&v); err == nil {
				t.UpdateMAR = append(t.UpdateMAR, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.CREATE_MAR_TYPE:
			var v ie.CreateMAR
			if err = de.UnmarshalValue(&v); err == nil {
				t.CreateMAR = append(t.CreateMAR, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.NODE_ID_TYPE:
			if seen[8] {
				continue
			}
			seen[8] = true
			var v ie.NodeID
			if err = de.UnmarshalValue(&v); err == nil {
				t.NodeID = &v
			}
		case ie.TSC_MANAGEMENT_INFORMATION_IE_WITHIN_PFCP_SESSION_MODIFICATION_REQUEST_TYPE:
			var v ie.TSCManagementInformationIEwithinPFCPSessionModificationRequest
			if err = de.UnmarshalValue(&v); err == nil {
				t.TSCManagementInformation = append(t.TSCManagementInformation, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.REMOVE_SRR_TYPE:
			var v ie.RemoveSRR
			if err = de.UnmarshalValue(&v); err == nil {
				t.RemoveSRR = append(t.RemoveSRR, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.CREATE_SRR_TYPE:
			var v ie.CreateSRR
			if err = de.UnmarshalValue(&v); err == nil {
				t.CreateSRR = append(t.CreateSRR, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.UPDATE_SRR_TYPE:
			var v ie.UpdateSRR
			if err = de.UnmarshalValue(&v); err == nil {
				t.UpdateSRR = append(t.UpdateSRR, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.PROVIDE_ATSSS_CONTROL_INFORMATION_TYPE:
			if seen[9] {
				continue
			}
			seen[9] = true
			var v ie.ProvideATSSSControlInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.ProvideATSSSControlInformation = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.ETHERNET_CONTEXT_INFORMATION_TYPE:
			if seen[10] {
				continue
			}
			seen[10] = true
			var v ie.EthernetContextInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.EthernetContextInformation = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.ACCESS_AVAILABILITY_INFORMATION_TYPE:
			opt = true
			t.AccessAvailabilityInformation = append(t.AccessAvailabilityInformation, de.Value)
		case ie.QUERY_PACKET_RATE_STATUS_IE_WITHIN_PFCP_SESSION_MODIFICATION_REQUEST_TYPE:
			var v ie.QueryPacketRateStatusIEwithinPFCPSessionModificationRequest
			if err = de.UnmarshalValue(&v); err == nil {
				t.QueryPacketRateStatus = append(t.QueryPacketRateStatus, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.S_NSSAI_TYPE:
			opt = true
			if seen[11] {
				continue
			}
			seen[11] = true
			t.SNSSAI = de.Value
		case ie.RAT_TYPE_TYPE:
			if seen[12] {
				continue
			}
			seen[12] = true
			t.RATType = de.Value
		case ie.GROUP_ID_TYPE:
			if seen[13] {
				continue
			}
			seen[13] = true
			t.GroupId = de.Value
		case ie.MBS_SESSION_N4_CONTROL_INFORMATION_TYPE:
			var v ie.MBSSessionN4ControlInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.MBSSessionN4ControlInformation = append(t.MBSSessionN4ControlInformation, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.DSCP_TO_PPI_CONTROL_INFORMATION_TYPE:
			if seen[14] {
				continue
			}
			seen[14] = true
			var v ie.DSCPtoPPIControlInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.DSCPtoPPIControlInformation = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.FQ_CSID_TYPE:
			var v ie.FQCSID
			if err = de.UnmarshalValue(&v); err == nil {
				t.FQCSIDList = append(t.FQCSIDList, v)
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
	return
}

func (t *PFCPSessionModificationRequest) UnmarshalTLV(typ uint32, buf []byte) (err error) {
	if typ != ie.MSG_PFCP_SESSION_MODIFICATION_REQUEST_TYPE {
		err = tlv.ErrCritical
		return
	}

	err = t.UnmarshalBinary(buf)
	return
}

func (t *PFCPSessionModificationRequest) MarshalJSON() (buf []byte, err error) {
	buf, err = jsonMarshal(t)
	return
}

func (t *PFCPSessionModificationRequest) UnmarshalJSON(buf []byte) (err error) {
	*t = PFCPSessionModificationRequest{}
	err = jsonUnmarshal(buf, t)
	return
}

/* 'PFCP Session Modification Response' data structure */
type PFCPSessionModificationResponse struct {
	Cause                             uint8                                                               //Cause
	OffendingIE                       *uint16                                                             //Offending IE
	CreatedPDR                        []ie.CreatedPDR                                                     //Created PDR
	LoadControlInformation            *ie.LoadControlInformation                                          //Load Control Information
	OverloadControlInformation        *ie.OverloadControlInformation                                      //Overload Control Information
	UsageReport                       []ie.UsageReportSessionModificationResponse                         //Usage Report Session Modification Response
	FailedRuleID                      []byte                                                              //Failed Rule ID
	AdditionalUsageReportsInformation []byte                                                              //Additional Usage Reports Information
	CreatedUpdatedTrafficEndpoint     []ie.CreatedTrafficEndpoint                                         //Created Traffic Endpoint
	TSCManagementInformation          []ie.TSCManagementInformationIEwithinPFCPSessionModificationRequest //TSC Management Information IE within PFCP Session Modification Request
	ATSSSControlParameters            *ie.ATSSSControlParameters                                          //ATSSS Control Parameters
	UpdatedPDR                        []ie.UpdatedPDR                                                     //Updated PDR
	PacketRateStatusReport            []ie.PacketRateStatusReport                                         //Packet Rate Status Report
	PartialFailureInformation         []ie.PartialFailureInformation                                      //Partial Failure Information
	MBSSessionN4Information           []ie.MBSSessionN4Information                                        //MBS Session N4 Information
	Unknown                           []uint32                                                            `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid                           []uint32                                                            `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
}

func (t *PFCPSessionModificationResponse) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(ie.CAUSE_TYPE, tlv.Bytes(pfcp.MarshalUint8(t.Cause))))
	if t.OffendingIE != nil {
		fields = append(fields, tlv.TLVFrom(ie.OFFENDING_IE_TYPE, tlv.Bytes(pfcp.MarshalUint16(*t.OffendingIE))))
	}
	for _, i := range t.CreatedPDR {
		fields = append(fields, i.Field())
	}
	if t.LoadControlInformation != nil {
		fields = append(fields, t.LoadControlInformation.Field())
	}
	if t.OverloadControlInformation != nil {
		fields = append(fields, t.OverloadControlInformation.Field())
	}
	for _, i := range t.UsageReport {
		fields = append(fields, i.Field())
	}
	if len(t.FailedRuleID) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.FAILED_RULE_ID_TYPE, tlv.Bytes(t.FailedRuleID)))
	}
	if len(t.AdditionalUsageReportsInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.ADDITIONAL_USAGE_REPORTS_INFORMATION_TYPE, tlv.Bytes(t.AdditionalUsageReportsInformation)))
	}
	for _, i := range t.CreatedUpdatedTrafficEndpoint {
		fields = append(fields, i.Field())
	}
	for _, i := range t.TSCManagementInformation {
		fields = append(fields, i.Field())
	}
	if t.ATSSSControlParameters != nil {
		fields = append(fields, t.ATSSSControlParameters.Field())
	}
	for _, i := range t.UpdatedPDR {
		fields = append(fields, i.Field())
	}
	for _, i := range t.PacketRateStatusReport {
		fields = append(fields, i.Field())
	}
	for _, i := range t.PartialFailureInformation {
		fields = append(fields, i.Field())
	}
	for _, i := range t.MBSSessionN4Information {
		fields = append(fields, i.Field())
	}
	return tlv.TLV(ie.MSG_PFCP_SESSION_MODIFICATION_RESPONSE_TYPE, fields...)
}

// Skipped is what decoding left out: IEs unknown or not defined here
// (TS 29.244 7.6.9), and optional ones discarded for a faulty value (7.6.8).
func (t *PFCPSessionModificationResponse) Skipped() (unknown, invalid []uint32) {
	return t.Unknown, t.Invalid
}

func (t *PFCPSessionModificationResponse) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPSessionModificationResponse{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [7]bool
	for _, de := range elements {
		opt := false
		switch de.Type {
		case ie.CAUSE_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			err = pfcp.UnmarshalUint8(de.Value, &t.Cause)
		case ie.OFFENDING_IE_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			var v uint16
			if err = pfcp.UnmarshalUint16(de.Value, &v); err == nil {
				t.OffendingIE = &v
			}
		case ie.CREATED_PDR_TYPE:
			var v ie.CreatedPDR
			if err = de.UnmarshalValue(&v); err == nil {
				t.CreatedPDR = append(t.CreatedPDR, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.LOAD_CONTROL_INFORMATION_TYPE:
			opt = true
			if seen[2] {
				continue
			}
			seen[2] = true
			var v ie.LoadControlInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.LoadControlInformation = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.OVERLOAD_CONTROL_INFORMATION_TYPE:
			opt = true
			if seen[3] {
				continue
			}
			seen[3] = true
			var v ie.OverloadControlInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.OverloadControlInformation = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.USAGE_REPORT_SESSION_MODIFICATION_RESPONSE_TYPE:
			var v ie.UsageReportSessionModificationResponse
			if err = de.UnmarshalValue(&v); err == nil {
				t.UsageReport = append(t.UsageReport, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.FAILED_RULE_ID_TYPE:
			if seen[4] {
				continue
			}
			seen[4] = true
			t.FailedRuleID = de.Value
		case ie.ADDITIONAL_USAGE_REPORTS_INFORMATION_TYPE:
			if seen[5] {
				continue
			}
			seen[5] = true
			t.AdditionalUsageReportsInformation = de.Value
		case ie.CREATED_TRAFFIC_ENDPOINT_TYPE:
			var v ie.CreatedTrafficEndpoint
			if err = de.UnmarshalValue(&v); err == nil {
				t.CreatedUpdatedTrafficEndpoint = append(t.CreatedUpdatedTrafficEndpoint, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.TSC_MANAGEMENT_INFORMATION_IE_WITHIN_PFCP_SESSION_MODIFICATION_REQUEST_TYPE:
			var v ie.TSCManagementInformationIEwithinPFCPSessionModificationRequest
			if err = de.UnmarshalValue(&v); err == nil {
				t.TSCManagementInformation = append(t.TSCManagementInformation, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.ATSSS_CONTROL_PARAMETERS_TYPE:
			if seen[6] {
				continue
			}
			seen[6] = true
			var v ie.ATSSSControlParameters
			if err = de.UnmarshalValue(&v); err == nil {
				t.ATSSSControlParameters = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.UPDATED_PDR_TYPE:
			var v ie.UpdatedPDR
			if err = de.UnmarshalValue(&v); err == nil {
				t.UpdatedPDR = append(t.UpdatedPDR, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.PACKET_RATE_STATUS_REPORT_TYPE:
			var v ie.PacketRateStatusReport
			if err = de.UnmarshalValue(&v); err == nil {
				t.PacketRateStatusReport = append(t.PacketRateStatusReport, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.PARTIAL_FAILURE_INFORMATION_TYPE:
			var v ie.PartialFailureInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.PartialFailureInformation = append(t.PartialFailureInformation, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.MBS_SESSION_N4_INFORMATION_TYPE:
			var v ie.MBSSessionN4Information
			if err = de.UnmarshalValue(&v); err == nil {
				t.MBSSessionN4Information = append(t.MBSSessionN4Information, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
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
	if !seen[0] {
		return &pfcp.MissingIE{Type: ie.CAUSE_TYPE}
	}
	return
}

func (t *PFCPSessionModificationResponse) UnmarshalTLV(typ uint32, buf []byte) (err error) {
	if typ != ie.MSG_PFCP_SESSION_MODIFICATION_RESPONSE_TYPE {
		err = tlv.ErrCritical
		return
	}

	err = t.UnmarshalBinary(buf)
	return
}

func (t *PFCPSessionModificationResponse) MarshalJSON() (buf []byte, err error) {
	buf, err = jsonMarshal(t)
	return
}

func (t *PFCPSessionModificationResponse) UnmarshalJSON(buf []byte) (err error) {
	*t = PFCPSessionModificationResponse{}
	err = jsonUnmarshal(buf, t)
	return
}
