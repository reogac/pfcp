package ie

import (
	"etrib5gc/pfcp"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* '3GPP Access Forwarding Action Information' data structure */
type AccessForwardingActionInformation struct {
	FARID    uint32   //FAR ID
	Weight   []byte   //Weight
	Priority []byte   //Priority
	URRID    []uint32 //URR ID
	RATType  []byte   //RAT Type
}

func (t *AccessForwardingActionInformation) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(FAR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(t.FARID))))
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
	return tlv.TLV(_ACCESS_FORWARDING_ACTION_INFORMATION_TYPE, fields...)
}

func (t *AccessForwardingActionInformation) UnmarshalBinary(buf []byte) (err error) {
	*t = AccessForwardingActionInformation{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case FAR_ID_TYPE:
			err = pfcp.UnmarshalUint32(de.Value, &t.FARID)
		case WEIGHT_TYPE:
			t.Weight = de.Value
		case PRIORITY_TYPE:
			t.Priority = de.Value
		case URR_ID_TYPE:
			var v uint32
			if err = pfcp.UnmarshalUint32(de.Value, &v); err == nil {
				t.URRID = append(t.URRID, v)
			}
		case RAT_TYPE_TYPE:
			t.RATType = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
