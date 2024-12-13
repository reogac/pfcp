package message

import (
	"github.com/reogac/pfcp"
	"github.com/reogac/pfcp/ie"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'PFCP Session Set Modification Request' data structure */
type PFCPSessionSetModificationRequest struct {
	NodeID                ie.NodeID                //Node ID
	PFCPSessionChangeInfo ie.PFCPSessionChangeInfo //PFCP Session Change Info
}

func (t *PFCPSessionSetModificationRequest) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(ie.NODE_ID_TYPE, tlv.Bytes(t.NodeID.Bytes())))
	fields = append(fields, t.PFCPSessionChangeInfo.Field())
	return tlv.TLV(ie.MSG_PFCP_SESSION_SET_MODIFICATION_REQUEST_TYPE, fields...)
}

func (t *PFCPSessionSetModificationRequest) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPSessionSetModificationRequest{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case ie.NODE_ID_TYPE:
			err = de.UnmarshalValue(&t.NodeID)
		case ie.PFCP_SESSION_CHANGE_INFO_TYPE:
			err = de.UnmarshalValue(&t.PFCPSessionChangeInfo)
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}

func (t *PFCPSessionSetModificationRequest) UnmarshalTLV(typ uint32, buf []byte) (err error) {
	if typ != ie.MSG_PFCP_SESSION_SET_MODIFICATION_REQUEST_TYPE {
		err = tlv.ErrCritical
		return
	}

	err = t.UnmarshalBinary(buf)
	return
}

func (t *PFCPSessionSetModificationRequest) MarshalJSON() (buf []byte, err error) {
	buf, err = jsonMarshal(t)
	return
}

func (t *PFCPSessionSetModificationRequest) UnmarshalJSON(buf []byte) (err error) {
	*t = PFCPSessionSetModificationRequest{}
	err = jsonUnmarshal(buf, t)
	return
}

/* 'PFCP Session Set Modification Response' data structure */
type PFCPSessionSetModificationResponse struct {
	NodeID      ie.NodeID //Node ID
	Cause       uint8     //Cause
	OffendingIE *uint16   //Offending IE
}

func (t *PFCPSessionSetModificationResponse) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(ie.NODE_ID_TYPE, tlv.Bytes(t.NodeID.Bytes())))
	fields = append(fields, tlv.TLVFrom(ie.CAUSE_TYPE, tlv.Bytes(pfcp.MarshalUint8(t.Cause))))
	if t.OffendingIE != nil {
		fields = append(fields, tlv.TLVFrom(ie.OFFENDING_IE_TYPE, tlv.Bytes(pfcp.MarshalUint16(*t.OffendingIE))))
	}
	return tlv.TLV(ie.MSG_PFCP_SESSION_SET_MODIFICATION_RESPONSE_TYPE, fields...)
}

func (t *PFCPSessionSetModificationResponse) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPSessionSetModificationResponse{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case ie.NODE_ID_TYPE:
			err = de.UnmarshalValue(&t.NodeID)
		case ie.CAUSE_TYPE:
			err = pfcp.UnmarshalUint8(de.Value, &t.Cause)
		case ie.OFFENDING_IE_TYPE:
			var v uint16
			if err = pfcp.UnmarshalUint16(de.Value, &v); err == nil {
				t.OffendingIE = &v
			}
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}

func (t *PFCPSessionSetModificationResponse) UnmarshalTLV(typ uint32, buf []byte) (err error) {
	if typ != ie.MSG_PFCP_SESSION_SET_MODIFICATION_RESPONSE_TYPE {
		err = tlv.ErrCritical
		return
	}

	err = t.UnmarshalBinary(buf)
	return
}

func (t *PFCPSessionSetModificationResponse) MarshalJSON() (buf []byte, err error) {
	buf, err = jsonMarshal(t)
	return
}

func (t *PFCPSessionSetModificationResponse) UnmarshalJSON(buf []byte) (err error) {
	*t = PFCPSessionSetModificationResponse{}
	err = jsonUnmarshal(buf, t)
	return
}
