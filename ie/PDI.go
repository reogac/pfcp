package ie

import (
	"github.com/reogac/pfcp"

	"github.com/reogac/pfcp/tlv"
)

/* 'PDI' data structure */
type PDI struct {
	SourceInterface                          uint8                                                            //Source Interface
	LocalFTEID                               *FTEID                                                           //F-TEID
	LocalIngressTunnel                       []byte                                                           //Local Ingress Tunnel
	NetworkInstance                          []byte                                                           //Network Instance
	RedundantTransmissionDetectionParameters *RedundantTransmissionParameters                                 //Redundant Transmission Parameters
	UEIPaddress                              []UEIPAddress                                                    //UE IP Address
	TrafficEndpointID                        [][]byte                                                         //Traffic Endpoint ID
	SDFFilter                                []SDFFilter                                                      //SDF Filter
	ApplicationID                            *ApplicationID                                                   //Application ID
	EthernetPDUSessionInformation            []byte                                                           //Ethernet PDU Session Information
	EthernetPacketFilter                     []EthernetPacketFilter                                           //Ethernet Packet Filter
	QFI                                      []uint8                                                          //QFI
	FramedRoute                              [][]byte                                                         //Framed-Route
	FramedRouting                            []byte                                                           //Framed-Routing
	FramedIPv6Route                          [][]byte                                                         //Framed-IPv6-Route
	SourceInterfaceType                      []byte                                                           //3GPP Interface Type
	IPMulticastAddressingInfo                []IPMulticastAddressingInfowithinPFCPSessionEstablishmentRequest //IP Multicast Addressing Info within PFCP Session Establishment Request
	DNSQueryFilter                           [][]byte                                                         //DNS Query Filter
	MBSSessionIdentifier                     []byte                                                           //MBS Session Identifier
	AreaSessionID                            []byte                                                           //Area Session ID
	Unknown                                  []uint32                                                         `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *PDI) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(SOURCE_INTERFACE_TYPE, tlv.Bytes(pfcp.MarshalUint8(t.SourceInterface))))
	if t.LocalFTEID != nil {
		fields = append(fields, tlv.TLVFrom(F_TEID_TYPE, tlv.Bytes(t.LocalFTEID.Bytes())))
	}
	if len(t.LocalIngressTunnel) > 0 {
		fields = append(fields, tlv.TLVFrom(LOCAL_INGRESS_TUNNEL_TYPE, tlv.Bytes(t.LocalIngressTunnel)))
	}
	if len(t.NetworkInstance) > 0 {
		fields = append(fields, tlv.TLVFrom(NETWORK_INSTANCE_TYPE, tlv.Bytes(t.NetworkInstance)))
	}
	if t.RedundantTransmissionDetectionParameters != nil {
		fields = append(fields, t.RedundantTransmissionDetectionParameters.Field())
	}
	for _, ie := range t.UEIPaddress {
		fields = append(fields, tlv.TLVFrom(UE_IP_ADDRESS_TYPE, tlv.Bytes(ie.Bytes())))
	}
	for _, ie := range t.TrafficEndpointID {
		fields = append(fields, tlv.TLVFrom(TRAFFIC_ENDPOINT_ID_TYPE, tlv.Bytes(ie)))
	}
	for _, ie := range t.SDFFilter {
		fields = append(fields, tlv.TLVFrom(SDF_FILTER_TYPE, tlv.Bytes(ie.Bytes())))
	}
	if t.ApplicationID != nil {
		fields = append(fields, tlv.TLVFrom(APPLICATION_ID_TYPE, tlv.Bytes(t.ApplicationID.Bytes())))
	}
	if len(t.EthernetPDUSessionInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(ETHERNET_PDU_SESSION_INFORMATION_TYPE, tlv.Bytes(t.EthernetPDUSessionInformation)))
	}
	for _, ie := range t.EthernetPacketFilter {
		fields = append(fields, ie.Field())
	}
	for _, ie := range t.QFI {
		fields = append(fields, tlv.TLVFrom(QFI_TYPE, tlv.Bytes(pfcp.MarshalUint8(ie))))
	}
	for _, ie := range t.FramedRoute {
		fields = append(fields, tlv.TLVFrom(FRAMED_ROUTE_TYPE, tlv.Bytes(ie)))
	}
	if len(t.FramedRouting) > 0 {
		fields = append(fields, tlv.TLVFrom(FRAMED_ROUTING_TYPE, tlv.Bytes(t.FramedRouting)))
	}
	for _, ie := range t.FramedIPv6Route {
		fields = append(fields, tlv.TLVFrom(FRAMED_IPV6_ROUTE_TYPE, tlv.Bytes(ie)))
	}
	if len(t.SourceInterfaceType) > 0 {
		fields = append(fields, tlv.TLVFrom(_INTERFACE_TYPE_TYPE, tlv.Bytes(t.SourceInterfaceType)))
	}
	for _, ie := range t.IPMulticastAddressingInfo {
		fields = append(fields, ie.Field())
	}
	for _, ie := range t.DNSQueryFilter {
		fields = append(fields, tlv.TLVFrom(DNS_QUERY_FILTER_TYPE, tlv.Bytes(ie)))
	}
	if len(t.MBSSessionIdentifier) > 0 {
		fields = append(fields, tlv.TLVFrom(MBS_SESSION_IDENTIFIER_TYPE, tlv.Bytes(t.MBSSessionIdentifier)))
	}
	if len(t.AreaSessionID) > 0 {
		fields = append(fields, tlv.TLVFrom(AREA_SESSION_ID_TYPE, tlv.Bytes(t.AreaSessionID)))
	}
	return tlv.TLV(PDI_TYPE, fields...)
}

func (t *PDI) UnmarshalBinary(buf []byte) (err error) {
	*t = PDI{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	var seen [11]bool
	for _, de := range elements {
		switch de.Type {
		case SOURCE_INTERFACE_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &t.SourceInterface)
		case F_TEID_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			var v FTEID
			if err = de.UnmarshalValue(&v); err == nil {
				t.LocalFTEID = &v
			}
		case LOCAL_INGRESS_TUNNEL_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			t.LocalIngressTunnel = de.Value
		case NETWORK_INSTANCE_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			t.NetworkInstance = de.Value
		case REDUNDANT_TRANSMISSION_PARAMETERS_TYPE:
			if seen[4] {
				continue
			}
			seen[4] = true
			var v RedundantTransmissionParameters
			if err = de.UnmarshalValue(&v); err == nil {
				t.RedundantTransmissionDetectionParameters = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
			}
		case UE_IP_ADDRESS_TYPE:
			var v UEIPAddress
			if err = de.UnmarshalValue(&v); err == nil {
				t.UEIPaddress = append(t.UEIPaddress, v)
			}
		case TRAFFIC_ENDPOINT_ID_TYPE:
			t.TrafficEndpointID = append(t.TrafficEndpointID, de.Value)
		case SDF_FILTER_TYPE:
			var v SDFFilter
			if err = de.UnmarshalValue(&v); err == nil {
				t.SDFFilter = append(t.SDFFilter, v)
			}
		case APPLICATION_ID_TYPE:
			if seen[5] {
				continue
			}
			seen[5] = true
			var v ApplicationID
			if err = de.UnmarshalValue(&v); err == nil {
				t.ApplicationID = &v
			}
		case ETHERNET_PDU_SESSION_INFORMATION_TYPE:
			if seen[6] {
				continue
			}
			seen[6] = true
			t.EthernetPDUSessionInformation = de.Value
		case ETHERNET_PACKET_FILTER_TYPE:
			var v EthernetPacketFilter
			if err = de.UnmarshalValue(&v); err == nil {
				t.EthernetPacketFilter = append(t.EthernetPacketFilter, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
			}
		case QFI_TYPE:
			var v uint8
			if err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &v); err == nil {
				t.QFI = append(t.QFI, v)
			}
		case FRAMED_ROUTE_TYPE:
			t.FramedRoute = append(t.FramedRoute, de.Value)
		case FRAMED_ROUTING_TYPE:
			if seen[7] {
				continue
			}
			seen[7] = true
			t.FramedRouting = de.Value
		case FRAMED_IPV6_ROUTE_TYPE:
			t.FramedIPv6Route = append(t.FramedIPv6Route, de.Value)
		case _INTERFACE_TYPE_TYPE:
			if seen[8] {
				continue
			}
			seen[8] = true
			t.SourceInterfaceType = de.Value
		case IP_MULTICAST_ADDRESSING_INFO_WITHIN_PFCP_SESSION_ESTABLISHMENT_REQUEST_TYPE:
			var v IPMulticastAddressingInfowithinPFCPSessionEstablishmentRequest
			if err = de.UnmarshalValue(&v); err == nil {
				t.IPMulticastAddressingInfo = append(t.IPMulticastAddressingInfo, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
			}
		case DNS_QUERY_FILTER_TYPE:
			t.DNSQueryFilter = append(t.DNSQueryFilter, de.Value)
		case MBS_SESSION_IDENTIFIER_TYPE:
			if seen[9] {
				continue
			}
			seen[9] = true
			t.MBSSessionIdentifier = de.Value
		case AREA_SESSION_ID_TYPE:
			if seen[10] {
				continue
			}
			seen[10] = true
			t.AreaSessionID = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: SOURCE_INTERFACE_TYPE}
	}
	return
}
