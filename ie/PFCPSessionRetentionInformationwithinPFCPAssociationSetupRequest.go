package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'PFCP Session Retention Information within PFCP Association Setup Request' data structure */
type PFCPSessionRetentionInformationwithinPFCPAssociationSetupRequest struct {
	CPPFCPEntityIPAddress []byte //CP PFCP Entity IP Address
}

func (t *PFCPSessionRetentionInformationwithinPFCPAssociationSetupRequest) Field() tlv.Field {
	fields := []tlv.Field{}
	if len(t.CPPFCPEntityIPAddress) > 0 {
		fields = append(fields, tlv.TLVFrom(CP_PFCP_ENTITY_IP_ADDRESS_TYPE, tlv.Bytes(t.CPPFCPEntityIPAddress)))
	}
	return tlv.TLV(PFCP_SESSION_RETENTION_INFORMATION_WITHIN_PFCP_ASSOCIATION_SETUP_REQUEST_TYPE, fields...)
}

func (t *PFCPSessionRetentionInformationwithinPFCPAssociationSetupRequest) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPSessionRetentionInformationwithinPFCPAssociationSetupRequest{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case CP_PFCP_ENTITY_IP_ADDRESS_TYPE:
			t.CPPFCPEntityIPAddress = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
