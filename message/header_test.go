package message

import (
	"bytes"
	"encoding/hex"
	"errors"
	"testing"

	"github.com/reogac/pfcp/ie"
)

// goPfcpMessages are whole messages as github.com/wmnsk/go-pfcp v0.0.24 writes
// them, captured once rather than imported.
var goPfcpMessages = []struct {
	name     string
	wire     string
	seid     uint64
	seq      uint32
	priority uint8
}{
	{"heartbeat request", "2001000c1234560000600004e8fe6f80", 0, 0x123456, 0},
	{"heartbeat response", "2002000cffffff0000600004e8fe6f80", 0, 0xffffff, 0},
	{"association setup request", "2005001500000900003c0005000a00000100600004e8fe6f80", 0, 9, 0},
	{"establishment request", "21320026000000000000000000000700003c0005000a0000010039000d020000000000000abc0a000001", 0, 7, 0},
	{"establishment request with priority", "23320026000000000000000000000830003c0005000a0000010039000d020000000000000abc0a000001", 0, 8, 3},
	{"report request", "213800111122334455667788010203000027000102", 0x1122334455667788, 0x010203, 0},
	{"deletion response", "21370011000000000000dead000005000013000101", 0xdead, 5, 0},
}

func TestMessagesMatchAnIndependentEncoder(t *testing.T) {
	for _, c := range goPfcpMessages {
		wire, _ := hex.DecodeString(c.wire)
		h, m, rest, err := Decode(wire)
		if err != nil {
			t.Errorf("%s: %v", c.name, err)
			continue
		}
		if h.Seid != c.seid || h.Seq != c.seq || h.Priority != c.priority || len(rest) != 0 {
			t.Errorf("%s: %+v", c.name, h)
		}
		again, err := AppendMessage(nil, h, m)
		if err != nil || !bytes.Equal(again, wire) {
			t.Errorf("%s: encodes % x, want % x (%v)", c.name, again, wire, err)
		}
	}
}

func TestBodyStandsAloneFromTheHeader(t *testing.T) {
	//what an SMF does: no header on either side
	m := &PFCPSessionDeletionRequest{}
	body, err := EncodeBody(&PFCPHeartbeatRequest{RecoveryTimeStamp: 7})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(body, []byte{0x00, 0x60, 0x00, 0x04, 0, 0, 0, 7}) {
		t.Fatalf("% x", body)
	}
	got, err := DecodeBody(Type(ie.MSG_PFCP_HEARTBEAT_REQUEST_TYPE), body)
	if err != nil || got.(*PFCPHeartbeatRequest).RecoveryTimeStamp != 7 {
		t.Fatalf("%+v %v", got, err)
	}
	//and framing a body already encoded writes what AppendMessage does
	h := Header{Type: m.MsgType(), Seid: 3, Seq: 4}
	viaBody, _ := EncodeBody(m)
	raw, err := AppendRaw(nil, h, viaBody)
	whole, _ := AppendMessage(nil, h, m)
	if err != nil || !bytes.Equal(raw, whole) {
		t.Fatalf("% x vs % x (%v)", raw, whole, err)
	}
}

func TestDecodeHeaderRefuses(t *testing.T) {
	hb := []byte{0x20, 0x01, 0x00, 0x0c, 0, 0, 1, 0, 0x00, 0x60, 0x00, 0x04, 1, 2, 3, 4}
	set := func(i int, v byte) []byte {
		b := append([]byte{}, hb...)
		b[i] = v
		return b
	}
	for name, c := range map[string]struct {
		wire []byte
		want error
	}{
		"empty":            {nil, ErrShort},
		"version 2":        {set(0, 0x40), ErrVersion},
		"length past end":  {set(3, 0x0d), ErrShort},
		"length in header": {[]byte{0x20, 0x01, 0x00, 0x02, 0, 0}, ErrShort},
		"node with SEID":   {set(0, 0x21), ErrSFlag},
		"unknown type":     {set(1, 99), ErrUnknownType},
		"junk without FO":  {append(append([]byte{}, hb...), 0), ErrTail},
		"session no SEID":  {[]byte{0x20, 0x32, 0x00, 0x04, 0, 0, 1, 0}, ErrSFlag},
	} {
		if _, _, _, err := DecodeHeader(c.wire); !errors.Is(err, c.want) {
			t.Errorf("%s: %v, want %v", name, err, c.want)
		}
	}
}

func TestFollowOnLeavesTheNextMessage(t *testing.T) {
	first, _ := AppendMessage(nil, Header{Seq: 1, FollowOn: true}, &PFCPHeartbeatRequest{RecoveryTimeStamp: 1})
	both, _ := AppendMessage(first, Header{Seq: 2}, &PFCPHeartbeatRequest{RecoveryTimeStamp: 2})
	h, _, rest, err := Decode(both)
	if err != nil || h.Seq != 1 || !h.FollowOn {
		t.Fatalf("%+v %v", h, err)
	}
	h, m, rest, err := Decode(rest)
	if err != nil || h.Seq != 2 || len(rest) != 0 || m.(*PFCPHeartbeatRequest).RecoveryTimeStamp != 2 {
		t.Fatalf("%+v %v", h, err)
	}
}

func TestEveryTypeAgreesWithItsBody(t *testing.T) {
	known := 0
	for i := 0; i < 256; i++ {
		ty := Type(i)
		if !ty.Known() {
			continue
		}
		known++
		if got := New(ty).MsgType(); got != ty {
			t.Errorf("%v: body says %v", ty, got)
		}
		if ty.IsRequest() {
			if rsp := ty.Response(); rsp.IsRequest() || rsp.IsSession() != ty.IsSession() || types[rsp].pair != ty {
				t.Errorf("%v pairs with %v", ty, rsp)
			}
		}
	}
	if known != 25 {
		t.Errorf("%d types known", known)
	}
	//parity says nothing: 12 is a request and 11 a response
	if !Type(12).IsRequest() || Type(11).IsRequest() || !Type(50).IsRequest() {
		t.Error("request test follows parity")
	}
}

func FuzzDecode(f *testing.F) {
	for _, c := range goPfcpMessages {
		wire, _ := hex.DecodeString(c.wire)
		f.Add(wire)
	}
	f.Fuzz(func(t *testing.T, wire []byte) {
		h, m, _, err := Decode(wire)
		if err != nil {
			return
		}
		//whatever decodes encodes, and decodes again to the same header
		again, err := AppendMessage(nil, h, m)
		if err != nil {
			t.Fatal(err)
		}
		h2, _, _, err := Decode(again)
		if err != nil || h2 != h {
			t.Fatalf("%+v became %+v (%v)", h, h2, err)
		}
	})
}

func BenchmarkDecodeHeader(b *testing.B) {
	wire, _ := hex.DecodeString(goPfcpMessages[3].wire)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, _, _, err := DecodeHeader(wire); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAppendHeartbeat(b *testing.B) {
	m := &PFCPHeartbeatRequest{RecoveryTimeStamp: 1}
	buf := make([]byte, 0, 64)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := AppendMessage(buf[:0], Header{Seq: 1}, m); err != nil {
			b.Fatal(err)
		}
	}
}
