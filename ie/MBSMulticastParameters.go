package ie

import (
	"github.com/reogac/pfcp"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'MBS Multicast Parameters' data structure */
type MBSMulticastParameters struct {
	DestinationInterface     uint8               //Destination Interface
	NetworkInstance          []byte              //Network Instance
	OuterHeaderCreation      OuterHeaderCreation //Outer Header Creation
	TransportLevelmarking    []byte              //Transport Level Marking
	DestinationInterfaceType []byte              //3GPP Interface Type
}

func (t *MBSMulticastParameters) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(DESTINATION_INTERFACE_TYPE, tlv.Bytes(pfcp.MarshalUint8(t.DestinationInterface))))
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
	return tlv.TLV(MBS_MULTICAST_PARAMETERS_TYPE, fields...)
}

func (t *MBSMulticastParameters) UnmarshalBinary(buf []byte) (err error) {
	*t = MBSMulticastParameters{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case DESTINATION_INTERFACE_TYPE:
			err = pfcp.UnmarshalUint8(de.Value, &t.DestinationInterface)
		case NETWORK_INSTANCE_TYPE:
			t.NetworkInstance = de.Value
		case OUTER_HEADER_CREATION_TYPE:
			err = de.UnmarshalValue(&t.OuterHeaderCreation)
		case TRANSPORT_LEVEL_MARKING_TYPE:
			t.TransportLevelmarking = de.Value
		case _INTERFACE_TYPE_TYPE:
			t.DestinationInterfaceType = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
