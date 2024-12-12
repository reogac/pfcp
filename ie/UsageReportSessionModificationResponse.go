package ie

import (
	"etrib5gc/pfcp"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Usage Report Session Modification Response' data structure */
type UsageReportSessionModificationResponse struct {
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
	QueryURRReference          []byte                      //Query URR Reference
	EthernetTrafficInformation *EthernetTrafficInformation //Ethernet Traffic Information
}

func (t *UsageReportSessionModificationResponse) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(URR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(t.URRID))))
	fields = append(fields, tlv.TLVFrom(UR_SEQN_TYPE, tlv.Bytes(pfcp.MarshalUint32(t.URSEQN))))
	fields = append(fields, tlv.TLVFrom(USAGE_REPORT_TRIGGER_TYPE, tlv.Bytes(pfcp.MarshalUint32(t.UsageReportTrigger))))
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
	if len(t.QueryURRReference) > 0 {
		fields = append(fields, tlv.TLVFrom(QUERY_URR_REFERENCE_TYPE, tlv.Bytes(t.QueryURRReference)))
	}
	if t.EthernetTrafficInformation != nil {
		fields = append(fields, t.EthernetTrafficInformation.Field())
	}
	return tlv.TLV(USAGE_REPORT_SESSION_MODIFICATION_RESPONSE_TYPE, fields...)
}

func (t *UsageReportSessionModificationResponse) UnmarshalBinary(buf []byte) (err error) {
	*t = UsageReportSessionModificationResponse{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case URR_ID_TYPE:
			err = pfcp.UnmarshalUint32(de.Value, &t.URRID)
		case UR_SEQN_TYPE:
			err = pfcp.UnmarshalUint32(de.Value, &t.URSEQN)
		case USAGE_REPORT_TRIGGER_TYPE:
			err = pfcp.UnmarshalUint32(de.Value, &t.UsageReportTrigger)
		case START_TIME_TYPE:
			var v uint32
			if err = pfcp.UnmarshalUint32(de.Value, &v); err == nil {
				t.StartTime = &v
			}
		case END_TIME_TYPE:
			var v uint32
			if err = pfcp.UnmarshalUint32(de.Value, &v); err == nil {
				t.EndTime = &v
			}
		case VOLUME_MEASUREMENT_TYPE:
			t.VolumeMeasurement = de.Value
		case DURATION_MEASUREMENT_TYPE:
			var v uint32
			if err = pfcp.UnmarshalUint32(de.Value, &v); err == nil {
				t.DurationMeasurement = &v
			}
		case TIME_OF_FIRST_PACKET_TYPE:
			var v uint32
			if err = pfcp.UnmarshalUint32(de.Value, &v); err == nil {
				t.TimeofFirstPacket = &v
			}
		case TIME_OF_LAST_PACKET_TYPE:
			var v uint32
			if err = pfcp.UnmarshalUint32(de.Value, &v); err == nil {
				t.TimeofLastPacket = &v
			}
		case USAGE_INFORMATION_TYPE:
			t.UsageInformation = de.Value
		case QUERY_URR_REFERENCE_TYPE:
			t.QueryURRReference = de.Value
		case ETHERNET_TRAFFIC_INFORMATION_TYPE:
			var v EthernetTrafficInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.EthernetTrafficInformation = &v
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
