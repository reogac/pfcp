package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
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
	Unknown                        []uint32     `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
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
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	var seen [7]bool
	for _, de := range elements {
		switch de.Type {
		case SEQUENCE_NUMBER_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.UEIPAddressUsageSequenceNumber = de.Value
		case METRIC_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			t.UEIPAddressUsageMetric = de.Value
		case VALIDITY_TIMER_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			t.ValidityTimer = de.Value
		case NUMBER_OF_UE_IP_ADDRESSES_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			t.NumberofUEIPAddresses = de.Value
		case NETWORK_INSTANCE_TYPE:
			if seen[4] {
				continue
			}
			seen[4] = true
			t.NetworkInstance = de.Value
		case UE_IP_ADDRESS_TYPE:
			if seen[5] {
				continue
			}
			seen[5] = true
			var v UEIPAddress
			if err = de.UnmarshalValue(&v); err == nil {
				t.UEIPAddressPoolId = &v
			}
		case S_NSSAI_TYPE:
			if seen[6] {
				continue
			}
			seen[6] = true
			t.SNSSAI = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: SEQUENCE_NUMBER_TYPE}
	}
	if !seen[1] {
		return &pfcp.MissingIE{Type: METRIC_TYPE}
	}
	if !seen[2] {
		return &pfcp.MissingIE{Type: VALIDITY_TIMER_TYPE}
	}
	if !seen[3] {
		return &pfcp.MissingIE{Type: NUMBER_OF_UE_IP_ADDRESSES_TYPE}
	}
	if !seen[4] {
		return &pfcp.MissingIE{Type: NETWORK_INSTANCE_TYPE}
	}
	return
}
