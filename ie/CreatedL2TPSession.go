package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Created L2TP Session' data structure */
type CreatedL2TPSession struct {
	DNSServerAddress  [][]byte //DNS Server Address
	NBNSServerAddress [][]byte //NBNS Server Address
	LNSAddress        []byte   //LNS Address
	Unknown           []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *CreatedL2TPSession) Field() tlv.Field {
	fields := []tlv.Field{}
	for _, ie := range t.DNSServerAddress {
		fields = append(fields, tlv.TLVFrom(DNS_SERVER_ADDRESS_TYPE, tlv.Bytes(ie)))
	}
	for _, ie := range t.NBNSServerAddress {
		fields = append(fields, tlv.TLVFrom(NBNS_SERVER_ADDRESS_TYPE, tlv.Bytes(ie)))
	}
	if len(t.LNSAddress) > 0 {
		fields = append(fields, tlv.TLVFrom(LNS_ADDRESS_TYPE, tlv.Bytes(t.LNSAddress)))
	}
	return tlv.TLV(CREATED_L2TP_SESSION_TYPE, fields...)
}

func (t *CreatedL2TPSession) UnmarshalBinary(buf []byte) (err error) {
	*t = CreatedL2TPSession{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [1]bool
	for _, de := range elements {
		switch de.Type {
		case DNS_SERVER_ADDRESS_TYPE:
			t.DNSServerAddress = append(t.DNSServerAddress, de.Value)
		case NBNS_SERVER_ADDRESS_TYPE:
			t.NBNSServerAddress = append(t.NBNSServerAddress, de.Value)
		case LNS_ADDRESS_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.LNSAddress = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return pfcp.Faulty(de.Type, err)
		}
	}
	return
}
