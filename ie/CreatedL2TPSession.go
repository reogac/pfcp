package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Created L2TP Session' data structure */
type CreatedL2TPSession struct {
	DNSServerAddress  []byte //DNS Server Address
	NBNSServerAddress []byte //NBNS Server Address
	LNSAddress        []byte //LNS Address
}

func (t *CreatedL2TPSession) Field() tlv.Field {
	fields := []tlv.Field{}
	if len(t.DNSServerAddress) > 0 {
		fields = append(fields, tlv.TLVFrom(DNS_SERVER_ADDRESS_TYPE, tlv.Bytes(t.DNSServerAddress)))
	}
	if len(t.NBNSServerAddress) > 0 {
		fields = append(fields, tlv.TLVFrom(NBNS_SERVER_ADDRESS_TYPE, tlv.Bytes(t.NBNSServerAddress)))
	}
	if len(t.LNSAddress) > 0 {
		fields = append(fields, tlv.TLVFrom(LNS_ADDRESS_TYPE, tlv.Bytes(t.LNSAddress)))
	}
	return tlv.TLV(CREATED_L2TP_SESSION_TYPE, fields...)
}

func (t *CreatedL2TPSession) UnmarshalBinary(buf []byte) (err error) {
	*t = CreatedL2TPSession{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case DNS_SERVER_ADDRESS_TYPE:
			t.DNSServerAddress = de.Value
		case NBNS_SERVER_ADDRESS_TYPE:
			t.NBNSServerAddress = de.Value
		case LNS_ADDRESS_TYPE:
			t.LNSAddress = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
