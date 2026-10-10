package tlv

import "encoding"

// Unmarshaler is anything that decodes itself from an element's type and value.
type Unmarshaler interface {
	UnmarshalTLV(typ uint32, value []byte) error
}

// Element is one decoded element. Value aliases the input.
type Element struct {
	Type uint32
	// EnterpriseID is set for a vendor-specific type and is not part of Value.
	EnterpriseID uint16
	Value        []byte
}

// UnmarshalValue decodes the value into u.
func (e Element) UnmarshalValue(u encoding.BinaryUnmarshaler) error {
	return u.UnmarshalBinary(e.Value)
}

// Unmarshal decodes the type and value into u.
func (e Element) Unmarshal(u Unmarshaler) error {
	return u.UnmarshalTLV(e.Type, e.Value)
}

// Next decodes the element at the start of b and returns what follows it.
func Next(b []byte) (e Element, rest []byte, err error) {
	if len(b) < headerLen {
		return e, b, ErrIncomplete
	}
	e.Type = uint32(b[0])<<8 | uint32(b[1])
	if e.Type == 0 {
		return e, b, ErrType
	}
	n := int(b[2])<<8 | int(b[3])
	if len(b)-headerLen < n {
		return e, b, ErrIncomplete
	}
	e.Value, rest = b[headerLen:headerLen+n], b[headerLen+n:]
	if IsEnterprise(e.Type) {
		if len(e.Value) < enterpriseLen {
			return e, b, ErrIncomplete
		}
		e.EnterpriseID = uint16(e.Value[0])<<8 | uint16(e.Value[1])
		e.Value = e.Value[enterpriseLen:]
	}
	return e, rest, nil
}

// Elements decodes every element of b, which must hold nothing else.
//
// It fails on the first malformed element rather than returning those before
// it: the NDN decoder this replaced stopped there silently, so a truncated
// message decoded as a valid one missing its last IEs.
func Elements(b []byte) ([]Element, error) {
	//count first, so the slice is allocated once and only for valid input
	count := 0
	for rest := b; len(rest) > 0; count++ {
		var err error
		if _, rest, err = Next(rest); err != nil {
			return nil, err
		}
	}
	elements := make([]Element, count)
	for i := range elements {
		elements[i], b, _ = Next(b)
	}
	return elements, nil
}

// Decode decodes wire, which must hold exactly one element, into u.
func Decode(wire []byte, u Unmarshaler) error {
	e, rest, err := Next(wire)
	if err != nil {
		return err
	}
	if len(rest) > 0 {
		return ErrTail
	}
	return e.Unmarshal(u)
}
