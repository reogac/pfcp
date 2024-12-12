package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'MPTCP Parameters' data structure */
type MPTCPParameters struct {
	MPTCPAddressInformation []byte //MPTCP Address Information
	UELinkSpecificIPAddress []byte //UE Link-Specific IP Address
}

func (t *MPTCPParameters) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(MPTCP_ADDRESS_INFORMATION_TYPE, tlv.Bytes(t.MPTCPAddressInformation)))
	fields = append(fields, tlv.TLVFrom(UE_LINK_SPECIFIC_IP_ADDRESS_TYPE, tlv.Bytes(t.UELinkSpecificIPAddress)))
	return tlv.TLV(MPTCP_PARAMETERS_TYPE, fields...)
}

func (t *MPTCPParameters) UnmarshalBinary(buf []byte) (err error) {
	*t = MPTCPParameters{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case MPTCP_ADDRESS_INFORMATION_TYPE:
			t.MPTCPAddressInformation = de.Value
		case UE_LINK_SPECIFIC_IP_ADDRESS_TYPE:
			t.UELinkSpecificIPAddress = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
