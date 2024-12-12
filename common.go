package pfcp

import (
	"encoding/binary"
	"etrib5gc/logctx"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

var log logctx.LogWriter

func init() {
	log = logctx.WithFields(logctx.Fields{
		"mod": "pfcp",
	})
}
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
func Marshal3Bytes(v uint32) []byte {
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[1:], v)
	return buf[:]
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
