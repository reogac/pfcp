package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'User Plane Path Recovery Report' data structure */
type UserPlanePathRecoveryReport struct {
	RemoteGTPUPeer []byte //Remote GTP-U Peer
}

func (t *UserPlanePathRecoveryReport) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(REMOTE_GTP_U_PEER_TYPE, tlv.Bytes(t.RemoteGTPUPeer)))
	return tlv.TLV(USER_PLANE_PATH_RECOVERY_REPORT_TYPE, fields...)
}

func (t *UserPlanePathRecoveryReport) UnmarshalBinary(buf []byte) (err error) {
	*t = UserPlanePathRecoveryReport{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
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
