package ie

import (
	"etrib5gc/pfcp"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Create BAR' data structure */
type CreateBAR struct {
	BARID                          uint8  //BAR ID
	DownlinkDataNotificationDelay  *uint8 //Downlink Data Notification Delay
	SuggestedBufferingPacketsCount *uint8 //Suggested Buffering Packets Count
	MTEDTControlInformation        []byte //MT-EDT Control Information
}

func (t *CreateBAR) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(BAR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint8(t.BARID))))
	if t.DownlinkDataNotificationDelay != nil {
		fields = append(fields, tlv.TLVFrom(DOWNLINK_DATA_NOTIFICATION_DELAY_TYPE, tlv.Bytes(pfcp.MarshalUint8(*t.DownlinkDataNotificationDelay))))
	}
	if t.SuggestedBufferingPacketsCount != nil {
		fields = append(fields, tlv.TLVFrom(SUGGESTED_BUFFERING_PACKETS_COUNT_TYPE, tlv.Bytes(pfcp.MarshalUint8(*t.SuggestedBufferingPacketsCount))))
	}
	if len(t.MTEDTControlInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(MT_EDT_CONTROL_INFORMATION_TYPE, tlv.Bytes(t.MTEDTControlInformation)))
	}
	return tlv.TLV(CREATE_BAR_TYPE, fields...)
}

func (t *CreateBAR) UnmarshalBinary(buf []byte) (err error) {
	*t = CreateBAR{}
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
		case SUGGESTED_BUFFERING_PACKETS_COUNT_TYPE:
			var v uint8
			if err = pfcp.UnmarshalUint8(de.Value, &v); err == nil {
				t.SuggestedBufferingPacketsCount = &v
			}
		case MT_EDT_CONTROL_INFORMATION_TYPE:
			t.MTEDTControlInformation = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
