package ie

import (
	"etrib5gc/pfcp"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'GTP-U Path QoS Report PFCP Node Report Request' data structure */
type GTPUPathQoSReportPFCPNodeReportRequest struct {
	RemoteGTPUPeer              []byte  //Remote GTP-U Peer
	GTPUPathInterfaceType       []byte  //GTP-U Path Interface Type
	QoSReportTrigger            []byte  //QoS Report Trigger
	DSCP                        []byte  //Transport Level Marking
	MeasurementPeriod           *uint32 //Measurement Period
	AveragePacketDelayThreshold []byte  //Average Packet Delay
	MinimumPacketDelayThreshold []byte  //Minimum Packet Delay
	MaximumPacketDelayThreshold []byte  //Maximum Packet Delay
	MinimumWaitingTime          []byte  //Timer
}

func (t *GTPUPathQoSReportPFCPNodeReportRequest) Field() tlv.Field {
	fields := []tlv.Field{}
	if len(t.RemoteGTPUPeer) > 0 {
		fields = append(fields, tlv.TLVFrom(REMOTE_GTP_U_PEER_TYPE, tlv.Bytes(t.RemoteGTPUPeer)))
	}
	if len(t.GTPUPathInterfaceType) > 0 {
		fields = append(fields, tlv.TLVFrom(GTP_U_PATH_INTERFACE_TYPE_TYPE, tlv.Bytes(t.GTPUPathInterfaceType)))
	}
	fields = append(fields, tlv.TLVFrom(QOS_REPORT_TRIGGER_TYPE, tlv.Bytes(t.QoSReportTrigger)))
	if len(t.DSCP) > 0 {
		fields = append(fields, tlv.TLVFrom(TRANSPORT_LEVEL_MARKING_TYPE, tlv.Bytes(t.DSCP)))
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
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case REMOTE_GTP_U_PEER_TYPE:
			t.RemoteGTPUPeer = de.Value
		case GTP_U_PATH_INTERFACE_TYPE_TYPE:
			t.GTPUPathInterfaceType = de.Value
		case QOS_REPORT_TRIGGER_TYPE:
			t.QoSReportTrigger = de.Value
		case TRANSPORT_LEVEL_MARKING_TYPE:
			t.DSCP = de.Value
		case MEASUREMENT_PERIOD_TYPE:
			var v uint32
			if err = pfcp.UnmarshalUint32(de.Value, &v); err == nil {
				t.MeasurementPeriod = &v
			}
		case AVERAGE_PACKET_DELAY_TYPE:
			t.AveragePacketDelayThreshold = de.Value
		case MINIMUM_PACKET_DELAY_TYPE:
			t.MinimumPacketDelayThreshold = de.Value
		case MAXIMUM_PACKET_DELAY_TYPE:
			t.MaximumPacketDelayThreshold = de.Value
		case TIMER_TYPE:
			t.MinimumWaitingTime = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
