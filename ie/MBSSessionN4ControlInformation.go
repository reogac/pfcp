package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'MBS Session N4 Control Information' data structure */
type MBSSessionN4ControlInformation struct {
	MBSSessionIdentifier          []byte //MBS Session Identifier
	AreaSessionID                 []byte //Area Session ID
	MulticastTransportInformation []byte //Multicast Transport Information
}

func (t *MBSSessionN4ControlInformation) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(MBS_SESSION_IDENTIFIER_TYPE, tlv.Bytes(t.MBSSessionIdentifier)))
	if len(t.AreaSessionID) > 0 {
		fields = append(fields, tlv.TLVFrom(AREA_SESSION_ID_TYPE, tlv.Bytes(t.AreaSessionID)))
	}
	if len(t.MulticastTransportInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(MULTICAST_TRANSPORT_INFORMATION_TYPE, tlv.Bytes(t.MulticastTransportInformation)))
	}
	return tlv.TLV(MBS_SESSION_N4_CONTROL_INFORMATION_TYPE, fields...)
}

func (t *MBSSessionN4ControlInformation) UnmarshalBinary(buf []byte) (err error) {
	*t = MBSSessionN4ControlInformation{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case MBS_SESSION_IDENTIFIER_TYPE:
			t.MBSSessionIdentifier = de.Value
		case AREA_SESSION_ID_TYPE:
			t.AreaSessionID = de.Value
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
