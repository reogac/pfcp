package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'IP Multicast Addressing Info within PFCP Session Establishment Request' data structure */
type IPMulticastAddressingInfowithinPFCPSessionEstablishmentRequest struct {
	IPMulticastAddress []byte //IP Multicast Address
	SourceIPAddress    []byte //Source IP Address
}

func (t *IPMulticastAddressingInfowithinPFCPSessionEstablishmentRequest) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(IP_MULTICAST_ADDRESS_TYPE, tlv.Bytes(t.IPMulticastAddress)))
	if len(t.SourceIPAddress) > 0 {
		fields = append(fields, tlv.TLVFrom(SOURCE_IP_ADDRESS_TYPE, tlv.Bytes(t.SourceIPAddress)))
	}
	return tlv.TLV(IP_MULTICAST_ADDRESSING_INFO_WITHIN_PFCP_SESSION_ESTABLISHMENT_REQUEST_TYPE, fields...)
}

func (t *IPMulticastAddressingInfowithinPFCPSessionEstablishmentRequest) UnmarshalBinary(buf []byte) (err error) {
	*t = IPMulticastAddressingInfowithinPFCPSessionEstablishmentRequest{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case IP_MULTICAST_ADDRESS_TYPE:
			t.IPMulticastAddress = de.Value
		case SOURCE_IP_ADDRESS_TYPE:
			t.SourceIPAddress = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
