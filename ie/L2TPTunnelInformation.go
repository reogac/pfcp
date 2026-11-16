package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'L2TP Tunnel Information' data structure */
type L2TPTunnelInformation struct {
	LNSAddress       []byte   //LNS Address
	TunnelPassword   []byte   //Tunnel Password
	TunnelPreference []byte   //Tunnel Preference
	Unknown          []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid          []uint32 `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
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
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [3]bool
	for _, de := range elements {
		opt := false
		switch de.Type {
		case LNS_ADDRESS_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.LNSAddress = de.Value
		case TUNNEL_PASSWORD_TYPE:
			opt = true
			if seen[1] {
				continue
			}
			seen[1] = true
			t.TunnelPassword = de.Value
		case TUNNEL_PREFERENCE_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			t.TunnelPreference = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			if err = pfcp.Faulty(de.Type, err); opt && pfcp.Discardable(err) {
				t.Invalid = append(t.Invalid, de.Type)
				err = nil
				continue
			}
			return err
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: LNS_ADDRESS_TYPE}
	}
	return
}
