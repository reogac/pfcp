package ie

import (
	"etrib5gc/pfcp"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Usage Report Session Report Request' data structure */
type UsageReportSessionReportRequest struct {
	URRID                           uint32                                          //URR ID
	URSEQN                          uint32                                          //UR-SEQN
	UsageReportTrigger              uint32                                          //Usage Report Trigger
	StartTime                       *uint32                                         //Start Time
	EndTime                         *uint32                                         //End Time
	VolumeMeasurement               []byte                                          //Volume Measurement
	DurationMeasurement             *uint32                                         //Duration Measurement
	ApplicationDetectionInformation *ApplicationDetectionInformation                //Application Detection Information
	UEIPaddress                     *UEIPAddress                                    //UE IP Address
	NetworkInstance                 []byte                                          //Network Instance
	TimeofFirstPacket               *uint32                                         //Time of First Packet
	TimeofLastPacket                *uint32                                         //Time of Last Packet
	UsageInformation                []byte                                          //Usage Information
	QueryURRReference               []byte                                          //Query URR Reference
	EventTimeStamp                  []byte                                          //Time Stamp
	EthernetTrafficInformation      *EthernetTrafficInformation                     //Ethernet Traffic Information
	JoinIPMuticastInformation       *JoinIPMulticastInformationIEwithinUsageReport  //Join IP Multicast Information IE within Usage Report
	LeaveIPMuticastInformation      *LeaveIPMulticastInformationIEwithinUsageReport //Leave IP Multicast Information IE within Usage Report
	PredefinedRulesName             []byte                                          //Predefined Rules Name
}

func (t *UsageReportSessionReportRequest) Field() tlv.Field {
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
	if t.ApplicationDetectionInformation != nil {
		fields = append(fields, t.ApplicationDetectionInformation.Field())
	}
	if t.UEIPaddress != nil {
		fields = append(fields, tlv.TLVFrom(UE_IP_ADDRESS_TYPE, tlv.Bytes(t.UEIPaddress.Bytes())))
	}
	if len(t.NetworkInstance) > 0 {
		fields = append(fields, tlv.TLVFrom(NETWORK_INSTANCE_TYPE, tlv.Bytes(t.NetworkInstance)))
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
	if len(t.EventTimeStamp) > 0 {
		fields = append(fields, tlv.TLVFrom(TIME_STAMP_TYPE, tlv.Bytes(t.EventTimeStamp)))
	}
	if t.EthernetTrafficInformation != nil {
		fields = append(fields, t.EthernetTrafficInformation.Field())
	}
	if t.JoinIPMuticastInformation != nil {
		fields = append(fields, t.JoinIPMuticastInformation.Field())
	}
	if t.LeaveIPMuticastInformation != nil {
		fields = append(fields, t.LeaveIPMuticastInformation.Field())
	}
	if len(t.PredefinedRulesName) > 0 {
		fields = append(fields, tlv.TLVFrom(PREDEFINED_RULES_NAME_TYPE, tlv.Bytes(t.PredefinedRulesName)))
	}
	return tlv.TLV(USAGE_REPORT_SESSION_REPORT_REQUEST_TYPE, fields...)
}

func (t *UsageReportSessionReportRequest) UnmarshalBinary(buf []byte) (err error) {
	*t = UsageReportSessionReportRequest{}
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
		case APPLICATION_DETECTION_INFORMATION_TYPE:
			var v ApplicationDetectionInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.ApplicationDetectionInformation = &v
			}
		case UE_IP_ADDRESS_TYPE:
			var v UEIPAddress
			if err = de.UnmarshalValue(&v); err == nil {
				t.UEIPaddress = &v
			}
		case NETWORK_INSTANCE_TYPE:
			t.NetworkInstance = de.Value
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
		case TIME_STAMP_TYPE:
			t.EventTimeStamp = de.Value
		case ETHERNET_TRAFFIC_INFORMATION_TYPE:
			var v EthernetTrafficInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.EthernetTrafficInformation = &v
			}
		case JOIN_IP_MULTICAST_INFORMATION_IE_WITHIN_USAGE_REPORT_TYPE:
			var v JoinIPMulticastInformationIEwithinUsageReport
			if err = de.UnmarshalValue(&v); err == nil {
				t.JoinIPMuticastInformation = &v
			}
		case LEAVE_IP_MULTICAST_INFORMATION_IE_WITHIN_USAGE_REPORT_TYPE:
			var v LeaveIPMulticastInformationIEwithinUsageReport
			if err = de.UnmarshalValue(&v); err == nil {
				t.LeaveIPMuticastInformation = &v
			}
		case PREDEFINED_RULES_NAME_TYPE:
			t.PredefinedRulesName = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
