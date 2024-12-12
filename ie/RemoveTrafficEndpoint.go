package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Remove Traffic Endpoint' data structure */
type RemoveTrafficEndpoint struct {
	TrafficEndpointID []byte //Traffic Endpoint ID
}

func (t *RemoveTrafficEndpoint) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(TRAFFIC_ENDPOINT_ID_TYPE, tlv.Bytes(t.TrafficEndpointID)))
	return tlv.TLV(REMOVE_TRAFFIC_ENDPOINT_TYPE, fields...)
}

func (t *RemoveTrafficEndpoint) UnmarshalBinary(buf []byte) (err error) {
	*t = RemoveTrafficEndpoint{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case TRAFFIC_ENDPOINT_ID_TYPE:
			t.TrafficEndpointID = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
