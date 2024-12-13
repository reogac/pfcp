package ie

import (
	"github.com/reogac/pfcp"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Remove QER' data structure */
type RemoveQER struct {
	QERID uint32 //QER ID
}

func (t *RemoveQER) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(QER_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(t.QERID))))
	return tlv.TLV(REMOVE_QER_TYPE, fields...)
}

func (t *RemoveQER) UnmarshalBinary(buf []byte) (err error) {
	*t = RemoveQER{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case QER_ID_TYPE:
			err = pfcp.UnmarshalUint32(de.Value, &t.QERID)
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
