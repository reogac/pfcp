package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'UE IP Address Usage Information' data structure */
type UEIPAddressUsageInformation struct {
	UEIPAddressUsageSequenceNumber []byte       //Sequence Number
	UEIPAddressUsageMetric         []byte       //Metric
	ValidityTimer                  []byte       //Validity Timer
	NumberofUEIPAddresses          []byte       //Number of UE IP Addresses
	NetworkInstance                []byte       //Network Instance
	UEIPAddressPoolId              *UEIPAddress //UE IP Address
	SNSSAI                         []byte       //S-NSSAI
}

func (t *UEIPAddressUsageInformation) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(SEQUENCE_NUMBER_TYPE, tlv.Bytes(t.UEIPAddressUsageSequenceNumber)))
	fields = append(fields, tlv.TLVFrom(METRIC_TYPE, tlv.Bytes(t.UEIPAddressUsageMetric)))
	fields = append(fields, tlv.TLVFrom(VALIDITY_TIMER_TYPE, tlv.Bytes(t.ValidityTimer)))
	fields = append(fields, tlv.TLVFrom(NUMBER_OF_UE_IP_ADDRESSES_TYPE, tlv.Bytes(t.NumberofUEIPAddresses)))
	fields = append(fields, tlv.TLVFrom(NETWORK_INSTANCE_TYPE, tlv.Bytes(t.NetworkInstance)))
	if t.UEIPAddressPoolId != nil {
		fields = append(fields, tlv.TLVFrom(UE_IP_ADDRESS_TYPE, tlv.Bytes(t.UEIPAddressPoolId.Bytes())))
	}
	if len(t.SNSSAI) > 0 {
		fields = append(fields, tlv.TLVFrom(S_NSSAI_TYPE, tlv.Bytes(t.SNSSAI)))
	}
	return tlv.TLV(UE_IP_ADDRESS_USAGE_INFORMATION_TYPE, fields...)
}

func (t *UEIPAddressUsageInformation) UnmarshalBinary(buf []byte) (err error) {
	*t = UEIPAddressUsageInformation{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case SEQUENCE_NUMBER_TYPE:
			t.UEIPAddressUsageSequenceNumber = de.Value
		case METRIC_TYPE:
			t.UEIPAddressUsageMetric = de.Value
		case VALIDITY_TIMER_TYPE:
			t.ValidityTimer = de.Value
		case NUMBER_OF_UE_IP_ADDRESSES_TYPE:
			t.NumberofUEIPAddresses = de.Value
		case NETWORK_INSTANCE_TYPE:
			t.NetworkInstance = de.Value
		case UE_IP_ADDRESS_TYPE:
			var v UEIPAddress
			if err = de.UnmarshalValue(&v); err == nil {
				t.UEIPAddressPoolId = &v
			}
		case S_NSSAI_TYPE:
			t.SNSSAI = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
