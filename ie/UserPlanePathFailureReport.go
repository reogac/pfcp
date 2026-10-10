package ie

import (
	"github.com/reogac/pfcp/tlv"
)

/* 'User Plane Path Failure Report' data structure */
type UserPlanePathFailureReport struct {
	RemoteGTPUPeer []byte //Remote GTP-U Peer
}

func (t *UserPlanePathFailureReport) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(REMOTE_GTP_U_PEER_TYPE, tlv.Bytes(t.RemoteGTPUPeer)))
	return tlv.TLV(USER_PLANE_PATH_FAILURE_REPORT_TYPE, fields...)
}

func (t *UserPlanePathFailureReport) UnmarshalBinary(buf []byte) (err error) {
	*t = UserPlanePathFailureReport{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	for _, de := range elements {
		switch de.Type {
		case REMOTE_GTP_U_PEER_TYPE:
			t.RemoteGTPUPeer = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
