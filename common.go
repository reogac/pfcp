package pfcp

import (
	"encoding/binary"

	"github.com/reogac/pfcp/v2/tlv"
)

func MarshalUint8(v uint8) []byte {
	return []byte{v}
}

func UnmarshalUint8(buf []byte, v *uint8) error {
	if len(buf) < 1 {
		return tlv.ErrIncomplete
	} else if len(buf) > 1 {
		return tlv.ErrTail
	}
	*v = buf[0]
	return nil
}

func MarshalUint16(v uint16) []byte {
	var buf [2]byte
	binary.BigEndian.PutUint16(buf[:], v)
	return buf[:]
}

func UnmarshalUint16(buf []byte, v *uint16) error {
	if len(buf) < 2 {
		return tlv.ErrIncomplete
	} else if len(buf) > 2 {
		return tlv.ErrTail
	}
	*v = binary.BigEndian.Uint16(buf[:])
	return nil
}
// Marshal3Bytes writes the low three octets of v, big endian. It wrote four
// octets into the last three of its buffer, and so panicked on every call
func Marshal3Bytes(v uint32) []byte {
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:], v)
	return buf[1:]
}
func Unmarshal3Bytes(buf []byte, v *uint32) error {
	if len(buf) < 3 {
		return tlv.ErrIncomplete
	} else if len(buf) > 3 {
		return tlv.ErrTail
	}
	newBuf := make([]byte, 4)
	copy(newBuf[1:], buf)
	*v = binary.BigEndian.Uint32(newBuf[:])
	return nil
}

func MarshalUint32(v uint32) []byte {
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:], v)
	return buf[:]
}
func UnmarshalUint32(buf []byte, v *uint32) error {
	if len(buf) < 4 {
		return tlv.ErrIncomplete
	} else if len(buf) > 4 {
		return tlv.ErrTail
	}
	*v = binary.BigEndian.Uint32(buf[:])
	return nil
}

// UnmarshalExtendable reads an Extendable IE of `known` octets into v, big
// endian. TS 29.244 8.1.2: an Extendable IE has a variable number of octets, at
// least the `min` it had when first defined, and "the legacy receiving entity
// shall ignore the unknown octets". So a value shorter than `known` but at least
// `min` long, an older sender's, leaves the octets it lacks 0, and one longer
// than `known`, a newer sender's, has its further octets ignored.
func UnmarshalExtendable[T uint8 | uint16 | uint32](buf []byte, min, known int, v *T) error {
	if len(buf) < min {
		return tlv.ErrIncomplete
	}
	var x uint64
	for i := 0; i < known; i++ {
		x <<= 8
		if i < len(buf) {
			x |= uint64(buf[i])
		}
	}
	*v = T(x)
	return nil
}
