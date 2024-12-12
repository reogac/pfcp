package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'MBS Session N4mb Information' data structure */
type MBSSessionN4mbInformation struct {
	MulticastTransportInformation []byte //Multicast Transport Information
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
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case MULTICAST_TRANSPORT_INFORMATION_TYPE:
			t.MulticastTransportInformation = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
