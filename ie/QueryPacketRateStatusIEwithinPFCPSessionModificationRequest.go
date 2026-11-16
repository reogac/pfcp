package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Query Packet Rate Status IE within PFCP Session Modification Request' data structure */
type QueryPacketRateStatusIEwithinPFCPSessionModificationRequest struct {
	QERID   uint32   //QER ID
	Unknown []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid []uint32 `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
}

func (t *QueryPacketRateStatusIEwithinPFCPSessionModificationRequest) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(QER_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(t.QERID))))
	return tlv.TLV(QUERY_PACKET_RATE_STATUS_IE_WITHIN_PFCP_SESSION_MODIFICATION_REQUEST_TYPE, fields...)
}

func (t *QueryPacketRateStatusIEwithinPFCPSessionModificationRequest) UnmarshalBinary(buf []byte) (err error) {
	*t = QueryPacketRateStatusIEwithinPFCPSessionModificationRequest{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [1]bool
	for _, de := range elements {
		opt := false
		switch de.Type {
		case QER_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &t.QERID)
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
