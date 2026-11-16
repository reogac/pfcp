package message

import (
	"github.com/reogac/pfcp/v2"
	"github.com/reogac/pfcp/v2/ie"

	"github.com/reogac/pfcp/v2/tlv"
)

/* 'PFCP Association Update Request' data structure */
type PFCPAssociationUpdateRequest struct {
	NodeID                         ie.NodeID                         //Node ID
	UPFunctionFeatures             []byte                            //UP Function Features
	CPFunctionFeatures             *uint8                            //CP Function Features
	UserPlaneIPResourceInformation [][]byte                          //User Plane IP Resource Information
	PFCPAssociationReleaseRequest  []byte                            //PFCP Association Release Request
	GracefulReleasePeriod          []byte                            //Graceful Release Period
	PFCPAUReqFlags                 *uint8                            //PFCPAUReq-Flags
	AlternativeSMFIPAddress        [][]byte                          //Alternative SMF IP Address
	SMFSetID                       []byte                            //SMF Set ID
	ClockDriftControlInformation   []ie.ClockDriftControlInformation //Clock Drift Control Information
	UEIPaddressPoolInformation     []ie.UEIPAddress                  //UE IP Address
	GTPUPathQoSControlInformation  [][]byte                          //GTP-U Path QoS Control Information
	UEIPAddressUsageInformation    []ie.UEIPAddressUsageInformation  //UE IP Address Usage Information
	Unknown                        []uint32                          `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid                        []uint32                          `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
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
	for _, i := range t.UserPlaneIPResourceInformation {
		fields = append(fields, tlv.TLVFrom(ie.USER_PLANE_IP_RESOURCE_INFORMATION_TYPE, tlv.Bytes(i)))
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
	for _, i := range t.AlternativeSMFIPAddress {
		fields = append(fields, tlv.TLVFrom(ie.ALTERNATIVE_SMF_IP_ADDRESS_TYPE, tlv.Bytes(i)))
	}
	if len(t.SMFSetID) > 0 {
		fields = append(fields, tlv.TLVFrom(ie.SMF_SET_ID_TYPE, tlv.Bytes(t.SMFSetID)))
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
	for _, i := range t.UEIPAddressUsageInformation {
		fields = append(fields, i.Field())
	}
	return tlv.TLV(ie.MSG_PFCP_ASSOCIATION_UPDATE_REQUEST_TYPE, fields...)
}

// Skipped is what decoding left out: IEs unknown or not defined here
// (TS 29.244 7.6.9), and optional ones discarded for a faulty value (7.6.8).
func (t *PFCPAssociationUpdateRequest) Skipped() (unknown, invalid []uint32) {
	return t.Unknown, t.Invalid
}

func (t *PFCPAssociationUpdateRequest) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPAssociationUpdateRequest{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [7]bool
	for _, de := range elements {
		opt := false
		switch de.Type {
		case ie.NODE_ID_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			if err = de.UnmarshalValue(&t.NodeID); err == nil {
			}
		case ie.UP_FUNCTION_FEATURES_TYPE:
			opt = true
			if seen[1] {
				continue
			}
			seen[1] = true
			t.UPFunctionFeatures = de.Value
		case ie.CP_FUNCTION_FEATURES_TYPE:
			opt = true
			if seen[2] {
				continue
			}
			seen[2] = true
			var v uint8
			if err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &v); err == nil {
				t.CPFunctionFeatures = &v
			}
		case ie.USER_PLANE_IP_RESOURCE_INFORMATION_TYPE:
			opt = true
			t.UserPlaneIPResourceInformation = append(t.UserPlaneIPResourceInformation, de.Value)
		case ie.PFCP_ASSOCIATION_RELEASE_REQUEST_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			t.PFCPAssociationReleaseRequest = de.Value
		case ie.GRACEFUL_RELEASE_PERIOD_TYPE:
			if seen[4] {
				continue
			}
			seen[4] = true
			t.GracefulReleasePeriod = de.Value
		case ie.PFCPAUREQ_FLAGS_TYPE:
			opt = true
			if seen[5] {
				continue
			}
			seen[5] = true
			var v uint8
			if err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &v); err == nil {
				t.PFCPAUReqFlags = &v
			}
		case ie.ALTERNATIVE_SMF_IP_ADDRESS_TYPE:
			opt = true
			t.AlternativeSMFIPAddress = append(t.AlternativeSMFIPAddress, de.Value)
		case ie.SMF_SET_ID_TYPE:
			opt = true
			if seen[6] {
				continue
			}
			seen[6] = true
			t.SMFSetID = de.Value
		case ie.CLOCK_DRIFT_CONTROL_INFORMATION_TYPE:
			var v ie.ClockDriftControlInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.ClockDriftControlInformation = append(t.ClockDriftControlInformation, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
		case ie.UE_IP_ADDRESS_TYPE:
			opt = true
			var v ie.UEIPAddress
			if err = de.UnmarshalValue(&v); err == nil {
				t.UEIPaddressPoolInformation = append(t.UEIPaddressPoolInformation, v)
			}
		case ie.GTP_U_PATH_QOS_CONTROL_INFORMATION_TYPE:
			t.GTPUPathQoSControlInformation = append(t.GTPUPathQoSControlInformation, de.Value)
		case ie.UE_IP_ADDRESS_USAGE_INFORMATION_TYPE:
			opt = true
			var v ie.UEIPAddressUsageInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.UEIPAddressUsageInformation = append(t.UEIPAddressUsageInformation, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
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
		return &pfcp.MissingIE{Type: ie.NODE_ID_TYPE}
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
	NodeID                         ie.NodeID                        //Node ID
	Cause                          uint8                            //Cause
	UPFunctionFeatures             []byte                           //UP Function Features
	CPFunctionFeatures             *uint8                           //CP Function Features
	UserPlaneIPResourceInformation [][]byte                         //User Plane IP Resource Information
	UEIPAddressUsageInformation    []ie.UEIPAddressUsageInformation //UE IP Address Usage Information
	Unknown                        []uint32                         `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
	Invalid                        []uint32                         `json:"-"` //optional IEs discarded, their value faulty (TS 29.244 7.6.8)
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
	for _, i := range t.UserPlaneIPResourceInformation {
		fields = append(fields, tlv.TLVFrom(ie.USER_PLANE_IP_RESOURCE_INFORMATION_TYPE, tlv.Bytes(i)))
	}
	for _, i := range t.UEIPAddressUsageInformation {
		fields = append(fields, i.Field())
	}
	return tlv.TLV(ie.MSG_PFCP_ASSOCIATION_UPDATE_RESPONSE_TYPE, fields...)
}

// Skipped is what decoding left out: IEs unknown or not defined here
// (TS 29.244 7.6.9), and optional ones discarded for a faulty value (7.6.8).
func (t *PFCPAssociationUpdateResponse) Skipped() (unknown, invalid []uint32) {
	return t.Unknown, t.Invalid
}

func (t *PFCPAssociationUpdateResponse) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPAssociationUpdateResponse{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return pfcp.Faulty(0, err)
	}
	var seen [4]bool
	for _, de := range elements {
		opt := false
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
		case ie.UP_FUNCTION_FEATURES_TYPE:
			opt = true
			if seen[2] {
				continue
			}
			seen[2] = true
			t.UPFunctionFeatures = de.Value
		case ie.CP_FUNCTION_FEATURES_TYPE:
			opt = true
			if seen[3] {
				continue
			}
			seen[3] = true
			var v uint8
			if err = pfcp.UnmarshalExtendable(de.Value, 1, 1, &v); err == nil {
				t.CPFunctionFeatures = &v
			}
		case ie.USER_PLANE_IP_RESOURCE_INFORMATION_TYPE:
			opt = true
			t.UserPlaneIPResourceInformation = append(t.UserPlaneIPResourceInformation, de.Value)
		case ie.UE_IP_ADDRESS_USAGE_INFORMATION_TYPE:
			opt = true
			var v ie.UEIPAddressUsageInformation
			if err = de.UnmarshalValue(&v); err == nil {
				t.UEIPAddressUsageInformation = append(t.UEIPAddressUsageInformation, v)
				t.Unknown = append(t.Unknown, v.Unknown...)
				t.Invalid = append(t.Invalid, v.Invalid...)
			}
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
	if !seen[1] {
		return &pfcp.MissingIE{Type: ie.CAUSE_TYPE}
	}
	if t.Cause < ie.CauseRequestRejected {
		if !seen[0] {
			return &pfcp.MissingIE{Type: ie.NODE_ID_TYPE}
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
