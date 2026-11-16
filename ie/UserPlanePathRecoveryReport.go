package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'User Plane Path Recovery Report' data structure */
type UserPlanePathRecoveryReport struct {
	RemoteGTPUPeer [][]byte //Remote GTP-U Peer
	Unknown        []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid        []uint32 `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
}

func (t *UserPlanePathRecoveryReport) Field() tlv.Field {
	fields := []tlv.Field{}
	for _, ie := range t.RemoteGTPUPeer {
		fields = append(fields, tlv.TLVFrom(REMOTE_GTP_U_PEER_TYPE, tlv.Bytes(ie)))
	}
	return tlv.TLV(USER_PLANE_PATH_RECOVERY_REPORT_TYPE, fields...)
}

func (t *UserPlanePathRecoveryReport) UnmarshalBinary(buf []byte) (err error) {
	*t = UserPlanePathRecoveryReport{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	for _, de := range elements {
		opt := false
		switch de.Type {
		case REMOTE_GTP_U_PEER_TYPE:
			t.RemoteGTPUPeer = append(t.RemoteGTPUPeer, de.Value)
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
	if len(t.RemoteGTPUPeer) == 0 {
		return &pfcp.MissingIE{Type: REMOTE_GTP_U_PEER_TYPE}
	}
	return
}
