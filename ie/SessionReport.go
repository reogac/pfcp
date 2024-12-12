package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Session Report' data structure */
type SessionReport struct {
	SRRID                    []byte                    //SRR ID
	AccessAvailabilityReport *AccessAvailabilityReport //Access Availability Report
	QoSMonitoringReport      *QoSMonitoringReport      //QoS Monitoring Report
}

func (t *SessionReport) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(SRR_ID_TYPE, tlv.Bytes(t.SRRID)))
	if t.AccessAvailabilityReport != nil {
		fields = append(fields, t.AccessAvailabilityReport.Field())
	}
	if t.QoSMonitoringReport != nil {
		fields = append(fields, t.QoSMonitoringReport.Field())
	}
	return tlv.TLV(SESSION_REPORT_TYPE, fields...)
}

func (t *SessionReport) UnmarshalBinary(buf []byte) (err error) {
	*t = SessionReport{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case SRR_ID_TYPE:
			t.SRRID = de.Value
		case ACCESS_AVAILABILITY_REPORT_TYPE:
			var v AccessAvailabilityReport
			if err = de.UnmarshalValue(&v); err == nil {
				t.AccessAvailabilityReport = &v
			}
		case QOS_MONITORING_REPORT_TYPE:
			var v QoSMonitoringReport
			if err = de.UnmarshalValue(&v); err == nil {
				t.QoSMonitoringReport = &v
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
