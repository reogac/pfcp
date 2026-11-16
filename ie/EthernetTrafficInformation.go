package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Ethernet Traffic Information' data structure */
type EthernetTrafficInformation struct {
	MACAddressesDetected [][]byte //MAC Addresses Detected
	MACAddressesRemoved  [][]byte //MAC Addresses Removed
	Unknown              []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *EthernetTrafficInformation) Field() tlv.Field {
	fields := []tlv.Field{}
	for _, ie := range t.MACAddressesDetected {
		fields = append(fields, tlv.TLVFrom(MAC_ADDRESSES_DETECTED_TYPE, tlv.Bytes(ie)))
	}
	for _, ie := range t.MACAddressesRemoved {
		fields = append(fields, tlv.TLVFrom(MAC_ADDRESSES_REMOVED_TYPE, tlv.Bytes(ie)))
	}
	return tlv.TLV(ETHERNET_TRAFFIC_INFORMATION_TYPE, fields...)
}

func (t *EthernetTrafficInformation) UnmarshalBinary(buf []byte) (err error) {
	*t = EthernetTrafficInformation{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	for _, de := range elements {
		switch de.Type {
		case MAC_ADDRESSES_DETECTED_TYPE:
			t.MACAddressesDetected = append(t.MACAddressesDetected, de.Value)
		case MAC_ADDRESSES_REMOVED_TYPE:
			t.MACAddressesRemoved = append(t.MACAddressesRemoved, de.Value)
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return pfcp.Faulty(de.Type, err)
		}
	}
	return
}
