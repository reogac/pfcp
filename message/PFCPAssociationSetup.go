package message

import (
	"etrib5gc/pfcp"
	"etrib5gc/pfcp/ie"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'PFCP Association Setup Request' data structure */
type PFCPAssociationSetupRequest struct {
	NodeID                          ie.NodeID                                                            //Node ID
	RecoveryTimeStamp               uint32                                                               //Recovery Time Stamp
	UPFunctionFeatures              []byte                                                               //UP Function Features
	CPFunctionFeatures              *uint8                                                               //CP Function Features
	UserPlaneIPResourceInformation  []byte                                                               //User Plane IP Resource Information
	AlternativeSMFIPAddress         []byte                                                               //Alternative SMF IP Address
	SMFSetID                        []byte                                                               //SMF Set ID
	PFCPSessionRetentionInformation *ie.PFCPSessionRetentionInformationwithinPFCPAssociationSetupRequest //PFCP Session Retention Information within PFCP Association Setup Request
	UEIPaddressPoolInformation      *ie.UEIPAddress                                                      //UE IP Address
	GTPUPathQoSControlInformation   []byte                                                               //GTP-U Path QoS Control Information
	ClockDriftControlInformation    *ie.ClockDriftControlInformation                                     //Clock Drift Control Information
	UPFInstanceID                   []byte                                                               //NF Instance ID
	PFCPASReqFlags                  []byte                                                               //PFCPASReq-Flags
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
	for len(t.UserPlaneIPResourceInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.USER_PLANE_IP_RESOURCE_INFORMATION_TYPE, tlv.Bytes(t.UserPlaneIPResourceInformation)))
	}

	if len(t.AlternativeSMFIPAddress) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.ALTERNATIVE_SMF_IP_ADDRESS_TYPE, tlv.Bytes(t.AlternativeSMFIPAddress)))
	}
	if len(t.SMFSetID) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.SMF_SET_ID_TYPE, tlv.Bytes(t.SMFSetID)))
	}
	if t.PFCPSessionRetentionInformation != nil {
		fields = append(fields, t.PFCPSessionRetentionInformation.Field())
	}
	if t.UEIPaddressPoolInformation != nil {
		fields = append(fields, tlv.TLVFrom(ie.UE_IP_ADDRESS_TYPE, tlv.Bytes(t.UEIPaddressPoolInformation.Bytes())))
	}
	if len(t.GTPUPathQoSControlInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.GTP_U_PATH_QOS_CONTROL_INFORMATION_TYPE, tlv.Bytes(t.GTPUPathQoSControlInformation)))
	}
	if t.ClockDriftControlInformation != nil {
		fields = append(fields, t.ClockDriftControlInformation.Field())
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
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case ie.NODE_ID_TYPE:
			err = de.UnmarshalValue(&t.NodeID)
		case ie.RECOVERY_TIME_STAMP_TYPE:
			err = pfcp.UnmarshalUint32(de.Value, &t.RecoveryTimeStamp)
		case ie.UP_FUNCTION_FEATURES_TYPE:
			t.UPFunctionFeatures = de.Value
		case ie.CP_FUNCTION_FEATURES_TYPE:
			var v uint8
			if err = pfcp.UnmarshalUint8(de.Value, &v); err == nil {
				t.CPFunctionFeatures = &v
			}
		case ie.USER_PLANE_IP_RESOURCE_INFORMATION_TYPE:
			t.UserPlaneIPResourceInformation = de.Value
		case ie.ALTERNATIVE_SMF_IP_ADDRESS_TYPE:
			t.AlternativeSMFIPAddress = de.Value
		case ie.SMF_SET_ID_TYPE:
			t.SMFSetID = de.Value
		case ie.PFCP_SESSION_RETENTION_INFORMATION_WITHIN_PFCP_ASSOCIATION_SETUP_REQUEST_TYPE:
			var v ie.PFCPSessionRetentionInformationwithinPFCPAssociationSetupRequest
			if err = de.UnmarshalValue(&v); err == nil {
				t.PFCPSessionRetentionInformation = &v
			}
		case ie.UE_IP_ADDRESS_TYPE:
			var v ie.UEIPAddress
			if err = de.UnmarshalValue(&v); err == nil {
				t.UEIPaddressPoolInformation = &v
			}
		case ie.GTP_U_PATH_QOS_CONTROL_INFORMATION_TYPE:
			t.GTPUPathQoSControlInformation = de.Value
		case ie.CLOCK_DRIFT_CONTROL_INFORMATION_TYPE:
			var v ie.ClockDriftControlInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.ClockDriftControlInformation = &v
			}
		case ie.NF_INSTANCE_ID_TYPE:
			t.UPFInstanceID = de.Value
		case ie.PFCPASREQ_FLAGS_TYPE:
			t.PFCPASReqFlags = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
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
	NodeID                         ie.NodeID                        //Node ID
	Cause                          uint8                            //Cause
	RecoveryTimeStamp              uint32                           //Recovery Time Stamp
	UPFunctionFeatures             []byte                           //UP Function Features
	CPFunctionFeatures             *uint8                           //CP Function Features
	UserPlaneIPResourceInformation []byte                           //User Plane IP Resource Information
	AlternativeSMFIPAddress        []byte                           //Alternative SMF IP Address
	SMFSetID                       []byte                           //SMF Set ID
	PFCPASRspFlags                 []byte                           //PFCPASRsp-Flags
	ClockDriftControlInformation   *ie.ClockDriftControlInformation //Clock Drift Control Information
	UEIPaddressPoolInformation     *ie.UEIPAddress                  //UE IP Address
	GTPUPathQoSControlInformation  []byte                           //GTP-U Path QoS Control Information
	UPFInstanceID                  []byte                           //NF Instance ID
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
	for len(t.UserPlaneIPResourceInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.USER_PLANE_IP_RESOURCE_INFORMATION_TYPE, tlv.Bytes(t.UserPlaneIPResourceInformation)))
	}
	if len(t.AlternativeSMFIPAddress) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.ALTERNATIVE_SMF_IP_ADDRESS_TYPE, tlv.Bytes(t.AlternativeSMFIPAddress)))
	}
	if len(t.SMFSetID) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.SMF_SET_ID_TYPE, tlv.Bytes(t.SMFSetID)))
	}
	if len(t.PFCPASRspFlags) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.PFCPASRSP_FLAGS_TYPE, tlv.Bytes(t.PFCPASRspFlags)))
	}
	if t.ClockDriftControlInformation != nil {
		fields = append(fields, t.ClockDriftControlInformation.Field())
	}
	if t.UEIPaddressPoolInformation != nil {
		fields = append(fields, tlv.TLVFrom(ie.UE_IP_ADDRESS_TYPE, tlv.Bytes(t.UEIPaddressPoolInformation.Bytes())))
	}
	if len(t.GTPUPathQoSControlInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.GTP_U_PATH_QOS_CONTROL_INFORMATION_TYPE, tlv.Bytes(t.GTPUPathQoSControlInformation)))
	}
	if len(t.UPFInstanceID) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.NF_INSTANCE_ID_TYPE, tlv.Bytes(t.UPFInstanceID)))
	}
	return tlv.TLV(ie.MSG_PFCP_ASSOCIATION_SETUP_RESPONSE_TYPE, fields...)
}

func (t *PFCPAssociationSetupResponse) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPAssociationSetupResponse{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case ie.NODE_ID_TYPE:
			err = de.UnmarshalValue(&t.NodeID)
		case ie.CAUSE_TYPE:
			err = pfcp.UnmarshalUint8(de.Value, &t.Cause)
		case ie.RECOVERY_TIME_STAMP_TYPE:
			err = pfcp.UnmarshalUint32(de.Value, &t.RecoveryTimeStamp)
		case ie.UP_FUNCTION_FEATURES_TYPE:
			t.UPFunctionFeatures = de.Value
		case ie.CP_FUNCTION_FEATURES_TYPE:
			var v uint8
			if err = pfcp.UnmarshalUint8(de.Value, &v); err == nil {
				t.CPFunctionFeatures = &v
			}
		case ie.USER_PLANE_IP_RESOURCE_INFORMATION_TYPE:
			t.UserPlaneIPResourceInformation = de.Value
		case ie.ALTERNATIVE_SMF_IP_ADDRESS_TYPE:
			t.AlternativeSMFIPAddress = de.Value
		case ie.SMF_SET_ID_TYPE:
			t.SMFSetID = de.Value
		case ie.PFCPASRSP_FLAGS_TYPE:
			t.PFCPASRspFlags = de.Value
		case ie.CLOCK_DRIFT_CONTROL_INFORMATION_TYPE:
			var v ie.ClockDriftControlInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.ClockDriftControlInformation = &v
			}
		case ie.UE_IP_ADDRESS_TYPE:
			var v ie.UEIPAddress
			if err = de.UnmarshalValue(&v); err == nil {
				t.UEIPaddressPoolInformation = &v
			}
		case ie.GTP_U_PATH_QOS_CONTROL_INFORMATION_TYPE:
			t.GTPUPathQoSControlInformation = de.Value
		case ie.NF_INSTANCE_ID_TYPE:
			t.UPFInstanceID = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
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
