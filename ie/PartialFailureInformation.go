package ie

import (
	"etrib5gc/pfcp"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Partial Failure Information' data structure */
type PartialFailureInformation struct {
	FailedRuleID           []byte //Failed Rule ID
	FailureCause           uint8  //Cause
	OffendingIEInformation []byte //Offending IE Information
}

func (t *PartialFailureInformation) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(FAILED_RULE_ID_TYPE, tlv.Bytes(t.FailedRuleID)))
	fields = append(fields, tlv.TLVFrom(CAUSE_TYPE, tlv.Bytes(pfcp.MarshalUint8(t.FailureCause))))
	fields = append(fields, tlv.TLVFrom(OFFENDING_IE_INFORMATION_TYPE, tlv.Bytes(t.OffendingIEInformation)))
	return tlv.TLV(PARTIAL_FAILURE_INFORMATION_TYPE, fields...)
}

func (t *PartialFailureInformation) UnmarshalBinary(buf []byte) (err error) {
	*t = PartialFailureInformation{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case FAILED_RULE_ID_TYPE:
			t.FailedRuleID = de.Value
		case CAUSE_TYPE:
			err = pfcp.UnmarshalUint8(de.Value, &t.FailureCause)
		case OFFENDING_IE_INFORMATION_TYPE:
			t.OffendingIEInformation = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
