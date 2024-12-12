package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Transport Delay Reporting' data structure */
type TransportDelayReporting struct {
	PrecedingULGTPUPeer []byte //Remote GTP-U Peer
	DSCP                []byte //Transport Level Marking
}

func (t *TransportDelayReporting) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(REMOTE_GTP_U_PEER_TYPE, tlv.Bytes(t.PrecedingULGTPUPeer)))
	if len(t.DSCP) > 0 {
		fields = append(fields, tlv.TLVFrom(TRANSPORT_LEVEL_MARKING_TYPE, tlv.Bytes(t.DSCP)))
	}
	return tlv.TLV(TRANSPORT_DELAY_REPORTING_TYPE, fields...)
}

func (t *TransportDelayReporting) UnmarshalBinary(buf []byte) (err error) {
	*t = TransportDelayReporting{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case REMOTE_GTP_U_PEER_TYPE:
			t.PrecedingULGTPUPeer = de.Value
		case TRANSPORT_LEVEL_MARKING_TYPE:
			t.DSCP = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
