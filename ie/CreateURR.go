package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
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
	LinkedURRID                         [][]byte                //Linked URR ID
	MeasurementInformation              *MeasurementInformation //Measurement Information
	TimeQuotaMechanism                  []byte                  //Time Quota Mechanism
	AggregatedURRs                      [][]byte                //Aggregated URRs
	FARIDforQuotaAction                 *uint32                 //FAR ID
	EthernetInactivityTimer             []byte                  //Ethernet Inactivity Timer
	AdditionalMonitoringTime            [][]byte                //Additional Monitoring Time
	NumberofReports                     []byte                  //Number of Reports
	ExemptedApplicationIDforQuotaAction []ApplicationID         //Application ID
	ExemptedSDFFilterforQuotaAction     []SDFFilter             //SDF Filter
	UserPlaneInactivityTimer            []byte                  //User Plane Inactivity Timer
	Unknown                             []uint32                `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid                             []uint32                `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
}

func (t *CreateURR) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(URR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(t.URRID))))
	fields = append(fields, tlv.TLVFrom(MEASUREMENT_METHOD_TYPE, tlv.Bytes(t.MeasurementMethod.Bytes())))
	fields = append(fields, tlv.TLVFrom(REPORTING_TRIGGERS_TYPE, tlv.Bytes(pfcp.Marshal3Bytes(t.ReportingTriggers))))
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
	for _, ie := range t.LinkedURRID {
		fields = append(fields, tlv.TLVFrom(LINKED_URR_ID_TYPE, tlv.Bytes(ie)))
	}
	if t.MeasurementInformation != nil {
		fields = append(fields, tlv.TLVFrom(MEASUREMENT_INFORMATION_TYPE, tlv.Bytes(t.MeasurementInformation.Bytes())))
	}
	if len(t.TimeQuotaMechanism) > 0 {
		fields = append(fields, tlv.TLVFrom(TIME_QUOTA_MECHANISM_TYPE, tlv.Bytes(t.TimeQuotaMechanism)))
	}
	for _, ie := range t.AggregatedURRs {
		fields = append(fields, tlv.TLVFrom(AGGREGATED_URRS_TYPE, tlv.Bytes(ie)))
	}
	if t.FARIDforQuotaAction != nil {
		fields = append(fields, tlv.TLVFrom(FAR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(*t.FARIDforQuotaAction))))
	}
	if len(t.EthernetInactivityTimer) > 0 {
		fields = append(fields, tlv.TLVFrom(ETHERNET_INACTIVITY_TIMER_TYPE, tlv.Bytes(t.EthernetInactivityTimer)))
	}
	for _, ie := range t.AdditionalMonitoringTime {
		fields = append(fields, tlv.TLVFrom(ADDITIONAL_MONITORING_TIME_TYPE, tlv.Bytes(ie)))
	}
	if len(t.NumberofReports) > 0 {
		fields = append(fields, tlv.TLVFrom(NUMBER_OF_REPORTS_TYPE, tlv.Bytes(t.NumberofReports)))
	}
	for _, ie := range t.ExemptedApplicationIDforQuotaAction {
		fields = append(fields, tlv.TLVFrom(APPLICATION_ID_TYPE, tlv.Bytes(ie.Bytes())))
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
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [27]bool
	for _, de := range elements {
		opt := false
		switch de.Type {
		case URR_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &t.URRID)
		case MEASUREMENT_METHOD_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			if err = de.UnmarshalValue(&t.MeasurementMethod); err == nil {
			}
		case REPORTING_TRIGGERS_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			err = pfcp.UnmarshalExtendable(de.Value, 2, 3, &t.ReportingTriggers)
		case MEASUREMENT_PERIOD_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			var v uint32
			if err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &v); err == nil {
				t.MeasurementPeriod = &v
			}
		case VOLUME_THRESHOLD_TYPE:
			if seen[4] {
				continue
			}
			seen[4] = true
			var v VolumeThreshold
			if err = de.UnmarshalValue(&v); err == nil {
				t.VolumeThreshold = &v
			}
		case VOLUME_QUOTA_TYPE:
			if seen[5] {
				continue
			}
			seen[5] = true
			var v VolumeQuota
			if err = de.UnmarshalValue(&v); err == nil {
				t.VolumeQuota = &v
			}
		case EVENT_THRESHOLD_TYPE:
			if seen[6] {
				continue
			}
			seen[6] = true
			var v uint32
			if err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &v); err == nil {
				t.EventThreshold = &v
			}
		case EVENT_QUOTA_TYPE:
			if seen[7] {
				continue
			}
			seen[7] = true
			var v uint32
			if err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &v); err == nil {
				t.EventQuota = &v
			}
		case TIME_THRESHOLD_TYPE:
			if seen[8] {
				continue
			}
			seen[8] = true
			var v uint32
			if err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &v); err == nil {
				t.TimeThreshold = &v
			}
		case TIME_QUOTA_TYPE:
			if seen[9] {
				continue
			}
			seen[9] = true
			var v uint32
			if err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &v); err == nil {
				t.TimeQuota = &v
			}
		case QUOTA_HOLDING_TIME_TYPE:
			if seen[10] {
				continue
			}
			seen[10] = true
			var v uint32
			if err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &v); err == nil {
				t.QuotaHoldingTime = &v
			}
		case DROPPED_DL_TRAFFIC_THRESHOLD_TYPE:
			if seen[11] {
				continue
			}
			seen[11] = true
			t.DroppedDLTrafficThreshold = de.Value
		case QUOTA_VALIDITY_TIME_TYPE:
			if seen[12] {
				continue
			}
			seen[12] = true
			var v uint32
			if err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &v); err == nil {
				t.QuotaValidityTime = &v
			}
		case MONITORING_TIME_TYPE:
			opt = true
			if seen[13] {
				continue
			}
			seen[13] = true
			t.MonitoringTime = de.Value
		case SUBSEQUENT_VOLUME_THRESHOLD_TYPE:
			opt = true
			if seen[14] {
				continue
			}
			seen[14] = true
			t.SubsequentVolumeThreshold = de.Value
		case SUBSEQUENT_TIME_THRESHOLD_TYPE:
			opt = true
			if seen[15] {
				continue
			}
			seen[15] = true
			t.SubsequentTimeThreshold = de.Value
		case SUBSEQUENT_VOLUME_QUOTA_TYPE:
			opt = true
			if seen[16] {
				continue
			}
			seen[16] = true
			t.SubsequentVolumeQuota = de.Value
		case SUBSEQUENT_TIME_QUOTA_TYPE:
			opt = true
			if seen[17] {
				continue
			}
			seen[17] = true
			t.SubsequentTimeQuota = de.Value
		case SUBSEQUENT_EVENT_THRESHOLD_TYPE:
			opt = true
			if seen[18] {
				continue
			}
			seen[18] = true
			t.SubsequentEventThreshold = de.Value
		case SUBSEQUENT_EVENT_QUOTA_TYPE:
			opt = true
			if seen[19] {
				continue
			}
			seen[19] = true
			t.SubsequentEventQuota = de.Value
		case INACTIVITY_DETECTION_TIME_TYPE:
			if seen[20] {
				continue
			}
			seen[20] = true
			t.InactivityDetectionTime = de.Value
		case LINKED_URR_ID_TYPE:
			t.LinkedURRID = append(t.LinkedURRID, de.Value)
		case MEASUREMENT_INFORMATION_TYPE:
			if seen[21] {
				continue
			}
			seen[21] = true
			var v MeasurementInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.MeasurementInformation = &v
			}
		case TIME_QUOTA_MECHANISM_TYPE:
			if seen[22] {
				continue
			}
			seen[22] = true
			t.TimeQuotaMechanism = de.Value
		case AGGREGATED_URRS_TYPE:
			t.AggregatedURRs = append(t.AggregatedURRs, de.Value)
		case FAR_ID_TYPE:
			if seen[23] {
				continue
			}
			seen[23] = true
			var v uint32
			if err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &v); err == nil {
				t.FARIDforQuotaAction = &v
			}
		case ETHERNET_INACTIVITY_TIMER_TYPE:
			if seen[24] {
				continue
			}
			seen[24] = true
			t.EthernetInactivityTimer = de.Value
		case ADDITIONAL_MONITORING_TIME_TYPE:
			opt = true
			t.AdditionalMonitoringTime = append(t.AdditionalMonitoringTime, de.Value)
		case NUMBER_OF_REPORTS_TYPE:
			opt = true
			if seen[25] {
				continue
			}
			seen[25] = true
			t.NumberofReports = de.Value
		case APPLICATION_ID_TYPE:
			opt = true
			var v ApplicationID
			if err = de.UnmarshalValue(&v); err == nil {
				t.ExemptedApplicationIDforQuotaAction = append(t.ExemptedApplicationIDforQuotaAction, v)
			}
		case SDF_FILTER_TYPE:
			opt = true
			var v SDFFilter
			if err = de.UnmarshalValue(&v); err == nil {
				t.ExemptedSDFFilterforQuotaAction = append(t.ExemptedSDFFilterforQuotaAction, v)
			}
		case USER_PLANE_INACTIVITY_TIMER_TYPE:
			if seen[26] {
				continue
			}
			seen[26] = true
			t.UserPlaneInactivityTimer = de.Value
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
		return &pfcp.MissingIE{Type: URR_ID_TYPE}
	}
	if !seen[1] {
		return &pfcp.MissingIE{Type: MEASUREMENT_METHOD_TYPE}
	}
	if !seen[2] {
		return &pfcp.MissingIE{Type: REPORTING_TRIGGERS_TYPE}
	}
	return
}
