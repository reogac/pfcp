package ie

import (
	"github.com/reogac/pfcp"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'Update PDR' data structure */
type UpdatePDR struct {
	PDRID                     uint16                                                          //PDR ID
	OuterHeaderRemoval        *OuterHeaderRemoval                                             //Outer Header Removal
	Precedence                *uint32                                                         //Precedence
	PDI                       *PDI                                                            //PDI
	FARID                     *uint32                                                         //FAR ID
	URRID                     *uint32                                                         //URR ID
	QERID                     *uint32                                                         //QER ID
	ActivatePredefinedRules   []byte                                                          //Activate Predefined Rules
	DeactivatePredefinedRules []byte                                                          //Deactivate Predefined Rules
	ActivationTime            []byte                                                          //Activation Time
	DeactivationTime          []byte                                                          //Deactivation Time
	IPMulticastAddressingInfo *IPMulticastAddressingInfowithinPFCPSessionEstablishmentRequest //IP Multicast Addressing Info within PFCP Session Establishment Request
	TransportDelayReporting   *TransportDelayReporting                                        //Transport Delay Reporting
	RATType                   []byte                                                          //RAT Type
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
	if t.URRID != nil {
		fields = append(fields, tlv.TLVFrom(URR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(*t.URRID))))
	}
	if t.QERID != nil {
		fields = append(fields, tlv.TLVFrom(QER_ID_TYPE, tlv.Bytes(pfcp.MarshalUint32(*t.QERID))))
	}
	if len(t.ActivatePredefinedRules) > 0 {
		fields = append(fields, tlv.TLVFrom(ACTIVATE_PREDEFINED_RULES_TYPE, tlv.Bytes(t.ActivatePredefinedRules)))
	}
	if len(t.DeactivatePredefinedRules) > 0 {
		fields = append(fields, tlv.TLVFrom(DEACTIVATE_PREDEFINED_RULES_TYPE, tlv.Bytes(t.DeactivatePredefinedRules)))
	}
	if len(t.ActivationTime) > 0 {
		fields = append(fields, tlv.TLVFrom(ACTIVATION_TIME_TYPE, tlv.Bytes(t.ActivationTime)))
	}
	if len(t.DeactivationTime) > 0 {
		fields = append(fields, tlv.TLVFrom(DEACTIVATION_TIME_TYPE, tlv.Bytes(t.DeactivationTime)))
	}
	if t.IPMulticastAddressingInfo != nil {
		fields = append(fields, t.IPMulticastAddressingInfo.Field())
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
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case PDR_ID_TYPE:
			err = pfcp.UnmarshalUint16(de.Value, &t.PDRID)
		case OUTER_HEADER_REMOVAL_TYPE:
			var v OuterHeaderRemoval
			if err = de.UnmarshalValue(&v); err == nil {
				t.OuterHeaderRemoval = &v
			}
		case PRECEDENCE_TYPE:
			var v uint32
			if err = pfcp.UnmarshalUint32(de.Value, &v); err == nil {
				t.Precedence = &v
			}
		case PDI_TYPE:
			var v PDI
			if err = de.UnmarshalValue(&v); err == nil {
				t.PDI = &v
			}
		case FAR_ID_TYPE:
			var v uint32
			if err = pfcp.UnmarshalUint32(de.Value, &v); err == nil {
				t.FARID = &v
			}
		case URR_ID_TYPE:
			var v uint32
			if err = pfcp.UnmarshalUint32(de.Value, &v); err == nil {
				t.URRID = &v
			}
		case QER_ID_TYPE:
			var v uint32
			if err = pfcp.UnmarshalUint32(de.Value, &v); err == nil {
				t.QERID = &v
			}
		case ACTIVATE_PREDEFINED_RULES_TYPE:
			t.ActivatePredefinedRules = de.Value
		case DEACTIVATE_PREDEFINED_RULES_TYPE:
			t.DeactivatePredefinedRules = de.Value
		case ACTIVATION_TIME_TYPE:
			t.ActivationTime = de.Value
		case DEACTIVATION_TIME_TYPE:
			t.DeactivationTime = de.Value
		case IP_MULTICAST_ADDRESSING_INFO_WITHIN_PFCP_SESSION_ESTABLISHMENT_REQUEST_TYPE:
			var v IPMulticastAddressingInfowithinPFCPSessionEstablishmentRequest
			if err = de.UnmarshalValue(&v); err == nil {
				t.IPMulticastAddressingInfo = &v
			}
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
