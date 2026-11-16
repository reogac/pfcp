package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Packet Rate Status Report IE within PFCP Session Modification Response' data structure */
type PacketRateStatusReportIEwithinPFCPSessionModificationResponse struct {
	QERID            uint32   //QER ID
	PacketRateStatus []byte   //Packet Rate Status
	Unknown          []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
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
	var seen [2]bool
	for _, de := range elements {
		switch de.Type {
		case QER_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &t.QERID)
		case PACKET_RATE_STATUS_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			t.PacketRateStatus = de.Value
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
	if !seen[1] {
		return &pfcp.MissingIE{Type: PACKET_RATE_STATUS_TYPE}
	}
	return
}
