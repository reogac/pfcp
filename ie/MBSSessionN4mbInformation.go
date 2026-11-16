package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'MBS Session N4mb Information' data structure */
type MBSSessionN4mbInformation struct {
	MulticastTransportInformation []byte   //Multicast Transport Information
	Unknown                       []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid                       []uint32 `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
}

func (t *MBSSessionN4mbInformation) Field() tlv.Field {
	fields := []tlv.Field{}
	if len(t.MulticastTransportInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(MULTICAST_TRANSPORT_INFORMATION_TYPE, tlv.Bytes(t.MulticastTransportInformation)))
	}
	return tlv.TLV(MBS_SESSION_N4MB_INFORMATION_TYPE, fields...)
}

func (t *MBSSessionN4mbInformation) UnmarshalBinary(buf []byte) (err error) {
	*t = MBSSessionN4mbInformation{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [1]bool
	for _, de := range elements {
		opt := false
		switch de.Type {
		case MULTICAST_TRANSPORT_INFORMATION_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.MulticastTransportInformation = de.Value
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
	return
}
