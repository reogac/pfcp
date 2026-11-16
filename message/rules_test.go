package message_test

import (
	"errors"
	"net"
	"slices"
	"testing"

	"github.com/reogac/pfcp/v2"
	"github.com/reogac/pfcp/v2/ie"
	"github.com/reogac/pfcp/v2/message"
	"github.com/reogac/pfcp/v2/tlv"
)

// The decoding rules of TS 29.244 7.6 and 8.1.2 the generated decoders follow,
// on messages built by hand.

func body(t *testing.T, fields ...tlv.Field) []byte {
	t.Helper()
	b, err := tlv.Encode(fields...)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func u32(typ uint32, v uint32) tlv.Field { return tlv.TLVFrom(typ, tlv.Bytes(pfcp.MarshalUint32(v))) }

func missing(err error) uint32 {
	var m *pfcp.MissingIE
	if errors.As(err, &m) {
		return m.Type
	}
	return 0
}

// 7.6.10: a single IE repeated where repetition is not specified is taken at
// its first occurrence.
func TestTheFirstOfARepeatedSingleIEIsKept(t *testing.T) {
	var m message.PFCPHeartbeatRequest
	if err := m.UnmarshalBinary(body(t, u32(ie.RECOVERY_TIME_STAMP_TYPE, 1), u32(ie.RECOVERY_TIME_STAMP_TYPE, 2))); err != nil {
		t.Fatal(err)
	}
	if m.RecoveryTimeStamp != 1 {
		t.Errorf("kept %d, want the first, 1", m.RecoveryTimeStamp)
	}
}

// 7.6.9: an IE of an unknown type, or of a known type the message does not
// define, is discarded and recorded, and the message processed; a grouped IE's
// reach the message's list.
func TestAnUnknownOrUnexpectedIEIsSkippedAndRecorded(t *testing.T) {
	var hb message.PFCPHeartbeatRequest
	err := hb.UnmarshalBinary(body(t, u32(ie.RECOVERY_TIME_STAMP_TYPE, 7),
		tlv.TLVFrom(31000, tlv.Bytes([]byte{1})), tlv.TLVFrom(ie.PDR_ID_TYPE, tlv.Bytes(pfcp.MarshalUint16(3)))))
	if err != nil || hb.RecoveryTimeStamp != 7 || !slices.Equal(hb.Unknown, []uint32{31000, ie.PDR_ID_TYPE}) {
		t.Errorf("decoded %d, skipped %v, %v", hb.RecoveryTimeStamp, hb.Unknown, err)
	}
	var rsp message.PFCPSessionReportResponse
	err = rsp.UnmarshalBinary(body(t, tlv.TLVFrom(ie.CAUSE_TYPE, tlv.Bytes([]byte{1})),
		tlv.TLV(ie.UPDATE_BAR_PFCP_SESSION_REPORT_RESPONSE_TYPE,
			tlv.TLVFrom(ie.BAR_ID_TYPE, tlv.Bytes([]byte{1})), tlv.TLVFrom(31001, tlv.Bytes([]byte{9})))))
	if err != nil || !slices.Equal(rsp.Unknown, []uint32{31001}) {
		t.Errorf("a grouped IE's unknown IE: message skipped %v, %v", rsp.Unknown, err)
	}
}

// 7.6.6: a missing mandatory IE is reported by its type, but in a Heartbeat
// Request, and in a response with a rejection Cause.
func TestAMissingMandatoryIEIsReported(t *testing.T) {
	var hb message.PFCPHeartbeatRequest
	if err := hb.UnmarshalBinary(nil); err != nil {
		t.Errorf("an empty Heartbeat Request: %v", err)
	}
	var hbr message.PFCPHeartbeatResponse
	if got := missing(hbr.UnmarshalBinary(nil)); got != ie.RECOVERY_TIME_STAMP_TYPE {
		t.Errorf("an empty Heartbeat Response named %d, want the Recovery Time Stamp", got)
	}
	cause := func(c uint8) []byte { return body(t, tlv.TLVFrom(ie.CAUSE_TYPE, tlv.Bytes([]byte{c}))) }
	var rsp message.PFCPSessionEstablishmentResponse
	if err := rsp.UnmarshalBinary(cause(ie.CauseRequestRejected)); err != nil {
		t.Errorf("a rejection without its Node ID: %v", err)
	}
	if got := missing(rsp.UnmarshalBinary(cause(1))); got != ie.NODE_ID_TYPE {
		t.Errorf("an acceptance without its Node ID named %d", got)
	}
	var req message.PFCPSessionEstablishmentRequest
	if got := missing(req.UnmarshalBinary(body(t))); got == ie.MBS_SESSION_N4MB_CONTROL_INFORMATION_TYPE || got == 0 {
		t.Errorf("an empty Session Establishment Request named %d, want an IE mandatory on N4", got)
	}
}

// 8.1.2: an Extendable IE is read from its first-defined length up, the
// octets beyond the known ones ignored.
func TestAnExtendableIEMayCarryMoreOctets(t *testing.T) {
	var hb message.PFCPHeartbeatRequest
	if err := hb.UnmarshalBinary(body(t, tlv.TLVFrom(ie.RECOVERY_TIME_STAMP_TYPE, tlv.Bytes([]byte{0, 0, 0, 9, 0xff})))); err != nil || hb.RecoveryTimeStamp != 9 {
		t.Errorf("a Recovery Time Stamp of 5 octets: %d, %v", hb.RecoveryTimeStamp, err)
	}
}

// 7.6.8: an IE whose value does not decode is named, the innermost one when
// it sits inside a grouped IE, so the answer can carry it in the Offending IE.
func TestAFaultyIEIsNamed(t *testing.T) {
	invalid := func(err error) uint32 {
		var e *pfcp.InvalidIE
		if errors.As(err, &e) {
			return e.Type
		}
		return 0
	}
	var hbr message.PFCPHeartbeatResponse
	short := body(t, tlv.TLVFrom(ie.RECOVERY_TIME_STAMP_TYPE, tlv.Bytes([]byte{0, 1})))
	if got := invalid(hbr.UnmarshalBinary(short)); got != ie.RECOVERY_TIME_STAMP_TYPE {
		t.Errorf("a Recovery Time Stamp of 2 octets named %d", got)
	}
	var rsp message.PFCPSessionReportResponse
	err := rsp.UnmarshalBinary(body(t, tlv.TLVFrom(ie.CAUSE_TYPE, tlv.Bytes([]byte{1})),
		tlv.TLV(ie.UPDATE_BAR_PFCP_SESSION_REPORT_RESPONSE_TYPE, tlv.TLVFrom(ie.BAR_ID_TYPE, tlv.Bytes(nil)))))
	if got := invalid(err); got != ie.BAR_ID_TYPE {
		t.Errorf("an empty BAR ID inside Update BAR named %d (%v), want the BAR ID", got, err)
	}
}

// 7.6.7: an IE whose Length runs past what holds it is an invalid length,
// naming that IE, in a message or inside a grouped IE.
func TestAnIELongerThanItsInputIsAnInvalidLength(t *testing.T) {
	length := func(err error) uint32 {
		var e *pfcp.InvalidLength
		if errors.As(err, &e) {
			return e.Type
		}
		return 0
	}
	//a Recovery Time Stamp claiming 8 octets, of which 4 came
	over := []byte{0, byte(ie.RECOVERY_TIME_STAMP_TYPE), 0, 8, 0, 0, 0, 1}
	var hbr message.PFCPHeartbeatResponse
	if got := length(hbr.UnmarshalBinary(over)); got != ie.RECOVERY_TIME_STAMP_TYPE {
		t.Errorf("in a message: named %d", got)
	}
	inner := []byte{0, byte(ie.BAR_ID_TYPE), 0, 5, 1}
	var rsp message.PFCPSessionReportResponse
	err := rsp.UnmarshalBinary(body(t, tlv.TLVFrom(ie.CAUSE_TYPE, tlv.Bytes([]byte{1})),
		tlv.TLVFrom(ie.UPDATE_BAR_PFCP_SESSION_REPORT_RESPONSE_TYPE, tlv.Bytes(inner))))
	if got := length(err); got != ie.BAR_ID_TYPE {
		t.Errorf("inside a grouped IE: named %d (%v), want the BAR ID", got, err)
	}
}

// 7.6.8: an optional IE whose value is faulty is discarded, the message
// processed as if it were absent and the IE's type recorded in Invalid; an
// optional grouped IE whose mandatory IE is faulty goes whole. A mandatory or
// conditional one still refuses the message, as TestAFaultyIEIsNamed has it.
func TestAFaultyOptionalIEIsDiscarded(t *testing.T) {
	var rsp message.PFCPSessionReportResponse
	err := rsp.UnmarshalBinary(body(t, tlv.TLVFrom(ie.CAUSE_TYPE, tlv.Bytes([]byte{1})),
		tlv.TLVFrom(ie.F_SEID_TYPE, tlv.Bytes([]byte{0x02}))))
	if err != nil || rsp.Cause != 1 || rsp.CPFSEID != nil || !slices.Equal(rsp.Invalid, []uint32{ie.F_SEID_TYPE}) {
		t.Errorf("a CP F-SEID of 1 octet: %v, F-SEID %v, discarded %v", err, rsp.CPFSEID, rsp.Invalid)
	}

	ip := net.IPv4(10, 0, 0, 1).To4()
	valid, err := message.EncodeBody(&message.PFCPSessionEstablishmentRequest{
		NodeID:    ie.NewNodeIDv4(ip),
		CPFSEID:   *ie.NewFSEID(&ip, nil, 1),
		CreatePDR: []ie.CreatePDR{{PDRID: 1, Precedence: 1, PDI: ie.PDI{SourceInterface: 0}}},
		CreateFAR: []ie.CreateFAR{{FARID: 1, ApplyAction: *ie.NewApplyAction(ie.APPLY_ACT_FORW)}},
	})
	if err != nil {
		t.Fatal(err)
	}
	withBar := append(valid, body(t, tlv.TLV(ie.CREATE_BAR_TYPE, tlv.TLVFrom(ie.BAR_ID_TYPE, tlv.Bytes(nil))))...)
	var req message.PFCPSessionEstablishmentRequest
	err = req.UnmarshalBinary(withBar)
	if err != nil || req.CreateBAR != nil || len(req.CreatePDR) != 1 || !slices.Equal(req.Invalid, []uint32{ie.CREATE_BAR_TYPE}) {
		t.Errorf("a Create BAR with an empty BAR ID: %v, Create BAR %v, discarded %v", err, req.CreateBAR, req.Invalid)
	}

	//inside a grouped IE, the discarded IE reaches the message's list: Create
	//QER's Averaging Window is optional, and four octets
	withWindow := append(valid, body(t, tlv.TLV(ie.CREATE_QER_TYPE,
		tlv.TLVFrom(ie.QER_ID_TYPE, tlv.Bytes([]byte{0, 0, 0, 1})),
		tlv.TLVFrom(ie.GATE_STATUS_TYPE, tlv.Bytes([]byte{0})),
		tlv.TLVFrom(ie.AVERAGING_WINDOW_TYPE, tlv.Bytes([]byte{1}))))...)
	req = message.PFCPSessionEstablishmentRequest{}
	err = req.UnmarshalBinary(withWindow)
	if err != nil || len(req.CreateQER) != 1 || req.CreateQER[0].AveragingWindow != nil || !slices.Equal(req.Invalid, []uint32{ie.AVERAGING_WINDOW_TYPE}) {
		t.Errorf("an Averaging Window of 1 octet in a Create QER: %v, discarded %v", err, req.Invalid)
	}
}
