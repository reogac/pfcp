package ie

import (
	"github.com/reogac/pfcp"

	"github.com/reogac/pfcp/tlv"
)

/* 'Packet Rate Status Report IE within PFCP Session Modification Response' data structure */
type PacketRateStatusReportIEwithinPFCPSessionModificationResponse struct {
	QERID            uint32 //QER ID
	PacketRateStatus []byte //Packet Rate Status
}

func (t *PacketRateStatusReportIEwithinPFCPSessionModificationResponse) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(QER_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(t.QERID))))
	fields = append(fields, tlv.TLVFrom(PACKET_RATE_STATUS_TYPE, tlv.Bytes(t.PacketRateStatus)))
	return tlv.TLV(PACKET_RATE_STATUS_REPORT_IE_WITHIN_PFCP_SESSION_MODIFICATION_RESPONSE_TYPE, fields...)
}

func (t *PacketRateStatusReportIEwithinPFCPSessionModificationResponse) UnmarshalBinary(buf []byte) (err error) {
	*t = PacketRateStatusReportIEwithinPFCPSessionModificationResponse{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	for _, de := range elements {
		switch de.Type {
		case QER_ID_TYPE:
			err = pfcp.UnmarshalUint32(de.Value, &t.QERID)
		case PACKET_RATE_STATUS_TYPE:
			t.PacketRateStatus = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
