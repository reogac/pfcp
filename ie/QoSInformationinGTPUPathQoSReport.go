package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'QoS Information in GTP-U Path QoS Report' data structure */
type QoSInformationinGTPUPathQoSReport struct {
	AveragePacketDelay []byte //Average Packet Delay
	MinimumPacketDelay []byte //Minimum Packet Delay
	MaximumPacketDelay []byte //Maximum Packet Delay
	DSCP               []byte //Transport Level Marking
}

func (t *QoSInformationinGTPUPathQoSReport) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(AVERAGE_PACKET_DELAY_TYPE, tlv.Bytes(t.AveragePacketDelay)))
	if len(t.MinimumPacketDelay) > 0 {
		fields = append(fields, tlv.TLVFrom(MINIMUM_PACKET_DELAY_TYPE, tlv.Bytes(t.MinimumPacketDelay)))
	}
	if len(t.MaximumPacketDelay) > 0 {
		fields = append(fields, tlv.TLVFrom(MAXIMUM_PACKET_DELAY_TYPE, tlv.Bytes(t.MaximumPacketDelay)))
	}
	if len(t.DSCP) > 0 {
		fields = append(fields, tlv.TLVFrom(TRANSPORT_LEVEL_MARKING_TYPE, tlv.Bytes(t.DSCP)))
	}
	return tlv.TLV(QOS_INFORMATION_IN_GTP_U_PATH_QOS_REPORT_TYPE, fields...)
}

func (t *QoSInformationinGTPUPathQoSReport) UnmarshalBinary(buf []byte) (err error) {
	*t = QoSInformationinGTPUPathQoSReport{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case AVERAGE_PACKET_DELAY_TYPE:
			t.AveragePacketDelay = de.Value
		case MINIMUM_PACKET_DELAY_TYPE:
			t.MinimumPacketDelay = de.Value
		case MAXIMUM_PACKET_DELAY_TYPE:
			t.MaximumPacketDelay = de.Value
		case TRANSPORT_LEVEL_MARKING_TYPE:
			t.DSCP = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
