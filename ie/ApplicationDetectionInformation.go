package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Application Detection Information' data structure */
type ApplicationDetectionInformation struct {
	ApplicationID         ApplicationID //Application ID
	ApplicationInstanceID []byte        //Application Instance ID
	FlowInformation       []byte        //Flow Information
	PDRID                 *uint16       //PDR ID
	Unknown               []uint32      `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *ApplicationDetectionInformation) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(APPLICATION_ID_TYPE, tlv.Bytes(t.ApplicationID.Bytes())))
	if len(t.ApplicationInstanceID) > 0 {
		fields = append(fields, tlv.TLVFrom(APPLICATION_INSTANCE_ID_TYPE, tlv.Bytes(t.ApplicationInstanceID)))
	}
	if len(t.FlowInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(FLOW_INFORMATION_TYPE, tlv.Bytes(t.FlowInformation)))
	}
	if t.PDRID != nil {
		fields = append(fields, tlv.TLVFrom(PDR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint16(*t.PDRID))))
	}
	return tlv.TLV(APPLICATION_DETECTION_INFORMATION_TYPE, fields...)
}

func (t *ApplicationDetectionInformation) UnmarshalBinary(buf []byte) (err error) {
	*t = ApplicationDetectionInformation{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	var seen [4]bool
	for _, de := range elements {
		switch de.Type {
		case APPLICATION_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			if err = de.UnmarshalValue(&t.ApplicationID); err == nil {
			}
		case APPLICATION_INSTANCE_ID_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			t.ApplicationInstanceID = de.Value
		case FLOW_INFORMATION_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			t.FlowInformation = de.Value
		case PDR_ID_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			var v uint16
			if err = pfcp.UnmarshalExtendable(de.Value, 2, 2, &v); err == nil {
				t.PDRID = &v
			}
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: APPLICATION_ID_TYPE}
	}
	return
}
