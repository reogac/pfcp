package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'L2TP Tunnel Information' data structure */
type L2TPTunnelInformation struct {
	LNSAddress       []byte //LNS Address
	TunnelPassword   []byte //Tunnel Password
	TunnelPreference []byte //Tunnel Preference
}

func (t *L2TPTunnelInformation) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(LNS_ADDRESS_TYPE, tlv.Bytes(t.LNSAddress)))
	if len(t.TunnelPassword) > 0 {
		fields = append(fields, tlv.TLVFrom(TUNNEL_PASSWORD_TYPE, tlv.Bytes(t.TunnelPassword)))
	}
	if len(t.TunnelPreference) > 0 {
		fields = append(fields, tlv.TLVFrom(TUNNEL_PREFERENCE_TYPE, tlv.Bytes(t.TunnelPreference)))
	}
	return tlv.TLV(L2TP_TUNNEL_INFORMATION_TYPE, fields...)
}

func (t *L2TPTunnelInformation) UnmarshalBinary(buf []byte) (err error) {
	*t = L2TPTunnelInformation{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case LNS_ADDRESS_TYPE:
			t.LNSAddress = de.Value
		case TUNNEL_PASSWORD_TYPE:
			t.TunnelPassword = de.Value
		case TUNNEL_PREFERENCE_TYPE:
			t.TunnelPreference = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
