package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Update SRR' data structure */
type UpdateSRR struct {
	SRRID                                     []byte   //SRR ID
	AccessAvailabilityControlInformation      []byte   //Access Availability Control Information
	QoSMonitoringperQoSflowControlInformation [][]byte //QoS Monitoring per QoS flow Control Information
	DirectReportingInformation                []byte   //Direct Reporting Information
	Unknown                                   []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *UpdateSRR) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(SRR_ID_TYPE, tlv.Bytes(t.SRRID)))
	if len(t.AccessAvailabilityControlInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(ACCESS_AVAILABILITY_CONTROL_INFORMATION_TYPE, tlv.Bytes(t.AccessAvailabilityControlInformation)))
	}
	for _, ie := range t.QoSMonitoringperQoSflowControlInformation {
		fields = append(fields, tlv.TLVFrom(QOS_MONITORING_PER_QOS_FLOW_CONTROL_INFORMATION_TYPE, tlv.Bytes(ie)))
	}
	if len(t.DirectReportingInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(DIRECT_REPORTING_INFORMATION_TYPE, tlv.Bytes(t.DirectReportingInformation)))
	}
	return tlv.TLV(UPDATE_SRR_TYPE, fields...)
}

func (t *UpdateSRR) UnmarshalBinary(buf []byte) (err error) {
	*t = UpdateSRR{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [3]bool
	for _, de := range elements {
		switch de.Type {
		case SRR_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.SRRID = de.Value
		case ACCESS_AVAILABILITY_CONTROL_INFORMATION_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			t.AccessAvailabilityControlInformation = de.Value
		case QOS_MONITORING_PER_QOS_FLOW_CONTROL_INFORMATION_TYPE:
			t.QoSMonitoringperQoSflowControlInformation = append(t.QoSMonitoringperQoSflowControlInformation, de.Value)
		case DIRECT_REPORTING_INFORMATION_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			t.DirectReportingInformation = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return pfcp.Faulty(de.Type, err)
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: SRR_ID_TYPE}
	}
	return
}
