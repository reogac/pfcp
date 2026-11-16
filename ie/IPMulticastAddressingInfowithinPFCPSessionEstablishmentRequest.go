package ie

import (
	"github.com/reogac/pfcp"

	"github.com/reogac/pfcp/tlv"
)

/* 'IP Multicast Addressing Info within PFCP Session Establishment Request' data structure */
type IPMulticastAddressingInfowithinPFCPSessionEstablishmentRequest struct {
	IPMulticastAddress []byte   //IP Multicast Address
	SourceIPAddress    [][]byte //Source IP Address
	Unknown            []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *IPMulticastAddressingInfowithinPFCPSessionEstablishmentRequest) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(IP_MULTICAST_ADDRESS_TYPE, tlv.Bytes(t.IPMulticastAddress)))
	for _, ie := range t.SourceIPAddress {
		fields = append(fields, tlv.TLVFrom(SOURCE_IP_ADDRESS_TYPE, tlv.Bytes(ie)))
	}
	return tlv.TLV(IP_MULTICAST_ADDRESSING_INFO_WITHIN_PFCP_SESSION_ESTABLISHMENT_REQUEST_TYPE, fields...)
}

func (t *IPMulticastAddressingInfowithinPFCPSessionEstablishmentRequest) UnmarshalBinary(buf []byte) (err error) {
	*t = IPMulticastAddressingInfowithinPFCPSessionEstablishmentRequest{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	var seen [1]bool
	for _, de := range elements {
		switch de.Type {
		case IP_MULTICAST_ADDRESS_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.IPMulticastAddress = de.Value
		case SOURCE_IP_ADDRESS_TYPE:
			t.SourceIPAddress = append(t.SourceIPAddress, de.Value)
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: IP_MULTICAST_ADDRESS_TYPE}
	}
	return
}
