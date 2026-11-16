package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Update QER' data structure */
type UpdateQER struct {
	QERID                 uint32      //QER ID
	QERCorrelationID      *uint32     //QER Correlation ID
	GateStatus            *GateStatus //Gate Status
	MaximumBitrate        *MBR        //MBR
	GuaranteedBitrate     *GBR        //GBR
	PacketRate            []byte      //Packet Rate
	DLFlowLevelMarking    []byte      //DL Flow Level Marking
	QoSflowidentifier     *uint8      //QFI
	ReflectiveQoS         *uint8      //RQI
	PagingPolicyIndicator *uint8      //Paging Policy Indicator
	AveragingWindow       *uint32     //Averaging Window
	QERControlIndications []byte      //QER Control Indications
	Unknown               []uint32    `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid               []uint32    `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
}

func (t *UpdateQER) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(QER_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(t.QERID))))
	if t.QERCorrelationID != nil {
		fields = append(fields, tlv.TLVFrom(QER_CORRELATION_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(*t.QERCorrelationID))))
	}
	if t.GateStatus != nil {
		fields = append(fields, tlv.TLVFrom(GATE_STATUS_TYPE, tlv.Bytes(t.GateStatus.Bytes())))
	}
	if t.MaximumBitrate != nil {
		fields = append(fields, tlv.TLVFrom(MBR_TYPE, tlv.Bytes(t.MaximumBitrate.Bytes())))
	}
	if t.GuaranteedBitrate != nil {
		fields = append(fields, tlv.TLVFrom(GBR_TYPE, tlv.Bytes(t.GuaranteedBitrate.Bytes())))
	}
	if len(t.PacketRate) > 0 {
		fields = append(fields, tlv.TLVFrom(PACKET_RATE_TYPE, tlv.Bytes(t.PacketRate)))
	}
	if len(t.DLFlowLevelMarking) > 0 {
		fields = append(fields, tlv.TLVFrom(DL_FLOW_LEVEL_MARKING_TYPE, tlv.Bytes(t.DLFlowLevelMarking)))
	}
	if t.QoSflowidentifier != nil {
		fields = append(fields, tlv.TLVFrom(QFI_TYPE, tlv.Bytes(pfcp.MarshalUint8(*t.QoSflowidentifier))))
	}
	if t.ReflectiveQoS != nil {
		fields = append(fields, tlv.TLVFrom(RQI_TYPE, tlv.Bytes(pfcp.MarshalUint8(*t.ReflectiveQoS))))
	}
	if t.PagingPolicyIndicator != nil {
		fields = append(fields, tlv.TLVFrom(PAGING_POLICY_INDICATOR_TYPE, tlv.Bytes(pfcp.MarshalUint8(*t.PagingPolicyIndicator))))
	}
	if t.AveragingWindow != nil {
		fields = append(fields, tlv.TLVFrom(AVERAGING_WINDOW_TYPE, tlv.Bytes(pfcp.MarshalUint32(*t.AveragingWindow))))
	}
	if len(t.QERControlIndications) > 0 {
		fields = append(fields, tlv.TLVFrom(QER_CONTROL_INDICATIONS_TYPE, tlv.Bytes(t.QERControlIndications)))
	}
	return tlv.TLV(UPDATE_QER_TYPE, fields...)
}

func (t *UpdateQER) UnmarshalBinary(buf []byte) (err error) {
	*t = UpdateQER{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [12]bool
	for _, de := range elements {
		opt := false
		switch de.Type {
		case QER_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &t.QERID)
		case QER_CORRELATION_ID_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			var v uint32
			if err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &v); err == nil {
				t.QERCorrelationID = &v
			}
		case GATE_STATUS_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			var v GateStatus
			if err = de.UnmarshalValue(&v); err == nil {
				t.GateStatus = &v
			}
		case MBR_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			var v MBR
			if err = de.UnmarshalValue(&v); err == nil {
				t.MaximumBitrate = &v
			}
		case GBR_TYPE:
			if seen[4] {
				continue
			}
			seen[4] = true
			var v GBR
			if err = de.UnmarshalValue(&v); err == nil {
				t.GuaranteedBitrate = &v
			}
		case PACKET_RATE_TYPE:
			if seen[5] {
				continue
			}
			seen[5] = true
			t.PacketRate = de.Value
		case DL_FLOW_LEVEL_MARKING_TYPE:
			if seen[6] {
				continue
			}
			seen[6] = true
			t.DLFlowLevelMarking = de.Value
		case QFI_TYPE:
			if seen[7] {
				continue
			}
			seen[7] = true
			var v uint8
			if err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &v); err == nil {
				t.QoSflowidentifier = &v
			}
		case RQI_TYPE:
			if seen[8] {
				continue
			}
			seen[8] = true
			var v uint8
			if err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &v); err == nil {
				t.ReflectiveQoS = &v
			}
		case PAGING_POLICY_INDICATOR_TYPE:
			if seen[9] {
				continue
			}
			seen[9] = true
			var v uint8
			if err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &v); err == nil {
				t.PagingPolicyIndicator = &v
			}
		case AVERAGING_WINDOW_TYPE:
			opt = true
			if seen[10] {
				continue
			}
			seen[10] = true
			var v uint32
			if err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &v); err == nil {
				t.AveragingWindow = &v
			}
		case QER_CONTROL_INDICATIONS_TYPE:
			if seen[11] {
				continue
			}
			seen[11] = true
			t.QERControlIndications = de.Value
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
		return &pfcp.MissingIE{Type: QER_ID_TYPE}
	}
	return
}
