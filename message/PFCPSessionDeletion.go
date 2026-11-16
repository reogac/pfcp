package message

import (
	"github.com/reogac/pfcp/v2"
	"github.com/reogac/pfcp/v2/ie"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'PFCP Session Deletion Request' data structure */
type PFCPSessionDeletionRequest struct {
	Unknown []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *PFCPSessionDeletionRequest) Field() tlv.Field {
	fields := []tlv.Field{}
	return tlv.TLV(ie.MSG_PFCP_SESSION_DELETION_REQUEST_TYPE, fields...)
}

func (t *PFCPSessionDeletionRequest) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPSessionDeletionRequest{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	for _, de := range elements {
		switch de.Type {
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	return
}

func (t *PFCPSessionDeletionRequest) UnmarshalTLV(typ uint32, buf []byte) (err error) {
	if typ != ie.MSG_PFCP_SESSION_DELETION_REQUEST_TYPE {
		err = tlv.ErrCritical
		return
	}

	err = t.UnmarshalBinary(buf)
	return
}

func (t *PFCPSessionDeletionRequest) MarshalJSON() (buf []byte, err error) {
	buf, err = jsonMarshal(t)
	return
}

func (t *PFCPSessionDeletionRequest) UnmarshalJSON(buf []byte) (err error) {
	*t = PFCPSessionDeletionRequest{}
	err = jsonUnmarshal(buf, t)
	return
}

/* 'PFCP Session Deletion Response' data structure */
type PFCPSessionDeletionResponse struct {
	Cause                             uint8                                   //Cause
	OffendingIE                       *uint16                                 //Offending IE
	LoadControlInformation            *ie.LoadControlInformation              //Load Control Information
	OverloadControlInformation        *ie.OverloadControlInformation          //Overload Control Information
	UsageReport                       []ie.UsageReportSessionDeletionResponse //Usage Report Session Deletion Response
	AdditionalUsageReportsInformation []byte                                  //Additional Usage Reports Information
	PacketRateStatusReport            *ie.PacketRateStatusReport              //Packet Rate Status Report
	SessionReport                     []ie.SessionReport                      //Session Report
	MBSSessionN4Information           []ie.MBSSessionN4Information            //MBS Session N4 Information
	PFCPSDRspFlags                    []byte                                  //PFCPSDRsp-Flags
	Unknown                           []uint32                                `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *PFCPSessionDeletionResponse) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(ie.CAUSE_TYPE, tlv.Bytes(pfcp.MarshalUint8(t.Cause))))
	if t.OffendingIE != nil {
		fields = append(fields, tlv.TLVFrom(ie.OFFENDING_IE_TYPE, tlv.Bytes(pfcp.MarshalUint16(*t.OffendingIE))))
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
	if len(t.AdditionalUsageReportsInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.ADDITIONAL_USAGE_REPORTS_INFORMATION_TYPE, tlv.Bytes(t.AdditionalUsageReportsInformation)))
	}
	if t.PacketRateStatusReport != nil {
		fields = append(fields, t.PacketRateStatusReport.Field())
	}
	for _, i := range t.SessionReport {
		fields = append(fields, i.Field())
	}
	for _, i := range t.MBSSessionN4Information {
		fields = append(fields, i.Field())
	}
	if len(t.PFCPSDRspFlags) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.PFCPSDRSP_FLAGS_TYPE, tlv.Bytes(t.PFCPSDRspFlags)))
	}
	return tlv.TLV(ie.MSG_PFCP_SESSION_DELETION_RESPONSE_TYPE, fields...)
}

func (t *PFCPSessionDeletionResponse) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPSessionDeletionResponse{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	var seen [7]bool
	for _, de := range elements {
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
		case ie.LOAD_CONTROL_INFORMATION_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			var v ie.LoadControlInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.LoadControlInformation = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
			}
		case ie.OVERLOAD_CONTROL_INFORMATION_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			var v ie.OverloadControlInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.OverloadControlInformation = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
			}
		case ie.USAGE_REPORT_SESSION_DELETION_RESPONSE_TYPE:
			var v ie.UsageReportSessionDeletionResponse
			if err = de.UnmarshalValue(&v); err == nil {
				t.UsageReport = append(t.UsageReport, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
			}
		case ie.ADDITIONAL_USAGE_REPORTS_INFORMATION_TYPE:
			if seen[4] {
				continue
			}
			seen[4] = true
			t.AdditionalUsageReportsInformation = de.Value
		case ie.PACKET_RATE_STATUS_REPORT_TYPE:
			if seen[5] {
				continue
			}
			seen[5] = true
			var v ie.PacketRateStatusReport
			if err = de.UnmarshalValue(&v); err == nil {
				t.PacketRateStatusReport = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
			}
		case ie.SESSION_REPORT_TYPE:
			var v ie.SessionReport
			if err = de.UnmarshalValue(&v); err == nil {
				t.SessionReport = append(t.SessionReport, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
			}
		case ie.MBS_SESSION_N4_INFORMATION_TYPE:
			var v ie.MBSSessionN4Information
			if err = de.UnmarshalValue(&v); err == nil {
				t.MBSSessionN4Information = append(t.MBSSessionN4Information, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
			}
		case ie.PFCPSDRSP_FLAGS_TYPE:
			if seen[6] {
				continue
			}
			seen[6] = true
			t.PFCPSDRspFlags = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: ie.CAUSE_TYPE}
	}
	return
}

func (t *PFCPSessionDeletionResponse) UnmarshalTLV(typ uint32, buf []byte) (err error) {
	if typ != ie.MSG_PFCP_SESSION_DELETION_RESPONSE_TYPE {
		err = tlv.ErrCritical
		return
	}

	err = t.UnmarshalBinary(buf)
	return
}

func (t *PFCPSessionDeletionResponse) MarshalJSON() (buf []byte, err error) {
	buf, err = jsonMarshal(t)
	return
}

func (t *PFCPSessionDeletionResponse) UnmarshalJSON(buf []byte) (err error) {
	*t = PFCPSessionDeletionResponse{}
	err = jsonUnmarshal(buf, t)
	return
}
