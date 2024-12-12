package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'UE IP address Pool Information' data structure */
type UEIPaddressPoolInformation struct {
	UEIPaddressPoolIdentity UEIPAddress //UE IP Address
	NetworkInstance         []byte      //Network Instance
	SNSSAI                  []byte      //S-NSSAI
	IPversion               []byte      //IP version
}

func (t *UEIPaddressPoolInformation) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(UE_IP_ADDRESS_TYPE, tlv.Bytes(t.UEIPaddressPoolIdentity.Bytes())))
	if len(t.NetworkInstance) > 0 {
		fields = append(fields, tlv.TLVFrom(NETWORK_INSTANCE_TYPE, tlv.Bytes(t.NetworkInstance)))
	}
	if len(t.SNSSAI) > 0 {
		fields = append(fields, tlv.TLVFrom(S_NSSAI_TYPE, tlv.Bytes(t.SNSSAI)))
	}
	if len(t.IPversion) > 0 {
		fields = append(fields, tlv.TLVFrom(IP_VERSION_TYPE, tlv.Bytes(t.IPversion)))
	}
	return tlv.TLV(UE_IP_ADDRESS_POOL_INFORMATION_TYPE, fields...)
}

func (t *UEIPaddressPoolInformation) UnmarshalBinary(buf []byte) (err error) {
	*t = UEIPaddressPoolInformation{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case UE_IP_ADDRESS_TYPE:
			err = de.UnmarshalValue(&t.UEIPaddressPoolIdentity)
		case NETWORK_INSTANCE_TYPE:
			t.NetworkInstance = de.Value
		case S_NSSAI_TYPE:
			t.SNSSAI = de.Value
		case IP_VERSION_TYPE:
			t.IPversion = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
