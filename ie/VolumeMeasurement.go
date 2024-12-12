package ie

import (
	"encoding/binary"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

type VolumeMeasurement struct {
	Flags                   uint8
	TotalVolume             uint64
	UplinkVolume            uint64
	DownlinkVolume          uint64
	TotalNumberOfPackets    uint64
	UplinkNumberOfPackets   uint64
	DownlinkNumberOfPackets uint64
}

func (ie *VolumeMeasurement) HasDLNOP() bool {
	return has6thBit(ie.Flags)
}

func (ie *VolumeMeasurement) HasULNOP() bool {
	return has5thBit(ie.Flags)
}

func (ie *VolumeMeasurement) HasTONOP() bool {
	return has4thBit(ie.Flags)
}

func (ie *VolumeMeasurement) HasDLVOL() bool {
	return has3rdBit(ie.Flags)
}

func (ie *VolumeMeasurement) HasULVOL() bool {
	return has2ndBit(ie.Flags)
}

func (ie *VolumeMeasurement) HasTOVOL() bool {
	return has1stBit(ie.Flags)
}

func (ie *VolumeMeasurement) Bytes() (wire []byte) {
	l := ie.bytesLen()
	if l < 1 {
		return
	}

	wire = make([]byte, l)
	wire[0] = ie.Flags
	offset := 1

	if ie.HasTOVOL() {
		if l < offset+8 {
			return
		}
		binary.BigEndian.PutUint64(wire[offset:offset+8], ie.TotalVolume)
		offset += 8
	}

	if ie.HasULVOL() {
		if l < offset+8 {
			return
		}
		binary.BigEndian.PutUint64(wire[offset:offset+8], ie.UplinkVolume)
		offset += 8
	}

	if ie.HasDLVOL() {
		if l < offset+8 {
			return
		}
		binary.BigEndian.PutUint64(wire[offset:offset+8], ie.DownlinkVolume)
		offset += 8
	}

	if ie.HasTONOP() {
		if l < offset+8 {
			return
		}
		binary.BigEndian.PutUint64(wire[offset:offset+8], ie.TotalNumberOfPackets)
		offset += 8
	}

	if ie.HasULNOP() {
		if l < offset+8 {
			return
		}
		binary.BigEndian.PutUint64(wire[offset:offset+8], ie.UplinkNumberOfPackets)
		offset += 8
	}

	if ie.HasDLNOP() {
		if l < offset+8 {
			return
		}
		binary.BigEndian.PutUint64(wire[offset:offset+8], ie.DownlinkNumberOfPackets)
	}

	return nil
}
func (ie *VolumeMeasurement) UnmarshalBinary(wire []byte) (err error) {
	l := len(wire)
	if l < 2 {
		return tlv.ErrIncomplete
	}

	ie.Flags = wire[0]
	offset := 1

	if ie.HasTOVOL() {
		if l < offset+8 {
			return tlv.ErrIncomplete
		}
		ie.TotalVolume = binary.BigEndian.Uint64(wire[offset : offset+8])
		offset += 8
	}

	if ie.HasULVOL() {
		if l < offset+8 {
			return tlv.ErrIncomplete
		}
		ie.UplinkVolume = binary.BigEndian.Uint64(wire[offset : offset+8])
		offset += 8
	}

	if ie.HasDLVOL() {
		if l < offset+8 {
			return tlv.ErrIncomplete
		}
		ie.DownlinkVolume = binary.BigEndian.Uint64(wire[offset : offset+8])
		offset += 8
	}

	if ie.HasTONOP() {
		if l < offset+8 {
			return tlv.ErrIncomplete
		}
		ie.TotalNumberOfPackets = binary.BigEndian.Uint64(wire[offset : offset+8])
		offset += 8
	}

	if ie.HasULNOP() {
		if l < offset+8 {
			return tlv.ErrIncomplete
		}
		ie.UplinkNumberOfPackets = binary.BigEndian.Uint64(wire[offset : offset+8])
		offset += 8
	}

	if ie.HasDLNOP() {
		if l < offset+8 {
			return tlv.ErrIncomplete
		}
		ie.DownlinkNumberOfPackets = binary.BigEndian.Uint64(wire[offset : offset+8])
	}

	return nil
}

func (ie *VolumeMeasurement) bytesLen() int {
	l := 1
	if ie.HasTOVOL() {
		l += 8
	}
	if ie.HasULVOL() {
		l += 8
	}
	if ie.HasDLVOL() {
		l += 8
	}
	if ie.HasTONOP() {
		l += 8
	}
	if ie.HasULNOP() {
		l += 8
	}
	if ie.HasDLNOP() {
		l += 8
	}

	return l
}
