package ie

import (
	"github.com/reogac/pfcp"

	"github.com/reogac/pfcp/tlv"
)

/* 'Session Report' data structure */
type SessionReport struct {
	SRRID                    []byte                    //SRR ID
	AccessAvailabilityReport *AccessAvailabilityReport //Access Availability Report
	QoSMonitoringReport      []QoSMonitoringReport     //QoS Monitoring Report
	Unknown                  []uint32                  `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *SessionReport) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(SRR_ID_TYPE, tlv.Bytes(t.SRRID)))
	if t.AccessAvailabilityReport != nil {
		fields = append(fields, t.AccessAvailabilityReport.Field())
	}
	for _, ie := range t.QoSMonitoringReport {
		fields = append(fields, ie.Field())
	}
	return tlv.TLV(SESSION_REPORT_TYPE, fields...)
}

func (t *SessionReport) UnmarshalBinary(buf []byte) (err error) {
	*t = SessionReport{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	var seen [2]bool
	for _, de := range elements {
		switch de.Type {
		case SRR_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.SRRID = de.Value
		case ACCESS_AVAILABILITY_REPORT_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			var v AccessAvailabilityReport
			if err = de.UnmarshalValue(&v); err == nil {
				t.AccessAvailabilityReport = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
			}
		case QOS_MONITORING_REPORT_TYPE:
			var v QoSMonitoringReport
			if err = de.UnmarshalValue(&v); err == nil {
				t.QoSMonitoringReport = append(t.QoSMonitoringReport, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
			}
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: SRR_ID_TYPE}
	}
	return
}
