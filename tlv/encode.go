package tlv

type fieldKind uint8

const (
	kindEmpty fieldKind = iota
	kindError
	kindBytes
	kindTLV
	kindTLVFrom
)

// Field is something that encodes. The zero value encodes to nothing.
type Field struct {
	kind fieldKind
	typ  uint32
	//enterprise is the Enterprise ID of a vendor-specific element
	enterprise uint16
	object     any
}

// Fielder is anything that can say how it encodes.
type Fielder interface {
	Field() Field
}

// Field implements Fielder.
func (f Field) Field() Field {
	return f
}

// FieldError is a Field that fails to encode with e.
func FieldError(e error) Field {
	if e == nil {
		e = ErrErrorField
	}
	return Field{kind: kindError, object: e}
}

// Bytes is a Field that encodes to b as it is.
func Bytes(b []byte) Field {
	return Field{kind: kindBytes, object: b}
}

// TLV is an element of type typ whose value is values, in order.
func TLV(typ uint32, values ...Field) Field {
	if typ == 0 || typ >= enterpriseBit {
		return FieldError(ErrType)
	}
	return Field{kind: kindTLV, typ: typ, object: values}
}

// TLVFrom is TLV over Fielders.
func TLVFrom(typ uint32, values ...Fielder) Field {
	if typ == 0 || typ >= enterpriseBit {
		return FieldError(ErrType)
	}
	return Field{kind: kindTLVFrom, typ: typ, object: values}
}

// TLVBytes is an element of type typ whose value is value.
func TLVBytes(typ uint32, value []byte) Field {
	return TLV(typ, Bytes(value))
}

// Enterprise is a vendor-specific element: typ must have its top bit set, and
// enterprise goes on the wire in front of values, inside the length.
func Enterprise(typ uint32, enterprise uint16, values ...Field) Field {
	if !IsEnterprise(typ) || typ > MaxType {
		return FieldError(ErrType)
	}
	return Field{kind: kindTLV, typ: typ, enterprise: enterprise, object: values}
}

// Encode appends f to b.
//
// An element is written in place: its header is reserved, its children
// appended after it, and the length filled in last, so nesting costs no buffer
// per level.
func (f Field) Encode(b []byte) ([]byte, error) {
	switch f.kind {
	case kindEmpty:
		return b, nil
	case kindError:
		return b, f.object.(error)
	case kindBytes:
		return append(b, f.object.([]byte)...), nil
	case kindTLV, kindTLVFrom:
		start := len(b)
		b = append(b, byte(f.typ>>8), byte(f.typ), 0, 0)
		if IsEnterprise(f.typ) {
			b = append(b, byte(f.enterprise>>8), byte(f.enterprise))
		}
		var err error
		if b, err = f.appendValue(b); err != nil {
			return b, err
		}
		n := len(b) - start - headerLen
		if n > MaxLength {
			return b, ErrLength
		}
		b[start+2], b[start+3] = byte(n>>8), byte(n)
		return b, nil
	default:
		panic(f.kind)
	}
}

// appendValue appends what an element holds, without its header.
func (f Field) appendValue(b []byte) (_ []byte, err error) {
	if f.kind == kindTLVFrom {
		for _, v := range f.object.([]Fielder) {
			if b, err = v.Field().Encode(b); err != nil {
				return
			}
		}
		return b, nil
	}
	for _, v := range f.object.([]Field) {
		if b, err = v.Encode(b); err != nil {
			return
		}
	}
	return b, nil
}

// Encode encodes fields one after another.
func Encode(fields ...Field) (wire []byte, err error) {
	for _, f := range fields {
		if wire, err = f.Encode(wire); err != nil {
			return nil, err
		}
	}
	return
}

// EncodeFrom is Encode over Fielders.
func EncodeFrom(fields ...Fielder) (wire []byte, err error) {
	for _, f := range fields {
		if wire, err = f.Field().Encode(wire); err != nil {
			return nil, err
		}
	}
	return
}

// AppendValue appends the value of the element f encodes to, without its
// header. A PFCP message is an element of its message type whose value is the
// message body, and on the wire the header in front of that body is the PFCP
// header rather than an element header (TS 29.244 §7.2.1): this is how a body
// is written into a buffer already holding one.
func AppendValue(b []byte, f Fielder) ([]byte, error) {
	field := f.Field()
	switch field.kind {
	case kindError:
		return b, field.object.(error)
	case kindTLV, kindTLVFrom:
		if IsEnterprise(field.typ) {
			b = append(b, byte(field.enterprise>>8), byte(field.enterprise))
		}
		return field.appendValue(b)
	default:
		return b, ErrType
	}
}

// EncodeValueOnly is AppendValue into a new buffer.
func EncodeValueOnly(f Fielder) ([]byte, error) {
	return AppendValue(nil, f)
}
