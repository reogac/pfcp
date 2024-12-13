package ie

import (
	"github.com/reogac/pfcp"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Query URR' data structure */
type QueryURR struct {
	URRID uint32 //URR ID
}

func (t *QueryURR) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(URR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(t.URRID))))
	return tlv.TLV(QUERY_URR_TYPE, fields...)
}

func (t *QueryURR) UnmarshalBinary(buf []byte) (err error) {
	*t = QueryURR{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case URR_ID_TYPE:
			err = pfcp.UnmarshalUint32(de.Value, &t.URRID)
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
