package message

import (
	"github.com/reogac/pfcp/v2"
	"github.com/reogac/pfcp/v2/ie"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'PFCP Heartbeat Request' data structure */
type PFCPHeartbeatRequest struct {
	RecoveryTimeStamp uint32   //Recovery Time Stamp
	SourceIPAddress   []byte   //Source IP Address
	Unknown           []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *PFCPHeartbeatRequest) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(ie.RECOVERY_TIME_STAMP_TYPE, tlv.Bytes(pfcp.MarshalUint32(t.RecoveryTimeStamp))))
	if len(t.SourceIPAddress) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.SOURCE_IP_ADDRESS_TYPE, tlv.Bytes(t.SourceIPAddress)))
	}
	return tlv.TLV(ie.MSG_PFCP_HEARTBEAT_REQUEST_TYPE, fields...)
}

func (t *PFCPHeartbeatRequest) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPHeartbeatRequest{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [2]bool
	for _, de := range elements {
		switch de.Type {
		case ie.RECOVERY_TIME_STAMP_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &t.RecoveryTimeStamp)
		case ie.SOURCE_IP_ADDRESS_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			t.SourceIPAddress = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return pfcp.Faulty(de.Type, err)
		}
	}
	return
}

func (t *PFCPHeartbeatRequest) UnmarshalTLV(typ uint32, buf []byte) (err error) {
	if typ != ie.MSG_PFCP_HEARTBEAT_REQUEST_TYPE {
		err = tlv.ErrCritical
		return
	}

	err = t.UnmarshalBinary(buf)
	return
}

func (t *PFCPHeartbeatRequest) MarshalJSON() (buf []byte, err error) {
	buf, err = jsonMarshal(t)
	return
}

func (t *PFCPHeartbeatRequest) UnmarshalJSON(buf []byte) (err error) {
	*t = PFCPHeartbeatRequest{}
	err = jsonUnmarshal(buf, t)
	return
}

/* 'PFCP Heartbeat Response' data structure */
type PFCPHeartbeatResponse struct {
	RecoveryTimeStamp uint32   //Recovery Time Stamp
	Unknown           []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *PFCPHeartbeatResponse) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(ie.RECOVERY_TIME_STAMP_TYPE, tlv.Bytes(pfcp.MarshalUint32(t.RecoveryTimeStamp))))
	return tlv.TLV(ie.MSG_PFCP_HEARTBEAT_RESPONSE_TYPE, fields...)
}

func (t *PFCPHeartbeatResponse) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPHeartbeatResponse{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [1]bool
	for _, de := range elements {
		switch de.Type {
		case ie.RECOVERY_TIME_STAMP_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &t.RecoveryTimeStamp)
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return pfcp.Faulty(de.Type, err)
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: ie.RECOVERY_TIME_STAMP_TYPE}
	}
	return
}

func (t *PFCPHeartbeatResponse) UnmarshalTLV(typ uint32, buf []byte) (err error) {
	if typ != ie.MSG_PFCP_HEARTBEAT_RESPONSE_TYPE {
		err = tlv.ErrCritical
		return
	}

	err = t.UnmarshalBinary(buf)
	return
}

func (t *PFCPHeartbeatResponse) MarshalJSON() (buf []byte, err error) {
	buf, err = jsonMarshal(t)
	return
}

func (t *PFCPHeartbeatResponse) UnmarshalJSON(buf []byte) (err error) {
	*t = PFCPHeartbeatResponse{}
	err = jsonUnmarshal(buf, t)
	return
}
