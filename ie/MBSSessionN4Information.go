package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'MBS Session N4 Information' data structure */
type MBSSessionN4Information struct {
	MBSSessionIdentifier []byte   //MBS Session Identifier
	AreaSessionID        []byte   //Area Session ID
	N19mbDLTunnelID      *FTEID   //F-TEID
	MBSN4RespFlags       []byte   //MBSN4Resp-Flags
	Unknown              []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid              []uint32 `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
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
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [4]bool
	for _, de := range elements {
		opt := false
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
		case F_TEID_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			var v FTEID
			if err = de.UnmarshalValue(&v); err == nil {
				t.N19mbDLTunnelID = &v
			}
		case MBSN4RESP_FLAGS_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			t.MBSN4RespFlags = de.Value
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
	if !seen[0] {
		return &pfcp.MissingIE{Type: MBS_SESSION_IDENTIFIER_TYPE}
	}
	return
}
