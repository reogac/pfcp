package message

import (
	"github.com/reogac/pfcp/v2"
	"github.com/reogac/pfcp/v2/ie"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'PFCP Association Release Request' data structure */
type PFCPAssociationReleaseRequest struct {
	NodeID  ie.NodeID //Node ID
	Unknown []uint32  `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *PFCPAssociationReleaseRequest) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(ie.NODE_ID_TYPE, tlv.Bytes(t.NodeID.Bytes())))
	return tlv.TLV(ie.MSG_PFCP_ASSOCIATION_RELEASE_REQUEST_TYPE, fields...)
}

func (t *PFCPAssociationReleaseRequest) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPAssociationReleaseRequest{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [1]bool
	for _, de := range elements {
		switch de.Type {
		case ie.NODE_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			if err = de.UnmarshalValue(&t.NodeID); err == nil {
			}
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return pfcp.Faulty(de.Type, err)
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: ie.NODE_ID_TYPE}
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
	NodeID  ie.NodeID //Node ID
	Cause   uint8     //Cause
	Unknown []uint32  `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *PFCPAssociationReleaseResponse) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(ie.NODE_ID_TYPE, tlv.Bytes(t.NodeID.Bytes())))
	fields = append(fields, tlv.TLVFrom(ie.CAUSE_TYPE, tlv.Bytes(pfcp.MarshalUint8(t.Cause))))
	return tlv.TLV(ie.MSG_PFCP_ASSOCIATION_RELEASE_RESPONSE_TYPE, fields...)
}

func (t *PFCPAssociationReleaseResponse) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPAssociationReleaseResponse{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [2]bool
	for _, de := range elements {
		switch de.Type {
		case ie.NODE_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			if err = de.UnmarshalValue(&t.NodeID); err == nil {
			}
		case ie.CAUSE_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			err = pfcp.UnmarshalUint8(de.Value, &t.Cause)
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return pfcp.Faulty(de.Type, err)
		}
	}
	if !seen[1] {
		return &pfcp.MissingIE{Type: ie.CAUSE_TYPE}
	}
	if t.Cause < ie.CauseRequestRejected {
		if !seen[0] {
			return &pfcp.MissingIE{Type: ie.NODE_ID_TYPE}
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
