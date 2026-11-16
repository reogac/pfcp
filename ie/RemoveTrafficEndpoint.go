package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Remove Traffic Endpoint' data structure */
type RemoveTrafficEndpoint struct {
	TrafficEndpointID []byte   //Traffic Endpoint ID
	Unknown           []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *RemoveTrafficEndpoint) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(TRAFFIC_ENDPOINT_ID_TYPE, tlv.Bytes(t.TrafficEndpointID)))
	return tlv.TLV(REMOVE_TRAFFIC_ENDPOINT_TYPE, fields...)
}

func (t *RemoveTrafficEndpoint) UnmarshalBinary(buf []byte) (err error) {
	*t = RemoveTrafficEndpoint{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [1]bool
	for _, de := range elements {
		switch de.Type {
		case TRAFFIC_ENDPOINT_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.TrafficEndpointID = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return pfcp.Faulty(de.Type, err)
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: TRAFFIC_ENDPOINT_ID_TYPE}
	}
	return
}
