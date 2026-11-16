package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Create SRR' data structure */
type CreateSRR struct {
	SRRID                                     []byte   //SRR ID
	AccessAvailabilityControlInformation      []byte   //Access Availability Control Information
	QoSMonitoringperQoSflowControlInformation [][]byte //QoS Monitoring per QoS flow Control Information
	DirectReportingInformation                []byte   //Direct Reporting Information
	Unknown                                   []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid                                   []uint32 `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
}

func (t *CreateSRR) Field() tlv.Field {
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
	return tlv.TLV(CREATE_SRR_TYPE, fields...)
}

func (t *CreateSRR) UnmarshalBinary(buf []byte) (err error) {
	*t = CreateSRR{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [3]bool
	for _, de := range elements {
		opt := false
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
			if err = pfcp.Faulty(de.Type, err); opt && pfcp.Discardable(err) {
				t.Invalid = append(t.Invalid, de.Type)
				err = nil
				continue
			}
			return err
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: SRR_ID_TYPE}
	}
	return
}
