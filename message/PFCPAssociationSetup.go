package message

import (
	"github.com/reogac/pfcp/v2"
	"github.com/reogac/pfcp/v2/ie"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'PFCP Association Setup Request' data structure */
type PFCPAssociationSetupRequest struct {
	NodeID                          ie.NodeID                                                            //Node ID
	RecoveryTimeStamp               uint32                                                               //Recovery Time Stamp
	UPFunctionFeatures              []byte                                                               //UP Function Features
	CPFunctionFeatures              *uint8                                                               //CP Function Features
	UserPlaneIPResourceInformation  [][]byte                                                             //User Plane IP Resource Information
	AlternativeSMFIPAddress         [][]byte                                                             //Alternative SMF IP Address
	SMFSetID                        []byte                                                               //SMF Set ID
	PFCPSessionRetentionInformation *ie.PFCPSessionRetentionInformationwithinPFCPAssociationSetupRequest //PFCP Session Retention Information within PFCP Association Setup Request
	UEIPaddressPoolInformation      []ie.UEIPAddress                                                     //UE IP Address
	GTPUPathQoSControlInformation   [][]byte                                                             //GTP-U Path QoS Control Information
	ClockDriftControlInformation    []ie.ClockDriftControlInformation                                    //Clock Drift Control Information
	UPFInstanceID                   []byte                                                               //NF Instance ID
	PFCPASReqFlags                  []byte                                                               //PFCPASReq-Flags
	Unknown                         []uint32                                                             `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *PFCPAssociationSetupRequest) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(ie.NODE_ID_TYPE, tlv.Bytes(t.NodeID.Bytes())))
	fields = append(fields, tlv.TLVFrom(ie.RECOVERY_TIME_STAMP_TYPE, tlv.Bytes(pfcp.MarshalUint32(t.RecoveryTimeStamp))))
	if len(t.UPFunctionFeatures) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.UP_FUNCTION_FEATURES_TYPE, tlv.Bytes(t.UPFunctionFeatures)))
	}
	if t.CPFunctionFeatures != nil {
		fields = append(fields, tlv.TLVFrom(ie.CP_FUNCTION_FEATURES_TYPE, tlv.Bytes(pfcp.MarshalUint8(*t.CPFunctionFeatures))))
	}
	for _, i := range t.UserPlaneIPResourceInformation {
		fields = append(fields, tlv.TLVFrom(ie.USER_PLANE_IP_RESOURCE_INFORMATION_TYPE, tlv.Bytes(i)))
	}
	for _, i := range t.AlternativeSMFIPAddress {
		fields = append(fields, tlv.TLVFrom(ie.ALTERNATIVE_SMF_IP_ADDRESS_TYPE, tlv.Bytes(i)))
	}
	if len(t.SMFSetID) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.SMF_SET_ID_TYPE, tlv.Bytes(t.SMFSetID)))
	}
	if t.PFCPSessionRetentionInformation != nil {
		fields = append(fields, t.PFCPSessionRetentionInformation.Field())
	}
	for _, i := range t.UEIPaddressPoolInformation {
		fields = append(fields, tlv.TLVFrom(ie.UE_IP_ADDRESS_TYPE, tlv.Bytes(i.Bytes())))
	}
	for _, i := range t.GTPUPathQoSControlInformation {
		fields = append(fields, tlv.TLVFrom(ie.GTP_U_PATH_QOS_CONTROL_INFORMATION_TYPE, tlv.Bytes(i)))
	}
	for _, i := range t.ClockDriftControlInformation {
		fields = append(fields, i.Field())
	}
	if len(t.UPFInstanceID) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.NF_INSTANCE_ID_TYPE, tlv.Bytes(t.UPFInstanceID)))
	}
	if len(t.PFCPASReqFlags) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.PFCPASREQ_FLAGS_TYPE, tlv.Bytes(t.PFCPASReqFlags)))
	}
	return tlv.TLV(ie.MSG_PFCP_ASSOCIATION_SETUP_REQUEST_TYPE, fields...)
}

func (t *PFCPAssociationSetupRequest) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPAssociationSetupRequest{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [8]bool
	for _, de := range elements {
		switch de.Type {
		case ie.NODE_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			if err = de.UnmarshalValue(&t.NodeID); err == nil {
			}
		case ie.RECOVERY_TIME_STAMP_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &t.RecoveryTimeStamp)
		case ie.UP_FUNCTION_FEATURES_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			t.UPFunctionFeatures = de.Value
		case ie.CP_FUNCTION_FEATURES_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			var v uint8
			if err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &v); err == nil {
				t.CPFunctionFeatures = &v
			}
		case ie.USER_PLANE_IP_RESOURCE_INFORMATION_TYPE:
			t.UserPlaneIPResourceInformation = append(t.UserPlaneIPResourceInformation, de.Value)
		case ie.ALTERNATIVE_SMF_IP_ADDRESS_TYPE:
			t.AlternativeSMFIPAddress = append(t.AlternativeSMFIPAddress, de.Value)
		case ie.SMF_SET_ID_TYPE:
			if seen[4] {
				continue
			}
			seen[4] = true
			t.SMFSetID = de.Value
		case ie.PFCP_SESSION_RETENTION_INFORMATION_WITHIN_PFCP_ASSOCIATION_SETUP_REQUEST_TYPE:
			if seen[5] {
				continue
			}
			seen[5] = true
			var v ie.PFCPSessionRetentionInformationwithinPFCPAssociationSetupRequest
			if err = de.UnmarshalValue(&v); err == nil {
				t.PFCPSessionRetentionInformation = &v
				t.Unknown = append(t.Unknown, v.Unknown...)
			}
		case ie.UE_IP_ADDRESS_TYPE:
			var v ie.UEIPAddress
			if err = de.UnmarshalValue(&v); err == nil {
				t.UEIPaddressPoolInformation = append(t.UEIPaddressPoolInformation, v)
			}
		case ie.GTP_U_PATH_QOS_CONTROL_INFORMATION_TYPE:
			t.GTPUPathQoSControlInformation = append(t.GTPUPathQoSControlInformation, de.Value)
		case ie.CLOCK_DRIFT_CONTROL_INFORMATION_TYPE:
			var v ie.ClockDriftControlInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.ClockDriftControlInformation = append(t.ClockDriftControlInformation, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
			}
		case ie.NF_INSTANCE_ID_TYPE:
			if seen[6] {
				continue
			}
			seen[6] = true
			t.UPFInstanceID = de.Value
		case ie.PFCPASREQ_FLAGS_TYPE:
			if seen[7] {
				continue
			}
			seen[7] = true
			t.PFCPASReqFlags = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return pfcp.Faulty(de.Type, err)
		}
	}
	if !seen[0] {
		return &pfcp.MissingIE{Type: ie.NODE_ID_TYPE}
	}
	if !seen[1] {
		return &pfcp.MissingIE{Type: ie.RECOVERY_TIME_STAMP_TYPE}
	}
	return
}

func (t *PFCPAssociationSetupRequest) UnmarshalTLV(typ uint32, buf []byte) (err error) {
	if typ != ie.MSG_PFCP_ASSOCIATION_SETUP_REQUEST_TYPE {
		err = tlv.ErrCritical
		return
	}

	err = t.UnmarshalBinary(buf)
	return
}

func (t *PFCPAssociationSetupRequest) MarshalJSON() (buf []byte, err error) {
	buf, err = jsonMarshal(t)
	return
}

func (t *PFCPAssociationSetupRequest) UnmarshalJSON(buf []byte) (err error) {
	*t = PFCPAssociationSetupRequest{}
	err = jsonUnmarshal(buf, t)
	return
}

/* 'PFCP Association Setup Response' data structure */
type PFCPAssociationSetupResponse struct {
	NodeID                         ie.NodeID                         //Node ID
	Cause                          uint8                             //Cause
	RecoveryTimeStamp              uint32                            //Recovery Time Stamp
	UPFunctionFeatures             []byte                            //UP Function Features
	CPFunctionFeatures             *uint8                            //CP Function Features
	UserPlaneIPResourceInformation [][]byte                          //User Plane IP Resource Information
	AlternativeSMFIPAddress        [][]byte                          //Alternative SMF IP Address
	SMFSetID                       []byte                            //SMF Set ID
	PFCPASRspFlags                 []byte                            //PFCPASRsp-Flags
	ClockDriftControlInformation   []ie.ClockDriftControlInformation //Clock Drift Control Information
	UEIPaddressPoolInformation     []ie.UEIPAddress                  //UE IP Address
	GTPUPathQoSControlInformation  [][]byte                          //GTP-U Path QoS Control Information
	UPFInstanceID                  []byte                            //NF Instance ID
	Unknown                        []uint32                          `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *PFCPAssociationSetupResponse) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(ie.NODE_ID_TYPE, tlv.Bytes(t.NodeID.Bytes())))
	fields = append(fields, tlv.TLVFrom(ie.CAUSE_TYPE, tlv.Bytes(pfcp.MarshalUint8(t.Cause))))
	fields = append(fields, tlv.TLVFrom(ie.RECOVERY_TIME_STAMP_TYPE, tlv.Bytes(pfcp.MarshalUint32(t.RecoveryTimeStamp))))
	if len(t.UPFunctionFeatures) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.UP_FUNCTION_FEATURES_TYPE, tlv.Bytes(t.UPFunctionFeatures)))
	}
	if t.CPFunctionFeatures != nil {
		fields = append(fields, tlv.TLVFrom(ie.CP_FUNCTION_FEATURES_TYPE, tlv.Bytes(pfcp.MarshalUint8(*t.CPFunctionFeatures))))
	}
	for _, i := range t.UserPlaneIPResourceInformation {
		fields = append(fields, tlv.TLVFrom(ie.USER_PLANE_IP_RESOURCE_INFORMATION_TYPE, tlv.Bytes(i)))
	}
	for _, i := range t.AlternativeSMFIPAddress {
		fields = append(fields, tlv.TLVFrom(ie.ALTERNATIVE_SMF_IP_ADDRESS_TYPE, tlv.Bytes(i)))
	}
	if len(t.SMFSetID) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.SMF_SET_ID_TYPE, tlv.Bytes(t.SMFSetID)))
	}
	if len(t.PFCPASRspFlags) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.PFCPASRSP_FLAGS_TYPE, tlv.Bytes(t.PFCPASRspFlags)))
	}
	for _, i := range t.ClockDriftControlInformation {
		fields = append(fields, i.Field())
	}
	for _, i := range t.UEIPaddressPoolInformation {
		fields = append(fields, tlv.TLVFrom(ie.UE_IP_ADDRESS_TYPE, tlv.Bytes(i.Bytes())))
	}
	for _, i := range t.GTPUPathQoSControlInformation {
		fields = append(fields, tlv.TLVFrom(ie.GTP_U_PATH_QOS_CONTROL_INFORMATION_TYPE, tlv.Bytes(i)))
	}
	if len(t.UPFInstanceID) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.NF_INSTANCE_ID_TYPE, tlv.Bytes(t.UPFInstanceID)))
	}
	return tlv.TLV(ie.MSG_PFCP_ASSOCIATION_SETUP_RESPONSE_TYPE, fields...)
}

func (t *PFCPAssociationSetupResponse) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPAssociationSetupResponse{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [8]bool
	for _, de := range elements {
		switch de.Type {
		case ie.NODE_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			if err = de.UnmarshalValue(&t.NodeID); err == nil {
			}
		case ie.CAUSE_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			err = pfcp.UnmarshalUint8(de.Value, &t.Cause)
		case ie.RECOVERY_TIME_STAMP_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			err = pfcp.UnmarshalExtendable(de.Value, 4, 4, &t.RecoveryTimeStamp)
		case ie.UP_FUNCTION_FEATURES_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			t.UPFunctionFeatures = de.Value
		case ie.CP_FUNCTION_FEATURES_TYPE:
			if seen[4] {
				continue
			}
			seen[4] = true
			var v uint8
			if err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &v); err == nil {
				t.CPFunctionFeatures = &v
			}
		case ie.USER_PLANE_IP_RESOURCE_INFORMATION_TYPE:
			t.UserPlaneIPResourceInformation = append(t.UserPlaneIPResourceInformation, de.Value)
		case ie.ALTERNATIVE_SMF_IP_ADDRESS_TYPE:
			t.AlternativeSMFIPAddress = append(t.AlternativeSMFIPAddress, de.Value)
		case ie.SMF_SET_ID_TYPE:
			if seen[5] {
				continue
			}
			seen[5] = true
			t.SMFSetID = de.Value
		case ie.PFCPASRSP_FLAGS_TYPE:
			if seen[6] {
				continue
			}
			seen[6] = true
			t.PFCPASRspFlags = de.Value
		case ie.CLOCK_DRIFT_CONTROL_INFORMATION_TYPE:
			var v ie.ClockDriftControlInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.ClockDriftControlInformation = append(t.ClockDriftControlInformation, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
			}
		case ie.UE_IP_ADDRESS_TYPE:
			var v ie.UEIPAddress
			if err = de.UnmarshalValue(&v); err == nil {
				t.UEIPaddressPoolInformation = append(t.UEIPaddressPoolInformation, v)
			}
		case ie.GTP_U_PATH_QOS_CONTROL_INFORMATION_TYPE:
			t.GTPUPathQoSControlInformation = append(t.GTPUPathQoSControlInformation, de.Value)
		case ie.NF_INSTANCE_ID_TYPE:
			if seen[7] {
				continue
			}
			seen[7] = true
			t.UPFInstanceID = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return pfcp.Faulty(de.Type, err)
		}
	}
	if !seen[1] {
		return &pfcp.MissingIE{Type: ie.CAUSE_TYPE}
	}
	if t.Cause < ie.CauseRequestRejected {
		if !seen[0] {
			return &pfcp.MissingIE{Type: ie.NODE_ID_TYPE}
		}
		if !seen[2] {
			return &pfcp.MissingIE{Type: ie.RECOVERY_TIME_STAMP_TYPE}
		}
	}
	return
}

func (t *PFCPAssociationSetupResponse) UnmarshalTLV(typ uint32, buf []byte) (err error) {
	if typ != ie.MSG_PFCP_ASSOCIATION_SETUP_RESPONSE_TYPE {
		err = tlv.ErrCritical
		return
	}

	err = t.UnmarshalBinary(buf)
	return
}

func (t *PFCPAssociationSetupResponse) MarshalJSON() (buf []byte, err error) {
	buf, err = jsonMarshal(t)
	return
}

func (t *PFCPAssociationSetupResponse) UnmarshalJSON(buf []byte) (err error) {
	*t = PFCPAssociationSetupResponse{}
	err = jsonUnmarshal(buf, t)
	return
}
