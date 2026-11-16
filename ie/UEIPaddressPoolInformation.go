package ie

import (
	"github.com/reogac/pfcp"

	"github.com/reogac/pfcp/tlv"
)

/* 'UE IP address Pool Information' data structure */
type UEIPaddressPoolInformation struct {
	UEIPaddressPoolIdentity []UEIPAddress //UE IP Address
	NetworkInstance         []byte        //Network Instance
	SNSSAI                  [][]byte      //S-NSSAI
	IPversion               []byte        //IP version
	Unknown                 []uint32      `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *UEIPaddressPoolInformation) Field() tlv.Field {
	fields := []tlv.Field{}
	for _, ie := range t.UEIPaddressPoolIdentity {
		fields = append(fields, tlv.TLVFrom(UE_IP_ADDRESS_TYPE, tlv.Bytes(ie.Bytes())))
	}
	if len(t.NetworkInstance) > 0 {
		fields = append(fields, tlv.TLVFrom(NETWORK_INSTANCE_TYPE, tlv.Bytes(t.NetworkInstance)))
	}
	for _, ie := range t.SNSSAI {
		fields = append(fields, tlv.TLVFrom(S_NSSAI_TYPE, tlv.Bytes(ie)))
	}
	if len(t.IPversion) > 0 {
		fields = append(fields, tlv.TLVFrom(IP_VERSION_TYPE, tlv.Bytes(t.IPversion)))
	}
	return tlv.TLV(UE_IP_ADDRESS_POOL_INFORMATION_TYPE, fields...)
}

func (t *UEIPaddressPoolInformation) UnmarshalBinary(buf []byte) (err error) {
	*t = UEIPaddressPoolInformation{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	var seen [2]bool
	for _, de := range elements {
		switch de.Type {
		case UE_IP_ADDRESS_TYPE:
			var v UEIPAddress
			if err = de.UnmarshalValue(&v); err == nil {
				t.UEIPaddressPoolIdentity = append(t.UEIPaddressPoolIdentity, v)
			}
		case NETWORK_INSTANCE_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.NetworkInstance = de.Value
		case S_NSSAI_TYPE:
			t.SNSSAI = append(t.SNSSAI, de.Value)
		case IP_VERSION_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			t.IPversion = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	if len(t.UEIPaddressPoolIdentity) == 0 {
		return &pfcp.MissingIE{Type: UE_IP_ADDRESS_TYPE}
	}
	return
}
