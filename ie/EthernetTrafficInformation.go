package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Ethernet Traffic Information' data structure */
type EthernetTrafficInformation struct {
	MACAddressesDetected []byte //MAC Addresses Detected
	MACAddressesRemoved  []byte //MAC Addresses Removed
}

func (t *EthernetTrafficInformation) Field() tlv.Field {
	fields := []tlv.Field{}
	if len(t.MACAddressesDetected) > 0 {
		fields = append(fields, tlv.TLVFrom(MAC_ADDRESSES_DETECTED_TYPE, tlv.Bytes(t.MACAddressesDetected)))
	}
	if len(t.MACAddressesRemoved) > 0 {
		fields = append(fields, tlv.TLVFrom(MAC_ADDRESSES_REMOVED_TYPE, tlv.Bytes(t.MACAddressesRemoved)))
	}
	return tlv.TLV(ETHERNET_TRAFFIC_INFORMATION_TYPE, fields...)
}

func (t *EthernetTrafficInformation) UnmarshalBinary(buf []byte) (err error) {
	*t = EthernetTrafficInformation{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case MAC_ADDRESSES_DETECTED_TYPE:
			t.MACAddressesDetected = de.Value
		case MAC_ADDRESSES_REMOVED_TYPE:
			t.MACAddressesRemoved = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
