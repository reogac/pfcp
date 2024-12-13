package message

import (
	"github.com/reogac/pfcp"
	"github.com/reogac/pfcp/ie"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'PFCP Association Release Request' data structure */
type PFCPAssociationReleaseRequest struct {
	NodeID ie.NodeID //Node ID
}

func (t *PFCPAssociationReleaseRequest) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(ie.NODE_ID_TYPE, tlv.Bytes(t.NodeID.Bytes())))
	return tlv.TLV(ie.MSG_PFCP_ASSOCIATION_RELEASE_REQUEST_TYPE, fields...)
}

func (t *PFCPAssociationReleaseRequest) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPAssociationReleaseRequest{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case ie.NODE_ID_TYPE:
			err = de.UnmarshalValue(&t.NodeID)
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}

func (t *PFCPAssociationReleaseRequest) UnmarshalTLV(typ uint32, buf []byte) (err error) {
	if typ != ie.MSG_PFCP_ASSOCIATION_RELEASE_REQUEST_TYPE {
		err = tlv.ErrCritical
		return
	}

	err = t.UnmarshalBinary(buf)
	return
}

func (t *PFCPAssociationReleaseRequest) MarshalJSON() (buf []byte, err error) {
	buf, err = jsonMarshal(t)
	return
}

func (t *PFCPAssociationReleaseRequest) UnmarshalJSON(buf []byte) (err error) {
	*t = PFCPAssociationReleaseRequest{}
	err = jsonUnmarshal(buf, t)
	return
}

/* 'PFCP Association Release Response' data structure */
type PFCPAssociationReleaseResponse struct {
	NodeID ie.NodeID //Node ID
	Cause  uint8     //Cause
}

func (t *PFCPAssociationReleaseResponse) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(ie.NODE_ID_TYPE, tlv.Bytes(t.NodeID.Bytes())))
	fields = append(fields, tlv.TLVFrom(ie.CAUSE_TYPE, tlv.Bytes(pfcp.MarshalUint8(t.Cause))))
	return tlv.TLV(ie.MSG_PFCP_ASSOCIATION_RELEASE_RESPONSE_TYPE, fields...)
}

func (t *PFCPAssociationReleaseResponse) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPAssociationReleaseResponse{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case ie.NODE_ID_TYPE:
			err = de.UnmarshalValue(&t.NodeID)
		case ie.CAUSE_TYPE:
			err = pfcp.UnmarshalUint8(de.Value, &t.Cause)
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}

func (t *PFCPAssociationReleaseResponse) UnmarshalTLV(typ uint32, buf []byte) (err error) {
	if typ != ie.MSG_PFCP_ASSOCIATION_RELEASE_RESPONSE_TYPE {
		err = tlv.ErrCritical
		return
	}

	err = t.UnmarshalBinary(buf)
	return
}

func (t *PFCPAssociationReleaseResponse) MarshalJSON() (buf []byte, err error) {
	buf, err = jsonMarshal(t)
	return
}

func (t *PFCPAssociationReleaseResponse) UnmarshalJSON(buf []byte) (err error) {
	*t = PFCPAssociationReleaseResponse{}
	err = jsonUnmarshal(buf, t)
	return
}
