package ie

import (
	"github.com/reogac/pfcp"

	"github.com/reogac/pfcp/tlv"
)

/* 'Leave IP Multicast Information IE within Usage Report' data structure */
type LeaveIPMulticastInformationIEwithinUsageReport struct {
	IPMulticastAddress []byte   //IP Multicast Address
	SourceIPAddress    [][]byte //Source IP Address
	Unknown            []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *LeaveIPMulticastInformationIEwithinUsageReport) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(IP_MULTICAST_ADDRESS_TYPE, tlv.Bytes(t.IPMulticastAddress)))
	for _, ie := range t.SourceIPAddress {
		fields = append(fields, tlv.TLVFrom(SOURCE_IP_ADDRESS_TYPE, tlv.Bytes(ie)))
	}
	return tlv.TLV(LEAVE_IP_MULTICAST_INFORMATION_IE_WITHIN_USAGE_REPORT_TYPE, fields...)
}

func (t *LeaveIPMulticastInformationIEwithinUsageReport) UnmarshalBinary(buf []byte) (err error) {
	*t = LeaveIPMulticastInformationIEwithinUsageReport{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	var seen [1]bool
	for _, de := range elements {
		switch de.Type {
		case IP_MULTICAST_ADDRESS_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.IPMulticastAddress = de.Value
		case SOURCE_IP_ADDRESS_TYPE:
			t.SourceIPAddress = append(t.SourceIPAddress, de.Value)
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: IP_MULTICAST_ADDRESS_TYPE}
	}
	return
}
