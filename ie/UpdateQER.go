package ie

import (
	"etrib5gc/pfcp"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
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
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case QER_ID_TYPE:
			err = pfcp.UnmarshalUint32(de.Value, &t.QERID)
		case QER_CORRELATION_ID_TYPE:
			var v uint32
			if err = pfcp.UnmarshalUint32(de.Value, &v); err == nil {
				t.QERCorrelationID = &v
			}
		case GATE_STATUS_TYPE:
			var v GateStatus
			if err = de.UnmarshalValue(&v); err == nil {
				t.GateStatus = &v
			}
		case MBR_TYPE:
			var v MBR
			if err = de.UnmarshalValue(&v); err == nil {
				t.MaximumBitrate = &v
			}
		case GBR_TYPE:
			var v GBR
			if err = de.UnmarshalValue(&v); err == nil {
				t.GuaranteedBitrate = &v
			}
		case PACKET_RATE_TYPE:
			t.PacketRate = de.Value
		case DL_FLOW_LEVEL_MARKING_TYPE:
			t.DLFlowLevelMarking = de.Value
		case QFI_TYPE:
			var v uint8
			if err = pfcp.UnmarshalUint8(de.Value, &v); err == nil {
				t.QoSflowidentifier = &v
			}
		case RQI_TYPE:
			var v uint8
			if err = pfcp.UnmarshalUint8(de.Value, &v); err == nil {
				t.ReflectiveQoS = &v
			}
		case PAGING_POLICY_INDICATOR_TYPE:
			var v uint8
			if err = pfcp.UnmarshalUint8(de.Value, &v); err == nil {
				t.PagingPolicyIndicator = &v
			}
		case AVERAGING_WINDOW_TYPE:
			var v uint32
			if err = pfcp.UnmarshalUint32(de.Value, &v); err == nil {
				t.AveragingWindow = &v
			}
		case QER_CONTROL_INDICATIONS_TYPE:
			t.QERControlIndications = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
