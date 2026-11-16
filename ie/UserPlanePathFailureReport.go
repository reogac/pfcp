package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'User Plane Path Failure Report' data structure */
type UserPlanePathFailureReport struct {
	RemoteGTPUPeer [][]byte //Remote GTP-U Peer
	Unknown        []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *UserPlanePathFailureReport) Field() tlv.Field {
	fields := []tlv.Field{}
	for _, ie := range t.RemoteGTPUPeer {
		fields = append(fields, tlv.TLVFrom(REMOTE_GTP_U_PEER_TYPE, tlv.Bytes(ie)))
	}
	return tlv.TLV(USER_PLANE_PATH_FAILURE_REPORT_TYPE, fields...)
}

func (t *UserPlanePathFailureReport) UnmarshalBinary(buf []byte) (err error) {
	*t = UserPlanePathFailureReport{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	for _, de := range elements {
		switch de.Type {
		case REMOTE_GTP_U_PEER_TYPE:
			t.RemoteGTPUPeer = append(t.RemoteGTPUPeer, de.Value)
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return pfcp.Faulty(de.Type, err)
		}
	}
	if len(t.RemoteGTPUPeer) == 0 {
		return &pfcp.MissingIE{Type: REMOTE_GTP_U_PEER_TYPE}
	}
	return
}
