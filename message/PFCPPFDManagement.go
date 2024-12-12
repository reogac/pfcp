package message

import (
	"etrib5gc/pfcp"
	"etrib5gc/pfcp/ie"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'PFCP PFD Management Request' data structure */
type PFCPPFDManagementRequest struct {
	ApplicationIDsPFDs *ie.ApplicationIDsPFDs //Application ID's PFDs
	NodeID             *ie.NodeID             //Node ID
}

func (t *PFCPPFDManagementRequest) Field() tlv.Field {
	fields := []tlv.Field{}
	if t.ApplicationIDsPFDs != nil {
		fields = append(fields, t.ApplicationIDsPFDs.Field())
	}
	if t.NodeID != nil {
		fields = append(fields, tlv.TLVFrom(ie.NODE_ID_TYPE, tlv.Bytes(t.NodeID.Bytes())))
	}
	return tlv.TLV(ie.MSG_PFCP_PFD_MANAGEMENT_REQUEST_TYPE, fields...)
}

func (t *PFCPPFDManagementRequest) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPPFDManagementRequest{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case ie.APPLICATION_ID_S_PFDS_TYPE:
			var v ie.ApplicationIDsPFDs
			if err = de.UnmarshalValue(&v); err == nil {
				t.ApplicationIDsPFDs = &v
			}
		case ie.NODE_ID_TYPE:
			var v ie.NodeID
			if err = de.UnmarshalValue(&v); err == nil {
				t.NodeID = &v
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

func (t *PFCPPFDManagementResponse) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPPFDManagementResponse{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case ie.CAUSE_TYPE:
			err = pfcp.UnmarshalUint8(de.Value, &t.Cause)
		case ie.OFFENDING_IE_TYPE:
			var v uint16
			if err = pfcp.UnmarshalUint16(de.Value, &v); err == nil {
				t.OffendingIE = &v
			}
		case ie.NODE_ID_TYPE:
			var v ie.NodeID
			if err = de.UnmarshalValue(&v); err == nil {
				t.NodeID = &v
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
