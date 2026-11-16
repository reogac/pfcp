package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Usage Report Session Deletion Response' data structure */
type UsageReportSessionDeletionResponse struct {
	URRID                      uint32                      //URR ID
	URSEQN                     uint32                      //UR-SEQN
	UsageReportTrigger         uint32                      //Usage Report Trigger
	StartTime                  *uint32                     //Start Time
	EndTime                    *uint32                     //End Time
	VolumeMeasurement          []byte                      //Volume Measurement
	DurationMeasurement        *uint32                     //Duration Measurement
	TimeofFirstPacket          *uint32                     //Time of First Packet
	TimeofLastPacket           *uint32                     //Time of Last Packet
	UsageInformation           []byte                      //Usage Information
	EthernetTrafficInformation *EthernetTrafficInformation //Ethernet Traffic Information
	Unknown                    []uint32                    `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid                    []uint32                    `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
}

func (t *UsageReportSessionDeletionResponse) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(URR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(t.URRID))))
	fields = append(fields, tlv.TLVFrom(UR_SEQN_TYPE, tlv.Bytes(pfcp.MarshalUint32(t.URSEQN))))
	fields = append(fields, tlv.TLVFrom(USAGE_REPORT_TRIGGER_TYPE, tlv.Bytes(pfcp.Marshal3Bytes(t.UsageReportTrigger))))
	if t.StartTime != nil {
		fields = append(fields, tlv.TLVFrom(START_TIME_TYPE, tlv.Bytes(pfcp.MarshalUint32(*t.StartTime))))
	}
	if t.EndTime != nil {
		fields = append(fields, tlv.TLVFrom(END_TIME_TYPE, tlv.Bytes(pfcp.MarshalUint32(*t.EndTime))))
	}
	if len(t.VolumeMeasurement) > 0 {
		fields = append(fields, tlv.TLVFrom(VOLUME_MEASUREMENT_TYPE, tlv.Bytes(t.VolumeMeasurement)))
	}
	if t.DurationMeasurement != nil {
		fields = append(fields, tlv.TLVFrom(DURATION_MEASUREMENT_TYPE, tlv.Bytes(pfcp.MarshalUint32(*t.DurationMeasurement))))
	}
	if t.TimeofFirstPacket != nil {
		fields = append(fields, tlv.TLVFrom(TIME_OF_FIRST_PACKET_TYPE, tlv.Bytes(pfcp.MarshalUint32(*t.TimeofFirstPacket))))
	}
	if t.TimeofLastPacket != nil {
		fields = append(fields, tlv.TLVFrom(TIME_OF_LAST_PACKET_TYPE, tlv.Bytes(pfcp.MarshalUint32(*t.TimeofLastPacket))))
	}
	if len(t.UsageInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(USAGE_INFORMATION_TYPE, tlv.Bytes(t.UsageInformation)))
	}
	if t.EthernetTrafficInformation != nil {
		fields = append(fields, t.EthernetTrafficInformation.Field())
	}
	return tlv.TLV(USAGE_REPORT_SESSION_DELETION_RESPONSE_TYPE, fields...)
}

func (t *UsageReportSessionDeletionResponse) UnmarshalBinary(buf []byte) (err error) {
	*t = UsageReportSessionDeletionResponse{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [11]bool
	for _, de := range elements {
		opt := false
		switch de.Type {
		case URR_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &t.URRID)
		case UR_SEQN_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			err = pfcp.UnmarshalUint32(de.Value, &t.URSEQN)
		case USAGE_REPORT_TRIGGER_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			err = pfcp.UnmarshalExtendable(de.Value, 2, 3, &t.UsageReportTrigger)
		case START_TIME_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			var v uint32
			if err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &v); err == nil {
				t.StartTime = &v
			}
		case END_TIME_TYPE:
			if seen[4] {
				continue
			}
			seen[4] = true
			var v uint32
			if err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &v); err == nil {
				t.EndTime = &v
			}
		case VOLUME_MEASUREMENT_TYPE:
			if seen[5] {
				continue
			}
			seen[5] = true
			t.VolumeMeasurement = de.Value
		case DURATION_MEASUREMENT_TYPE:
			if seen[6] {
				continue
			}
			seen[6] = true
			var v uint32
			if err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &v); err == nil {
				t.DurationMeasurement = &v
			}
		case TIME_OF_FIRST_PACKET_TYPE:
			if seen[7] {
				continue
			}
			seen[7] = true
			var v uint32
			if err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &v); err == nil {
				t.TimeofFirstPacket = &v
			}
		case TIME_OF_LAST_PACKET_TYPE:
			if seen[8] {
				continue
			}
			seen[8] = true
			var v uint32
			if err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &v); err == nil {
				t.TimeofLastPacket = &v
			}
		case USAGE_INFORMATION_TYPE:
			if seen[9] {
				continue
			}
			seen[9] = true
			t.UsageInformation = de.Value
		case ETHERNET_TRAFFIC_INFORMATION_TYPE:
			if seen[10] {
				continue
			}
			seen[10] = true
			var v EthernetTrafficInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.EthernetTrafficInformation = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
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
		return &pfcp.MissingIE{Type: URR_ID_TYPE}
	}
	if !seen[1] {
		return &pfcp.MissingIE{Type: UR_SEQN_TYPE}
	}
	if !seen[2] {
		return &pfcp.MissingIE{Type: USAGE_REPORT_TRIGGER_TYPE}
	}
	return
}
