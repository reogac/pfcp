package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Downlink Data Report' data structure */
type DownlinkDataReport struct {
	PDRID                          []uint16 //PDR ID
	DownlinkDataServiceInformation [][]byte //Downlink Data Service Information
	DLDataPacketsSize              []byte   //DL Data Packets Size
	DLDataStatus                   *uint8   //Data Status
	Unknown                        []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *DownlinkDataReport) Field() tlv.Field {
	fields := []tlv.Field{}
	for _, ie := range t.PDRID {
		fields = append(fields, tlv.TLVFrom(PDR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint16(ie))))
	}
	for _, ie := range t.DownlinkDataServiceInformation {
		fields = append(fields, tlv.TLVFrom(DOWNLINK_DATA_SERVICE_INFORMATION_TYPE, tlv.Bytes(ie)))
	}
	if len(t.DLDataPacketsSize) > 0 {
		fields = append(fields, tlv.TLVFrom(DL_DATA_PACKETS_SIZE_TYPE, tlv.Bytes(t.DLDataPacketsSize)))
	}
	if t.DLDataStatus != nil {
		fields = append(fields, tlv.TLVFrom(DATA_STATUS_TYPE, tlv.Bytes(pfcp.MarshalUint8(*t.DLDataStatus))))
	}
	return tlv.TLV(DOWNLINK_DATA_REPORT_TYPE, fields...)
}

func (t *DownlinkDataReport) UnmarshalBinary(buf []byte) (err error) {
	*t = DownlinkDataReport{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [2]bool
	for _, de := range elements {
		switch de.Type {
		case PDR_ID_TYPE:
			var v uint16
			if err = pfcp.UnmarshalExtendable(de.Value, 2, 2, &v); err == nil {
				t.PDRID = append(t.PDRID, v)
			}
		case DOWNLINK_DATA_SERVICE_INFORMATION_TYPE:
			t.DownlinkDataServiceInformation = append(t.DownlinkDataServiceInformation, de.Value)
		case DL_DATA_PACKETS_SIZE_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.DLDataPacketsSize = de.Value
		case DATA_STATUS_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			var v uint8
			if err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &v); err == nil {
				t.DLDataStatus = &v
			}
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return pfcp.Faulty(de.Type, err)
		}
	}
	if len(t.PDRID) == 0 {
		return &pfcp.MissingIE{Type: PDR_ID_TYPE}
	}
	return
}
