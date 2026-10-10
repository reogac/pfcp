package tlv

import (
	"bytes"
	"errors"
	"testing"
)

func TestEncodeFixesTypeAndLengthAtTwoOctets(t *testing.T) {
	//Recovery Time Stamp (96) as TS 29.244 §8.1.1 lays it out; the NDN encoding
	//this replaced wrote 60 04 01 02 03 04
	got, err := Encode(TLVBytes(96, []byte{1, 2, 3, 4}))
	if err != nil {
		t.Fatal(err)
	}
	want := []byte{0x00, 0x60, 0x00, 0x04, 1, 2, 3, 4}
	if !bytes.Equal(got, want) {
		t.Fatalf("got % x, want % x", got, want)
	}
}

func TestEncodeGroupedCountsItsChildren(t *testing.T) {
	//Create PDR (1) holding PDR ID (56) and a PDI (2) holding Source Interface (20)
	f := TLV(1,
		TLVBytes(56, []byte{0x00, 0x07}),
		TLV(2, TLVBytes(20, []byte{0x00})),
	)
	got, err := Encode(f)
	if err != nil {
		t.Fatal(err)
	}
	want := []byte{
		0x00, 0x01, 0x00, 0x0f,
		0x00, 0x38, 0x00, 0x02, 0x00, 0x07,
		0x00, 0x02, 0x00, 0x05,
		0x00, 0x14, 0x00, 0x01, 0x00,
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("got % x, want % x", got, want)
	}
}

func TestEncodeLengthPastOneOctet(t *testing.T) {
	got, err := Encode(TLVBytes(22, make([]byte, 300)))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got[:4], []byte{0x00, 0x16, 0x01, 0x2c}) || len(got) != 304 {
		t.Fatalf("header % x, len %d", got[:4], len(got))
	}
}

func TestEncodeEnterpriseCarriesItsIdInsideTheLength(t *testing.T) {
	got, err := Encode(Enterprise(0x8001, 0x1234, Bytes([]byte{0xaa})))
	if err != nil {
		t.Fatal(err)
	}
	want := []byte{0x80, 0x01, 0x00, 0x03, 0x12, 0x34, 0xaa}
	if !bytes.Equal(got, want) {
		t.Fatalf("got % x, want % x", got, want)
	}
}

func TestEncodeRefusesTypesPfcpDoesNotAllow(t *testing.T) {
	for _, f := range []Field{
		TLVBytes(0, nil),       //reserved
		TLVBytes(0x8001, nil),  //vendor-specific without an Enterprise ID
		Enterprise(0x0001, 1),  //not vendor-specific
		Enterprise(0x10000, 1), //does not fit two octets
		TLVBytes(1, make([]byte, MaxLength+1)),
	} {
		if _, err := Encode(f); err == nil {
			t.Errorf("%+v encoded", f)
		}
	}
	if _, err := Encode(TLVBytes(1, make([]byte, MaxLength+1))); !errors.Is(err, ErrLength) {
		t.Errorf("overlong value: %v", err)
	}
}

func TestAppendValueLeavesOutTheHeader(t *testing.T) {
	//a message is an element of its type whose value is the body; the body is
	//what follows a PFCP header on the wire
	msg := TLV(1, TLVBytes(96, []byte{1, 2, 3, 4}))
	got, err := AppendValue([]byte{0xff}, msg)
	if err != nil {
		t.Fatal(err)
	}
	want := []byte{0xff, 0x00, 0x60, 0x00, 0x04, 1, 2, 3, 4}
	if !bytes.Equal(got, want) {
		t.Fatalf("got % x, want % x", got, want)
	}
}

func BenchmarkEncodeGrouped(b *testing.B) {
	f := TLV(1,
		TLVBytes(56, []byte{0x00, 0x07}),
		TLVBytes(29, []byte{0, 0, 0, 100}),
		TLV(2, TLVBytes(20, []byte{0x00}), TLVBytes(21, make([]byte, 9))),
		TLVBytes(108, []byte{0, 0, 0, 1}),
	)
	buf := make([]byte, 0, 256)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := f.Encode(buf[:0]); err != nil {
			b.Fatal(err)
		}
	}
}
