package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Create PDR' data structure */
type CreatePDR struct {
	PDRID                                           uint16                                                           //PDR ID
	Precedence                                      uint32                                                           //Precedence
	PDI                                             PDI                                                              //PDI
	OuterHeaderRemoval                              *OuterHeaderRemoval                                              //Outer Header Removal
	FARID                                           *uint32                                                          //FAR ID
	URRID                                           []uint32                                                         //URR ID
	QERID                                           []uint32                                                         //QER ID
	ActivatePredefinedRules                         [][]byte                                                         //Activate Predefined Rules
	ActivationTime                                  []byte                                                           //Activation Time
	DeactivationTime                                []byte                                                           //Deactivation Time
	MARID                                           []byte                                                           //MAR ID
	PacketReplicationandDetectionCarryOnInformation []byte                                                           //Packet Replication and Detection Carry-On Information
	IPMulticastAddressingInfo                       []IPMulticastAddressingInfowithinPFCPSessionEstablishmentRequest //IP Multicast Addressing Info within PFCP Session Establishment Request
	UEIPaddressPoolIdentity                         []UEIPAddress                                                    //UE IP Address
	MPTCPApplicableIndication                       []byte                                                           //MPTCP Applicable Indication
	TransportDelayReporting                         *TransportDelayReporting                                         //Transport Delay Reporting
	RATType                                         []byte                                                           //RAT Type
	Unknown                                         []uint32                                                         `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *CreatePDR) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(PDR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint16(t.PDRID))))
	fields = append(fields, tlv.TLVFrom(PRECEDENCE_TYPE, tlv.Bytes(pfcp.MarshalUint32(t.Precedence))))
	fields = append(fields, t.PDI.Field())
	if t.OuterHeaderRemoval != nil {
		fields = append(fields, tlv.TLVFrom(OUTER_HEADER_REMOVAL_TYPE, tlv.Bytes(t.OuterHeaderRemoval.Bytes())))
	}
	if t.FARID != nil {
		fields = append(fields, tlv.TLVFrom(FAR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(*t.FARID))))
	}
	for _, ie := range t.URRID {
		fields = append(fields, tlv.TLVFrom(URR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(ie))))
	}
	for _, ie := range t.QERID {
		fields = append(fields, tlv.TLVFrom(QER_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(ie))))
	}
	for _, ie := range t.ActivatePredefinedRules {
		fields = append(fields, tlv.TLVFrom(ACTIVATE_PREDEFINED_RULES_TYPE, tlv.Bytes(ie)))
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
	for _, ie := range t.IPMulticastAddressingInfo {
		fields = append(fields, ie.Field())
	}
	for _, ie := range t.UEIPaddressPoolIdentity {
		fields = append(fields, tlv.TLVFrom(UE_IP_ADDRESS_TYPE, tlv.Bytes(ie.Bytes())))
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
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [12]bool
	for _, de := range elements {
		switch de.Type {
		case PDR_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			err = pfcp.UnmarshalExtendable(de.Value, 2, 2, &t.PDRID)
		case PRECEDENCE_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &t.Precedence)
		case PDI_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			if err = de.UnmarshalValue(&t.PDI); err == nil {
				t.Unknown = append(t.Unknown, t.PDI.Unknown...)
			}
		case OUTER_HEADER_REMOVAL_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			var v OuterHeaderRemoval
			if err = de.UnmarshalValue(&v); err == nil {
				t.OuterHeaderRemoval = &v
			}
		case FAR_ID_TYPE:
			if seen[4] {
				continue
			}
			seen[4] = true
			var v uint32
			if err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &v); err == nil {
				t.FARID = &v
			}
		case URR_ID_TYPE:
			var v uint32
			if err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &v); err == nil {
				t.URRID = append(t.URRID, v)
			}
		case QER_ID_TYPE:
			var v uint32
			if err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &v); err == nil {
				t.QERID = append(t.QERID, v)
			}
		case ACTIVATE_PREDEFINED_RULES_TYPE:
			t.ActivatePredefinedRules = append(t.ActivatePredefinedRules, de.Value)
		case ACTIVATION_TIME_TYPE:
			if seen[5] {
				continue
			}
			seen[5] = true
			t.ActivationTime = de.Value
		case DEACTIVATION_TIME_TYPE:
			if seen[6] {
				continue
			}
			seen[6] = true
			t.DeactivationTime = de.Value
		case MAR_ID_TYPE:
			if seen[7] {
				continue
			}
			seen[7] = true
			t.MARID = de.Value
		case PACKET_REPLICATION_AND_DETECTION_CARRY_ON_INFORMATION_TYPE:
			if seen[8] {
				continue
			}
			seen[8] = true
			t.PacketReplicationandDetectionCarryOnInformation = de.Value
		case IP_MULTICAST_ADDRESSING_INFO_WITHIN_PFCP_SESSION_ESTABLISHMENT_REQUEST_TYPE:
			var v IPMulticastAddressingInfowithinPFCPSessionEstablishmentRequest
			if err = de.UnmarshalValue(&v); err == nil {
				t.IPMulticastAddressingInfo = append(t.IPMulticastAddressingInfo, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
			}
		case UE_IP_ADDRESS_TYPE:
			var v UEIPAddress
			if err = de.UnmarshalValue(&v); err == nil {
				t.UEIPaddressPoolIdentity = append(t.UEIPaddressPoolIdentity, v)
			}
		case MPTCP_APPLICABLE_INDICATION_TYPE:
			if seen[9] {
				continue
			}
			seen[9] = true
			t.MPTCPApplicableIndication = de.Value
		case TRANSPORT_DELAY_REPORTING_TYPE:
			if seen[10] {
				continue
			}
			seen[10] = true
			var v TransportDelayReporting
			if err = de.UnmarshalValue(&v); err == nil {
				t.TransportDelayReporting = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
			}
		case RAT_TYPE_TYPE:
			if seen[11] {
				continue
			}
			seen[11] = true
			t.RATType = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return pfcp.Faulty(de.Type, err)
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: PDR_ID_TYPE}
	}
	if !seen[1] {
		return &pfcp.MissingIE{Type: PRECEDENCE_TYPE}
	}
	if !seen[2] {
		return &pfcp.MissingIE{Type: PDI_TYPE}
	}
	return
}
