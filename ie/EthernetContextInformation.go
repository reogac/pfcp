package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Ethernet Context Information' data structure */
type EthernetContextInformation struct {
	MACAddressesDetected []byte //MAC Addresses Detected
}

func (t *EthernetContextInformation) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(MAC_ADDRESSES_DETECTED_TYPE, tlv.Bytes(t.MACAddressesDetected)))
	return tlv.TLV(ETHERNET_CONTEXT_INFORMATION_TYPE, fields...)
}

func (t *EthernetContextInformation) UnmarshalBinary(buf []byte) (err error) {
	*t = EthernetContextInformation{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case MAC_ADDRESSES_DETECTED_TYPE:
			t.MACAddressesDetected = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
