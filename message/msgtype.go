package message

import (
	"encoding"
	"fmt"

	"github.com/reogac/pfcp/ie"
	"github.com/reogac/pfcp/tlv"
)

// Type is a PFCP message type (TS 29.244 §7.3).
type Type uint8

// Message is any PFCP message body. It encodes as an element of its own type
// whose value is the body, which is why a body alone is tlv.AppendValue of it.
type Message interface {
	tlv.Fielder
	encoding.BinaryUnmarshaler
	MsgType() Type
}

// typeInfo is what a message type fixes about every message of it.
type typeInfo struct {
	name string
	//request is a message that is answered; everything else is the answer
	request bool
	//session is a message about one PFCP session, which carries a SEID in
	//its header (S flag set); node messages carry none
	session bool
	//pair is a request's response type, and a response's request type
	pair Type
	new  func() Message
}

// types is indexed by message type: a lookup is one index, and a type with no
// entry is one this library does not know.
var types [256]*typeInfo

func define(t, pair uint32, request, session bool, name string, new func() Message) {
	types[t] = &typeInfo{name: name, request: request, session: session, pair: Type(pair), new: new}
}

func init() {
	node, session := false, true
	req, rsp := true, false
	define(ie.MSG_PFCP_HEARTBEAT_REQUEST_TYPE, ie.MSG_PFCP_HEARTBEAT_RESPONSE_TYPE, req, node, "Heartbeat Request",
		func() Message { return &PFCPHeartbeatRequest{} })
	define(ie.MSG_PFCP_HEARTBEAT_RESPONSE_TYPE, ie.MSG_PFCP_HEARTBEAT_REQUEST_TYPE, rsp, node, "Heartbeat Response",
		func() Message { return &PFCPHeartbeatResponse{} })
	define(ie.MSG_PFCP_PFD_MANAGEMENT_REQUEST_TYPE, ie.MSG_PFCP_PFD_MANAGEMENT_RESPONSE_TYPE, req, node, "PFD Management Request",
		func() Message { return &PFCPPFDManagementRequest{} })
	define(ie.MSG_PFCP_PFD_MANAGEMENT_RESPONSE_TYPE, ie.MSG_PFCP_PFD_MANAGEMENT_REQUEST_TYPE, rsp, node, "PFD Management Response",
		func() Message { return &PFCPPFDManagementResponse{} })
	define(ie.MSG_PFCP_ASSOCIATION_SETUP_REQUEST_TYPE, ie.MSG_PFCP_ASSOCIATION_SETUP_RESPONSE_TYPE, req, node, "Association Setup Request",
		func() Message { return &PFCPAssociationSetupRequest{} })
	define(ie.MSG_PFCP_ASSOCIATION_SETUP_RESPONSE_TYPE, ie.MSG_PFCP_ASSOCIATION_SETUP_REQUEST_TYPE, rsp, node, "Association Setup Response",
		func() Message { return &PFCPAssociationSetupResponse{} })
	define(ie.MSG_PFCP_ASSOCIATION_UPDATE_REQUEST_TYPE, ie.MSG_PFCP_ASSOCIATION_UPDATE_RESPONSE_TYPE, req, node, "Association Update Request",
		func() Message { return &PFCPAssociationUpdateRequest{} })
	define(ie.MSG_PFCP_ASSOCIATION_UPDATE_RESPONSE_TYPE, ie.MSG_PFCP_ASSOCIATION_UPDATE_REQUEST_TYPE, rsp, node, "Association Update Response",
		func() Message { return &PFCPAssociationUpdateResponse{} })
	define(ie.MSG_PFCP_ASSOCIATION_RELEASE_REQUEST_TYPE, ie.MSG_PFCP_ASSOCIATION_RELEASE_RESPONSE_TYPE, req, node, "Association Release Request",
		func() Message { return &PFCPAssociationReleaseRequest{} })
	define(ie.MSG_PFCP_ASSOCIATION_RELEASE_RESPONSE_TYPE, ie.MSG_PFCP_ASSOCIATION_RELEASE_REQUEST_TYPE, rsp, node, "Association Release Response",
		func() Message { return &PFCPAssociationReleaseResponse{} })
	//answers a request of any type, so it pairs with none
	define(ie.MSG_PFCP_VERSION_NOT_SUPPORTED_RESPONSE_TYPE, 0, rsp, node, "Version Not Supported Response",
		func() Message { return &PFCPVersionNotSupportedResponse{} })
	define(ie.MSG_PFCP_NODE_REPORT_REQUEST_TYPE, ie.MSG_PFCP_NODE_REPORT_RESPONSE_TYPE, req, node, "Node Report Request",
		func() Message { return &PFCPNodeReportRequest{} })
	define(ie.MSG_PFCP_NODE_REPORT_RESPONSE_TYPE, ie.MSG_PFCP_NODE_REPORT_REQUEST_TYPE, rsp, node, "Node Report Response",
		func() Message { return &PFCPNodeReportResponse{} })
	define(ie.MSG_PFCP_SESSION_SET_DELETION_REQUEST_TYPE, ie.MSG_PFCP_SESSION_SET_DELETION_RESPONSE_TYPE, req, node, "Session Set Deletion Request",
		func() Message { return &PFCPSessionSetDeletionRequest{} })
	define(ie.MSG_PFCP_SESSION_SET_DELETION_RESPONSE_TYPE, ie.MSG_PFCP_SESSION_SET_DELETION_REQUEST_TYPE, rsp, node, "Session Set Deletion Response",
		func() Message { return &PFCPSessionSetDeletionResponse{} })
	define(ie.MSG_PFCP_SESSION_SET_MODIFICATION_REQUEST_TYPE, ie.MSG_PFCP_SESSION_SET_MODIFICATION_RESPONSE_TYPE, req, node, "Session Set Modification Request",
		func() Message { return &PFCPSessionSetModificationRequest{} })
	define(ie.MSG_PFCP_SESSION_SET_MODIFICATION_RESPONSE_TYPE, ie.MSG_PFCP_SESSION_SET_MODIFICATION_REQUEST_TYPE, rsp, node, "Session Set Modification Response",
		func() Message { return &PFCPSessionSetModificationResponse{} })
	define(ie.MSG_PFCP_SESSION_ESTABLISHMENT_REQUEST_TYPE, ie.MSG_PFCP_SESSION_ESTABLISHMENT_RESPONSE_TYPE, req, session, "Session Establishment Request",
		func() Message { return &PFCPSessionEstablishmentRequest{} })
	define(ie.MSG_PFCP_SESSION_ESTABLISHMENT_RESPONSE_TYPE, ie.MSG_PFCP_SESSION_ESTABLISHMENT_REQUEST_TYPE, rsp, session, "Session Establishment Response",
		func() Message { return &PFCPSessionEstablishmentResponse{} })
	define(ie.MSG_PFCP_SESSION_MODIFICATION_REQUEST_TYPE, ie.MSG_PFCP_SESSION_MODIFICATION_RESPONSE_TYPE, req, session, "Session Modification Request",
		func() Message { return &PFCPSessionModificationRequest{} })
	define(ie.MSG_PFCP_SESSION_MODIFICATION_RESPONSE_TYPE, ie.MSG_PFCP_SESSION_MODIFICATION_REQUEST_TYPE, rsp, session, "Session Modification Response",
		func() Message { return &PFCPSessionModificationResponse{} })
	define(ie.MSG_PFCP_SESSION_DELETION_REQUEST_TYPE, ie.MSG_PFCP_SESSION_DELETION_RESPONSE_TYPE, req, session, "Session Deletion Request",
		func() Message { return &PFCPSessionDeletionRequest{} })
	define(ie.MSG_PFCP_SESSION_DELETION_RESPONSE_TYPE, ie.MSG_PFCP_SESSION_DELETION_REQUEST_TYPE, rsp, session, "Session Deletion Response",
		func() Message { return &PFCPSessionDeletionResponse{} })
	define(ie.MSG_PFCP_SESSION_REPORT_REQUEST_TYPE, ie.MSG_PFCP_SESSION_REPORT_RESPONSE_TYPE, req, session, "Session Report Request",
		func() Message { return &PFCPSessionReportRequest{} })
	define(ie.MSG_PFCP_SESSION_REPORT_RESPONSE_TYPE, ie.MSG_PFCP_SESSION_REPORT_REQUEST_TYPE, rsp, session, "Session Report Response",
		func() Message { return &PFCPSessionReportResponse{} })
}

// Known reports whether this library has a body for t.
func (t Type) Known() bool {
	return types[t] != nil
}

// IsRequest reports whether t is answered. It is false for a type this
// library does not know: the parity of a type number says nothing, since Node
// Report Request is 12 and Version Not Supported Response 11.
func (t Type) IsRequest() bool {
	return types[t] != nil && types[t].request
}

// IsSession reports whether t is about one PFCP session, whose header carries a
// SEID.
func (t Type) IsSession() bool {
	return types[t] != nil && types[t].session
}

// Response is the type that answers a request of type t, and 0 for anything
// that is not a known request.
func (t Type) Response() Type {
	if !t.IsRequest() {
		return 0
	}
	return types[t].pair
}

func (t Type) String() string {
	if info := types[t]; info != nil {
		return info.name
	}
	return fmt.Sprintf("message type %d", uint8(t))
}

// New is an empty body of type t, or nil for a type this library does not know.
func New(t Type) Message {
	if info := types[t]; info != nil {
		return info.new()
	}
	return nil
}

func (*PFCPHeartbeatRequest) MsgType() Type  { return Type(ie.MSG_PFCP_HEARTBEAT_REQUEST_TYPE) }
func (*PFCPHeartbeatResponse) MsgType() Type { return Type(ie.MSG_PFCP_HEARTBEAT_RESPONSE_TYPE) }
func (*PFCPPFDManagementRequest) MsgType() Type {
	return Type(ie.MSG_PFCP_PFD_MANAGEMENT_REQUEST_TYPE)
}
func (*PFCPPFDManagementResponse) MsgType() Type {
	return Type(ie.MSG_PFCP_PFD_MANAGEMENT_RESPONSE_TYPE)
}
func (*PFCPAssociationSetupRequest) MsgType() Type {
	return Type(ie.MSG_PFCP_ASSOCIATION_SETUP_REQUEST_TYPE)
}
func (*PFCPAssociationSetupResponse) MsgType() Type {
	return Type(ie.MSG_PFCP_ASSOCIATION_SETUP_RESPONSE_TYPE)
}
func (*PFCPAssociationUpdateRequest) MsgType() Type {
	return Type(ie.MSG_PFCP_ASSOCIATION_UPDATE_REQUEST_TYPE)
}
func (*PFCPAssociationUpdateResponse) MsgType() Type {
	return Type(ie.MSG_PFCP_ASSOCIATION_UPDATE_RESPONSE_TYPE)
}
func (*PFCPAssociationReleaseRequest) MsgType() Type {
	return Type(ie.MSG_PFCP_ASSOCIATION_RELEASE_REQUEST_TYPE)
}
func (*PFCPAssociationReleaseResponse) MsgType() Type {
	return Type(ie.MSG_PFCP_ASSOCIATION_RELEASE_RESPONSE_TYPE)
}
func (*PFCPVersionNotSupportedResponse) MsgType() Type {
	return Type(ie.MSG_PFCP_VERSION_NOT_SUPPORTED_RESPONSE_TYPE)
}
func (*PFCPNodeReportRequest) MsgType() Type { return Type(ie.MSG_PFCP_NODE_REPORT_REQUEST_TYPE) }
func (*PFCPNodeReportResponse) MsgType() Type {
	return Type(ie.MSG_PFCP_NODE_REPORT_RESPONSE_TYPE)
}
func (*PFCPSessionSetDeletionRequest) MsgType() Type {
	return Type(ie.MSG_PFCP_SESSION_SET_DELETION_REQUEST_TYPE)
}
func (*PFCPSessionSetDeletionResponse) MsgType() Type {
	return Type(ie.MSG_PFCP_SESSION_SET_DELETION_RESPONSE_TYPE)
}
func (*PFCPSessionSetModificationRequest) MsgType() Type {
	return Type(ie.MSG_PFCP_SESSION_SET_MODIFICATION_REQUEST_TYPE)
}
func (*PFCPSessionSetModificationResponse) MsgType() Type {
	return Type(ie.MSG_PFCP_SESSION_SET_MODIFICATION_RESPONSE_TYPE)
}
func (*PFCPSessionEstablishmentRequest) MsgType() Type {
	return Type(ie.MSG_PFCP_SESSION_ESTABLISHMENT_REQUEST_TYPE)
}
func (*PFCPSessionEstablishmentResponse) MsgType() Type {
	return Type(ie.MSG_PFCP_SESSION_ESTABLISHMENT_RESPONSE_TYPE)
}
func (*PFCPSessionModificationRequest) MsgType() Type {
	return Type(ie.MSG_PFCP_SESSION_MODIFICATION_REQUEST_TYPE)
}
func (*PFCPSessionModificationResponse) MsgType() Type {
	return Type(ie.MSG_PFCP_SESSION_MODIFICATION_RESPONSE_TYPE)
}
func (*PFCPSessionDeletionRequest) MsgType() Type {
	return Type(ie.MSG_PFCP_SESSION_DELETION_REQUEST_TYPE)
}
func (*PFCPSessionDeletionResponse) MsgType() Type {
	return Type(ie.MSG_PFCP_SESSION_DELETION_RESPONSE_TYPE)
}
func (*PFCPSessionReportRequest) MsgType() Type {
	return Type(ie.MSG_PFCP_SESSION_REPORT_REQUEST_TYPE)
}
func (*PFCPSessionReportResponse) MsgType() Type {
	return Type(ie.MSG_PFCP_SESSION_REPORT_RESPONSE_TYPE)
}
