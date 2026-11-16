package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Transport Delay Reporting' data structure */
type TransportDelayReporting struct {
	PrecedingULGTPUPeer []byte   //Remote GTP-U Peer
	DSCP                []byte   //Transport Level Marking
	Unknown             []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid             []uint32 `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
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
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [2]bool
	for _, de := range elements {
		opt := false
		switch de.Type {
		case REMOTE_GTP_U_PEER_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.PrecedingULGTPUPeer = de.Value
		case TRANSPORT_LEVEL_MARKING_TYPE:
			opt = true
			if seen[1] {
				continue
			}
			seen[1] = true
			t.DSCP = de.Value
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
		return &pfcp.MissingIE{Type: REMOTE_GTP_U_PEER_TYPE}
	}
	return
}
