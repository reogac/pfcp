package ie

import (
	"github.com/reogac/pfcp"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Create PDR' data structure */
type CreatePDR struct {
	PDRID                                           uint16                                                          //PDR ID
	Precedence                                      uint32                                                          //Precedence
	PDI                                             *PDI                                                            //PDI
	OuterHeaderRemoval                              *OuterHeaderRemoval                                             //Outer Header Removal
	FARID                                           *uint32                                                         //FAR ID
	URRID                                           []uint32                                                        //URR ID
	QERID                                           *uint32                                                         //QER ID
	ActivatePredefinedRules                         []byte                                                          //Activate Predefined Rules
	ActivationTime                                  []byte                                                          //Activation Time
	DeactivationTime                                []byte                                                          //Deactivation Time
	MARID                                           []byte                                                          //MAR ID
	PacketReplicationandDetectionCarryOnInformation []byte                                                          //Packet Replication and Detection Carry-On Information
	IPMulticastAddressingInfo                       *IPMulticastAddressingInfowithinPFCPSessionEstablishmentRequest //IP Multicast Addressing Info within PFCP Session Establishment Request
	UEIPaddressPoolIdentity                         *UEIPAddress                                                    //UE IP Address
	MPTCPApplicableIndication                       []byte                                                          //MPTCP Applicable Indication
	TransportDelayReporting                         *TransportDelayReporting                                        //Transport Delay Reporting
	RATType                                         []byte                                                          //RAT Type
}

func (t *CreatePDR) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(PDR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint16(t.PDRID))))
	fields = append(fields, tlv.TLVFrom(PRECEDENCE_TYPE, tlv.Bytes(pfcp.MarshalUint32(t.Precedence))))
	if t.PDI != nil {
		fields = append(fields, t.PDI.Field())
	}
	if t.OuterHeaderRemoval != nil {
		fields = append(fields, tlv.TLVFrom(OUTER_HEADER_REMOVAL_TYPE, tlv.Bytes(t.OuterHeaderRemoval.Bytes())))
	}
	if t.FARID != nil {
		fields = append(fields, tlv.TLVFrom(FAR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(*t.FARID))))
	}
	for _, ie := range t.URRID {
		fields = append(fields, tlv.TLVFrom(URR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(ie))))
	}
	if t.QERID != nil {
		fields = append(fields, tlv.TLVFrom(QER_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(*t.QERID))))
	}
	if len(t.ActivatePredefinedRules) > 0 {
		fields = append(fields, tlv.TLVFrom(ACTIVATE_PREDEFINED_RULES_TYPE, tlv.Bytes(t.ActivatePredefinedRules)))
	}
	if len(t.ActivationTime) > 0 {
		fields = append(fields, tlv.TLVFrom(ACTIVATION_TIME_TYPE, tlv.Bytes(t.ActivationTime)))
	}
	if len(t.DeactivationTime) > 0 {
		fields = append(fields, tlv.TLVFrom(DEACTIVATION_TIME_TYPE, tlv.Bytes(t.DeactivationTime)))
	}
	if len(t.MARID) > 0 {
		fields = append(fields, tlv.TLVFrom(MAR_ID_TYPE, tlv.Bytes(t.MARID)))
	}
	if len(t.PacketReplicationandDetectionCarryOnInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(PACKET_REPLICATION_AND_DETECTION_CARRY_ON_INFORMATION_TYPE, tlv.Bytes(t.PacketReplicationandDetectionCarryOnInformation)))
	}
	if t.IPMulticastAddressingInfo != nil {
		fields = append(fields, t.IPMulticastAddressingInfo.Field())
	}
	if t.UEIPaddressPoolIdentity != nil {
		fields = append(fields, tlv.TLVFrom(UE_IP_ADDRESS_TYPE, tlv.Bytes(t.UEIPaddressPoolIdentity.Bytes())))
	}
	if len(t.MPTCPApplicableIndication) > 0 {
		fields = append(fields, tlv.TLVFrom(MPTCP_APPLICABLE_INDICATION_TYPE, tlv.Bytes(t.MPTCPApplicableIndication)))
	}
	if t.TransportDelayReporting != nil {
		fields = append(fields, t.TransportDelayReporting.Field())
	}
	if len(t.RATType) > 0 {
		fields = append(fields, tlv.TLVFrom(RAT_TYPE_TYPE, tlv.Bytes(t.RATType)))
	}
	return tlv.TLV(CREATE_PDR_TYPE, fields...)
}

func (t *CreatePDR) UnmarshalBinary(buf []byte) (err error) {
	*t = CreatePDR{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case PDR_ID_TYPE:
			err = pfcp.UnmarshalUint16(de.Value, &t.PDRID)
		case PRECEDENCE_TYPE:
			err = pfcp.UnmarshalUint32(de.Value, &t.Precedence)
		case PDI_TYPE:
			var v PDI
			if err = de.UnmarshalValue(&v); err == nil {
				t.PDI = &v
			}
		case OUTER_HEADER_REMOVAL_TYPE:
			var v OuterHeaderRemoval
			if err = de.UnmarshalValue(&v); err == nil {
				t.OuterHeaderRemoval = &v
			}
		case FAR_ID_TYPE:
			var v uint32
			if err = pfcp.UnmarshalUint32(de.Value, &v); err == nil {
				t.FARID = &v
			}
		case URR_ID_TYPE:
			var v uint32
			if err = pfcp.UnmarshalUint32(de.Value, &v); err == nil {
				t.URRID = append(t.URRID, v)
			}
		case QER_ID_TYPE:
			var v uint32
			if err = pfcp.UnmarshalUint32(de.Value, &v); err == nil {
				t.QERID = &v
			}
		case ACTIVATE_PREDEFINED_RULES_TYPE:
			t.ActivatePredefinedRules = de.Value
		case ACTIVATION_TIME_TYPE:
			t.ActivationTime = de.Value
		case DEACTIVATION_TIME_TYPE:
			t.DeactivationTime = de.Value
		case MAR_ID_TYPE:
			t.MARID = de.Value
		case PACKET_REPLICATION_AND_DETECTION_CARRY_ON_INFORMATION_TYPE:
			t.PacketReplicationandDetectionCarryOnInformation = de.Value
		case IP_MULTICAST_ADDRESSING_INFO_WITHIN_PFCP_SESSION_ESTABLISHMENT_REQUEST_TYPE:
			var v IPMulticastAddressingInfowithinPFCPSessionEstablishmentRequest
			if err = de.UnmarshalValue(&v); err == nil {
				t.IPMulticastAddressingInfo = &v
			}
		case UE_IP_ADDRESS_TYPE:
			var v UEIPAddress
			if err = de.UnmarshalValue(&v); err == nil {
				t.UEIPaddressPoolIdentity = &v
			}
		case MPTCP_APPLICABLE_INDICATION_TYPE:
			t.MPTCPApplicableIndication = de.Value
		case TRANSPORT_DELAY_REPORTING_TYPE:
			var v TransportDelayReporting
			if err = de.UnmarshalValue(&v); err == nil {
				t.TransportDelayReporting = &v
			}
		case RAT_TYPE_TYPE:
			t.RATType = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
