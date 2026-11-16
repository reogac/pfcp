package ie

import (
	"github.com/reogac/pfcp/v2"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'Update PDR' data structure */
type UpdatePDR struct {
	PDRID                     uint16                                                           //PDR ID
	OuterHeaderRemoval        *OuterHeaderRemoval                                              //Outer Header Removal
	Precedence                *uint32                                                          //Precedence
	PDI                       *PDI                                                             //PDI
	FARID                     *uint32                                                          //FAR ID
	URRID                     []uint32                                                         //URR ID
	QERID                     []uint32                                                         //QER ID
	ActivatePredefinedRules   [][]byte                                                         //Activate Predefined Rules
	DeactivatePredefinedRules [][]byte                                                         //Deactivate Predefined Rules
	ActivationTime            []byte                                                           //Activation Time
	DeactivationTime          []byte                                                           //Deactivation Time
	IPMulticastAddressingInfo []IPMulticastAddressingInfowithinPFCPSessionEstablishmentRequest //IP Multicast Addressing Info within PFCP Session Establishment Request
	TransportDelayReporting   *TransportDelayReporting                                         //Transport Delay Reporting
	RATType                   []byte                                                           //RAT Type
	Unknown                   []uint32                                                         `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid                   []uint32                                                         `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
}

func (t *UpdatePDR) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(PDR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint16(t.PDRID))))
	if t.OuterHeaderRemoval != nil {
		fields = append(fields, tlv.TLVFrom(OUTER_HEADER_REMOVAL_TYPE, tlv.Bytes(t.OuterHeaderRemoval.Bytes())))
	}
	if t.Precedence != nil {
		fields = append(fields, tlv.TLVFrom(PRECEDENCE_TYPE, tlv.Bytes(pfcp.MarshalUint32(*t.Precedence))))
	}
	if t.PDI != nil {
		fields = append(fields, t.PDI.Field())
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
	for _, ie := range t.DeactivatePredefinedRules {
		fields = append(fields, tlv.TLVFrom(DEACTIVATE_PREDEFINED_RULES_TYPE, tlv.Bytes(ie)))
	}
	if len(t.ActivationTime) > 0 {
		fields = append(fields, tlv.TLVFrom(ACTIVATION_TIME_TYPE, tlv.Bytes(t.ActivationTime)))
	}
	if len(t.DeactivationTime) > 0 {
		fields = append(fields, tlv.TLVFrom(DEACTIVATION_TIME_TYPE, tlv.Bytes(t.DeactivationTime)))
	}
	for _, ie := range t.IPMulticastAddressingInfo {
		fields = append(fields, ie.Field())
	}
	if t.TransportDelayReporting != nil {
		fields = append(fields, t.TransportDelayReporting.Field())
	}
	if len(t.RATType) > 0 {
		fields = append(fields, tlv.TLVFrom(RAT_TYPE_TYPE, tlv.Bytes(t.RATType)))
	}
	return tlv.TLV(UPDATE_PDR_TYPE, fields...)
}

func (t *UpdatePDR) UnmarshalBinary(buf []byte) (err error) {
	*t = UpdatePDR{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [9]bool
	for _, de := range elements {
		opt := false
		switch de.Type {
		case PDR_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			err = pfcp.UnmarshalExtendable(de.Value, 2, 2, &t.PDRID)
		case OUTER_HEADER_REMOVAL_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			var v OuterHeaderRemoval
			if err = de.UnmarshalValue(&v); err == nil {
				t.OuterHeaderRemoval = &v
			}
		case PRECEDENCE_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			var v uint32
			if err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &v); err == nil {
				t.Precedence = &v
			}
		case PDI_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			var v PDI
			if err = de.UnmarshalValue(&v); err == nil {
				t.PDI = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
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
		case DEACTIVATE_PREDEFINED_RULES_TYPE:
			t.DeactivatePredefinedRules = append(t.DeactivatePredefinedRules, de.Value)
		case ACTIVATION_TIME_TYPE:
			opt = true
			if seen[5] {
				continue
			}
			seen[5] = true
			t.ActivationTime = de.Value
		case DEACTIVATION_TIME_TYPE:
			opt = true
			if seen[6] {
				continue
			}
			seen[6] = true
			t.DeactivationTime = de.Value
		case IP_MULTICAST_ADDRESSING_INFO_WITHIN_PFCP_SESSION_ESTABLISHMENT_REQUEST_TYPE:
			opt = true
			var v IPMulticastAddressingInfowithinPFCPSessionEstablishmentRequest
			if err = de.UnmarshalValue(&v); err == nil {
				t.IPMulticastAddressingInfo = append(t.IPMulticastAddressingInfo, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case TRANSPORT_DELAY_REPORTING_TYPE:
			if seen[7] {
				continue
			}
			seen[7] = true
			var v TransportDelayReporting
			if err = de.UnmarshalValue(&v); err == nil {
				t.TransportDelayReporting = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case RAT_TYPE_TYPE:
			opt = true
			if seen[8] {
				continue
			}
			seen[8] = true
			t.RATType = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			if err = pfcp.Faulty(de.Type, err); opt && pfcp.Discardable(err) {
				t.Invalid = append(t.Invalid, de.Type)
				err = nil
				continue
			}
			return err
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: PDR_ID_TYPE}
	}
	return
}
