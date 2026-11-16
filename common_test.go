package pfcp

import (
	"errors"
	"testing"

	"github.com/reogac/pfcp/tlv"
)

// An Extendable IE decodes from its first-defined length up: missing known
// octets are 0, octets beyond the known ones ignored (TS 29.244 8.1.2).
func TestAnExtendableIEDecodesAnyLengthFromItsMinimum(t *testing.T) {
	var v uint32
	for _, c := range []struct {
		buf  []byte
		want uint32
	}{
		{[]byte{0x12, 0x34}, 0x123400},             // Reporting Triggers as first defined, 2 octets
		{[]byte{0x12, 0x34, 0x56}, 0x123456},       // the 3 this library knows
		{[]byte{0x12, 0x34, 0x56, 0x78}, 0x123456}, // a newer sender's fourth, ignored
	} {
		if err := UnmarshalExtendable(c.buf, 2, 3, &v); err != nil || v != c.want {
			t.Errorf("% x: %06x, %v; want %06x", c.buf, v, err, c.want)
		}
	}
	if err := UnmarshalExtendable([]byte{0x12}, 2, 3, &v); !errors.Is(err, tlv.ErrIncomplete) {
		t.Errorf("shorter than its minimum: %v", err)
	}
	var id uint16
	if err := UnmarshalExtendable([]byte{0, 7, 0xff}, 2, 2, &id); err != nil || id != 7 {
		t.Errorf("a PDR ID with an octet more: %d, %v", id, err)
	}
}

// A 3-octet value is its low three octets, and reads back.
func TestThreeOctetsRoundTrip(t *testing.T) {
	b := Marshal3Bytes(0x123456)
	var v uint32
	if len(b) != 3 || b[0] != 0x12 || b[2] != 0x56 || Unmarshal3Bytes(b, &v) != nil || v != 0x123456 {
		t.Errorf("% x read back as %06x", b, v)
	}
}
