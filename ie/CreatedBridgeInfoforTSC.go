package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Created Bridge Info for TSC' data structure */
type CreatedBridgeInfoforTSC struct {
	DSTTPortNumber      []byte   //DS-TT Port Number
	FiveGSUserPlaneNode []byte   //5GS User Plane Node
	Unknown             []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *CreatedBridgeInfoforTSC) Field() tlv.Field {
	fields := []tlv.Field{}
	if len(t.DSTTPortNumber) > 0 {
		fields = append(fields, tlv.TLVFrom(DS_TT_PORT_NUMBER_TYPE, tlv.Bytes(t.DSTTPortNumber)))
	}
	if len(t.FiveGSUserPlaneNode) > 0 {
		fields = append(fields, tlv.TLVFrom(FIVEGS_USER_PLANE_NODE_TYPE, tlv.Bytes(t.FiveGSUserPlaneNode)))
	}
	return tlv.TLV(CREATED_BRIDGE_INFO_FOR_TSC_TYPE, fields...)
}

func (t *CreatedBridgeInfoforTSC) UnmarshalBinary(buf []byte) (err error) {
	*t = CreatedBridgeInfoforTSC{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [2]bool
	for _, de := range elements {
		switch de.Type {
		case DS_TT_PORT_NUMBER_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.DSTTPortNumber = de.Value
		case FIVEGS_USER_PLANE_NODE_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			t.FiveGSUserPlaneNode = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return pfcp.Faulty(de.Type, err)
		}
	}
	return
}
