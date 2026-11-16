package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Partial Failure Information' data structure */
type PartialFailureInformation struct {
	FailedRuleID           []byte   //Failed Rule ID
	FailureCause           uint8    //Cause
	OffendingIEInformation [][]byte //Offending IE Information
	Unknown                []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *PartialFailureInformation) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(FAILED_RULE_ID_TYPE, tlv.Bytes(t.FailedRuleID)))
	fields = append(fields, tlv.TLVFrom(CAUSE_TYPE, tlv.Bytes(pfcp.MarshalUint8(t.FailureCause))))
	for _, ie := range t.OffendingIEInformation {
		fields = append(fields, tlv.TLVFrom(OFFENDING_IE_INFORMATION_TYPE, tlv.Bytes(ie)))
	}
	return tlv.TLV(PARTIAL_FAILURE_INFORMATION_TYPE, fields...)
}

func (t *PartialFailureInformation) UnmarshalBinary(buf []byte) (err error) {
	*t = PartialFailureInformation{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [2]bool
	for _, de := range elements {
		switch de.Type {
		case FAILED_RULE_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.FailedRuleID = de.Value
		case CAUSE_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			err = pfcp.UnmarshalUint8(de.Value, &t.FailureCause)
		case OFFENDING_IE_INFORMATION_TYPE:
			t.OffendingIEInformation = append(t.OffendingIEInformation, de.Value)
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return pfcp.Faulty(de.Type, err)
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: FAILED_RULE_ID_TYPE}
	}
	if !seen[1] {
		return &pfcp.MissingIE{Type: CAUSE_TYPE}
	}
	if len(t.OffendingIEInformation) == 0 {
		return &pfcp.MissingIE{Type: OFFENDING_IE_INFORMATION_TYPE}
	}
	return
}
