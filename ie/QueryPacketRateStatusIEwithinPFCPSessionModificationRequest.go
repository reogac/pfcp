package ie

import (
	"github.com/reogac/pfcp"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Query Packet Rate Status IE within PFCP Session Modification Request' data structure */
type QueryPacketRateStatusIEwithinPFCPSessionModificationRequest struct {
	QERID uint32 //QER ID
}

func (t *QueryPacketRateStatusIEwithinPFCPSessionModificationRequest) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(QER_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(t.QERID))))
	return tlv.TLV(QUERY_PACKET_RATE_STATUS_IE_WITHIN_PFCP_SESSION_MODIFICATION_REQUEST_TYPE, fields...)
}

func (t *QueryPacketRateStatusIEwithinPFCPSessionModificationRequest) UnmarshalBinary(buf []byte) (err error) {
	*t = QueryPacketRateStatusIEwithinPFCPSessionModificationRequest{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case QER_ID_TYPE:
			err = pfcp.UnmarshalUint32(de.Value, &t.QERID)
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
