package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Created Traffic Endpoint' data structure */
type CreatedTrafficEndpoint struct {
	TrafficEndpointID  []byte       //Traffic Endpoint ID
	LocalFTEID         *FTEID       //F-TEID
	UEIPAddress        *UEIPAddress //UE IP Address
	LocalIngressTunnel []byte       //Local Ingress Tunnel
}

func (t *CreatedTrafficEndpoint) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(TRAFFIC_ENDPOINT_ID_TYPE, tlv.Bytes(t.TrafficEndpointID)))
	if t.LocalFTEID != nil {
		fields = append(fields, tlv.TLVFrom(F_TEID_TYPE, tlv.Bytes(t.LocalFTEID.Bytes())))
	}
	if t.UEIPAddress != nil {
		fields = append(fields, tlv.TLVFrom(UE_IP_ADDRESS_TYPE, tlv.Bytes(t.UEIPAddress.Bytes())))
	}
	if len(t.LocalIngressTunnel) > 0 {
		fields = append(fields, tlv.TLVFrom(LOCAL_INGRESS_TUNNEL_TYPE, tlv.Bytes(t.LocalIngressTunnel)))
	}
	return tlv.TLV(CREATED_TRAFFIC_ENDPOINT_TYPE, fields...)
}

func (t *CreatedTrafficEndpoint) UnmarshalBinary(buf []byte) (err error) {
	*t = CreatedTrafficEndpoint{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case TRAFFIC_ENDPOINT_ID_TYPE:
			t.TrafficEndpointID = de.Value
		case F_TEID_TYPE:
			var v FTEID
			if err = de.UnmarshalValue(&v); err == nil {
				t.LocalFTEID = &v
			}
		case UE_IP_ADDRESS_TYPE:
			var v UEIPAddress
			if err = de.UnmarshalValue(&v); err == nil {
				t.UEIPAddress = &v
			}
		case LOCAL_INGRESS_TUNNEL_TYPE:
			t.LocalIngressTunnel = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
