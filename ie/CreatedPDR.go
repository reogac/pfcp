package ie

import (
	"github.com/reogac/pfcp"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Created PDR' data structure */
type CreatedPDR struct {
	PDRID              uint16       //PDR ID
	LocalFTEID         *FTEID       //F-TEID
	UEIPAddress        *UEIPAddress //UE IP Address
	LocalIngressTunnel []byte       //Local Ingress Tunnel
}

func (t *CreatedPDR) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(PDR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint16(t.PDRID))))
	if t.LocalFTEID != nil {
		fields = append(fields, tlv.TLVFrom(F_TEID_TYPE, tlv.Bytes(t.LocalFTEID.Bytes())))
	}
	if t.UEIPAddress != nil {
		fields = append(fields, tlv.TLVFrom(UE_IP_ADDRESS_TYPE, tlv.Bytes(t.UEIPAddress.Bytes())))
	}
	if len(t.LocalIngressTunnel) > 0 {
		fields = append(fields, tlv.TLVFrom(LOCAL_INGRESS_TUNNEL_TYPE, tlv.Bytes(t.LocalIngressTunnel)))
	}
	return tlv.TLV(CREATED_PDR_TYPE, fields...)
}

func (t *CreatedPDR) UnmarshalBinary(buf []byte) (err error) {
	*t = CreatedPDR{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case PDR_ID_TYPE:
			err = pfcp.UnmarshalUint16(de.Value, &t.PDRID)
		case F_TEID_TYPE:
			var v FTEID
			if err = de.UnmarshalValue(&v); err == nil {
				t.LocalFTEID = &v
			}
		case UE_IP_ADDRESS_TYPE:
			var v UEIPAddress
			if err = de.UnmarshalValue(&v); err == nil {
				t.UEIPAddress = &v
			}
		case LOCAL_INGRESS_TUNNEL_TYPE:
			t.LocalIngressTunnel = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
