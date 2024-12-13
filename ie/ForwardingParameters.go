package ie

import (
	"github.com/reogac/pfcp"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

func NewForwardingParameters(dest uint8, in string, oc *OuterHeaderCreation, pid string) (far *ForwardingParameters) {
	far = &ForwardingParameters{}
	far.DestinationInterface = dest
	far.NetworkInstance = []byte(in)
	if oc != nil {
		far.OuterHeaderCreation = oc
	}
	far.ForwardingPolicy = []byte(pid)
	return
}

/* 'Forwarding Parameters' data structure */
type ForwardingParameters struct {
	DestinationInterface              uint8                //Destination Interface
	NetworkInstance                   []byte               //Network Instance
	RedirectInformation               []byte               //Redirect Information
	OuterHeaderCreation               *OuterHeaderCreation //Outer Header Creation
	TransportLevelMarking             []byte               //Transport Level Marking
	ForwardingPolicy                  []byte               //Forwarding Policy
	HeaderEnrichment                  []byte               //Header Enrichment
	LinkedTrafficEndpointID           []byte               //Traffic Endpoint ID
	Proxying                          []byte               //Proxying
	DestinationInterfaceType          []byte               //3GPP Interface Type
	DataNetworkAccessIdentifier       []byte               //Data Network Access Identifier
	IPAddressandPortNumberReplacement []byte               //IP Address and Port number Replacement
}

func (t *ForwardingParameters) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(DESTINATION_INTERFACE_TYPE, tlv.Bytes(pfcp.MarshalUint8(t.DestinationInterface))))
	if len(t.NetworkInstance) > 0 {
		fields = append(fields, tlv.TLVFrom(NETWORK_INSTANCE_TYPE, tlv.Bytes(t.NetworkInstance)))
	}
	if len(t.RedirectInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(REDIRECT_INFORMATION_TYPE, tlv.Bytes(t.RedirectInformation)))
	}
	if t.OuterHeaderCreation != nil {
		fields = append(fields, tlv.TLVFrom(OUTER_HEADER_CREATION_TYPE, tlv.Bytes(t.OuterHeaderCreation.Bytes())))
	}
	if len(t.TransportLevelMarking) > 0 {
		fields = append(fields, tlv.TLVFrom(TRANSPORT_LEVEL_MARKING_TYPE, tlv.Bytes(t.TransportLevelMarking)))
	}
	if len(t.ForwardingPolicy) > 0 {
		fields = append(fields, tlv.TLVFrom(FORWARDING_POLICY_TYPE, tlv.Bytes(t.ForwardingPolicy)))
	}
	if len(t.HeaderEnrichment) > 0 {
		fields = append(fields, tlv.TLVFrom(HEADER_ENRICHMENT_TYPE, tlv.Bytes(t.HeaderEnrichment)))
	}
	if len(t.LinkedTrafficEndpointID) > 0 {
		fields = append(fields, tlv.TLVFrom(TRAFFIC_ENDPOINT_ID_TYPE, tlv.Bytes(t.LinkedTrafficEndpointID)))
	}
	if len(t.Proxying) > 0 {
		fields = append(fields, tlv.TLVFrom(PROXYING_TYPE, tlv.Bytes(t.Proxying)))
	}
	if len(t.DestinationInterfaceType) > 0 {
		fields = append(fields, tlv.TLVFrom(_INTERFACE_TYPE_TYPE, tlv.Bytes(t.DestinationInterfaceType)))
	}
	if len(t.DataNetworkAccessIdentifier) > 0 {
		fields = append(fields, tlv.TLVFrom(DATA_NETWORK_ACCESS_IDENTIFIER_TYPE, tlv.Bytes(t.DataNetworkAccessIdentifier)))
	}
	if len(t.IPAddressandPortNumberReplacement) > 0 {
		fields = append(fields, tlv.TLVFrom(IP_ADDRESS_AND_PORT_NUMBER_REPLACEMENT_TYPE, tlv.Bytes(t.IPAddressandPortNumberReplacement)))
	}
	return tlv.TLV(FORWARDING_PARAMETERS_TYPE, fields...)
}

func (t *ForwardingParameters) UnmarshalBinary(buf []byte) (err error) {
	*t = ForwardingParameters{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case DESTINATION_INTERFACE_TYPE:
			err = pfcp.UnmarshalUint8(de.Value, &t.DestinationInterface)
		case NETWORK_INSTANCE_TYPE:
			t.NetworkInstance = de.Value
		case REDIRECT_INFORMATION_TYPE:
			t.RedirectInformation = de.Value
		case OUTER_HEADER_CREATION_TYPE:
			var v OuterHeaderCreation
			if err = de.UnmarshalValue(&v); err == nil {
				t.OuterHeaderCreation = &v
			}
		case TRANSPORT_LEVEL_MARKING_TYPE:
			t.TransportLevelMarking = de.Value
		case FORWARDING_POLICY_TYPE:
			t.ForwardingPolicy = de.Value
		case HEADER_ENRICHMENT_TYPE:
			t.HeaderEnrichment = de.Value
		case TRAFFIC_ENDPOINT_ID_TYPE:
			t.LinkedTrafficEndpointID = de.Value
		case PROXYING_TYPE:
			t.Proxying = de.Value
		case _INTERFACE_TYPE_TYPE:
			t.DestinationInterfaceType = de.Value
		case DATA_NETWORK_ACCESS_IDENTIFIER_TYPE:
			t.DataNetworkAccessIdentifier = de.Value
		case IP_ADDRESS_AND_PORT_NUMBER_REPLACEMENT_TYPE:
			t.IPAddressandPortNumberReplacement = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
