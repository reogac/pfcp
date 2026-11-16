package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Ethernet Packet Filter' data structure */
type EthernetPacketFilter struct {
	EthernetFilterID         []byte      //Ethernet Filter ID
	EthernetFilterProperties []byte      //Ethernet Filter Properties
	MACaddress               []byte      //MAC address
	Ethertype                []byte      //Ethertype
	CTAG                     []byte      //C-TAG
	STAG                     []byte      //S-TAG
	SDFFilter                []SDFFilter //SDF Filter
	Unknown                  []uint32    `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *EthernetPacketFilter) Field() tlv.Field {
	fields := []tlv.Field{}
	if len(t.EthernetFilterID) > 0 {
		fields = append(fields, tlv.TLVFrom(ETHERNET_FILTER_ID_TYPE, tlv.Bytes(t.EthernetFilterID)))
	}
	if len(t.EthernetFilterProperties) > 0 {
		fields = append(fields, tlv.TLVFrom(ETHERNET_FILTER_PROPERTIES_TYPE, tlv.Bytes(t.EthernetFilterProperties)))
	}
	if len(t.MACaddress) > 0 {
		fields = append(fields, tlv.TLVFrom(MAC_ADDRESS_TYPE, tlv.Bytes(t.MACaddress)))
	}
	if len(t.Ethertype) > 0 {
		fields = append(fields, tlv.TLVFrom(ETHERTYPE_TYPE, tlv.Bytes(t.Ethertype)))
	}
	if len(t.CTAG) > 0 {
		fields = append(fields, tlv.TLVFrom(C_TAG_TYPE, tlv.Bytes(t.CTAG)))
	}
	if len(t.STAG) > 0 {
		fields = append(fields, tlv.TLVFrom(S_TAG_TYPE, tlv.Bytes(t.STAG)))
	}
	for _, ie := range t.SDFFilter {
		fields = append(fields, tlv.TLVFrom(SDF_FILTER_TYPE, tlv.Bytes(ie.Bytes())))
	}
	return tlv.TLV(ETHERNET_PACKET_FILTER_TYPE, fields...)
}

func (t *EthernetPacketFilter) UnmarshalBinary(buf []byte) (err error) {
	*t = EthernetPacketFilter{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [6]bool
	for _, de := range elements {
		switch de.Type {
		case ETHERNET_FILTER_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.EthernetFilterID = de.Value
		case ETHERNET_FILTER_PROPERTIES_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			t.EthernetFilterProperties = de.Value
		case MAC_ADDRESS_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			t.MACaddress = de.Value
		case ETHERTYPE_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			t.Ethertype = de.Value
		case C_TAG_TYPE:
			if seen[4] {
				continue
			}
			seen[4] = true
			t.CTAG = de.Value
		case S_TAG_TYPE:
			if seen[5] {
				continue
			}
			seen[5] = true
			t.STAG = de.Value
		case SDF_FILTER_TYPE:
			var v SDFFilter
			if err = de.UnmarshalValue(&v); err == nil {
				t.SDFFilter = append(t.SDFFilter, v)
			}
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return pfcp.Faulty(de.Type, err)
		}
	}
	return
}
