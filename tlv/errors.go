package tlv

import (
	"errors"
	"fmt"
)

var (
	// ErrIncomplete is input that ends inside an element, or a value shorter
	// than its type requires.
	ErrIncomplete = errors.New("incomplete input")
	// ErrTail is input left over after what was to be decoded.
	ErrTail = errors.New("junk after end of TLV")
	// ErrType is a type PFCP does not allow: 0 is reserved, a type does not
	// fit two octets, or a vendor-specific type without an Enterprise ID.
	ErrType = errors.New("TLV type out of range")
	// ErrLength is a value too long for two octets of length.
	ErrLength = errors.New("TLV value longer than 65535 octets")
	// ErrCritical is an element its container does not recognize.
	ErrCritical = errors.New("unrecognized TLV type")
	// ErrErrorField is a FieldError made with a nil error.
	ErrErrorField = errors.New("Error(nil) field")
)

// LengthError is an element whose Length runs past the input that holds it:
// Type is the element's. It is ErrIncomplete to errors.Is, which is what it
// was before it named the element.
type LengthError struct {
	Type uint32
}

func (e *LengthError) Error() string {
	return fmt.Sprintf("TLV type %d: its length runs past the input", e.Type)
}

func (e *LengthError) Is(target error) bool { return target == ErrIncomplete }
