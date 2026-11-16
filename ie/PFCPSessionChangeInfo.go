package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'PFCP Session Change Info' data structure */
type PFCPSessionChangeInfo struct {
	PGWCSMFFQCSID               []FQCSID //FQ-CSID
	GroupId                     [][]byte //Group ID
	CPIPAddress                 [][]byte //CP IP Address
	AlternativeSMFPGWCIPAddress []byte   //Alternative SMF IP Address
	Unknown                     []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid                     []uint32 `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
}

func (t *PFCPSessionChangeInfo) Field() tlv.Field {
	fields := []tlv.Field{}
	for _, ie := range t.PGWCSMFFQCSID {
		fields = append(fields, tlv.TLVFrom(FQ_CSID_TYPE, tlv.Bytes(ie.Bytes())))
	}
	for _, ie := range t.GroupId {
		fields = append(fields, tlv.TLVFrom(GROUP_ID_TYPE, tlv.Bytes(ie)))
	}
	for _, ie := range t.CPIPAddress {
		fields = append(fields, tlv.TLVFrom(CP_IP_ADDRESS_TYPE, tlv.Bytes(ie)))
	}
	fields = append(fields, tlv.TLVFrom(ALTERNATIVE_SMF_IP_ADDRESS_TYPE, tlv.Bytes(t.AlternativeSMFPGWCIPAddress)))
	return tlv.TLV(PFCP_SESSION_CHANGE_INFO_TYPE, fields...)
}

func (t *PFCPSessionChangeInfo) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPSessionChangeInfo{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [1]bool
	for _, de := range elements {
		opt := false
		switch de.Type {
		case FQ_CSID_TYPE:
			var v FQCSID
			if err = de.UnmarshalValue(&v); err == nil {
				t.PGWCSMFFQCSID = append(t.PGWCSMFFQCSID, v)
			}
		case GROUP_ID_TYPE:
			t.GroupId = append(t.GroupId, de.Value)
		case CP_IP_ADDRESS_TYPE:
			t.CPIPAddress = append(t.CPIPAddress, de.Value)
		case ALTERNATIVE_SMF_IP_ADDRESS_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.AlternativeSMFPGWCIPAddress = de.Value
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
		return &pfcp.MissingIE{Type: ALTERNATIVE_SMF_IP_ADDRESS_TYPE}
	}
	return
}
