package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Join IP Multicast Information IE within Usage Report' data structure */
type JoinIPMulticastInformationIEwithinUsageReport struct {
	IPMulticastAddress []byte //IP Multicast Address
	SourceIPAddress    []byte //Source IP Address
}

func (t *JoinIPMulticastInformationIEwithinUsageReport) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(IP_MULTICAST_ADDRESS_TYPE, tlv.Bytes(t.IPMulticastAddress)))
	if len(t.SourceIPAddress) > 0 {
		fields = append(fields, tlv.TLVFrom(SOURCE_IP_ADDRESS_TYPE, tlv.Bytes(t.SourceIPAddress)))
	}
	return tlv.TLV(JOIN_IP_MULTICAST_INFORMATION_IE_WITHIN_USAGE_REPORT_TYPE, fields...)
}

func (t *JoinIPMulticastInformationIEwithinUsageReport) UnmarshalBinary(buf []byte) (err error) {
	*t = JoinIPMulticastInformationIEwithinUsageReport{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case IP_MULTICAST_ADDRESS_TYPE:
			t.IPMulticastAddress = de.Value
		case SOURCE_IP_ADDRESS_TYPE:
			t.SourceIPAddress = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
