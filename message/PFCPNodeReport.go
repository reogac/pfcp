package message

import (
	"github.com/reogac/pfcp"
	"github.com/reogac/pfcp/ie"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'PFCP Node Report Request' data structure */
type PFCPNodeReportRequest struct {
	NodeID                       ie.NodeID                                  //Node ID
	NodeReportType               []byte                                     //Node Report Type
	VendorSpecificNodeReportType []byte                                     //Vendor-Specific Node Report Type
	UserPlanePathFailureReport   *ie.UserPlanePathFailureReport             //User Plane Path Failure Report
	UserPlanePathRecoveryReport  *ie.UserPlanePathRecoveryReport            //User Plane Path Recovery Report
	ClockDriftReport             *ie.ClockDriftReport                       //Clock Drift Report
	GTPUPathQoSReport            *ie.GTPUPathQoSReportPFCPNodeReportRequest //GTP-U Path QoS Report PFCP Node Report Request
	PeerUPRestartReport          []byte                                     //Peer UP Restart Report
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
	if t.ClockDriftReport != nil {
		fields = append(fields, t.ClockDriftReport.Field())
	}
	if t.GTPUPathQoSReport != nil {
		fields = append(fields, t.GTPUPathQoSReport.Field())
	}
	if len(t.PeerUPRestartReport) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.PEER_UP_RESTART_REPORT_TYPE, tlv.Bytes(t.PeerUPRestartReport)))
	}
	return tlv.TLV(ie.MSG_PFCP_NODE_REPORT_REQUEST_TYPE, fields...)
}

func (t *PFCPNodeReportRequest) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPNodeReportRequest{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case ie.NODE_ID_TYPE:
			err = de.UnmarshalValue(&t.NodeID)
		case ie.NODE_REPORT_TYPE_TYPE:
			t.NodeReportType = de.Value
		case ie.VENDOR_SPECIFIC_NODE_REPORT_TYPE_TYPE:
			t.VendorSpecificNodeReportType = de.Value
		case ie.USER_PLANE_PATH_FAILURE_REPORT_TYPE:
			var v ie.UserPlanePathFailureReport
			if err = de.UnmarshalValue(&v); err == nil {
				t.UserPlanePathFailureReport = &v
			}
		case ie.USER_PLANE_PATH_RECOVERY_REPORT_TYPE:
			var v ie.UserPlanePathRecoveryReport
			if err = de.UnmarshalValue(&v); err == nil {
				t.UserPlanePathRecoveryReport = &v
			}
		case ie.CLOCK_DRIFT_REPORT_TYPE:
			var v ie.ClockDriftReport
			if err = de.UnmarshalValue(&v); err == nil {
				t.ClockDriftReport = &v
			}
		case ie.GTP_U_PATH_QOS_REPORT_PFCP_NODE_REPORT_REQUEST_TYPE:
			var v ie.GTPUPathQoSReportPFCPNodeReportRequest
			if err = de.UnmarshalValue(&v); err == nil {
				t.GTPUPathQoSReport = &v
			}
		case ie.PEER_UP_RESTART_REPORT_TYPE:
			t.PeerUPRestartReport = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
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
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
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
