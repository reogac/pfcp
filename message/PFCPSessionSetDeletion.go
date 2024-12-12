package message

import (
	"etrib5gc/pfcp"
	"etrib5gc/pfcp/ie"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'PFCP Session Set Deletion Request' data structure */
type PFCPSessionSetDeletionRequest struct {
	NodeID     ie.NodeID   //Node ID
	FQCSIDList []ie.FQCSID //FQ-CSID
}

func (t *PFCPSessionSetDeletionRequest) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(ie.NODE_ID_TYPE, tlv.Bytes(t.NodeID.Bytes())))
	for _, i := range t.FQCSIDList {
		fields = append(fields, tlv.TLVFrom(ie.FQ_CSID_TYPE, tlv.Bytes(i.Bytes())))
	}
	return tlv.TLV(ie.MSG_PFCP_SESSION_SET_DELETION_REQUEST_TYPE, fields...)
}

func (t *PFCPSessionSetDeletionRequest) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPSessionSetDeletionRequest{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case ie.NODE_ID_TYPE:
			err = de.UnmarshalValue(&t.NodeID)
		case ie.FQ_CSID_TYPE:
			var v ie.FQCSID
			if err = de.UnmarshalValue(&v); err == nil {
				t.FQCSIDList = append(t.FQCSIDList, v)
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

func (t *PFCPSessionSetDeletionRequest) UnmarshalTLV(typ uint32, buf []byte) (err error) {
	if typ != ie.MSG_PFCP_SESSION_SET_DELETION_REQUEST_TYPE {
		err = tlv.ErrCritical
		return
	}

	err = t.UnmarshalBinary(buf)
	return
}

func (t *PFCPSessionSetDeletionRequest) MarshalJSON() (buf []byte, err error) {
	buf, err = jsonMarshal(t)
	return
}

func (t *PFCPSessionSetDeletionRequest) UnmarshalJSON(buf []byte) (err error) {
	*t = PFCPSessionSetDeletionRequest{}
	err = jsonUnmarshal(buf, t)
	return
}

/* 'PFCP Session Set Deletion Response' data structure */
type PFCPSessionSetDeletionResponse struct {
	NodeID      ie.NodeID //Node ID
	Cause       uint8     //Cause
	OffendingIE *uint16   //Offending IE
}

func (t *PFCPSessionSetDeletionResponse) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(ie.NODE_ID_TYPE, tlv.Bytes(t.NodeID.Bytes())))
	fields = append(fields, tlv.TLVFrom(ie.CAUSE_TYPE, tlv.Bytes(pfcp.MarshalUint8(t.Cause))))
	if t.OffendingIE != nil {
		fields = append(fields, tlv.TLVFrom(ie.OFFENDING_IE_TYPE, tlv.Bytes(pfcp.MarshalUint16(*t.OffendingIE))))
	}
	return tlv.TLV(ie.MSG_PFCP_SESSION_SET_DELETION_RESPONSE_TYPE, fields...)
}

func (t *PFCPSessionSetDeletionResponse) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPSessionSetDeletionResponse{}
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

func (t *PFCPSessionSetDeletionResponse) UnmarshalTLV(typ uint32, buf []byte) (err error) {
	if typ != ie.MSG_PFCP_SESSION_SET_DELETION_RESPONSE_TYPE {
		err = tlv.ErrCritical
		return
	}

	err = t.UnmarshalBinary(buf)
	return
}

func (t *PFCPSessionSetDeletionResponse) MarshalJSON() (buf []byte, err error) {
	buf, err = jsonMarshal(t)
	return
}

func (t *PFCPSessionSetDeletionResponse) UnmarshalJSON(buf []byte) (err error) {
	*t = PFCPSessionSetDeletionResponse{}
	err = jsonUnmarshal(buf, t)
	return
}
