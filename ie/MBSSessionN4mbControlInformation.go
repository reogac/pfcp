package ie

import (
	"github.com/reogac/pfcp/tlv"
)

/* 'MBS Session N4mb Control Information' data structure */
type MBSSessionN4mbControlInformation struct {
	MBSSessionIdentifier                           []byte   //MBS Session Identifier
	AreaSessionID                                  []byte   //Area Session ID
	MBSN4mbReqFlags                                []byte   //MBSN4mbReq-Flags
	MulticastTransportInformationforN3mbandorN19mb []byte   //Multicast Transport Information
	Unknown                                        []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
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
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	var seen [4]bool
	for _, de := range elements {
		switch de.Type {
		case MBS_SESSION_IDENTIFIER_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.MBSSessionIdentifier = de.Value
		case AREA_SESSION_ID_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			t.AreaSessionID = de.Value
		case MBSN4MBREQ_FLAGS_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			t.MBSN4mbReqFlags = de.Value
		case MULTICAST_TRANSPORT_INFORMATION_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			t.MulticastTransportInformationforN3mbandorN19mb = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	return
}
