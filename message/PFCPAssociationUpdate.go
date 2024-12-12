package message

import (
	"etrib5gc/pfcp"
	"etrib5gc/pfcp/ie"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'PFCP Association Update Request' data structure */
type PFCPAssociationUpdateRequest struct {
	NodeID                         ie.NodeID                        //Node ID
	UPFunctionFeatures             []byte                           //UP Function Features
	CPFunctionFeatures             *uint8                           //CP Function Features
	UserPlaneIPResourceInformation []byte                           //User Plane IP Resource Information
	PFCPAssociationReleaseRequest  []byte                           //PFCP Association Release Request
	GracefulReleasePeriod          []byte                           //Graceful Release Period
	PFCPAUReqFlags                 *uint8                           //PFCPAUReq-Flags
	AlternativeSMFIPAddress        []byte                           //Alternative SMF IP Address
	SMFSetID                       []byte                           //SMF Set ID
	ClockDriftControlInformation   *ie.ClockDriftControlInformation //Clock Drift Control Information
	UEIPaddressPoolInformation     *ie.UEIPAddress                  //UE IP Address
	GTPUPathQoSControlInformation  []byte                           //GTP-U Path QoS Control Information
	UEIPAddressUsageInformation    *ie.UEIPAddressUsageInformation  //UE IP Address Usage Information
}

func (t *PFCPAssociationUpdateRequest) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(ie.NODE_ID_TYPE, tlv.Bytes(t.NodeID.Bytes())))
	if len(t.UPFunctionFeatures) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.UP_FUNCTION_FEATURES_TYPE, tlv.Bytes(t.UPFunctionFeatures)))
	}
	if t.CPFunctionFeatures != nil {
		fields = append(fields, tlv.TLVFrom(ie.CP_FUNCTION_FEATURES_TYPE, tlv.Bytes(pfcp.MarshalUint8(*t.CPFunctionFeatures))))
	}
	for len(t.UserPlaneIPResourceInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.USER_PLANE_IP_RESOURCE_INFORMATION_TYPE, tlv.Bytes(t.UserPlaneIPResourceInformation)))
	}
	if len(t.PFCPAssociationReleaseRequest) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.PFCP_ASSOCIATION_RELEASE_REQUEST_TYPE, tlv.Bytes(t.PFCPAssociationReleaseRequest)))
	}
	if len(t.GracefulReleasePeriod) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.GRACEFUL_RELEASE_PERIOD_TYPE, tlv.Bytes(t.GracefulReleasePeriod)))
	}
	if t.PFCPAUReqFlags != nil {
		fields = append(fields, tlv.TLVFrom(ie.PFCPAUREQ_FLAGS_TYPE, tlv.Bytes(pfcp.MarshalUint8(*t.PFCPAUReqFlags))))
	}
	if len(t.AlternativeSMFIPAddress) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.ALTERNATIVE_SMF_IP_ADDRESS_TYPE, tlv.Bytes(t.AlternativeSMFIPAddress)))
	}
	if len(t.SMFSetID) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.SMF_SET_ID_TYPE, tlv.Bytes(t.SMFSetID)))
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
	if t.UEIPAddressUsageInformation != nil {
		fields = append(fields, t.UEIPAddressUsageInformation.Field())
	}
	return tlv.TLV(ie.MSG_PFCP_ASSOCIATION_UPDATE_REQUEST_TYPE, fields...)
}

func (t *PFCPAssociationUpdateRequest) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPAssociationUpdateRequest{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case ie.NODE_ID_TYPE:
			err = de.UnmarshalValue(&t.NodeID)
		case ie.UP_FUNCTION_FEATURES_TYPE:
			t.UPFunctionFeatures = de.Value
		case ie.CP_FUNCTION_FEATURES_TYPE:
			var v uint8
			if err = pfcp.UnmarshalUint8(de.Value, &v); err == nil {
				t.CPFunctionFeatures = &v
			}
		case ie.USER_PLANE_IP_RESOURCE_INFORMATION_TYPE:
			t.UserPlaneIPResourceInformation = de.Value
		case ie.PFCP_ASSOCIATION_RELEASE_REQUEST_TYPE:
			t.PFCPAssociationReleaseRequest = de.Value
		case ie.GRACEFUL_RELEASE_PERIOD_TYPE:
			t.GracefulReleasePeriod = de.Value
		case ie.PFCPAUREQ_FLAGS_TYPE:
			var v uint8
			if err = pfcp.UnmarshalUint8(de.Value, &v); err == nil {
				t.PFCPAUReqFlags = &v
			}
		case ie.ALTERNATIVE_SMF_IP_ADDRESS_TYPE:
			t.AlternativeSMFIPAddress = de.Value
		case ie.SMF_SET_ID_TYPE:
			t.SMFSetID = de.Value
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
		case ie.UE_IP_ADDRESS_USAGE_INFORMATION_TYPE:
			var v ie.UEIPAddressUsageInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.UEIPAddressUsageInformation = &v
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

func (t *PFCPAssociationUpdateRequest) UnmarshalTLV(typ uint32, buf []byte) (err error) {
	if typ != ie.MSG_PFCP_ASSOCIATION_UPDATE_REQUEST_TYPE {
		err = tlv.ErrCritical
		return
	}

	err = t.UnmarshalBinary(buf)
	return
}

func (t *PFCPAssociationUpdateRequest) MarshalJSON() (buf []byte, err error) {
	buf, err = jsonMarshal(t)
	return
}

func (t *PFCPAssociationUpdateRequest) UnmarshalJSON(buf []byte) (err error) {
	*t = PFCPAssociationUpdateRequest{}
	err = jsonUnmarshal(buf, t)
	return
}

/* 'PFCP Association Update Response' data structure */
type PFCPAssociationUpdateResponse struct {
	NodeID                         ie.NodeID                       //Node ID
	Cause                          uint8                           //Cause
	UPFunctionFeatures             []byte                          //UP Function Features
	CPFunctionFeatures             *uint8                          //CP Function Features
	UserPlaneIPResourceInformation []byte                          //User Plane IP Resource Information
	UEIPAddressUsageInformation    *ie.UEIPAddressUsageInformation //UE IP Address Usage Information
}

func (t *PFCPAssociationUpdateResponse) Field() tlv.Field {
	fields := []tlv.Field{}
	fields = append(fields, tlv.TLVFrom(ie.NODE_ID_TYPE, tlv.Bytes(t.NodeID.Bytes())))
	fields = append(fields, tlv.TLVFrom(ie.CAUSE_TYPE, tlv.Bytes(pfcp.MarshalUint8(t.Cause))))
	if len(t.UPFunctionFeatures) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.UP_FUNCTION_FEATURES_TYPE, tlv.Bytes(t.UPFunctionFeatures)))
	}
	if t.CPFunctionFeatures != nil {
		fields = append(fields, tlv.TLVFrom(ie.CP_FUNCTION_FEATURES_TYPE, tlv.Bytes(pfcp.MarshalUint8(*t.CPFunctionFeatures))))
	}
	for len(t.UserPlaneIPResourceInformation) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.USER_PLANE_IP_RESOURCE_INFORMATION_TYPE, tlv.Bytes(t.UserPlaneIPResourceInformation)))
	}
	if t.UEIPAddressUsageInformation != nil {
		fields = append(fields, t.UEIPAddressUsageInformation.Field())
	}
	return tlv.TLV(ie.MSG_PFCP_ASSOCIATION_UPDATE_RESPONSE_TYPE, fields...)
}

func (t *PFCPAssociationUpdateResponse) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPAssociationUpdateResponse{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case ie.NODE_ID_TYPE:
			err = de.UnmarshalValue(&t.NodeID)
		case ie.CAUSE_TYPE:
			err = pfcp.UnmarshalUint8(de.Value, &t.Cause)
		case ie.UP_FUNCTION_FEATURES_TYPE:
			t.UPFunctionFeatures = de.Value
		case ie.CP_FUNCTION_FEATURES_TYPE:
			var v uint8
			if err = pfcp.UnmarshalUint8(de.Value, &v); err == nil {
				t.CPFunctionFeatures = &v
			}
		case ie.USER_PLANE_IP_RESOURCE_INFORMATION_TYPE:
			t.UserPlaneIPResourceInformation = de.Value
		case ie.UE_IP_ADDRESS_USAGE_INFORMATION_TYPE:
			var v ie.UEIPAddressUsageInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.UEIPAddressUsageInformation = &v
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

func (t *PFCPAssociationUpdateResponse) UnmarshalTLV(typ uint32, buf []byte) (err error) {
	if typ != ie.MSG_PFCP_ASSOCIATION_UPDATE_RESPONSE_TYPE {
		err = tlv.ErrCritical
		return
	}

	err = t.UnmarshalBinary(buf)
	return
}

func (t *PFCPAssociationUpdateResponse) MarshalJSON() (buf []byte, err error) {
	buf, err = jsonMarshal(t)
	return
}

func (t *PFCPAssociationUpdateResponse) UnmarshalJSON(buf []byte) (err error) {
	*t = PFCPAssociationUpdateResponse{}
	err = jsonUnmarshal(buf, t)
	return
}
