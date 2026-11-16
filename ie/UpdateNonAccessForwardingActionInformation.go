package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Update Non-3GPP Access Forwarding Action Information' data structure */
type UpdateNonAccessForwardingActionInformation struct {
	FARID    *uint32  //FAR ID
	Weight   []byte   //Weight
	Priority []byte   //Priority
	URRID    []uint32 //URR ID
	RATType  []byte   //RAT Type
	Unknown  []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *UpdateNonAccessForwardingActionInformation) Field() tlv.Field {
	fields := []tlv.Field{}
	if t.FARID != nil {
		fields = append(fields, tlv.TLVFrom(FAR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(*t.FARID))))
	}
	if len(t.Weight) > 0 {
		fields = append(fields, tlv.TLVFrom(WEIGHT_TYPE, tlv.Bytes(t.Weight)))
	}
	if len(t.Priority) > 0 {
		fields = append(fields, tlv.TLVFrom(PRIORITY_TYPE, tlv.Bytes(t.Priority)))
	}
	for _, ie := range t.URRID {
		fields = append(fields, tlv.TLVFrom(URR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(ie))))
	}
	if len(t.RATType) > 0 {
		fields = append(fields, tlv.TLVFrom(RAT_TYPE_TYPE, tlv.Bytes(t.RATType)))
	}
	return tlv.TLV(UPDATE_NON__ACCESS_FORWARDING_ACTION_INFORMATION_TYPE, fields...)
}

func (t *UpdateNonAccessForwardingActionInformation) UnmarshalBinary(buf []byte) (err error) {
	*t = UpdateNonAccessForwardingActionInformation{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [4]bool
	for _, de := range elements {
		switch de.Type {
		case FAR_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			var v uint32
			if err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &v); err == nil {
				t.FARID = &v
			}
		case WEIGHT_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			t.Weight = de.Value
		case PRIORITY_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			t.Priority = de.Value
		case URR_ID_TYPE:
			var v uint32
			if err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &v); err == nil {
				t.URRID = append(t.URRID, v)
			}
		case RAT_TYPE_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			t.RATType = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return pfcp.Faulty(de.Type, err)
		}
	}
	return
}
