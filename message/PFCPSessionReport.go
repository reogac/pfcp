package message

import (
	"github.com/reogac/pfcp/v2"
	"github.com/reogac/pfcp/v2/ie"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'PFCP Session Report Request' data structure */
type PFCPSessionReportRequest struct {
	ReportType                        uint8                                                               //Report Type
	DownlinkDataReport                *ie.DownlinkDataReport                                              //Downlink Data Report
	UsageReport                       []ie.UsageReportSessionReportRequest                                //Usage Report Session Report Request
	ErrorIndicationReport             *ie.ErrorIndicationReport                                           //Error Indication Report
	LoadControlInformation            *ie.LoadControlInformation                                          //Load Control Information
	OverloadControlInformation        *ie.OverloadControlInformation                                      //Overload Control Information
	AdditionalUsageReportsInformation []byte                                                              //Additional Usage Reports Information
	PFCPSRReqFlags                    *uint8                                                              //PFCPSRReq-Flags
	OldCPFSEID                        *ie.FSEID                                                           //F-SEID
	PacketRateStatusReport            *ie.PacketRateStatusReport                                          //Packet Rate Status Report
	TSCManagementInformation          []ie.TSCManagementInformationIEwithinPFCPSessionModificationRequest //TSC Management Information IE within PFCP Session Modification Request
	SessionReport                     []ie.SessionReport                                                  //Session Report
	Cause                             *uint8                                                              //Cause
	Unknown                           []uint32                                                            `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid                           []uint32                                                            `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
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
	for _, i := range t.TSCManagementInformation {
		fields = append(fields, i.Field())
	}
	for _, i := range t.SessionReport {
		fields = append(fields, i.Field())
	}
	if t.Cause != nil {
		fields = append(fields, tlv.TLVFrom(ie.CAUSE_TYPE, tlv.Bytes(pfcp.MarshalUint8(*t.Cause))))
	}
	return tlv.TLV(ie.MSG_PFCP_SESSION_REPORT_REQUEST_TYPE, fields...)
}

// Skipped is what decoding left out: IEs unknown or not defined here
// (TS 29.244 7.6.9), and optional ones discarded for a faulty value (7.6.8).
func (t *PFCPSessionReportRequest) Skipped() (unknown, invalid []uint32) {
	return t.Unknown, t.Invalid
}

func (t *PFCPSessionReportRequest) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPSessionReportRequest{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [10]bool
	for _, de := range elements {
		opt := false
		switch de.Type {
		case ie.REPORT_TYPE_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &t.ReportType)
		case ie.DOWNLINK_DATA_REPORT_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			var v ie.DownlinkDataReport
			if err = de.UnmarshalValue(&v); err == nil {
				t.DownlinkDataReport = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.USAGE_REPORT_SESSION_REPORT_REQUEST_TYPE:
			var v ie.UsageReportSessionReportRequest
			if err = de.UnmarshalValue(&v); err == nil {
				t.UsageReport = append(t.UsageReport, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.ERROR_INDICATION_REPORT_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			var v ie.ErrorIndicationReport
			if err = de.UnmarshalValue(&v); err == nil {
				t.ErrorIndicationReport = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.LOAD_CONTROL_INFORMATION_TYPE:
			opt = true
			if seen[3] {
				continue
			}
			seen[3] = true
			var v ie.LoadControlInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.LoadControlInformation = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.OVERLOAD_CONTROL_INFORMATION_TYPE:
			opt = true
			if seen[4] {
				continue
			}
			seen[4] = true
			var v ie.OverloadControlInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.OverloadControlInformation = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.ADDITIONAL_USAGE_REPORTS_INFORMATION_TYPE:
			if seen[5] {
				continue
			}
			seen[5] = true
			t.AdditionalUsageReportsInformation = de.Value
		case ie.PFCPSRREQ_FLAGS_TYPE:
			if seen[6] {
				continue
			}
			seen[6] = true
			var v uint8
			if err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &v); err == nil {
				t.PFCPSRReqFlags = &v
			}
		case ie.F_SEID_TYPE:
			if seen[7] {
				continue
			}
			seen[7] = true
			var v ie.FSEID
			if err = de.UnmarshalValue(&v); err == nil {
				t.OldCPFSEID = &v
			}
		case ie.PACKET_RATE_STATUS_REPORT_TYPE:
			if seen[8] {
				continue
			}
			seen[8] = true
			var v ie.PacketRateStatusReport
			if err = de.UnmarshalValue(&v); err == nil {
				t.PacketRateStatusReport = &v
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
		case ie.SESSION_REPORT_TYPE:
			var v ie.SessionReport
			if err = de.UnmarshalValue(&v); err == nil {
				t.SessionReport = append(t.SessionReport, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.CAUSE_TYPE:
			opt = true
			if seen[9] {
				continue
			}
			seen[9] = true
			var v uint8
			if err = pfcp.UnmarshalUint8(de.Value, &v); err == nil {
				t.Cause = &v
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
		return &pfcp.MissingIE{Type: ie.REPORT_TYPE_TYPE}
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
	Unknown                 []uint32                               `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid                 []uint32                               `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
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

// Skipped is what decoding left out: IEs unknown or not defined here
// (TS 29.244 7.6.9), and optional ones discarded for a faulty value (7.6.8).
func (t *PFCPSessionReportResponse) Skipped() (unknown, invalid []uint32) {
	return t.Unknown, t.Invalid
}

func (t *PFCPSessionReportResponse) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPSessionReportResponse{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [10]bool
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
		case ie.UPDATE_BAR_PFCP_SESSION_REPORT_RESPONSE_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			var v ie.UpdateBARPFCPSessionReportResponse
			if err = de.UnmarshalValue(&v); err == nil {
				t.UpdateBAR = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.PFCPSRRSP_FLAGS_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			var v uint8
			if err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &v); err == nil {
				t.PFCPSRRspFlags = &v
			}
		case ie.F_SEID_TYPE:
			opt = true
			if seen[4] {
				continue
			}
			seen[4] = true
			var v ie.FSEID
			if err = de.UnmarshalValue(&v); err == nil {
				t.CPFSEID = &v
			}
		case ie.F_TEID_TYPE:
			opt = true
			if seen[5] {
				continue
			}
			seen[5] = true
			var v ie.FTEID
			if err = de.UnmarshalValue(&v); err == nil {
				t.N4uFTEID = &v
			}
		case ie.ALTERNATIVE_SMF_IP_ADDRESS_TYPE:
			opt = true
			if seen[6] {
				continue
			}
			seen[6] = true
			t.AlternativeSMFIPAddress = de.Value
		case ie.FQ_CSID_TYPE:
			opt = true
			if seen[7] {
				continue
			}
			seen[7] = true
			var v ie.FQCSID
			if err = de.UnmarshalValue(&v); err == nil {
				t.PGWCSMFFQCSID = &v
			}
		case ie.GROUP_ID_TYPE:
			if seen[8] {
				continue
			}
			seen[8] = true
			t.GroupId = de.Value
		case ie.NODE_ID_TYPE:
			if seen[9] {
				continue
			}
			seen[9] = true
			var v ie.NodeID
			if err = de.UnmarshalValue(&v); err == nil {
				t.NodeID = &v
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
