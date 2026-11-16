package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'PFCP Session Retention Information within PFCP Association Setup Request' data structure */
type PFCPSessionRetentionInformationwithinPFCPAssociationSetupRequest struct {
	CPPFCPEntityIPAddress [][]byte //CP PFCP Entity IP Address
	Unknown               []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid               []uint32 `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
}

func (t *PFCPSessionRetentionInformationwithinPFCPAssociationSetupRequest) Field() tlv.Field {
	fields := []tlv.Field{}
	for _, ie := range t.CPPFCPEntityIPAddress {
		fields = append(fields, tlv.TLVFrom(CP_PFCP_ENTITY_IP_ADDRESS_TYPE, tlv.Bytes(ie)))
	}
	return tlv.TLV(PFCP_SESSION_RETENTION_INFORMATION_WITHIN_PFCP_ASSOCIATION_SETUP_REQUEST_TYPE, fields...)
}

func (t *PFCPSessionRetentionInformationwithinPFCPAssociationSetupRequest) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPSessionRetentionInformationwithinPFCPAssociationSetupRequest{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	for _, de := range elements {
		opt := false
		switch de.Type {
		case CP_PFCP_ENTITY_IP_ADDRESS_TYPE:
			opt = true
			t.CPPFCPEntityIPAddress = append(t.CPPFCPEntityIPAddress, de.Value)
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
