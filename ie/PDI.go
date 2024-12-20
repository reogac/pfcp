package ie

import (
	"github.com/reogac/pfcp"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'PDI' data structure */
type PDI struct {
	SourceInterface                          uint8                                                           //Source Interface
	LocalFTEID                               *FTEID                                                          //F-TEID
	LocalIngressTunnel                       []byte                                                          //Local Ingress Tunnel
	NetworkInstance                          []byte                                                          //Network Instance
	RedundantTransmissionDetectionParameters *RedundantTransmissionParameters                                //Redundant Transmission Parameters
	UEIPaddress                              *UEIPAddress                                                    //UE IP Address
	TrafficEndpointID                        []byte                                                          //Traffic Endpoint ID
	SDFFilter                                []SDFFilter                                                     //SDF Filter
	ApplicationID                            *ApplicationID                                                  //Application ID
	EthernetPDUSessionInformation            []byte                                                          //Ethernet PDU Session Information
	EthernetPacketFilter                     *EthernetPacketFilter                                           //Ethernet Packet Filter
	QFI                                      *uint8                                                          //QFI
	FramedRoute                              []byte                                                          //Framed-Route
	FramedRouting                            []byte                                                          //Framed-Routing
	FramedIPv6Route                          []byte                                                          //Framed-IPv6-Route
	SourceInterfaceType                      []byte                                                          //3GPP Interface Type
	IPMulticastAddressingInfo                *IPMulticastAddressingInfowithinPFCPSessionEstablishmentRequest //IP Multicast Addressing Info within PFCP Session Establishment Request
	DNSQueryFilter                           []byte                                                          //DNS Query Filter
	MBSSessionIdentifier                     []byte                                                          //MBS Session Identifier
	AreaSessionID                            []byte                                                          //Area Session ID
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
	if t.UEIPaddress != nil {
		fields = append(fields, tlv.TLVFrom(UE_IP_ADDRESS_TYPE, tlv.Bytes(t.UEIPaddress.Bytes())))
	}
	if len(t.TrafficEndpointID) > 0 {
		fields = append(fields, tlv.TLVFrom(TRAFFIC_ENDPOINT_ID_TYPE, tlv.Bytes(t.TrafficEndpointID)))
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
	if t.EthernetPacketFilter != nil {
		fields = append(fields, t.EthernetPacketFilter.Field())
	}
	if t.QFI != nil {
		fields = append(fields, tlv.TLVFrom(QFI_TYPE, tlv.Bytes(pfcp.MarshalUint8(*t.QFI))))
	}
	if len(t.FramedRoute) > 0 {
		fields = append(fields, tlv.TLVFrom(FRAMED_ROUTE_TYPE, tlv.Bytes(t.FramedRoute)))
	}
	if len(t.FramedRouting) > 0 {
		fields = append(fields, tlv.TLVFrom(FRAMED_ROUTING_TYPE, tlv.Bytes(t.FramedRouting)))
	}
	for len(t.FramedIPv6Route) > 0 {
		fields = append(fields, tlv.TLVFrom(FRAMED_IPV6_ROUTE_TYPE, tlv.Bytes(t.FramedIPv6Route)))
	}
	if len(t.SourceInterfaceType) > 0 {
		fields = append(fields, tlv.TLVFrom(_INTERFACE_TYPE_TYPE, tlv.Bytes(t.SourceInterfaceType)))
	}
	if t.IPMulticastAddressingInfo != nil {
		fields = append(fields, t.IPMulticastAddressingInfo.Field())
	}
	if len(t.DNSQueryFilter) > 0 {
		fields = append(fields, tlv.TLVFrom(DNS_QUERY_FILTER_TYPE, tlv.Bytes(t.DNSQueryFilter)))
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
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case SOURCE_INTERFACE_TYPE:
			err = pfcp.UnmarshalUint8(de.Value, &t.SourceInterface)
		case F_TEID_TYPE:
			var v FTEID
			if err = de.UnmarshalValue(&v); err == nil {
				t.LocalFTEID = &v
			}
		case LOCAL_INGRESS_TUNNEL_TYPE:
			t.LocalIngressTunnel = de.Value
		case NETWORK_INSTANCE_TYPE:
			t.NetworkInstance = de.Value
		case REDUNDANT_TRANSMISSION_PARAMETERS_TYPE:
			var v RedundantTransmissionParameters
			if err = de.UnmarshalValue(&v); err == nil {
				t.RedundantTransmissionDetectionParameters = &v
			}
		case UE_IP_ADDRESS_TYPE:
			var v UEIPAddress
			if err = de.UnmarshalValue(&v); err == nil {
				t.UEIPaddress = &v
			}
		case TRAFFIC_ENDPOINT_ID_TYPE:
			t.TrafficEndpointID = de.Value
		case SDF_FILTER_TYPE:
			var v SDFFilter
			if err = de.UnmarshalValue(&v); err == nil {
				t.SDFFilter = append(t.SDFFilter, v)
			}
		case APPLICATION_ID_TYPE:
			var v ApplicationID
			if err = de.UnmarshalValue(&v); err == nil {
				t.ApplicationID = &v
			}
		case ETHERNET_PDU_SESSION_INFORMATION_TYPE:
			t.EthernetPDUSessionInformation = de.Value
		case ETHERNET_PACKET_FILTER_TYPE:
			var v EthernetPacketFilter
			if err = de.UnmarshalValue(&v); err == nil {
				t.EthernetPacketFilter = &v
			}
		case QFI_TYPE:
			var v uint8
			if err = pfcp.UnmarshalUint8(de.Value, &v); err == nil {
				t.QFI = &v
			}
		case FRAMED_ROUTE_TYPE:
			t.FramedRoute = de.Value
		case FRAMED_ROUTING_TYPE:
			t.FramedRouting = de.Value
		case FRAMED_IPV6_ROUTE_TYPE:
			t.FramedIPv6Route = de.Value
		case _INTERFACE_TYPE_TYPE:
			t.SourceInterfaceType = de.Value
		case IP_MULTICAST_ADDRESSING_INFO_WITHIN_PFCP_SESSION_ESTABLISHMENT_REQUEST_TYPE:
			var v IPMulticastAddressingInfowithinPFCPSessionEstablishmentRequest
			if err = de.UnmarshalValue(&v); err == nil {
				t.IPMulticastAddressingInfo = &v
			}
		case DNS_QUERY_FILTER_TYPE:
			t.DNSQueryFilter = de.Value
		case MBS_SESSION_IDENTIFIER_TYPE:
			t.MBSSessionIdentifier = de.Value
		case AREA_SESSION_ID_TYPE:
			t.AreaSessionID = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
