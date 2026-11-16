package ie

import (
	"github.com/reogac/pfcp"

	"github.com/reogac/pfcp/tlv"
)

/* 'Created PDR' data structure */
type CreatedPDR struct {
	PDRID                              uint16        //PDR ID
	LocalFTEID                         *FTEID        //F-TEID
	LocalFTEIDforRedundantTransmission *FTEID        //F-TEID
	UEIPAddress                        []UEIPAddress //UE IP Address
	LocalIngressTunnel                 []byte        //Local Ingress Tunnel
	Unknown                            []uint32      `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *CreatedPDR) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(PDR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint16(t.PDRID))))
	if t.LocalFTEID != nil {
		fields = append(fields, tlv.TLVFrom(F_TEID_TYPE, tlv.Bytes(t.LocalFTEID.Bytes())))
	}
	if t.LocalFTEIDforRedundantTransmission != nil {
		fields = append(fields, tlv.TLVFrom(F_TEID_TYPE, tlv.Bytes(t.LocalFTEIDforRedundantTransmission.Bytes())))
	}
	for _, ie := range t.UEIPAddress {
		fields = append(fields, tlv.TLVFrom(UE_IP_ADDRESS_TYPE, tlv.Bytes(ie.Bytes())))
	}
	if len(t.LocalIngressTunnel) > 0 {
		fields = append(fields, tlv.TLVFrom(LOCAL_INGRESS_TUNNEL_TYPE, tlv.Bytes(t.LocalIngressTunnel)))
	}
	return tlv.TLV(CREATED_PDR_TYPE, fields...)
}

func (t *CreatedPDR) UnmarshalBinary(buf []byte) (err error) {
	*t = CreatedPDR{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	var seen [4]bool
	for _, de := range elements {
		switch de.Type {
		case PDR_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			err = pfcp.UnmarshalExtendable(de.Value, 2, 2, &t.PDRID)
		case F_TEID_TYPE:
			if !seen[1] {
				seen[1] = true
				var v FTEID
				if err = de.UnmarshalValue(&v); err == nil {
					t.LocalFTEID = &v
				}
			} else if !seen[2] {
				seen[2] = true
				var v FTEID
				if err = de.UnmarshalValue(&v); err == nil {
					t.LocalFTEIDforRedundantTransmission = &v
				}
			} else {
				continue
			}
		case UE_IP_ADDRESS_TYPE:
			var v UEIPAddress
			if err = de.UnmarshalValue(&v); err == nil {
				t.UEIPAddress = append(t.UEIPAddress, v)
			}
		case LOCAL_INGRESS_TUNNEL_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			t.LocalIngressTunnel = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: PDR_ID_TYPE}
	}
	return
}
