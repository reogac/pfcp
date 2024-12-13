package message

import (
	"github.com/reogac/pfcp"
	"github.com/reogac/pfcp/ie"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'PFCP Heartbeat Request' data structure */
type PFCPHeartbeatRequest struct {
	RecoveryTimeStamp uint32 //Recovery Time Stamp
	SourceIPAddress   []byte //Source IP Address
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
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case ie.RECOVERY_TIME_STAMP_TYPE:
			err = pfcp.UnmarshalUint32(de.Value, &t.RecoveryTimeStamp)
		case ie.SOURCE_IP_ADDRESS_TYPE:
			t.SourceIPAddress = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
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
	RecoveryTimeStamp uint32 //Recovery Time Stamp
}

func (t *PFCPHeartbeatResponse) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(ie.RECOVERY_TIME_STAMP_TYPE, tlv.Bytes(pfcp.MarshalUint32(t.RecoveryTimeStamp))))
	return tlv.TLV(ie.MSG_PFCP_HEARTBEAT_RESPONSE_TYPE, fields...)
}

func (t *PFCPHeartbeatResponse) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPHeartbeatResponse{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case ie.RECOVERY_TIME_STAMP_TYPE:
			err = pfcp.UnmarshalUint32(de.Value, &t.RecoveryTimeStamp)
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
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
