package ie

import (
	"etrib5gc/pfcp"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Remove FAR' data structure */
type RemoveFAR struct {
	FARID uint32 //FAR ID
}

func (t *RemoveFAR) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(FAR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(t.FARID))))
	return tlv.TLV(REMOVE_FAR_TYPE, fields...)
}

func (t *RemoveFAR) UnmarshalBinary(buf []byte) (err error) {
	*t = RemoveFAR{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case FAR_ID_TYPE:
			err = pfcp.UnmarshalUint32(de.Value, &t.FARID)
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
