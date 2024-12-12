package ie

import (
	"etrib5gc/pfcp"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'QoS Monitoring Report' data structure */
type QoSMonitoringReport struct {
	QFI                      uint8   //QFI
	QoSMonitoringMeasurement []byte  //QoS Monitoring Measurement
	TimeStamp                []byte  //Time Stamp
	StartTime                *uint32 //Start Time
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
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case QFI_TYPE:
			err = pfcp.UnmarshalUint8(de.Value, &t.QFI)
		case QOS_MONITORING_MEASUREMENT_TYPE:
			t.QoSMonitoringMeasurement = de.Value
		case TIME_STAMP_TYPE:
			t.TimeStamp = de.Value
		case START_TIME_TYPE:
			var v uint32
			if err = pfcp.UnmarshalUint32(de.Value, &v); err == nil {
				t.StartTime = &v
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
