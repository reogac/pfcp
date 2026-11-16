package ie

import (
	"github.com/reogac/pfcp"

	"github.com/reogac/pfcp/tlv"
)

/* 'Create QER' data structure */
type CreateQER struct {
	QERID                 uint32     //QER ID
	QERCorrelationID      *uint32    //QER Correlation ID
	GateStatus            GateStatus //Gate Status
	MaximumBitrate        *MBR       //MBR
	GuaranteedBitrate     *GBR       //GBR
	PacketRate            []byte     //Packet Rate
	PacketRateStatus      []byte     //Packet Rate Status
	DLFlowLevelMarking    []byte     //DL Flow Level Marking
	QoSflowidentifier     *uint8     //QFI
	ReflectiveQoS         *uint8     //RQI
	PagingPolicyIndicator *uint8     //Paging Policy Indicator
	AveragingWindow       *uint32    //Averaging Window
	QERControlIndications []byte     //QER Control Indications
	QERIndications        []byte     //QER Indications
	Unknown               []uint32   `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *CreateQER) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(QER_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(t.QERID))))
	if t.QERCorrelationID != nil {
		fields = append(fields, tlv.TLVFrom(QER_CORRELATION_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(*t.QERCorrelationID))))
	}
	fields = append(fields, tlv.TLVFrom(GATE_STATUS_TYPE, tlv.Bytes(t.GateStatus.Bytes())))
	if t.MaximumBitrate != nil {
		fields = append(fields, tlv.TLVFrom(MBR_TYPE, tlv.Bytes(t.MaximumBitrate.Bytes())))
	}
	if t.GuaranteedBitrate != nil {
		fields = append(fields, tlv.TLVFrom(GBR_TYPE, tlv.Bytes(t.GuaranteedBitrate.Bytes())))
	}
	if len(t.PacketRate) > 0 {
		fields = append(fields, tlv.TLVFrom(PACKET_RATE_TYPE, tlv.Bytes(t.PacketRate)))
	}
	if len(t.PacketRateStatus) > 0 {
		fields = append(fields, tlv.TLVFrom(PACKET_RATE_STATUS_TYPE, tlv.Bytes(t.PacketRateStatus)))
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
	if len(t.QERIndications) > 0 {
		fields = append(fields, tlv.TLVFrom(QER_INDICATIONS_TYPE, tlv.Bytes(t.QERIndications)))
	}
	return tlv.TLV(CREATE_QER_TYPE, fields...)
}

func (t *CreateQER) UnmarshalBinary(buf []byte) (err error) {
	*t = CreateQER{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	var seen [14]bool
	for _, de := range elements {
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
			if err = de.UnmarshalValue(&t.GateStatus); err == nil {
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
		case PACKET_RATE_STATUS_TYPE:
			if seen[6] {
				continue
			}
			seen[6] = true
			t.PacketRateStatus = de.Value
		case DL_FLOW_LEVEL_MARKING_TYPE:
			if seen[7] {
				continue
			}
			seen[7] = true
			t.DLFlowLevelMarking = de.Value
		case QFI_TYPE:
			if seen[8] {
				continue
			}
			seen[8] = true
			var v uint8
			if err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &v); err == nil {
				t.QoSflowidentifier = &v
			}
		case RQI_TYPE:
			if seen[9] {
				continue
			}
			seen[9] = true
			var v uint8
			if err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &v); err == nil {
				t.ReflectiveQoS = &v
			}
		case PAGING_POLICY_INDICATOR_TYPE:
			if seen[10] {
				continue
			}
			seen[10] = true
			var v uint8
			if err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &v); err == nil {
				t.PagingPolicyIndicator = &v
			}
		case AVERAGING_WINDOW_TYPE:
			if seen[11] {
				continue
			}
			seen[11] = true
			var v uint32
			if err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &v); err == nil {
				t.AveragingWindow = &v
			}
		case QER_CONTROL_INDICATIONS_TYPE:
			if seen[12] {
				continue
			}
			seen[12] = true
			t.QERControlIndications = de.Value
		case QER_INDICATIONS_TYPE:
			if seen[13] {
				continue
			}
			seen[13] = true
			t.QERIndications = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: QER_ID_TYPE}
	}
	if !seen[2] {
		return &pfcp.MissingIE{Type: GATE_STATUS_TYPE}
	}
	return
}
