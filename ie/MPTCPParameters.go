package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'MPTCP Parameters' data structure */
type MPTCPParameters struct {
	MPTCPAddressInformation []byte   //MPTCP Address Information
	UELinkSpecificIPAddress []byte   //UE Link-Specific IP Address
	Unknown                 []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid                 []uint32 `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
}

func (t *MPTCPParameters) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(MPTCP_ADDRESS_INFORMATION_TYPE, tlv.Bytes(t.MPTCPAddressInformation)))
	fields = append(fields, tlv.TLVFrom(UE_LINK_SPECIFIC_IP_ADDRESS_TYPE, tlv.Bytes(t.UELinkSpecificIPAddress)))
	return tlv.TLV(MPTCP_PARAMETERS_TYPE, fields...)
}

func (t *MPTCPParameters) UnmarshalBinary(buf []byte) (err error) {
	*t = MPTCPParameters{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [2]bool
	for _, de := range elements {
		opt := false
		switch de.Type {
		case MPTCP_ADDRESS_INFORMATION_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.MPTCPAddressInformation = de.Value
		case UE_LINK_SPECIFIC_IP_ADDRESS_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			t.UELinkSpecificIPAddress = de.Value
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
		return &pfcp.MissingIE{Type: MPTCP_ADDRESS_INFORMATION_TYPE}
	}
	if !seen[1] {
		return &pfcp.MissingIE{Type: UE_LINK_SPECIFIC_IP_ADDRESS_TYPE}
	}
	return
}
