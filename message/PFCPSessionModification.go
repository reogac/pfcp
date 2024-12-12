package message

import (
	"etrib5gc/pfcp"
	"etrib5gc/pfcp/ie"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'PFCP Session Modification Request' data structure */
type PFCPSessionModificationRequest struct {
	CPFSEID                        *ie.FSEID                                                          //F-SEID
	RemovePDR                      []ie.RemovePDR                                                     //Remove PDR
	RemoveFAR                      []ie.RemoveFAR                                                     //Remove FAR
	RemoveURR                      []ie.RemoveURR                                                     //Remove URR
	RemoveQER                      []ie.RemoveQER                                                     //Remove QER
	RemoveBAR                      *ie.RemoveBAR                                                      //Remove BAR
	RemoveTrafficEndpoint          *ie.RemoveTrafficEndpoint                                          //Remove Traffic Endpoint
	CreatePDR                      []ie.CreatePDR                                                     //Create PDR
	CreateFAR                      []ie.CreateFAR                                                     //Create FAR
	CreateURR                      []ie.CreateURR                                                     //Create URR
	CreateQER                      []ie.CreateQER                                                     //Create QER
	CreateBAR                      *ie.CreateBAR                                                      //Create BAR
	CreateTrafficEndpoint          *ie.CreateTrafficEndpoint                                          //Create Traffic Endpoint
	UpdatePDR                      []ie.UpdatePDR                                                     //Update PDR
	UpdateFAR                      []ie.UpdateFAR                                                     //Update FAR
	UpdateURR                      []ie.UpdateURR                                                     //Update URR
	UpdateQER                      []ie.UpdateQER                                                     //Update QER
	UpdateBAR                      *ie.UpdateBARSessionModificationRequest                            //Update BAR Session Modification Request
	UpdateTrafficEndpoint          []byte                                                             //Update Traffic Endpoint
	PFCPSMReqFlags                 *uint8                                                             //PFCPSMReq-Flags
	QueryURR                       *ie.QueryURR                                                       //Query URR
	UserPlaneInactivityTimer       []byte                                                             //User Plane Inactivity Timer
	QueryURRReference              []byte                                                             //Query URR Reference
	TraceInformation               []byte                                                             //Trace Information
	RemoveMAR                      *ie.RemoveMAR                                                      //Remove MAR
	UpdateMAR                      *ie.UpdateMAR                                                      //Update MAR
	CreateMAR                      *ie.CreateMAR                                                      //Create MAR
	NodeID                         *ie.NodeID                                                         //Node ID
	TSCManagementInformation       *ie.TSCManagementInformationIEwithinPFCPSessionModificationRequest //TSC Management Information IE within PFCP Session Modification Request
	RemoveSRR                      *ie.RemoveSRR                                                      //Remove SRR
	CreateSRR                      *ie.CreateSRR                                                      //Create SRR
	UpdateSRR                      *ie.UpdateSRR                                                      //Update SRR
	ProvideATSSSControlInformation *ie.ProvideATSSSControlInformation                                 //Provide ATSSS Control Information
	EthernetContextInformation     *ie.EthernetContextInformation                                     //Ethernet Context Information
	AccessAvailabilityInformation  []byte                                                             //Access Availability Information
	QueryPacketRateStatus          *ie.QueryPacketRateStatusIEwithinPFCPSessionModificationRequest    //Query Packet Rate Status IE within PFCP Session Modification Request
	SNSSAI                         []byte                                                             //S-NSSAI
	RATType                        []byte                                                             //RAT Type
	GroupId                        []byte                                                             //Group ID
	MBSSessionN4ControlInformation *ie.MBSSessionN4ControlInformation                                 //MBS Session N4 Control Information
	DSCPtoPPIControlInformation    *ie.DSCPtoPPIControlInformation                                    //DSCP to PPI Control Information
	FQCSIDList                     []ie.FQCSID                                                        //FQ-CSID
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
	if t.RemoveTrafficEndpoint != nil {
		fields = append(fields, t.RemoveTrafficEndpoint.Field())
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
	if t.CreateTrafficEndpoint != nil {
		fields = append(fields, t.CreateTrafficEndpoint.Field())
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
	if len(t.UpdateTrafficEndpoint) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.UPDATE_TRAFFIC_ENDPOINT_TYPE, tlv.Bytes(t.UpdateTrafficEndpoint)))
	}
	if t.PFCPSMReqFlags != nil {
		fields = append(fields, tlv.TLVFrom(ie.PFCPSMREQ_FLAGS_TYPE, tlv.Bytes(pfcp.MarshalUint8(*t.PFCPSMReqFlags))))
	}
	if t.QueryURR != nil {
		fields = append(fields, t.QueryURR.Field())
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
	if t.RemoveMAR != nil {
		fields = append(fields, t.RemoveMAR.Field())
	}
	if t.UpdateMAR != nil {
		fields = append(fields, t.UpdateMAR.Field())
	}
	if t.CreateMAR != nil {
		fields = append(fields, t.CreateMAR.Field())
	}
	if t.NodeID != nil {
		fields = append(fields, tlv.TLVFrom(ie.NODE_ID_TYPE, tlv.Bytes(t.NodeID.Bytes())))
	}
	if t.TSCManagementInformation != nil {
		fields = append(fields, t.TSCManagementInformation.Field())
	}
	if t.RemoveSRR != nil {
		fields = append(fields, t.RemoveSRR.Field())
	}
	if t.CreateSRR != nil {
		fields = append(fields, t.CreateSRR.Field())
	}
	if t.UpdateSRR != nil {
		fields = append(fields, t.UpdateSRR.Field())
	}
	if t.ProvideATSSSControlInformation != nil {
		fields = append(fields, t.ProvideATSSSControlInformation.Field())
	}
	if t.EthernetContextInformation != nil {
		fields = append(fields, t.EthernetContextInformation.Field())
	}
	if len(t.AccessAvailabilityInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.ACCESS_AVAILABILITY_INFORMATION_TYPE, tlv.Bytes(t.AccessAvailabilityInformation)))
	}
	if t.QueryPacketRateStatus != nil {
		fields = append(fields, t.QueryPacketRateStatus.Field())
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
	if t.MBSSessionN4ControlInformation != nil {
		fields = append(fields, t.MBSSessionN4ControlInformation.Field())
	}
	if t.DSCPtoPPIControlInformation != nil {
		fields = append(fields, t.DSCPtoPPIControlInformation.Field())
	}
	for _, i := range t.FQCSIDList {
		fields = append(fields, tlv.TLVFrom(ie.FQ_CSID_TYPE, tlv.Bytes(i.Bytes())))
	}
	return tlv.TLV(ie.MSG_PFCP_SESSION_MODIFICATION_REQUEST_TYPE, fields...)
}

func (t *PFCPSessionModificationRequest) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPSessionModificationRequest{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case ie.F_SEID_TYPE:
			var v ie.FSEID
			if err = de.UnmarshalValue(&v); err == nil {
				t.CPFSEID = &v
			}
		case ie.REMOVE_PDR_TYPE:
			var v ie.RemovePDR
			if err = de.UnmarshalValue(&v); err == nil {
				t.RemovePDR = append(t.RemovePDR, v)
			}
		case ie.REMOVE_FAR_TYPE:
			var v ie.RemoveFAR
			if err = de.UnmarshalValue(&v); err == nil {
				t.RemoveFAR = append(t.RemoveFAR, v)
			}
		case ie.REMOVE_URR_TYPE:
			var v ie.RemoveURR
			if err = de.UnmarshalValue(&v); err == nil {
				t.RemoveURR = append(t.RemoveURR, v)
			}
		case ie.REMOVE_QER_TYPE:
			var v ie.RemoveQER
			if err = de.UnmarshalValue(&v); err == nil {
				t.RemoveQER = append(t.RemoveQER, v)
			}
		case ie.REMOVE_BAR_TYPE:
			var v ie.RemoveBAR
			if err = de.UnmarshalValue(&v); err == nil {
				t.RemoveBAR = &v
			}
		case ie.REMOVE_TRAFFIC_ENDPOINT_TYPE:
			var v ie.RemoveTrafficEndpoint
			if err = de.UnmarshalValue(&v); err == nil {
				t.RemoveTrafficEndpoint = &v
			}
		case ie.CREATE_PDR_TYPE:
			var v ie.CreatePDR
			if err = de.UnmarshalValue(&v); err == nil {
				t.CreatePDR = append(t.CreatePDR, v)
			}
		case ie.CREATE_FAR_TYPE:
			var v ie.CreateFAR
			if err = de.UnmarshalValue(&v); err == nil {
				t.CreateFAR = append(t.CreateFAR, v)
			}
		case ie.CREATE_URR_TYPE:
			var v ie.CreateURR
			if err = de.UnmarshalValue(&v); err == nil {
				t.CreateURR = append(t.CreateURR, v)
			}
		case ie.CREATE_QER_TYPE:
			var v ie.CreateQER
			if err = de.UnmarshalValue(&v); err == nil {
				t.CreateQER = append(t.CreateQER, v)
			}
		case ie.CREATE_BAR_TYPE:
			var v ie.CreateBAR
			if err = de.UnmarshalValue(&v); err == nil {
				t.CreateBAR = &v
			}
		case ie.CREATE_TRAFFIC_ENDPOINT_TYPE:
			var v ie.CreateTrafficEndpoint
			if err = de.UnmarshalValue(&v); err == nil {
				t.CreateTrafficEndpoint = &v
			}
		case ie.UPDATE_PDR_TYPE:
			var v ie.UpdatePDR
			if err = de.UnmarshalValue(&v); err == nil {
				t.UpdatePDR = append(t.UpdatePDR, v)
			}
		case ie.UPDATE_FAR_TYPE:
			var v ie.UpdateFAR
			if err = de.UnmarshalValue(&v); err == nil {
				t.UpdateFAR = append(t.UpdateFAR, v)
			}
		case ie.UPDATE_URR_TYPE:
			var v ie.UpdateURR
			if err = de.UnmarshalValue(&v); err == nil {
				t.UpdateURR = append(t.UpdateURR, v)
			}
		case ie.UPDATE_QER_TYPE:
			var v ie.UpdateQER
			if err = de.UnmarshalValue(&v); err == nil {
				t.UpdateQER = append(t.UpdateQER, v)
			}
		case ie.UPDATE_BAR_SESSION_MODIFICATION_REQUEST_TYPE:
			var v ie.UpdateBARSessionModificationRequest
			if err = de.UnmarshalValue(&v); err == nil {
				t.UpdateBAR = &v
			}
		case ie.UPDATE_TRAFFIC_ENDPOINT_TYPE:
			t.UpdateTrafficEndpoint = de.Value
		case ie.PFCPSMREQ_FLAGS_TYPE:
			var v uint8
			if err = pfcp.UnmarshalUint8(de.Value, &v); err == nil {
				t.PFCPSMReqFlags = &v
			}
		case ie.QUERY_URR_TYPE:
			var v ie.QueryURR
			if err = de.UnmarshalValue(&v); err == nil {
				t.QueryURR = &v
			}
		case ie.USER_PLANE_INACTIVITY_TIMER_TYPE:
			t.UserPlaneInactivityTimer = de.Value
		case ie.QUERY_URR_REFERENCE_TYPE:
			t.QueryURRReference = de.Value
		case ie.TRACE_INFORMATION_TYPE:
			t.TraceInformation = de.Value
		case ie.REMOVE_MAR_TYPE:
			var v ie.RemoveMAR
			if err = de.UnmarshalValue(&v); err == nil {
				t.RemoveMAR = &v
			}
		case ie.UPDATE_MAR_TYPE:
			var v ie.UpdateMAR
			if err = de.UnmarshalValue(&v); err == nil {
				t.UpdateMAR = &v
			}
		case ie.CREATE_MAR_TYPE:
			var v ie.CreateMAR
			if err = de.UnmarshalValue(&v); err == nil {
				t.CreateMAR = &v
			}
		case ie.NODE_ID_TYPE:
			var v ie.NodeID
			if err = de.UnmarshalValue(&v); err == nil {
				t.NodeID = &v
			}
		case ie.TSC_MANAGEMENT_INFORMATION_IE_WITHIN_PFCP_SESSION_MODIFICATION_REQUEST_TYPE:
			var v ie.TSCManagementInformationIEwithinPFCPSessionModificationRequest
			if err = de.UnmarshalValue(&v); err == nil {
				t.TSCManagementInformation = &v
			}
		case ie.REMOVE_SRR_TYPE:
			var v ie.RemoveSRR
			if err = de.UnmarshalValue(&v); err == nil {
				t.RemoveSRR = &v
			}
		case ie.CREATE_SRR_TYPE:
			var v ie.CreateSRR
			if err = de.UnmarshalValue(&v); err == nil {
				t.CreateSRR = &v
			}
		case ie.UPDATE_SRR_TYPE:
			var v ie.UpdateSRR
			if err = de.UnmarshalValue(&v); err == nil {
				t.UpdateSRR = &v
			}
		case ie.PROVIDE_ATSSS_CONTROL_INFORMATION_TYPE:
			var v ie.ProvideATSSSControlInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.ProvideATSSSControlInformation = &v
			}
		case ie.ETHERNET_CONTEXT_INFORMATION_TYPE:
			var v ie.EthernetContextInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.EthernetContextInformation = &v
			}
		case ie.ACCESS_AVAILABILITY_INFORMATION_TYPE:
			t.AccessAvailabilityInformation = de.Value
		case ie.QUERY_PACKET_RATE_STATUS_IE_WITHIN_PFCP_SESSION_MODIFICATION_REQUEST_TYPE:
			var v ie.QueryPacketRateStatusIEwithinPFCPSessionModificationRequest
			if err = de.UnmarshalValue(&v); err == nil {
				t.QueryPacketRateStatus = &v
			}
		case ie.S_NSSAI_TYPE:
			t.SNSSAI = de.Value
		case ie.RAT_TYPE_TYPE:
			t.RATType = de.Value
		case ie.GROUP_ID_TYPE:
			t.GroupId = de.Value
		case ie.MBS_SESSION_N4_CONTROL_INFORMATION_TYPE:
			var v ie.MBSSessionN4ControlInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.MBSSessionN4ControlInformation = &v
			}
		case ie.DSCP_TO_PPI_CONTROL_INFORMATION_TYPE:
			var v ie.DSCPtoPPIControlInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.DSCPtoPPIControlInformation = &v
			}
		case ie.FQ_CSID_TYPE:
			var v ie.FQCSID
			if err = de.UnmarshalValue(&v); err == nil {
				t.FQCSIDList = append(t.FQCSIDList, v)
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
	Cause                             uint8                                                              //Cause
	OffendingIE                       *uint16                                                            //Offending IE
	CreatedPDR                        []ie.CreatedPDR                                                    //Created PDR
	LoadControlInformation            *ie.LoadControlInformation                                         //Load Control Information
	OverloadControlInformation        *ie.OverloadControlInformation                                     //Overload Control Information
	UsageReport                       []ie.UsageReportSessionModificationResponse                        //Usage Report Session Modification Response
	FailedRuleID                      []byte                                                             //Failed Rule ID
	AdditionalUsageReportsInformation []byte                                                             //Additional Usage Reports Information
	CreatedUpdatedTrafficEndpoint     *ie.CreatedTrafficEndpoint                                         //Created Traffic Endpoint
	TSCManagementInformation          *ie.TSCManagementInformationIEwithinPFCPSessionModificationRequest //TSC Management Information IE within PFCP Session Modification Request
	ATSSSControlParameters            *ie.ATSSSControlParameters                                         //ATSSS Control Parameters
	UpdatedPDR                        *ie.UpdatedPDR                                                     //Updated PDR
	PacketRateStatusReport            *ie.PacketRateStatusReport                                         //Packet Rate Status Report
	PartialFailureInformation         *ie.PartialFailureInformation                                      //Partial Failure Information
	MBSSessionN4Information           *ie.MBSSessionN4Information                                        //MBS Session N4 Information
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
	if t.CreatedUpdatedTrafficEndpoint != nil {
		fields = append(fields, t.CreatedUpdatedTrafficEndpoint.Field())
	}
	if t.TSCManagementInformation != nil {
		fields = append(fields, t.TSCManagementInformation.Field())
	}
	if t.ATSSSControlParameters != nil {
		fields = append(fields, t.ATSSSControlParameters.Field())
	}
	if t.UpdatedPDR != nil {
		fields = append(fields, t.UpdatedPDR.Field())
	}
	if t.PacketRateStatusReport != nil {
		fields = append(fields, t.PacketRateStatusReport.Field())
	}
	if t.PartialFailureInformation != nil {
		fields = append(fields, t.PartialFailureInformation.Field())
	}
	if t.MBSSessionN4Information != nil {
		fields = append(fields, t.MBSSessionN4Information.Field())
	}
	return tlv.TLV(ie.MSG_PFCP_SESSION_MODIFICATION_RESPONSE_TYPE, fields...)
}

func (t *PFCPSessionModificationResponse) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPSessionModificationResponse{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case ie.CAUSE_TYPE:
			err = pfcp.UnmarshalUint8(de.Value, &t.Cause)
		case ie.OFFENDING_IE_TYPE:
			var v uint16
			if err = pfcp.UnmarshalUint16(de.Value, &v); err == nil {
				t.OffendingIE = &v
			}
		case ie.CREATED_PDR_TYPE:
			var v ie.CreatedPDR
			if err = de.UnmarshalValue(&v); err == nil {
				t.CreatedPDR = append(t.CreatedPDR, v)
			}
		case ie.LOAD_CONTROL_INFORMATION_TYPE:
			var v ie.LoadControlInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.LoadControlInformation = &v
			}
		case ie.OVERLOAD_CONTROL_INFORMATION_TYPE:
			var v ie.OverloadControlInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.OverloadControlInformation = &v
			}
		case ie.USAGE_REPORT_SESSION_MODIFICATION_RESPONSE_TYPE:
			var v ie.UsageReportSessionModificationResponse
			if err = de.UnmarshalValue(&v); err == nil {
				t.UsageReport = append(t.UsageReport, v)
			}
		case ie.FAILED_RULE_ID_TYPE:
			t.FailedRuleID = de.Value
		case ie.ADDITIONAL_USAGE_REPORTS_INFORMATION_TYPE:
			t.AdditionalUsageReportsInformation = de.Value
		case ie.CREATED_TRAFFIC_ENDPOINT_TYPE:
			var v ie.CreatedTrafficEndpoint
			if err = de.UnmarshalValue(&v); err == nil {
				t.CreatedUpdatedTrafficEndpoint = &v
			}
		case ie.TSC_MANAGEMENT_INFORMATION_IE_WITHIN_PFCP_SESSION_MODIFICATION_REQUEST_TYPE:
			var v ie.TSCManagementInformationIEwithinPFCPSessionModificationRequest
			if err = de.UnmarshalValue(&v); err == nil {
				t.TSCManagementInformation = &v
			}
		case ie.ATSSS_CONTROL_PARAMETERS_TYPE:
			var v ie.ATSSSControlParameters
			if err = de.UnmarshalValue(&v); err == nil {
				t.ATSSSControlParameters = &v
			}
		case ie.UPDATED_PDR_TYPE:
			var v ie.UpdatedPDR
			if err = de.UnmarshalValue(&v); err == nil {
				t.UpdatedPDR = &v
			}
		case ie.PACKET_RATE_STATUS_REPORT_TYPE:
			var v ie.PacketRateStatusReport
			if err = de.UnmarshalValue(&v); err == nil {
				t.PacketRateStatusReport = &v
			}
		case ie.PARTIAL_FAILURE_INFORMATION_TYPE:
			var v ie.PartialFailureInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.PartialFailureInformation = &v
			}
		case ie.MBS_SESSION_N4_INFORMATION_TYPE:
			var v ie.MBSSessionN4Information
			if err = de.UnmarshalValue(&v); err == nil {
				t.MBSSessionN4Information = &v
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
