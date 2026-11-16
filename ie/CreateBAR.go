package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Create BAR' data structure */
type CreateBAR struct {
	BARID                          uint8    //BAR ID
	DownlinkDataNotificationDelay  *uint8   //Downlink Data Notification Delay
	SuggestedBufferingPacketsCount *uint8   //Suggested Buffering Packets Count
	MTEDTControlInformation        []byte   //MT-EDT Control Information
	Unknown                        []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
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
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	var seen [4]bool
	for _, de := range elements {
		switch de.Type {
		case BAR_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &t.BARID)
		case DOWNLINK_DATA_NOTIFICATION_DELAY_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			var v uint8
			if err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &v); err == nil {
				t.DownlinkDataNotificationDelay = &v
			}
		case SUGGESTED_BUFFERING_PACKETS_COUNT_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			var v uint8
			if err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &v); err == nil {
				t.SuggestedBufferingPacketsCount = &v
			}
		case MT_EDT_CONTROL_INFORMATION_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			t.MTEDTControlInformation = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: BAR_ID_TYPE}
	}
	return
}
