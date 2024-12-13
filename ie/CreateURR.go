package ie

import (
	"github.com/reogac/pfcp"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Create URR' data structure */
type CreateURR struct {
	URRID                               uint32                  //URR ID
	MeasurementMethod                   MeasurementMethod       //Measurement Method
	ReportingTriggers                   uint32                  //Reporting Triggers
	MeasurementPeriod                   *uint32                 //Measurement Period
	VolumeThreshold                     *VolumeThreshold        //Volume Threshold
	VolumeQuota                         *VolumeQuota            //Volume Quota
	EventThreshold                      *uint32                 //Event Threshold
	EventQuota                          *uint32                 //Event Quota
	TimeThreshold                       *uint32                 //Time Threshold
	TimeQuota                           *uint32                 //Time Quota
	QuotaHoldingTime                    *uint32                 //Quota Holding Time
	DroppedDLTrafficThreshold           []byte                  //Dropped DL Traffic Threshold
	QuotaValidityTime                   *uint32                 //Quota Validity Time
	MonitoringTime                      []byte                  //Monitoring Time
	SubsequentVolumeThreshold           []byte                  //Subsequent Volume Threshold
	SubsequentTimeThreshold             []byte                  //Subsequent Time Threshold
	SubsequentVolumeQuota               []byte                  //Subsequent Volume Quota
	SubsequentTimeQuota                 []byte                  //Subsequent Time Quota
	SubsequentEventThreshold            []byte                  //Subsequent Event Threshold
	SubsequentEventQuota                []byte                  //Subsequent Event Quota
	InactivityDetectionTime             []byte                  //Inactivity Detection Time
	LinkedURRID                         []byte                  //Linked URR ID
	MeasurementInformation              *MeasurementInformation //Measurement Information
	TimeQuotaMechanism                  []byte                  //Time Quota Mechanism
	AggregatedURRs                      []byte                  //Aggregated URRs
	FARIDforQuotaAction                 *uint32                 //FAR ID
	EthernetInactivityTimer             []byte                  //Ethernet Inactivity Timer
	AdditionalMonitoringTime            []byte                  //Additional Monitoring Time
	NumberofReports                     []byte                  //Number of Reports
	ExemptedApplicationIDforQuotaAction *ApplicationID          //Application ID
	ExemptedSDFFilterforQuotaAction     []SDFFilter             //SDF Filter
	UserPlaneInactivityTimer            []byte                  //User Plane Inactivity Timer
}

func (t *CreateURR) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(URR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(t.URRID))))
	fields = append(fields, tlv.TLVFrom(MEASUREMENT_METHOD_TYPE, tlv.Bytes(t.MeasurementMethod.Bytes())))
	fields = append(fields, tlv.TLVFrom(REPORTING_TRIGGERS_TYPE, tlv.Bytes(pfcp.MarshalUint32(t.ReportingTriggers))))
	if t.MeasurementPeriod != nil {
		fields = append(fields, tlv.TLVFrom(MEASUREMENT_PERIOD_TYPE, tlv.Bytes(pfcp.MarshalUint32(*t.MeasurementPeriod))))
	}
	if t.VolumeThreshold != nil {
		fields = append(fields, tlv.TLVFrom(VOLUME_THRESHOLD_TYPE, tlv.Bytes(t.VolumeThreshold.Bytes())))
	}
	if t.VolumeQuota != nil {
		fields = append(fields, tlv.TLVFrom(VOLUME_QUOTA_TYPE, tlv.Bytes(t.VolumeQuota.Bytes())))
	}
	if t.EventThreshold != nil {
		fields = append(fields, tlv.TLVFrom(EVENT_THRESHOLD_TYPE, tlv.Bytes(pfcp.MarshalUint32(*t.EventThreshold))))
	}
	if t.EventQuota != nil {
		fields = append(fields, tlv.TLVFrom(EVENT_QUOTA_TYPE, tlv.Bytes(pfcp.MarshalUint32(*t.EventQuota))))
	}
	if t.TimeThreshold != nil {
		fields = append(fields, tlv.TLVFrom(TIME_THRESHOLD_TYPE, tlv.Bytes(pfcp.MarshalUint32(*t.TimeThreshold))))
	}
	if t.TimeQuota != nil {
		fields = append(fields, tlv.TLVFrom(TIME_QUOTA_TYPE, tlv.Bytes(pfcp.MarshalUint32(*t.TimeQuota))))
	}
	if t.QuotaHoldingTime != nil {
		fields = append(fields, tlv.TLVFrom(QUOTA_HOLDING_TIME_TYPE, tlv.Bytes(pfcp.MarshalUint32(*t.QuotaHoldingTime))))
	}
	if len(t.DroppedDLTrafficThreshold) > 0 {
		fields = append(fields, tlv.TLVFrom(DROPPED_DL_TRAFFIC_THRESHOLD_TYPE, tlv.Bytes(t.DroppedDLTrafficThreshold)))
	}
	if t.QuotaValidityTime != nil {
		fields = append(fields, tlv.TLVFrom(QUOTA_VALIDITY_TIME_TYPE, tlv.Bytes(pfcp.MarshalUint32(*t.QuotaValidityTime))))
	}
	if len(t.MonitoringTime) > 0 {
		fields = append(fields, tlv.TLVFrom(MONITORING_TIME_TYPE, tlv.Bytes(t.MonitoringTime)))
	}
	if len(t.SubsequentVolumeThreshold) > 0 {
		fields = append(fields, tlv.TLVFrom(SUBSEQUENT_VOLUME_THRESHOLD_TYPE, tlv.Bytes(t.SubsequentVolumeThreshold)))
	}
	if len(t.SubsequentTimeThreshold) > 0 {
		fields = append(fields, tlv.TLVFrom(SUBSEQUENT_TIME_THRESHOLD_TYPE, tlv.Bytes(t.SubsequentTimeThreshold)))
	}
	if len(t.SubsequentVolumeQuota) > 0 {
		fields = append(fields, tlv.TLVFrom(SUBSEQUENT_VOLUME_QUOTA_TYPE, tlv.Bytes(t.SubsequentVolumeQuota)))
	}
	if len(t.SubsequentTimeQuota) > 0 {
		fields = append(fields, tlv.TLVFrom(SUBSEQUENT_TIME_QUOTA_TYPE, tlv.Bytes(t.SubsequentTimeQuota)))
	}
	if len(t.SubsequentEventThreshold) > 0 {
		fields = append(fields, tlv.TLVFrom(SUBSEQUENT_EVENT_THRESHOLD_TYPE, tlv.Bytes(t.SubsequentEventThreshold)))
	}
	if len(t.SubsequentEventQuota) > 0 {
		fields = append(fields, tlv.TLVFrom(SUBSEQUENT_EVENT_QUOTA_TYPE, tlv.Bytes(t.SubsequentEventQuota)))
	}
	if len(t.InactivityDetectionTime) > 0 {
		fields = append(fields, tlv.TLVFrom(INACTIVITY_DETECTION_TIME_TYPE, tlv.Bytes(t.InactivityDetectionTime)))
	}
	if len(t.LinkedURRID) > 0 {
		fields = append(fields, tlv.TLVFrom(LINKED_URR_ID_TYPE, tlv.Bytes(t.LinkedURRID)))
	}
	if t.MeasurementInformation != nil {
		fields = append(fields, tlv.TLVFrom(MEASUREMENT_INFORMATION_TYPE, tlv.Bytes(t.MeasurementInformation.Bytes())))
	}
	if len(t.TimeQuotaMechanism) > 0 {
		fields = append(fields, tlv.TLVFrom(TIME_QUOTA_MECHANISM_TYPE, tlv.Bytes(t.TimeQuotaMechanism)))
	}
	if len(t.AggregatedURRs) > 0 {
		fields = append(fields, tlv.TLVFrom(AGGREGATED_URRS_TYPE, tlv.Bytes(t.AggregatedURRs)))
	}
	if t.FARIDforQuotaAction != nil {
		fields = append(fields, tlv.TLVFrom(FAR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(*t.FARIDforQuotaAction))))
	}
	if len(t.EthernetInactivityTimer) > 0 {
		fields = append(fields, tlv.TLVFrom(ETHERNET_INACTIVITY_TIMER_TYPE, tlv.Bytes(t.EthernetInactivityTimer)))
	}
	if len(t.AdditionalMonitoringTime) > 0 {
		fields = append(fields, tlv.TLVFrom(ADDITIONAL_MONITORING_TIME_TYPE, tlv.Bytes(t.AdditionalMonitoringTime)))
	}
	if len(t.NumberofReports) > 0 {
		fields = append(fields, tlv.TLVFrom(NUMBER_OF_REPORTS_TYPE, tlv.Bytes(t.NumberofReports)))
	}
	if t.ExemptedApplicationIDforQuotaAction != nil {
		fields = append(fields, tlv.TLVFrom(APPLICATION_ID_TYPE, tlv.Bytes(t.ExemptedApplicationIDforQuotaAction.Bytes())))
	}
	for _, ie := range t.ExemptedSDFFilterforQuotaAction {
		fields = append(fields, tlv.TLVFrom(SDF_FILTER_TYPE, tlv.Bytes(ie.Bytes())))
	}
	if len(t.UserPlaneInactivityTimer) > 0 {
		fields = append(fields, tlv.TLVFrom(USER_PLANE_INACTIVITY_TIMER_TYPE, tlv.Bytes(t.UserPlaneInactivityTimer)))
	}
	return tlv.TLV(CREATE_URR_TYPE, fields...)
}

func (t *CreateURR) UnmarshalBinary(buf []byte) (err error) {
	*t = CreateURR{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case URR_ID_TYPE:
			err = pfcp.UnmarshalUint32(de.Value, &t.URRID)
		case MEASUREMENT_METHOD_TYPE:
			err = de.UnmarshalValue(&t.MeasurementMethod)
		case REPORTING_TRIGGERS_TYPE:
			err = pfcp.UnmarshalUint32(de.Value, &t.ReportingTriggers)
		case MEASUREMENT_PERIOD_TYPE:
			var v uint32
			if err = pfcp.UnmarshalUint32(de.Value, &v); err == nil {
				t.MeasurementPeriod = &v
			}
		case VOLUME_THRESHOLD_TYPE:
			var v VolumeThreshold
			if err = de.UnmarshalValue(&v); err == nil {
				t.VolumeThreshold = &v
			}
		case VOLUME_QUOTA_TYPE:
			var v VolumeQuota
			if err = de.UnmarshalValue(&v); err == nil {
				t.VolumeQuota = &v
			}
		case EVENT_THRESHOLD_TYPE:
			var v uint32
			if err = pfcp.UnmarshalUint32(de.Value, &v); err == nil {
				t.EventThreshold = &v
			}
		case EVENT_QUOTA_TYPE:
			var v uint32
			if err = pfcp.UnmarshalUint32(de.Value, &v); err == nil {
				t.EventQuota = &v
			}
		case TIME_THRESHOLD_TYPE:
			var v uint32
			if err = pfcp.UnmarshalUint32(de.Value, &v); err == nil {
				t.TimeThreshold = &v
			}
		case TIME_QUOTA_TYPE:
			var v uint32
			if err = pfcp.UnmarshalUint32(de.Value, &v); err == nil {
				t.TimeQuota = &v
			}
		case QUOTA_HOLDING_TIME_TYPE:
			var v uint32
			if err = pfcp.UnmarshalUint32(de.Value, &v); err == nil {
				t.QuotaHoldingTime = &v
			}
		case DROPPED_DL_TRAFFIC_THRESHOLD_TYPE:
			t.DroppedDLTrafficThreshold = de.Value
		case QUOTA_VALIDITY_TIME_TYPE:
			var v uint32
			if err = pfcp.UnmarshalUint32(de.Value, &v); err == nil {
				t.QuotaValidityTime = &v
			}
		case MONITORING_TIME_TYPE:
			t.MonitoringTime = de.Value
		case SUBSEQUENT_VOLUME_THRESHOLD_TYPE:
			t.SubsequentVolumeThreshold = de.Value
		case SUBSEQUENT_TIME_THRESHOLD_TYPE:
			t.SubsequentTimeThreshold = de.Value
		case SUBSEQUENT_VOLUME_QUOTA_TYPE:
			t.SubsequentVolumeQuota = de.Value
		case SUBSEQUENT_TIME_QUOTA_TYPE:
			t.SubsequentTimeQuota = de.Value
		case SUBSEQUENT_EVENT_THRESHOLD_TYPE:
			t.SubsequentEventThreshold = de.Value
		case SUBSEQUENT_EVENT_QUOTA_TYPE:
			t.SubsequentEventQuota = de.Value
		case INACTIVITY_DETECTION_TIME_TYPE:
			t.InactivityDetectionTime = de.Value
		case LINKED_URR_ID_TYPE:
			t.LinkedURRID = de.Value
		case MEASUREMENT_INFORMATION_TYPE:
			var v MeasurementInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.MeasurementInformation = &v
			}
		case TIME_QUOTA_MECHANISM_TYPE:
			t.TimeQuotaMechanism = de.Value
		case AGGREGATED_URRS_TYPE:
			t.AggregatedURRs = de.Value
		case FAR_ID_TYPE:
			var v uint32
			if err = pfcp.UnmarshalUint32(de.Value, &v); err == nil {
				t.FARIDforQuotaAction = &v
			}
		case ETHERNET_INACTIVITY_TIMER_TYPE:
			t.EthernetInactivityTimer = de.Value
		case ADDITIONAL_MONITORING_TIME_TYPE:
			t.AdditionalMonitoringTime = de.Value
		case NUMBER_OF_REPORTS_TYPE:
			t.NumberofReports = de.Value
		case APPLICATION_ID_TYPE:
			var v ApplicationID
			if err = de.UnmarshalValue(&v); err == nil {
				t.ExemptedApplicationIDforQuotaAction = &v
			}
		case SDF_FILTER_TYPE:
			var v SDFFilter
			if err = de.UnmarshalValue(&v); err == nil {
				t.ExemptedSDFFilterforQuotaAction = append(t.ExemptedSDFFilterforQuotaAction, v)
			}
		case USER_PLANE_INACTIVITY_TIMER_TYPE:
			t.UserPlaneInactivityTimer = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
