package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Update Forwarding Parameters' data structure */
type UpdateForwardingParameters struct {
	DestinationInterface              *uint8               //Destination Interface
	NetworkInstance                   []byte               //Network Instance
	RedirectInformation               []byte               //Redirect Information
	OuterHeaderCreation               *OuterHeaderCreation //Outer Header Creation
	TransportLevelMarking             []byte               //Transport Level Marking
	ForwardingPolicy                  []byte               //Forwarding Policy
	HeaderEnrichment                  []byte               //Header Enrichment
	PFCPSMReqFlags                    *uint8               //PFCPSMReq-Flags
	LinkedTrafficEndpointID           []byte               //Traffic Endpoint ID
	DestinationInterfaceType          []byte               //3GPP Interface Type
	DataNetworkAccessIdentifier       []byte               //Data Network Access Identifier
	IPAddressandPortNumberReplacement []byte               //IP Address and Port number Replacement
	Unknown                           []uint32             `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *UpdateForwardingParameters) Field() tlv.Field {
	fields := []tlv.Field{}
	if t.DestinationInterface != nil {
		fields = append(fields, tlv.TLVFrom(DESTINATION_INTERFACE_TYPE, tlv.Bytes(pfcp.MarshalUint8(*t.DestinationInterface))))
	}
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
	if t.PFCPSMReqFlags != nil {
		fields = append(fields, tlv.TLVFrom(PFCPSMREQ_FLAGS_TYPE, tlv.Bytes(pfcp.MarshalUint8(*t.PFCPSMReqFlags))))
	}
	if len(t.LinkedTrafficEndpointID) > 0 {
		fields = append(fields, tlv.TLVFrom(TRAFFIC_ENDPOINT_ID_TYPE, tlv.Bytes(t.LinkedTrafficEndpointID)))
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
	return tlv.TLV(UPDATE_FORWARDING_PARAMETERS_TYPE, fields...)
}

func (t *UpdateForwardingParameters) UnmarshalBinary(buf []byte) (err error) {
	*t = UpdateForwardingParameters{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	var seen [12]bool
	for _, de := range elements {
		switch de.Type {
		case DESTINATION_INTERFACE_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			var v uint8
			if err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &v); err == nil {
				t.DestinationInterface = &v
			}
		case NETWORK_INSTANCE_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			t.NetworkInstance = de.Value
		case REDIRECT_INFORMATION_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			t.RedirectInformation = de.Value
		case OUTER_HEADER_CREATION_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			var v OuterHeaderCreation
			if err = de.UnmarshalValue(&v); err == nil {
				t.OuterHeaderCreation = &v
			}
		case TRANSPORT_LEVEL_MARKING_TYPE:
			if seen[4] {
				continue
			}
			seen[4] = true
			t.TransportLevelMarking = de.Value
		case FORWARDING_POLICY_TYPE:
			if seen[5] {
				continue
			}
			seen[5] = true
			t.ForwardingPolicy = de.Value
		case HEADER_ENRICHMENT_TYPE:
			if seen[6] {
				continue
			}
			seen[6] = true
			t.HeaderEnrichment = de.Value
		case PFCPSMREQ_FLAGS_TYPE:
			if seen[7] {
				continue
			}
			seen[7] = true
			var v uint8
			if err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &v); err == nil {
				t.PFCPSMReqFlags = &v
			}
		case TRAFFIC_ENDPOINT_ID_TYPE:
			if seen[8] {
				continue
			}
			seen[8] = true
			t.LinkedTrafficEndpointID = de.Value
		case _INTERFACE_TYPE_TYPE:
			if seen[9] {
				continue
			}
			seen[9] = true
			t.DestinationInterfaceType = de.Value
		case DATA_NETWORK_ACCESS_IDENTIFIER_TYPE:
			if seen[10] {
				continue
			}
			seen[10] = true
			t.DataNetworkAccessIdentifier = de.Value
		case IP_ADDRESS_AND_PORT_NUMBER_REPLACEMENT_TYPE:
			if seen[11] {
				continue
			}
			seen[11] = true
			t.IPAddressandPortNumberReplacement = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	return
}
