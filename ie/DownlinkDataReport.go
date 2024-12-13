package ie

import (
	"github.com/reogac/pfcp"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Downlink Data Report' data structure */
type DownlinkDataReport struct {
	PDRID                          uint16 //PDR ID
	DownlinkDataServiceInformation []byte //Downlink Data Service Information
	DLDataPacketsSize              []byte //DL Data Packets Size
	DLDataStatus                   *uint8 //Data Status
}

func (t *DownlinkDataReport) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(PDR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint16(t.PDRID))))
	if len(t.DownlinkDataServiceInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(DOWNLINK_DATA_SERVICE_INFORMATION_TYPE, tlv.Bytes(t.DownlinkDataServiceInformation)))
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
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case PDR_ID_TYPE:
			err = pfcp.UnmarshalUint16(de.Value, &t.PDRID)
		case DOWNLINK_DATA_SERVICE_INFORMATION_TYPE:
			t.DownlinkDataServiceInformation = de.Value
		case DL_DATA_PACKETS_SIZE_TYPE:
			t.DLDataPacketsSize = de.Value
		case DATA_STATUS_TYPE:
			var v uint8
			if err = pfcp.UnmarshalUint8(de.Value, &v); err == nil {
				t.DLDataStatus = &v
			}
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
