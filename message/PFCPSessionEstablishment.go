package message

import (
	"github.com/reogac/pfcp"
	"github.com/reogac/pfcp/ie"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'PFCP Session Establishment Request' data structure */
type PFCPSessionEstablishmentRequest struct {
	NodeID                             ie.NodeID                              //Node ID
	CPFSEID                            ie.FSEID                               //F-SEID
	CreatePDR                          []ie.CreatePDR                         //Create PDR
	CreateFAR                          []ie.CreateFAR                         //Create FAR
	CreateURR                          []ie.CreateURR                         //Create URR
	CreateQER                          []ie.CreateQER                         //Create QER
	CreateBAR                          *ie.CreateBAR                          //Create BAR
	CreateTrafficEndpoint              *ie.CreateTrafficEndpoint              //Create Traffic Endpoint
	PDNType                            *uint8                                 //PDN Type
	UserPlaneInactivityTimer           []byte                                 //User Plane Inactivity Timer
	UserID                             []byte                                 //User ID
	TraceInformation                   []byte                                 //Trace Information
	APNDNN                             []byte                                 //APN/DNN
	CreateMAR                          *ie.CreateMAR                          //Create MAR
	PFCPSEReqFlags                     *uint8                                 //PFCPSEReq-Flags
	CreateBridgeInfoforTSC             []byte                                 //Create Bridge Info for TSC
	CreateSRR                          *ie.CreateSRR                          //Create SRR
	ProvideATSSSControlInformation     *ie.ProvideATSSSControlInformation     //Provide ATSSS Control Information
	RecoveryTimeStamp                  *uint32                                //Recovery Time Stamp
	SNSSAI                             []byte                                 //S-NSSAI
	ProvideRDSconfigurationinformation *ie.ProvideRDSConfigurationInformation //Provide RDS Configuration Information
	RATType                            []byte                                 //RAT Type
	L2TPTunnelInformation              *ie.L2TPTunnelInformation              //L2TP Tunnel Information
	L2TPSessionInformation             *ie.L2TPSessionInformation             //L2TP Session Information
	GroupId                            []byte                                 //Group ID
	MBSSessionN4mbControlInformation   ie.MBSSessionN4mbControlInformation    //MBS Session N4mb Control Information
	MBSSessionN4ControlInformation     *ie.MBSSessionN4ControlInformation     //MBS Session N4 Control Information
	DSCPtoPPIControlInformation        *ie.DSCPtoPPIControlInformation        //DSCP to PPI Control Information
	FQCSIDList                         []ie.FQCSID                            //FQ-CSID
}

func (t *PFCPSessionEstablishmentRequest) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(ie.NODE_ID_TYPE, tlv.Bytes(t.NodeID.Bytes())))
	fields = append(fields, tlv.TLVFrom(ie.F_SEID_TYPE, tlv.Bytes(t.CPFSEID.Bytes())))
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
	if t.PDNType != nil {
		fields = append(fields, tlv.TLVFrom(ie.PDN_TYPE_TYPE, tlv.Bytes(pfcp.MarshalUint8(*t.PDNType))))
	}
	if len(t.UserPlaneInactivityTimer) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.USER_PLANE_INACTIVITY_TIMER_TYPE, tlv.Bytes(t.UserPlaneInactivityTimer)))
	}
	if len(t.UserID) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.USER_ID_TYPE, tlv.Bytes(t.UserID)))
	}
	if len(t.TraceInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.TRACE_INFORMATION_TYPE, tlv.Bytes(t.TraceInformation)))
	}
	if len(t.APNDNN) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.APN_DNN_TYPE, tlv.Bytes(t.APNDNN)))
	}
	if t.CreateMAR != nil {
		fields = append(fields, t.CreateMAR.Field())
	}
	if t.PFCPSEReqFlags != nil {
		fields = append(fields, tlv.TLVFrom(ie.PFCPSEREQ_FLAGS_TYPE, tlv.Bytes(pfcp.MarshalUint8(*t.PFCPSEReqFlags))))
	}
	if len(t.CreateBridgeInfoforTSC) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.CREATE_BRIDGE_INFO_FOR_TSC_TYPE, tlv.Bytes(t.CreateBridgeInfoforTSC)))
	}
	if t.CreateSRR != nil {
		fields = append(fields, t.CreateSRR.Field())
	}
	if t.ProvideATSSSControlInformation != nil {
		fields = append(fields, t.ProvideATSSSControlInformation.Field())
	}
	if t.RecoveryTimeStamp != nil {
		fields = append(fields, tlv.TLVFrom(ie.RECOVERY_TIME_STAMP_TYPE, tlv.Bytes(pfcp.MarshalUint32(*t.RecoveryTimeStamp))))
	}
	if len(t.SNSSAI) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.S_NSSAI_TYPE, tlv.Bytes(t.SNSSAI)))
	}
	if t.ProvideRDSconfigurationinformation != nil {
		fields = append(fields, t.ProvideRDSconfigurationinformation.Field())
	}
	if len(t.RATType) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.RAT_TYPE_TYPE, tlv.Bytes(t.RATType)))
	}
	if t.L2TPTunnelInformation != nil {
		fields = append(fields, t.L2TPTunnelInformation.Field())
	}
	if t.L2TPSessionInformation != nil {
		fields = append(fields, t.L2TPSessionInformation.Field())
	}
	if len(t.GroupId) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.GROUP_ID_TYPE, tlv.Bytes(t.GroupId)))
	}
	fields = append(fields, t.MBSSessionN4mbControlInformation.Field())
	if t.MBSSessionN4ControlInformation != nil {
		fields = append(fields, t.MBSSessionN4ControlInformation.Field())
	}
	if t.DSCPtoPPIControlInformation != nil {
		fields = append(fields, t.DSCPtoPPIControlInformation.Field())
	}
	for _, i := range t.FQCSIDList {
		fields = append(fields, tlv.TLVFrom(ie.FQ_CSID_TYPE, tlv.Bytes(i.Bytes())))
	}
	return tlv.TLV(ie.MSG_PFCP_SESSION_ESTABLISHMENT_REQUEST_TYPE, fields...)
}

func (t *PFCPSessionEstablishmentRequest) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPSessionEstablishmentRequest{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case ie.NODE_ID_TYPE:
			err = de.UnmarshalValue(&t.NodeID)
		case ie.F_SEID_TYPE:
			err = de.UnmarshalValue(&t.CPFSEID)
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
		case ie.PDN_TYPE_TYPE:
			var v uint8
			if err = pfcp.UnmarshalUint8(de.Value, &v); err == nil {
				t.PDNType = &v
			}
		case ie.USER_PLANE_INACTIVITY_TIMER_TYPE:
			t.UserPlaneInactivityTimer = de.Value
		case ie.USER_ID_TYPE:
			t.UserID = de.Value
		case ie.TRACE_INFORMATION_TYPE:
			t.TraceInformation = de.Value
		case ie.APN_DNN_TYPE:
			t.APNDNN = de.Value
		case ie.CREATE_MAR_TYPE:
			var v ie.CreateMAR
			if err = de.UnmarshalValue(&v); err == nil {
				t.CreateMAR = &v
			}
		case ie.PFCPSEREQ_FLAGS_TYPE:
			var v uint8
			if err = pfcp.UnmarshalUint8(de.Value, &v); err == nil {
				t.PFCPSEReqFlags = &v
			}
		case ie.CREATE_BRIDGE_INFO_FOR_TSC_TYPE:
			t.CreateBridgeInfoforTSC = de.Value
		case ie.CREATE_SRR_TYPE:
			var v ie.CreateSRR
			if err = de.UnmarshalValue(&v); err == nil {
				t.CreateSRR = &v
			}
		case ie.PROVIDE_ATSSS_CONTROL_INFORMATION_TYPE:
			var v ie.ProvideATSSSControlInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.ProvideATSSSControlInformation = &v
			}
		case ie.RECOVERY_TIME_STAMP_TYPE:
			var v uint32
			if err = pfcp.UnmarshalUint32(de.Value, &v); err == nil {
				t.RecoveryTimeStamp = &v
			}
		case ie.S_NSSAI_TYPE:
			t.SNSSAI = de.Value
		case ie.PROVIDE_RDS_CONFIGURATION_INFORMATION_TYPE:
			var v ie.ProvideRDSConfigurationInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.ProvideRDSconfigurationinformation = &v
			}
		case ie.RAT_TYPE_TYPE:
			t.RATType = de.Value
		case ie.L2TP_TUNNEL_INFORMATION_TYPE:
			var v ie.L2TPTunnelInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.L2TPTunnelInformation = &v
			}
		case ie.L2TP_SESSION_INFORMATION_TYPE:
			var v ie.L2TPSessionInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.L2TPSessionInformation = &v
			}
		case ie.GROUP_ID_TYPE:
			t.GroupId = de.Value
		case ie.MBS_SESSION_N4MB_CONTROL_INFORMATION_TYPE:
			err = de.UnmarshalValue(&t.MBSSessionN4mbControlInformation)
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

func (t *PFCPSessionEstablishmentRequest) UnmarshalTLV(typ uint32, buf []byte) (err error) {
	if typ != ie.MSG_PFCP_SESSION_ESTABLISHMENT_REQUEST_TYPE {
		err = tlv.ErrCritical
		return
	}

	err = t.UnmarshalBinary(buf)
	return
}

func (t *PFCPSessionEstablishmentRequest) MarshalJSON() (buf []byte, err error) {
	buf, err = jsonMarshal(t)
	return
}

func (t *PFCPSessionEstablishmentRequest) UnmarshalJSON(buf []byte) (err error) {
	*t = PFCPSessionEstablishmentRequest{}
	err = jsonUnmarshal(buf, t)
	return
}

/* 'PFCP Session Establishment Response' data structure */
type PFCPSessionEstablishmentResponse struct {
	NodeID                      ie.NodeID                      //Node ID
	Cause                       uint8                          //Cause
	OffendingIE                 *uint16                        //Offending IE
	UPFSEID                     *ie.FSEID                      //F-SEID
	CreatedPDR                  []ie.CreatedPDR                //Created PDR
	LoadControlInformation      *ie.LoadControlInformation     //Load Control Information
	OverloadControlInformation  *ie.OverloadControlInformation //Overload Control Information
	PGWUSGWUUPFFQCSID           *ie.FQCSID                     //FQ-CSID
	FailedRuleID                []byte                         //Failed Rule ID
	CreatedTrafficEndpoint      *ie.CreatedTrafficEndpoint     //Created Traffic Endpoint
	CreatedBridgeInfoforTSC     *ie.CreatedBridgeInfoforTSC    //Created Bridge Info for TSC
	ATSSSControlParameters      *ie.ATSSSControlParameters     //ATSSS Control Parameters
	RDSconfigurationinformation []byte                         //RDS Configuration Information
	PartialFailureInformation   *ie.PartialFailureInformation  //Partial Failure Information
	CreatedL2TPSession          *ie.CreatedL2TPSession         //Created L2TP Session
	MBSSessionN4mbInformation   *ie.MBSSessionN4mbInformation  //MBS Session N4mb Information
	MBSSessionN4Information     *ie.MBSSessionN4Information    //MBS Session N4 Information
}

func (t *PFCPSessionEstablishmentResponse) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(ie.NODE_ID_TYPE, tlv.Bytes(t.NodeID.Bytes())))
	fields = append(fields, tlv.TLVFrom(ie.CAUSE_TYPE, tlv.Bytes(pfcp.MarshalUint8(t.Cause))))
	if t.OffendingIE != nil {
		fields = append(fields, tlv.TLVFrom(ie.OFFENDING_IE_TYPE, tlv.Bytes(pfcp.MarshalUint16(*t.OffendingIE))))
	}
	if t.UPFSEID != nil {
		fields = append(fields, tlv.TLVFrom(ie.F_SEID_TYPE, tlv.Bytes(t.UPFSEID.Bytes())))
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
	if t.PGWUSGWUUPFFQCSID != nil {
		fields = append(fields, tlv.TLVFrom(ie.FQ_CSID_TYPE, tlv.Bytes(t.PGWUSGWUUPFFQCSID.Bytes())))
	}
	if len(t.FailedRuleID) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.FAILED_RULE_ID_TYPE, tlv.Bytes(t.FailedRuleID)))
	}
	if t.CreatedTrafficEndpoint != nil {
		fields = append(fields, t.CreatedTrafficEndpoint.Field())
	}
	if t.CreatedBridgeInfoforTSC != nil {
		fields = append(fields, t.CreatedBridgeInfoforTSC.Field())
	}
	if t.ATSSSControlParameters != nil {
		fields = append(fields, t.ATSSSControlParameters.Field())
	}
	if len(t.RDSconfigurationinformation) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.RDS_CONFIGURATION_INFORMATION_TYPE, tlv.Bytes(t.RDSconfigurationinformation)))
	}
	if t.PartialFailureInformation != nil {
		fields = append(fields, t.PartialFailureInformation.Field())
	}
	if t.CreatedL2TPSession != nil {
		fields = append(fields, t.CreatedL2TPSession.Field())
	}
	if t.MBSSessionN4mbInformation != nil {
		fields = append(fields, t.MBSSessionN4mbInformation.Field())
	}
	if t.MBSSessionN4Information != nil {
		fields = append(fields, t.MBSSessionN4Information.Field())
	}
	return tlv.TLV(ie.MSG_PFCP_SESSION_ESTABLISHMENT_RESPONSE_TYPE, fields...)
}

func (t *PFCPSessionEstablishmentResponse) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPSessionEstablishmentResponse{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case ie.NODE_ID_TYPE:
			err = de.UnmarshalValue(&t.NodeID)
		case ie.CAUSE_TYPE:
			err = pfcp.UnmarshalUint8(de.Value, &t.Cause)
		case ie.OFFENDING_IE_TYPE:
			var v uint16
			if err = pfcp.UnmarshalUint16(de.Value, &v); err == nil {
				t.OffendingIE = &v
			}
		case ie.F_SEID_TYPE:
			var v ie.FSEID
			if err = de.UnmarshalValue(&v); err == nil {
				t.UPFSEID = &v
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
		case ie.FQ_CSID_TYPE:
			var v ie.FQCSID
			if err = de.UnmarshalValue(&v); err == nil {
				t.PGWUSGWUUPFFQCSID = &v
			}
		case ie.FAILED_RULE_ID_TYPE:
			t.FailedRuleID = de.Value
		case ie.CREATED_TRAFFIC_ENDPOINT_TYPE:
			var v ie.CreatedTrafficEndpoint
			if err = de.UnmarshalValue(&v); err == nil {
				t.CreatedTrafficEndpoint = &v
			}
		case ie.CREATED_BRIDGE_INFO_FOR_TSC_TYPE:
			var v ie.CreatedBridgeInfoforTSC
			if err = de.UnmarshalValue(&v); err == nil {
				t.CreatedBridgeInfoforTSC = &v
			}
		case ie.ATSSS_CONTROL_PARAMETERS_TYPE:
			var v ie.ATSSSControlParameters
			if err = de.UnmarshalValue(&v); err == nil {
				t.ATSSSControlParameters = &v
			}
		case ie.RDS_CONFIGURATION_INFORMATION_TYPE:
			t.RDSconfigurationinformation = de.Value
		case ie.PARTIAL_FAILURE_INFORMATION_TYPE:
			var v ie.PartialFailureInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.PartialFailureInformation = &v
			}
		case ie.CREATED_L2TP_SESSION_TYPE:
			var v ie.CreatedL2TPSession
			if err = de.UnmarshalValue(&v); err == nil {
				t.CreatedL2TPSession = &v
			}
		case ie.MBS_SESSION_N4MB_INFORMATION_TYPE:
			var v ie.MBSSessionN4mbInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.MBSSessionN4mbInformation = &v
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

func (t *PFCPSessionEstablishmentResponse) UnmarshalTLV(typ uint32, buf []byte) (err error) {
	if typ != ie.MSG_PFCP_SESSION_ESTABLISHMENT_RESPONSE_TYPE {
		err = tlv.ErrCritical
		return
	}

	err = t.UnmarshalBinary(buf)
	return
}

func (t *PFCPSessionEstablishmentResponse) MarshalJSON() (buf []byte, err error) {
	buf, err = jsonMarshal(t)
	return
}

func (t *PFCPSessionEstablishmentResponse) UnmarshalJSON(buf []byte) (err error) {
	*t = PFCPSessionEstablishmentResponse{}
	err = jsonUnmarshal(buf, t)
	return
}
