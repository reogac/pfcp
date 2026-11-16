package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'QoS Information in GTP-U Path QoS Report' data structure */
type QoSInformationinGTPUPathQoSReport struct {
	AveragePacketDelay []byte   //Average Packet Delay
	MinimumPacketDelay []byte   //Minimum Packet Delay
	MaximumPacketDelay []byte   //Maximum Packet Delay
	DSCP               []byte   //Transport Level Marking
	Unknown            []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid            []uint32 `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
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
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [4]bool
	for _, de := range elements {
		opt := false
		switch de.Type {
		case AVERAGE_PACKET_DELAY_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.AveragePacketDelay = de.Value
		case MINIMUM_PACKET_DELAY_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			t.MinimumPacketDelay = de.Value
		case MAXIMUM_PACKET_DELAY_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			t.MaximumPacketDelay = de.Value
		case TRANSPORT_LEVEL_MARKING_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			t.DSCP = de.Value
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
		return &pfcp.MissingIE{Type: AVERAGE_PACKET_DELAY_TYPE}
	}
	return
}
