package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Created Traffic Endpoint' data structure */
type CreatedTrafficEndpoint struct {
	TrafficEndpointID                  []byte        //Traffic Endpoint ID
	LocalFTEID                         *FTEID        //F-TEID
	LocalFTEIDforRedundantTransmission *FTEID        //F-TEID
	UEIPAddress                        []UEIPAddress //UE IP Address
	LocalIngressTunnel                 []byte        //Local Ingress Tunnel
	Unknown                            []uint32      `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *CreatedTrafficEndpoint) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(TRAFFIC_ENDPOINT_ID_TYPE, tlv.Bytes(t.TrafficEndpointID)))
	if t.LocalFTEID != nil {
		fields = append(fields, tlv.TLVFrom(F_TEID_TYPE, tlv.Bytes(t.LocalFTEID.Bytes())))
	}
	if t.LocalFTEIDforRedundantTransmission != nil {
		fields = append(fields, tlv.TLVFrom(F_TEID_TYPE, tlv.Bytes(t.LocalFTEIDforRedundantTransmission.Bytes())))
	}
	for _, ie := range t.UEIPAddress {
		fields = append(fields, tlv.TLVFrom(UE_IP_ADDRESS_TYPE, tlv.Bytes(ie.Bytes())))
	}
	if len(t.LocalIngressTunnel) > 0 {
		fields = append(fields, tlv.TLVFrom(LOCAL_INGRESS_TUNNEL_TYPE, tlv.Bytes(t.LocalIngressTunnel)))
	}
	return tlv.TLV(CREATED_TRAFFIC_ENDPOINT_TYPE, fields...)
}

func (t *CreatedTrafficEndpoint) UnmarshalBinary(buf []byte) (err error) {
	*t = CreatedTrafficEndpoint{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	var seen [4]bool
	for _, de := range elements {
		switch de.Type {
		case TRAFFIC_ENDPOINT_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.TrafficEndpointID = de.Value
		case F_TEID_TYPE:
			if !seen[1] {
				seen[1] = true
				var v FTEID
				if err = de.UnmarshalValue(&v); err == nil {
					t.LocalFTEID = &v
				}
			} else if !seen[2] {
				seen[2] = true
				var v FTEID
				if err = de.UnmarshalValue(&v); err == nil {
					t.LocalFTEIDforRedundantTransmission = &v
				}
			} else {
				continue
			}
		case UE_IP_ADDRESS_TYPE:
			var v UEIPAddress
			if err = de.UnmarshalValue(&v); err == nil {
				t.UEIPAddress = append(t.UEIPAddress, v)
			}
		case LOCAL_INGRESS_TUNNEL_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			t.LocalIngressTunnel = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: TRAFFIC_ENDPOINT_ID_TYPE}
	}
	return
}
