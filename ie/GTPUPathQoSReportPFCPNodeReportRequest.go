package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'GTP-U Path QoS Report PFCP Node Report Request' data structure */
type GTPUPathQoSReportPFCPNodeReportRequest struct {
	RemoteGTPUPeer              [][]byte //Remote GTP-U Peer
	GTPUPathInterfaceType       []byte   //GTP-U Path Interface Type
	QoSReportTrigger            []byte   //QoS Report Trigger
	DSCP                        [][]byte //Transport Level Marking
	MeasurementPeriod           *uint32  //Measurement Period
	AveragePacketDelayThreshold []byte   //Average Packet Delay
	MinimumPacketDelayThreshold []byte   //Minimum Packet Delay
	MaximumPacketDelayThreshold []byte   //Maximum Packet Delay
	MinimumWaitingTime          []byte   //Timer
	Unknown                     []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *GTPUPathQoSReportPFCPNodeReportRequest) Field() tlv.Field {
	fields := []tlv.Field{}
	for _, ie := range t.RemoteGTPUPeer {
		fields = append(fields, tlv.TLVFrom(REMOTE_GTP_U_PEER_TYPE, tlv.Bytes(ie)))
	}
	if len(t.GTPUPathInterfaceType) > 0 {
		fields = append(fields, tlv.TLVFrom(GTP_U_PATH_INTERFACE_TYPE_TYPE, tlv.Bytes(t.GTPUPathInterfaceType)))
	}
	fields = append(fields, tlv.TLVFrom(QOS_REPORT_TRIGGER_TYPE, tlv.Bytes(t.QoSReportTrigger)))
	for _, ie := range t.DSCP {
		fields = append(fields, tlv.TLVFrom(TRANSPORT_LEVEL_MARKING_TYPE, tlv.Bytes(ie)))
	}
	if t.MeasurementPeriod != nil {
		fields = append(fields, tlv.TLVFrom(MEASUREMENT_PERIOD_TYPE, tlv.Bytes(pfcp.MarshalUint32(*t.MeasurementPeriod))))
	}
	if len(t.AveragePacketDelayThreshold) > 0 {
		fields = append(fields, tlv.TLVFrom(AVERAGE_PACKET_DELAY_TYPE, tlv.Bytes(t.AveragePacketDelayThreshold)))
	}
	if len(t.MinimumPacketDelayThreshold) > 0 {
		fields = append(fields, tlv.TLVFrom(MINIMUM_PACKET_DELAY_TYPE, tlv.Bytes(t.MinimumPacketDelayThreshold)))
	}
	if len(t.MaximumPacketDelayThreshold) > 0 {
		fields = append(fields, tlv.TLVFrom(MAXIMUM_PACKET_DELAY_TYPE, tlv.Bytes(t.MaximumPacketDelayThreshold)))
	}
	if len(t.MinimumWaitingTime) > 0 {
		fields = append(fields, tlv.TLVFrom(TIMER_TYPE, tlv.Bytes(t.MinimumWaitingTime)))
	}
	return tlv.TLV(GTP_U_PATH_QOS_REPORT_PFCP_NODE_REPORT_REQUEST_TYPE, fields...)
}

func (t *GTPUPathQoSReportPFCPNodeReportRequest) UnmarshalBinary(buf []byte) (err error) {
	*t = GTPUPathQoSReportPFCPNodeReportRequest{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [7]bool
	for _, de := range elements {
		switch de.Type {
		case REMOTE_GTP_U_PEER_TYPE:
			t.RemoteGTPUPeer = append(t.RemoteGTPUPeer, de.Value)
		case GTP_U_PATH_INTERFACE_TYPE_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.GTPUPathInterfaceType = de.Value
		case QOS_REPORT_TRIGGER_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			t.QoSReportTrigger = de.Value
		case TRANSPORT_LEVEL_MARKING_TYPE:
			t.DSCP = append(t.DSCP, de.Value)
		case MEASUREMENT_PERIOD_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			var v uint32
			if err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &v); err == nil {
				t.MeasurementPeriod = &v
			}
		case AVERAGE_PACKET_DELAY_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			t.AveragePacketDelayThreshold = de.Value
		case MINIMUM_PACKET_DELAY_TYPE:
			if seen[4] {
				continue
			}
			seen[4] = true
			t.MinimumPacketDelayThreshold = de.Value
		case MAXIMUM_PACKET_DELAY_TYPE:
			if seen[5] {
				continue
			}
			seen[5] = true
			t.MaximumPacketDelayThreshold = de.Value
		case TIMER_TYPE:
			if seen[6] {
				continue
			}
			seen[6] = true
			t.MinimumWaitingTime = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return pfcp.Faulty(de.Type, err)
		}
	}
	if !seen[1] {
		return &pfcp.MissingIE{Type: QOS_REPORT_TRIGGER_TYPE}
	}
	return
}
