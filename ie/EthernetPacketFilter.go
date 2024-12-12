package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
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
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case ETHERNET_FILTER_ID_TYPE:
			t.EthernetFilterID = de.Value
		case ETHERNET_FILTER_PROPERTIES_TYPE:
			t.EthernetFilterProperties = de.Value
		case MAC_ADDRESS_TYPE:
			t.MACaddress = de.Value
		case ETHERTYPE_TYPE:
			t.Ethertype = de.Value
		case C_TAG_TYPE:
			t.CTAG = de.Value
		case S_TAG_TYPE:
			t.STAG = de.Value
		case SDF_FILTER_TYPE:
			var v SDFFilter
			if err = de.UnmarshalValue(&v); err == nil {
				t.SDFFilter = append(t.SDFFilter, v)
			}
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
