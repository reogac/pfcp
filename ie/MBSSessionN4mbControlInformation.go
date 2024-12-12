package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'MBS Session N4mb Control Information' data structure */
type MBSSessionN4mbControlInformation struct {
	MBSSessionIdentifier                           []byte //MBS Session Identifier
	AreaSessionID                                  []byte //Area Session ID
	MBSN4mbReqFlags                                []byte //MBSN4mbReq-Flags
	MulticastTransportInformationforN3mbandorN19mb []byte //Multicast Transport Information
}

func (t *MBSSessionN4mbControlInformation) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(MBS_SESSION_IDENTIFIER_TYPE, tlv.Bytes(t.MBSSessionIdentifier)))
	if len(t.AreaSessionID) > 0 {
		fields = append(fields, tlv.TLVFrom(AREA_SESSION_ID_TYPE, tlv.Bytes(t.AreaSessionID)))
	}
	if len(t.MBSN4mbReqFlags) > 0 {
		fields = append(fields, tlv.TLVFrom(MBSN4MBREQ_FLAGS_TYPE, tlv.Bytes(t.MBSN4mbReqFlags)))
	}
	if len(t.MulticastTransportInformationforN3mbandorN19mb) > 0 {
		fields = append(fields, tlv.TLVFrom(MULTICAST_TRANSPORT_INFORMATION_TYPE, tlv.Bytes(t.MulticastTransportInformationforN3mbandorN19mb)))
	}
	return tlv.TLV(MBS_SESSION_N4MB_CONTROL_INFORMATION_TYPE, fields...)
}

func (t *MBSSessionN4mbControlInformation) UnmarshalBinary(buf []byte) (err error) {
	*t = MBSSessionN4mbControlInformation{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case MBS_SESSION_IDENTIFIER_TYPE:
			t.MBSSessionIdentifier = de.Value
		case AREA_SESSION_ID_TYPE:
			t.AreaSessionID = de.Value
		case MBSN4MBREQ_FLAGS_TYPE:
			t.MBSN4mbReqFlags = de.Value
		case MULTICAST_TRANSPORT_INFORMATION_TYPE:
			t.MulticastTransportInformationforN3mbandorN19mb = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
