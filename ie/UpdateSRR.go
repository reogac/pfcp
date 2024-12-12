package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Update SRR' data structure */
type UpdateSRR struct {
	SRRID                                     []byte //SRR ID
	AccessAvailabilityControlInformation      []byte //Access Availability Control Information
	QoSMonitoringperQoSflowControlInformation []byte //QoS Monitoring per QoS flow Control Information
	DirectReportingInformation                []byte //Direct Reporting Information
}

func (t *UpdateSRR) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(SRR_ID_TYPE, tlv.Bytes(t.SRRID)))
	if len(t.AccessAvailabilityControlInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(ACCESS_AVAILABILITY_CONTROL_INFORMATION_TYPE, tlv.Bytes(t.AccessAvailabilityControlInformation)))
	}
	if len(t.QoSMonitoringperQoSflowControlInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(QOS_MONITORING_PER_QOS_FLOW_CONTROL_INFORMATION_TYPE, tlv.Bytes(t.QoSMonitoringperQoSflowControlInformation)))
	}
	if len(t.DirectReportingInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(DIRECT_REPORTING_INFORMATION_TYPE, tlv.Bytes(t.DirectReportingInformation)))
	}
	return tlv.TLV(UPDATE_SRR_TYPE, fields...)
}

func (t *UpdateSRR) UnmarshalBinary(buf []byte) (err error) {
	*t = UpdateSRR{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case SRR_ID_TYPE:
			t.SRRID = de.Value
		case ACCESS_AVAILABILITY_CONTROL_INFORMATION_TYPE:
			t.AccessAvailabilityControlInformation = de.Value
		case QOS_MONITORING_PER_QOS_FLOW_CONTROL_INFORMATION_TYPE:
			t.QoSMonitoringperQoSflowControlInformation = de.Value
		case DIRECT_REPORTING_INFORMATION_TYPE:
			t.DirectReportingInformation = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
