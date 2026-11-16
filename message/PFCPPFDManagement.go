package message

import (
	"github.com/reogac/pfcp/v2"
	"github.com/reogac/pfcp/v2/ie"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'PFCP PFD Management Request' data structure */
type PFCPPFDManagementRequest struct {
	ApplicationIDsPFDs []ie.ApplicationIDsPFDs //Application ID's PFDs
	NodeID             *ie.NodeID              //Node ID
	Unknown            []uint32                `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid            []uint32                `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
}

func (t *PFCPPFDManagementRequest) Field() tlv.Field {
	fields := []tlv.Field{}
	for _, i := range t.ApplicationIDsPFDs {
		fields = append(fields, i.Field())
	}
	if t.NodeID != nil {
		fields = append(fields, tlv.TLVFrom(ie.NODE_ID_TYPE, tlv.Bytes(t.NodeID.Bytes())))
	}
	return tlv.TLV(ie.MSG_PFCP_PFD_MANAGEMENT_REQUEST_TYPE, fields...)
}

// Skipped is what decoding left out: IEs unknown or not defined here
// (TS 29.244 7.6.9), and optional ones discarded for a faulty value (7.6.8).
func (t *PFCPPFDManagementRequest) Skipped() (unknown, invalid []uint32) {
	return t.Unknown, t.Invalid
}

func (t *PFCPPFDManagementRequest) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPPFDManagementRequest{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [1]bool
	for _, de := range elements {
		opt := false
		switch de.Type {
		case ie.APPLICATION_ID_S_PFDS_TYPE:
			var v ie.ApplicationIDsPFDs
			if err = de.UnmarshalValue(&v); err == nil {
				t.ApplicationIDsPFDs = append(t.ApplicationIDsPFDs, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.NODE_ID_TYPE:
			opt = true
			if seen[0] {
				continue
			}
			seen[0] = true
			var v ie.NodeID
			if err = de.UnmarshalValue(&v); err == nil {
				t.NodeID = &v
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
	return
}

func (t *PFCPPFDManagementRequest) UnmarshalTLV(typ uint32, buf []byte) (err error) {
	if typ != ie.MSG_PFCP_PFD_MANAGEMENT_REQUEST_TYPE {
		err = tlv.ErrCritical
		return
	}

	err = t.UnmarshalBinary(buf)
	return
}

func (t *PFCPPFDManagementRequest) MarshalJSON() (buf []byte, err error) {
	buf, err = jsonMarshal(t)
	return
}

func (t *PFCPPFDManagementRequest) UnmarshalJSON(buf []byte) (err error) {
	*t = PFCPPFDManagementRequest{}
	err = jsonUnmarshal(buf, t)
	return
}

/* 'PFCP PFD Management Response' data structure */
type PFCPPFDManagementResponse struct {
	Cause       uint8      //Cause
	OffendingIE *uint16    //Offending IE
	NodeID      *ie.NodeID //Node ID
	Unknown     []uint32   `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid     []uint32   `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
}

func (t *PFCPPFDManagementResponse) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(ie.CAUSE_TYPE, tlv.Bytes(pfcp.MarshalUint8(t.Cause))))
	if t.OffendingIE != nil {
		fields = append(fields, tlv.TLVFrom(ie.OFFENDING_IE_TYPE, tlv.Bytes(pfcp.MarshalUint16(*t.OffendingIE))))
	}
	if t.NodeID != nil {
		fields = append(fields, tlv.TLVFrom(ie.NODE_ID_TYPE, tlv.Bytes(t.NodeID.Bytes())))
	}
	return tlv.TLV(ie.MSG_PFCP_PFD_MANAGEMENT_RESPONSE_TYPE, fields...)
}

// Skipped is what decoding left out: IEs unknown or not defined here
// (TS 29.244 7.6.9), and optional ones discarded for a faulty value (7.6.8).
func (t *PFCPPFDManagementResponse) Skipped() (unknown, invalid []uint32) {
	return t.Unknown, t.Invalid
}

func (t *PFCPPFDManagementResponse) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPPFDManagementResponse{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [3]bool
	for _, de := range elements {
		opt := false
		switch de.Type {
		case ie.CAUSE_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			err = pfcp.UnmarshalUint8(de.Value, &t.Cause)
		case ie.OFFENDING_IE_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			var v uint16
			if err = pfcp.UnmarshalUint16(de.Value, &v); err == nil {
				t.OffendingIE = &v
			}
		case ie.NODE_ID_TYPE:
			opt = true
			if seen[2] {
				continue
			}
			seen[2] = true
			var v ie.NodeID
			if err = de.UnmarshalValue(&v); err == nil {
				t.NodeID = &v
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
		return &pfcp.MissingIE{Type: ie.CAUSE_TYPE}
	}
	return
}

func (t *PFCPPFDManagementResponse) UnmarshalTLV(typ uint32, buf []byte) (err error) {
	if typ != ie.MSG_PFCP_PFD_MANAGEMENT_RESPONSE_TYPE {
		err = tlv.ErrCritical
		return
	}

	err = t.UnmarshalBinary(buf)
	return
}

func (t *PFCPPFDManagementResponse) MarshalJSON() (buf []byte, err error) {
	buf, err = jsonMarshal(t)
	return
}

func (t *PFCPPFDManagementResponse) UnmarshalJSON(buf []byte) (err error) {
	*t = PFCPPFDManagementResponse{}
	err = jsonUnmarshal(buf, t)
	return
}
