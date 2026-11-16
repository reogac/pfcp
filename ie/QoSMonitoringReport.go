package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'QoS Monitoring Report' data structure */
type QoSMonitoringReport struct {
	QFI                      uint8    //QFI
	QoSMonitoringMeasurement []byte   //QoS Monitoring Measurement
	TimeStamp                []byte   //Time Stamp
	StartTime                *uint32  //Start Time
	Unknown                  []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *QoSMonitoringReport) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(QFI_TYPE, tlv.Bytes(pfcp.MarshalUint8(t.QFI))))
	fields = append(fields, tlv.TLVFrom(QOS_MONITORING_MEASUREMENT_TYPE, tlv.Bytes(t.QoSMonitoringMeasurement)))
	fields = append(fields, tlv.TLVFrom(TIME_STAMP_TYPE, tlv.Bytes(t.TimeStamp)))
	if t.StartTime != nil {
		fields = append(fields, tlv.TLVFrom(START_TIME_TYPE, tlv.Bytes(pfcp.MarshalUint32(*t.StartTime))))
	}
	return tlv.TLV(QOS_MONITORING_REPORT_TYPE, fields...)
}

func (t *QoSMonitoringReport) UnmarshalBinary(buf []byte) (err error) {
	*t = QoSMonitoringReport{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	var seen [4]bool
	for _, de := range elements {
		switch de.Type {
		case QFI_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &t.QFI)
		case QOS_MONITORING_MEASUREMENT_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			t.QoSMonitoringMeasurement = de.Value
		case TIME_STAMP_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			t.TimeStamp = de.Value
		case START_TIME_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			var v uint32
			if err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &v); err == nil {
				t.StartTime = &v
			}
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: QFI_TYPE}
	}
	if !seen[1] {
		return &pfcp.MissingIE{Type: QOS_MONITORING_MEASUREMENT_TYPE}
	}
	if !seen[2] {
		return &pfcp.MissingIE{Type: TIME_STAMP_TYPE}
	}
	return
}
