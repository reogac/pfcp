package message

import (
	"github.com/reogac/pfcp"
	"github.com/reogac/pfcp/ie"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'PFCP Session Report Request' data structure */
type PFCPSessionReportRequest struct {
	ReportType                        uint8                                                              //Report Type
	DownlinkDataReport                *ie.DownlinkDataReport                                             //Downlink Data Report
	UsageReport                       []ie.UsageReportSessionReportRequest                               //Usage Report Session Report Request
	ErrorIndicationReport             *ie.ErrorIndicationReport                                          //Error Indication Report
	LoadControlInformation            *ie.LoadControlInformation                                         //Load Control Information
	OverloadControlInformation        *ie.OverloadControlInformation                                     //Overload Control Information
	AdditionalUsageReportsInformation []byte                                                             //Additional Usage Reports Information
	PFCPSRReqFlags                    *uint8                                                             //PFCPSRReq-Flags
	OldCPFSEID                        *ie.FSEID                                                          //F-SEID
	PacketRateStatusReport            *ie.PacketRateStatusReport                                         //Packet Rate Status Report
	TSCManagementInformation          *ie.TSCManagementInformationIEwithinPFCPSessionModificationRequest //TSC Management Information IE within PFCP Session Modification Request
	SessionReport                     *ie.SessionReport                                                  //Session Report
	Cause                             *uint8                                                             //Cause
}

func (t *PFCPSessionReportRequest) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(ie.REPORT_TYPE_TYPE, tlv.Bytes(pfcp.MarshalUint8(t.ReportType))))
	if t.DownlinkDataReport != nil {
		fields = append(fields, t.DownlinkDataReport.Field())
	}
	for _, i := range t.UsageReport {
		fields = append(fields, i.Field())
	}
	if t.ErrorIndicationReport != nil {
		fields = append(fields, t.ErrorIndicationReport.Field())
	}
	if t.LoadControlInformation != nil {
		fields = append(fields, t.LoadControlInformation.Field())
	}
	if t.OverloadControlInformation != nil {
		fields = append(fields, t.OverloadControlInformation.Field())
	}
	if len(t.AdditionalUsageReportsInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.ADDITIONAL_USAGE_REPORTS_INFORMATION_TYPE, tlv.Bytes(t.AdditionalUsageReportsInformation)))
	}
	if t.PFCPSRReqFlags != nil {
		fields = append(fields, tlv.TLVFrom(ie.PFCPSRREQ_FLAGS_TYPE, tlv.Bytes(pfcp.MarshalUint8(*t.PFCPSRReqFlags))))
	}
	if t.OldCPFSEID != nil {
		fields = append(fields, tlv.TLVFrom(ie.F_SEID_TYPE, tlv.Bytes(t.OldCPFSEID.Bytes())))
	}
	if t.PacketRateStatusReport != nil {
		fields = append(fields, t.PacketRateStatusReport.Field())
	}
	if t.TSCManagementInformation != nil {
		fields = append(fields, t.TSCManagementInformation.Field())
	}
	if t.SessionReport != nil {
		fields = append(fields, t.SessionReport.Field())
	}
	if t.Cause != nil {
		fields = append(fields, tlv.TLVFrom(ie.CAUSE_TYPE, tlv.Bytes(pfcp.MarshalUint8(*t.Cause))))
	}
	return tlv.TLV(ie.MSG_PFCP_SESSION_REPORT_REQUEST_TYPE, fields...)
}

func (t *PFCPSessionReportRequest) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPSessionReportRequest{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case ie.REPORT_TYPE_TYPE:
			err = pfcp.UnmarshalUint8(de.Value, &t.ReportType)
		case ie.DOWNLINK_DATA_REPORT_TYPE:
			var v ie.DownlinkDataReport
			if err = de.UnmarshalValue(&v); err == nil {
				t.DownlinkDataReport = &v
			}
		case ie.USAGE_REPORT_SESSION_REPORT_REQUEST_TYPE:
			var v ie.UsageReportSessionReportRequest
			if err = de.UnmarshalValue(&v); err == nil {
				t.UsageReport = append(t.UsageReport, v)
			}
		case ie.ERROR_INDICATION_REPORT_TYPE:
			var v ie.ErrorIndicationReport
			if err = de.UnmarshalValue(&v); err == nil {
				t.ErrorIndicationReport = &v
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
		case ie.ADDITIONAL_USAGE_REPORTS_INFORMATION_TYPE:
			t.AdditionalUsageReportsInformation = de.Value
		case ie.PFCPSRREQ_FLAGS_TYPE:
			var v uint8
			if err = pfcp.UnmarshalUint8(de.Value, &v); err == nil {
				t.PFCPSRReqFlags = &v
			}
		case ie.F_SEID_TYPE:
			var v ie.FSEID
			if err = de.UnmarshalValue(&v); err == nil {
				t.OldCPFSEID = &v
			}
		case ie.PACKET_RATE_STATUS_REPORT_TYPE:
			var v ie.PacketRateStatusReport
			if err = de.UnmarshalValue(&v); err == nil {
				t.PacketRateStatusReport = &v
			}
		case ie.TSC_MANAGEMENT_INFORMATION_IE_WITHIN_PFCP_SESSION_MODIFICATION_REQUEST_TYPE:
			var v ie.TSCManagementInformationIEwithinPFCPSessionModificationRequest
			if err = de.UnmarshalValue(&v); err == nil {
				t.TSCManagementInformation = &v
			}
		case ie.SESSION_REPORT_TYPE:
			var v ie.SessionReport
			if err = de.UnmarshalValue(&v); err == nil {
				t.SessionReport = &v
			}
		case ie.CAUSE_TYPE:
			var v uint8
			if err = pfcp.UnmarshalUint8(de.Value, &v); err == nil {
				t.Cause = &v
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

func (t *PFCPSessionReportRequest) UnmarshalTLV(typ uint32, buf []byte) (err error) {
	if typ != ie.MSG_PFCP_SESSION_REPORT_REQUEST_TYPE {
		err = tlv.ErrCritical
		return
	}

	err = t.UnmarshalBinary(buf)
	return
}

func (t *PFCPSessionReportRequest) MarshalJSON() (buf []byte, err error) {
	buf, err = jsonMarshal(t)
	return
}

func (t *PFCPSessionReportRequest) UnmarshalJSON(buf []byte) (err error) {
	*t = PFCPSessionReportRequest{}
	err = jsonUnmarshal(buf, t)
	return
}

/* 'PFCP Session Report Response' data structure */
type PFCPSessionReportResponse struct {
	Cause                   uint8                                  //Cause
	OffendingIE             *uint16                                //Offending IE
	UpdateBAR               *ie.UpdateBARPFCPSessionReportResponse //Update BAR PFCP Session Report Response
	PFCPSRRspFlags          *uint8                                 //PFCPSRRsp-Flags
	CPFSEID                 *ie.FSEID                              //F-SEID
	N4uFTEID                *ie.FTEID                              //F-TEID
	AlternativeSMFIPAddress []byte                                 //Alternative SMF IP Address
	PGWCSMFFQCSID           *ie.FQCSID                             //FQ-CSID
	GroupId                 []byte                                 //Group ID
	NodeID                  *ie.NodeID                             //Node ID
}

func (t *PFCPSessionReportResponse) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(ie.CAUSE_TYPE, tlv.Bytes(pfcp.MarshalUint8(t.Cause))))
	if t.OffendingIE != nil {
		fields = append(fields, tlv.TLVFrom(ie.OFFENDING_IE_TYPE, tlv.Bytes(pfcp.MarshalUint16(*t.OffendingIE))))
	}
	if t.UpdateBAR != nil {
		fields = append(fields, t.UpdateBAR.Field())
	}
	if t.PFCPSRRspFlags != nil {
		fields = append(fields, tlv.TLVFrom(ie.PFCPSRRSP_FLAGS_TYPE, tlv.Bytes(pfcp.MarshalUint8(*t.PFCPSRRspFlags))))
	}
	if t.CPFSEID != nil {
		fields = append(fields, tlv.TLVFrom(ie.F_SEID_TYPE, tlv.Bytes(t.CPFSEID.Bytes())))
	}
	if t.N4uFTEID != nil {
		fields = append(fields, tlv.TLVFrom(ie.F_TEID_TYPE, tlv.Bytes(t.N4uFTEID.Bytes())))
	}
	if len(t.AlternativeSMFIPAddress) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.ALTERNATIVE_SMF_IP_ADDRESS_TYPE, tlv.Bytes(t.AlternativeSMFIPAddress)))
	}
	if t.PGWCSMFFQCSID != nil {
		fields = append(fields, tlv.TLVFrom(ie.FQ_CSID_TYPE, tlv.Bytes(t.PGWCSMFFQCSID.Bytes())))
	}
	if len(t.GroupId) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.GROUP_ID_TYPE, tlv.Bytes(t.GroupId)))
	}
	if t.NodeID != nil {
		fields = append(fields, tlv.TLVFrom(ie.NODE_ID_TYPE, tlv.Bytes(t.NodeID.Bytes())))
	}
	return tlv.TLV(ie.MSG_PFCP_SESSION_REPORT_RESPONSE_TYPE, fields...)
}

func (t *PFCPSessionReportResponse) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPSessionReportResponse{}
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
		case ie.UPDATE_BAR_PFCP_SESSION_REPORT_RESPONSE_TYPE:
			var v ie.UpdateBARPFCPSessionReportResponse
			if err = de.UnmarshalValue(&v); err == nil {
				t.UpdateBAR = &v
			}
		case ie.PFCPSRRSP_FLAGS_TYPE:
			var v uint8
			if err = pfcp.UnmarshalUint8(de.Value, &v); err == nil {
				t.PFCPSRRspFlags = &v
			}
		case ie.F_SEID_TYPE:
			var v ie.FSEID
			if err = de.UnmarshalValue(&v); err == nil {
				t.CPFSEID = &v
			}
		case ie.F_TEID_TYPE:
			var v ie.FTEID
			if err = de.UnmarshalValue(&v); err == nil {
				t.N4uFTEID = &v
			}
		case ie.ALTERNATIVE_SMF_IP_ADDRESS_TYPE:
			t.AlternativeSMFIPAddress = de.Value
		case ie.FQ_CSID_TYPE:
			var v ie.FQCSID
			if err = de.UnmarshalValue(&v); err == nil {
				t.PGWCSMFFQCSID = &v
			}
		case ie.GROUP_ID_TYPE:
			t.GroupId = de.Value
		case ie.NODE_ID_TYPE:
			var v ie.NodeID
			if err = de.UnmarshalValue(&v); err == nil {
				t.NodeID = &v
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

func (t *PFCPSessionReportResponse) UnmarshalTLV(typ uint32, buf []byte) (err error) {
	if typ != ie.MSG_PFCP_SESSION_REPORT_RESPONSE_TYPE {
		err = tlv.ErrCritical
		return
	}

	err = t.UnmarshalBinary(buf)
	return
}

func (t *PFCPSessionReportResponse) MarshalJSON() (buf []byte, err error) {
	buf, err = jsonMarshal(t)
	return
}

func (t *PFCPSessionReportResponse) UnmarshalJSON(buf []byte) (err error) {
	*t = PFCPSessionReportResponse{}
	err = jsonUnmarshal(buf, t)
	return
}
