package message_test

import (
	"errors"
	"slices"
	"testing"

	"github.com/reogac/pfcp"
	"github.com/reogac/pfcp/ie"
	"github.com/reogac/pfcp/message"
	"github.com/reogac/pfcp/tlv"
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
