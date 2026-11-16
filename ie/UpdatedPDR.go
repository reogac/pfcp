package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Updated PDR' data structure */
type UpdatedPDR struct {
	PDRID                              uint16        //PDR ID
	LocalFTEIDforRedundantTransmission *FTEID        //F-TEID
	LocalFTEID                         *FTEID        //F-TEID
	UEIPAddress                        []UEIPAddress //UE IP Address
	Unknown                            []uint32      `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *UpdatedPDR) Field() tlv.Field {
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
	return tlv.TLV(UPDATED_PDR_TYPE, fields...)
}

func (t *UpdatedPDR) UnmarshalBinary(buf []byte) (err error) {
	*t = UpdatedPDR{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	var seen [3]bool
	for _, de := range elements {
		switch de.Type {
		case PDR_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			err = pfcp.UnmarshalExtendable(de.Value, 2, 2, &t.PDRID)
		case F_TEID_TYPE:
			if !seen[2] {
				seen[2] = true
				var v FTEID
				if err = de.UnmarshalValue(&v); err == nil {
					t.LocalFTEID = &v
				}
			} else if !seen[1] {
				seen[1] = true
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
