package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'MBS Session N4 Information' data structure */
type MBSSessionN4Information struct {
	MBSSessionIdentifier []byte //MBS Session Identifier
	AreaSessionID        []byte //Area Session ID
	N19mbDLTunnelID      *FTEID //F-TEID
	MBSN4RespFlags       []byte //MBSN4Resp-Flags
}

func (t *MBSSessionN4Information) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(MBS_SESSION_IDENTIFIER_TYPE, tlv.Bytes(t.MBSSessionIdentifier)))
	if len(t.AreaSessionID) > 0 {
		fields = append(fields, tlv.TLVFrom(AREA_SESSION_ID_TYPE, tlv.Bytes(t.AreaSessionID)))
	}
	if t.N19mbDLTunnelID != nil {
		fields = append(fields, tlv.TLVFrom(F_TEID_TYPE, tlv.Bytes(t.N19mbDLTunnelID.Bytes())))
	}
	if len(t.MBSN4RespFlags) > 0 {
		fields = append(fields, tlv.TLVFrom(MBSN4RESP_FLAGS_TYPE, tlv.Bytes(t.MBSN4RespFlags)))
	}
	return tlv.TLV(MBS_SESSION_N4_INFORMATION_TYPE, fields...)
}

func (t *MBSSessionN4Information) UnmarshalBinary(buf []byte) (err error) {
	*t = MBSSessionN4Information{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case MBS_SESSION_IDENTIFIER_TYPE:
			t.MBSSessionIdentifier = de.Value
		case AREA_SESSION_ID_TYPE:
			t.AreaSessionID = de.Value
		case F_TEID_TYPE:
			var v FTEID
			if err = de.UnmarshalValue(&v); err == nil {
				t.N19mbDLTunnelID = &v
			}
		case MBSN4RESP_FLAGS_TYPE:
			t.MBSN4RespFlags = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
