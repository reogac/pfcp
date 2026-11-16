package message

import (
	"github.com/reogac/pfcp/v2"
	"github.com/reogac/pfcp/v2/ie"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'PFCP Session Set Modification Request' data structure */
type PFCPSessionSetModificationRequest struct {
	NodeID                ie.NodeID                  //Node ID
	PFCPSessionChangeInfo []ie.PFCPSessionChangeInfo //PFCP Session Change Info
	Unknown               []uint32                   `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid               []uint32                   `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
}

func (t *PFCPSessionSetModificationRequest) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(ie.NODE_ID_TYPE, tlv.Bytes(t.NodeID.Bytes())))
	for _, i := range t.PFCPSessionChangeInfo {
		fields = append(fields, i.Field())
	}
	return tlv.TLV(ie.MSG_PFCP_SESSION_SET_MODIFICATION_REQUEST_TYPE, fields...)
}

// Skipped is what decoding left out: IEs unknown or not defined here
// (TS 29.244 7.6.9), and optional ones discarded for a faulty value (7.6.8).
func (t *PFCPSessionSetModificationRequest) Skipped() (unknown, invalid []uint32) {
	return t.Unknown, t.Invalid
}

func (t *PFCPSessionSetModificationRequest) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPSessionSetModificationRequest{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [1]bool
	for _, de := range elements {
		opt := false
		switch de.Type {
		case ie.NODE_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			if err = de.UnmarshalValue(&t.NodeID); err == nil {
			}
		case ie.PFCP_SESSION_CHANGE_INFO_TYPE:
			var v ie.PFCPSessionChangeInfo
			if err = de.UnmarshalValue(&v); err == nil {
				t.PFCPSessionChangeInfo = append(t.PFCPSessionChangeInfo, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
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
		return &pfcp.MissingIE{Type: ie.NODE_ID_TYPE}
	}
	if len(t.PFCPSessionChangeInfo) == 0 {
		return &pfcp.MissingIE{Type: ie.PFCP_SESSION_CHANGE_INFO_TYPE}
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
	Unknown     []uint32  `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid     []uint32  `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
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

// Skipped is what decoding left out: IEs unknown or not defined here
// (TS 29.244 7.6.9), and optional ones discarded for a faulty value (7.6.8).
func (t *PFCPSessionSetModificationResponse) Skipped() (unknown, invalid []uint32) {
	return t.Unknown, t.Invalid
}

func (t *PFCPSessionSetModificationResponse) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPSessionSetModificationResponse{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [3]bool
	for _, de := range elements {
		opt := false
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
		case ie.OFFENDING_IE_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			var v uint16
			if err = pfcp.UnmarshalUint16(de.Value, &v); err == nil {
				t.OffendingIE = &v
			}
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
