package ie

import (
	"github.com/reogac/pfcp"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Update BAR PFCP Session Report Response' data structure */
type UpdateBARPFCPSessionReportResponse struct {
	BARID                           uint8  //BAR ID
	DownlinkDataNotificationDelay   *uint8 //Downlink Data Notification Delay
	DLBufferingDuration             []byte //DL Buffering Duration
	DLBufferingSuggestedPacketCount []byte //DL Buffering Suggested Packet Count
	SuggestedBufferingPacketsCount  *uint8 //Suggested Buffering Packets Count
}

func (t *UpdateBARPFCPSessionReportResponse) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(BAR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint8(t.BARID))))
	if t.DownlinkDataNotificationDelay != nil {
		fields = append(fields, tlv.TLVFrom(DOWNLINK_DATA_NOTIFICATION_DELAY_TYPE, tlv.Bytes(pfcp.MarshalUint8(*t.DownlinkDataNotificationDelay))))
	}
	if len(t.DLBufferingDuration) > 0 {
		fields = append(fields, tlv.TLVFrom(DL_BUFFERING_DURATION_TYPE, tlv.Bytes(t.DLBufferingDuration)))
	}
	if len(t.DLBufferingSuggestedPacketCount) > 0 {
		fields = append(fields, tlv.TLVFrom(DL_BUFFERING_SUGGESTED_PACKET_COUNT_TYPE, tlv.Bytes(t.DLBufferingSuggestedPacketCount)))
	}
	if t.SuggestedBufferingPacketsCount != nil {
		fields = append(fields, tlv.TLVFrom(SUGGESTED_BUFFERING_PACKETS_COUNT_TYPE, tlv.Bytes(pfcp.MarshalUint8(*t.SuggestedBufferingPacketsCount))))
	}
	return tlv.TLV(UPDATE_BAR_PFCP_SESSION_REPORT_RESPONSE_TYPE, fields...)
}

func (t *UpdateBARPFCPSessionReportResponse) UnmarshalBinary(buf []byte) (err error) {
	*t = UpdateBARPFCPSessionReportResponse{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case BAR_ID_TYPE:
			err = pfcp.UnmarshalUint8(de.Value, &t.BARID)
		case DOWNLINK_DATA_NOTIFICATION_DELAY_TYPE:
			var v uint8
			if err = pfcp.UnmarshalUint8(de.Value, &v); err == nil {
				t.DownlinkDataNotificationDelay = &v
			}
		case DL_BUFFERING_DURATION_TYPE:
			t.DLBufferingDuration = de.Value
		case DL_BUFFERING_SUGGESTED_PACKET_COUNT_TYPE:
			t.DLBufferingSuggestedPacketCount = de.Value
		case SUGGESTED_BUFFERING_PACKETS_COUNT_TYPE:
			var v uint8
			if err = pfcp.UnmarshalUint8(de.Value, &v); err == nil {
				t.SuggestedBufferingPacketsCount = &v
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
