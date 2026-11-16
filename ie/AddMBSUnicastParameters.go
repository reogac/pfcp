package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Add MBS Unicast Parameters' data structure */
type AddMBSUnicastParameters struct {
	DestinationInterface     uint8               //Destination Interface
	MBSUnicastParametersID   []byte              //MBS Unicast Parameters ID
	NetworkInstance          []byte              //Network Instance
	OuterHeaderCreation      OuterHeaderCreation //Outer Header Creation
	TransportLevelmarking    []byte              //Transport Level Marking
	DestinationInterfaceType []byte              //3GPP Interface Type
	Unknown                  []uint32            `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *AddMBSUnicastParameters) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(DESTINATION_INTERFACE_TYPE, tlv.Bytes(pfcp.MarshalUint8(t.DestinationInterface))))
	fields = append(fields, tlv.TLVFrom(MBS_UNICAST_PARAMETERS_ID_TYPE, tlv.Bytes(t.MBSUnicastParametersID)))
	if len(t.NetworkInstance) > 0 {
		fields = append(fields, tlv.TLVFrom(NETWORK_INSTANCE_TYPE, tlv.Bytes(t.NetworkInstance)))
	}
	fields = append(fields, tlv.TLVFrom(OUTER_HEADER_CREATION_TYPE, tlv.Bytes(t.OuterHeaderCreation.Bytes())))
	if len(t.TransportLevelmarking) > 0 {
		fields = append(fields, tlv.TLVFrom(TRANSPORT_LEVEL_MARKING_TYPE, tlv.Bytes(t.TransportLevelmarking)))
	}
	if len(t.DestinationInterfaceType) > 0 {
		fields = append(fields, tlv.TLVFrom(_INTERFACE_TYPE_TYPE, tlv.Bytes(t.DestinationInterfaceType)))
	}
	return tlv.TLV(ADD_MBS_UNICAST_PARAMETERS_TYPE, fields...)
}

func (t *AddMBSUnicastParameters) UnmarshalBinary(buf []byte) (err error) {
	*t = AddMBSUnicastParameters{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	var seen [6]bool
	for _, de := range elements {
		switch de.Type {
		case DESTINATION_INTERFACE_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &t.DestinationInterface)
		case MBS_UNICAST_PARAMETERS_ID_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			t.MBSUnicastParametersID = de.Value
		case NETWORK_INSTANCE_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			t.NetworkInstance = de.Value
		case OUTER_HEADER_CREATION_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			if err = de.UnmarshalValue(&t.OuterHeaderCreation); err == nil {
			}
		case TRANSPORT_LEVEL_MARKING_TYPE:
			if seen[4] {
				continue
			}
			seen[4] = true
			t.TransportLevelmarking = de.Value
		case _INTERFACE_TYPE_TYPE:
			if seen[5] {
				continue
			}
			seen[5] = true
			t.DestinationInterfaceType = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	return
}
