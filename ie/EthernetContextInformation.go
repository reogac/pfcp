package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Ethernet Context Information' data structure */
type EthernetContextInformation struct {
	MACAddressesDetected [][]byte //MAC Addresses Detected
	Unknown              []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid              []uint32 `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
}

func (t *EthernetContextInformation) Field() tlv.Field {
	fields := []tlv.Field{}
	for _, ie := range t.MACAddressesDetected {
		fields = append(fields, tlv.TLVFrom(MAC_ADDRESSES_DETECTED_TYPE, tlv.Bytes(ie)))
	}
	return tlv.TLV(ETHERNET_CONTEXT_INFORMATION_TYPE, fields...)
}

func (t *EthernetContextInformation) UnmarshalBinary(buf []byte) (err error) {
	*t = EthernetContextInformation{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	for _, de := range elements {
		opt := false
		switch de.Type {
		case MAC_ADDRESSES_DETECTED_TYPE:
			t.MACAddressesDetected = append(t.MACAddressesDetected, de.Value)
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
	if len(t.MACAddressesDetected) == 0 {
		return &pfcp.MissingIE{Type: MAC_ADDRESSES_DETECTED_TYPE}
	}
	return
}
