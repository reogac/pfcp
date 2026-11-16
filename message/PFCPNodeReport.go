package message

import (
	"github.com/reogac/pfcp/v2"
	"github.com/reogac/pfcp/v2/ie"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'PFCP Node Report Request' data structure */
type PFCPNodeReportRequest struct {
	NodeID                       ie.NodeID                                   //Node ID
	NodeReportType               []byte                                      //Node Report Type
	VendorSpecificNodeReportType []byte                                      //Vendor-Specific Node Report Type
	UserPlanePathFailureReport   *ie.UserPlanePathFailureReport              //User Plane Path Failure Report
	UserPlanePathRecoveryReport  *ie.UserPlanePathRecoveryReport             //User Plane Path Recovery Report
	ClockDriftReport             []ie.ClockDriftReport                       //Clock Drift Report
	GTPUPathQoSReport            []ie.GTPUPathQoSReportPFCPNodeReportRequest //GTP-U Path QoS Report PFCP Node Report Request
	PeerUPRestartReport          []byte                                      //Peer UP Restart Report
	Unknown                      []uint32                                    `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *PFCPNodeReportRequest) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(ie.NODE_ID_TYPE, tlv.Bytes(t.NodeID.Bytes())))
	fields = append(fields, tlv.TLVFrom(ie.NODE_REPORT_TYPE_TYPE, tlv.Bytes(t.NodeReportType)))
	if len(t.VendorSpecificNodeReportType) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.VENDOR_SPECIFIC_NODE_REPORT_TYPE_TYPE, tlv.Bytes(t.VendorSpecificNodeReportType)))
	}
	if t.UserPlanePathFailureReport != nil {
		fields = append(fields, t.UserPlanePathFailureReport.Field())
	}
	if t.UserPlanePathRecoveryReport != nil {
		fields = append(fields, t.UserPlanePathRecoveryReport.Field())
	}
	for _, i := range t.ClockDriftReport {
		fields = append(fields, i.Field())
	}
	for _, i := range t.GTPUPathQoSReport {
		fields = append(fields, i.Field())
	}
	if len(t.PeerUPRestartReport) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.PEER_UP_RESTART_REPORT_TYPE, tlv.Bytes(t.PeerUPRestartReport)))
	}
	return tlv.TLV(ie.MSG_PFCP_NODE_REPORT_REQUEST_TYPE, fields...)
}

func (t *PFCPNodeReportRequest) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPNodeReportRequest{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [6]bool
	for _, de := range elements {
		switch de.Type {
		case ie.NODE_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			if err = de.UnmarshalValue(&t.NodeID); err == nil {
			}
		case ie.NODE_REPORT_TYPE_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			t.NodeReportType = de.Value
		case ie.VENDOR_SPECIFIC_NODE_REPORT_TYPE_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			t.VendorSpecificNodeReportType = de.Value
		case ie.USER_PLANE_PATH_FAILURE_REPORT_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			var v ie.UserPlanePathFailureReport
			if err = de.UnmarshalValue(&v); err == nil {
				t.UserPlanePathFailureReport = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
			}
		case ie.USER_PLANE_PATH_RECOVERY_REPORT_TYPE:
			if seen[4] {
				continue
			}
			seen[4] = true
			var v ie.UserPlanePathRecoveryReport
			if err = de.UnmarshalValue(&v); err == nil {
				t.UserPlanePathRecoveryReport = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
			}
		case ie.CLOCK_DRIFT_REPORT_TYPE:
			var v ie.ClockDriftReport
			if err = de.UnmarshalValue(&v); err == nil {
				t.ClockDriftReport = append(t.ClockDriftReport, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
			}
		case ie.GTP_U_PATH_QOS_REPORT_PFCP_NODE_REPORT_REQUEST_TYPE:
			var v ie.GTPUPathQoSReportPFCPNodeReportRequest
			if err = de.UnmarshalValue(&v); err == nil {
				t.GTPUPathQoSReport = append(t.GTPUPathQoSReport, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
			}
		case ie.PEER_UP_RESTART_REPORT_TYPE:
			if seen[5] {
				continue
			}
			seen[5] = true
			t.PeerUPRestartReport = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return pfcp.Faulty(de.Type, err)
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: ie.NODE_ID_TYPE}
	}
	if !seen[1] {
		return &pfcp.MissingIE{Type: ie.NODE_REPORT_TYPE_TYPE}
	}
	return
}

func (t *PFCPNodeReportRequest) UnmarshalTLV(typ uint32, buf []byte) (err error) {
	if typ != ie.MSG_PFCP_NODE_REPORT_REQUEST_TYPE {
		err = tlv.ErrCritical
		return
	}

	err = t.UnmarshalBinary(buf)
	return
}

func (t *PFCPNodeReportRequest) MarshalJSON() (buf []byte, err error) {
	buf, err = jsonMarshal(t)
	return
}

func (t *PFCPNodeReportRequest) UnmarshalJSON(buf []byte) (err error) {
	*t = PFCPNodeReportRequest{}
	err = jsonUnmarshal(buf, t)
	return
}

/* 'PFCP Node Report Response' data structure */
type PFCPNodeReportResponse struct {
	NodeID      ie.NodeID //Node ID
	Cause       uint8     //Cause
	OffendingIE *uint16   //Offending IE
	Unknown     []uint32  `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *PFCPNodeReportResponse) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(ie.NODE_ID_TYPE, tlv.Bytes(t.NodeID.Bytes())))
	fields = append(fields, tlv.TLVFrom(ie.CAUSE_TYPE, tlv.Bytes(pfcp.MarshalUint8(t.Cause))))
	if t.OffendingIE != nil {
		fields = append(fields, tlv.TLVFrom(ie.OFFENDING_IE_TYPE, tlv.Bytes(pfcp.MarshalUint16(*t.OffendingIE))))
	}
	return tlv.TLV(ie.MSG_PFCP_NODE_REPORT_RESPONSE_TYPE, fields...)
}

func (t *PFCPNodeReportResponse) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPNodeReportResponse{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [3]bool
	for _, de := range elements {
		switch de.Type {
		case ie.NODE_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			if err = de.UnmarshalValue(&t.NodeID); err == nil {
			}
		case ie.CAUSE_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			err = pfcp.UnmarshalUint8(de.Value, &t.Cause)
		case ie.OFFENDING_IE_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			var v uint16
			if err = pfcp.UnmarshalUint16(de.Value, &v); err == nil {
				t.OffendingIE = &v
			}
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return pfcp.Faulty(de.Type, err)
		}
	}
	if !seen[1] {
		return &pfcp.MissingIE{Type: ie.CAUSE_TYPE}
	}
	if t.Cause < ie.CauseRequestRejected {
		if !seen[0] {
			return &pfcp.MissingIE{Type: ie.NODE_ID_TYPE}
		}
	}
	return
}

func (t *PFCPNodeReportResponse) UnmarshalTLV(typ uint32, buf []byte) (err error) {
	if typ != ie.MSG_PFCP_NODE_REPORT_RESPONSE_TYPE {
		err = tlv.ErrCritical
		return
	}

	err = t.UnmarshalBinary(buf)
	return
}

func (t *PFCPNodeReportResponse) MarshalJSON() (buf []byte, err error) {
	buf, err = jsonMarshal(t)
	return
}

func (t *PFCPNodeReportResponse) UnmarshalJSON(buf []byte) (err error) {
	*t = PFCPNodeReportResponse{}
	err = jsonUnmarshal(buf, t)
	return
}
