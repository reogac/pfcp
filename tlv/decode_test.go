package tlv

import (
	"bytes"
	"errors"
	"testing"
)

func TestElementsReadsTwoOctetTypeAndLength(t *testing.T) {
	wire := []byte{
		0x00, 0x60, 0x00, 0x04, 1, 2, 3, 4,
		0x00, 0x16, 0x01, 0x2c,
	}
	wire = append(wire, make([]byte, 300)...)
	es, err := Elements(wire)
	if err != nil {
		t.Fatal(err)
	}
	if len(es) != 2 || es[0].Type != 96 || !bytes.Equal(es[0].Value, []byte{1, 2, 3, 4}) ||
		es[1].Type != 22 || len(es[1].Value) != 300 {
		t.Fatalf("%+v", es)
	}
}

func TestElementsOfNothingIsNoElements(t *testing.T) {
	es, err := Elements(nil)
	if err != nil || len(es) != 0 {
		t.Fatalf("%v %v", es, err)
	}
}

func TestElementsFailsRatherThanStopping(t *testing.T) {
	valid := []byte{0x00, 0x60, 0x00, 0x04, 1, 2, 3, 4}
	for name, tail := range map[string][]byte{
		"truncated header": {0x00, 0x60, 0x00},
		"truncated value":  {0x00, 0x60, 0x00, 0x04, 1, 2},
		"reserved type":    {0x00, 0x00, 0x00, 0x00},
		"enterprise no id": {0x80, 0x01, 0x00, 0x01, 0xaa},
	} {
		if es, err := Elements(append(append([]byte{}, valid...), tail...)); err == nil {
			t.Errorf("%s: decoded %+v", name, es)
		}
	}
}

func TestElementsSplitsOffTheEnterpriseId(t *testing.T) {
	es, err := Elements([]byte{0x80, 0x01, 0x00, 0x03, 0x12, 0x34, 0xaa})
	if err != nil {
		t.Fatal(err)
	}
	if es[0].Type != 0x8001 || es[0].EnterpriseID != 0x1234 || !bytes.Equal(es[0].Value, []byte{0xaa}) {
		t.Fatalf("%+v", es[0])
	}
}

type element struct {
	typ   uint32
	value []byte
}

func (e *element) UnmarshalTLV(typ uint32, value []byte) error {
	e.typ, e.value = typ, value
	return nil
}

func TestDecodeTakesExactlyOneElement(t *testing.T) {
	var e element
	if err := Decode([]byte{0x00, 0x01, 0x00, 0x01, 0x07}, &e); err != nil || e.typ != 1 || e.value[0] != 7 {
		t.Fatalf("%+v %v", e, err)
	}
	if err := Decode([]byte{0x00, 0x01, 0x00, 0x01, 0x07, 0x00}, &e); !errors.Is(err, ErrTail) {
		t.Fatalf("trailing octet: %v", err)
	}
}

func TestEncodeThenDecodeIsIdentity(t *testing.T) {
	f := TLV(1,
		TLVBytes(56, []byte{0x00, 0x07}),
		Enterprise(0x8002, 99, Bytes([]byte("upf"))),
		TLV(2, TLVBytes(20, []byte{0x00})),
	)
	wire, err := Encode(f)
	if err != nil {
		t.Fatal(err)
	}
	var outer element
	if err = Decode(wire, &outer); err != nil {
		t.Fatal(err)
	}
	es, err := Elements(outer.value)
	if err != nil || len(es) != 3 || es[1].EnterpriseID != 99 || string(es[1].Value) != "upf" {
		t.Fatalf("%+v %v", es, err)
	}
	inner, err := Elements(es[2].Value)
	if err != nil || len(inner) != 1 || inner[0].Type != 20 {
		t.Fatalf("%+v %v", inner, err)
	}
}

func FuzzElements(f *testing.F) {
	f.Add([]byte{0x00, 0x60, 0x00, 0x04, 1, 2, 3, 4})
	f.Add([]byte{0x80, 0x01, 0x00, 0x03, 0x12, 0x34, 0xaa})
	f.Fuzz(func(t *testing.T, wire []byte) {
		es, err := Elements(wire)
		if err != nil {
			return
		}
		//whatever decodes must encode back to the same octets
		var out []byte
		for _, e := range es {
			var field Field
			if IsEnterprise(e.Type) {
				field = Enterprise(e.Type, e.EnterpriseID, Bytes(e.Value))
			} else {
				field = TLVBytes(e.Type, e.Value)
			}
			if out, err = field.Encode(out); err != nil {
				t.Fatal(err)
			}
		}
		if !bytes.Equal(out, wire) {
			t.Fatalf("% x re-encoded as % x", wire, out)
		}
	})
}
