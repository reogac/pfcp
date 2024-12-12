package ie

import (
	"encoding/binary"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

func NewVolumeThreshold(flags uint8, tvol, uvol, dvol uint64) *VolumeThreshold {
	ie := &VolumeThreshold{Flags: flags}

	if ie.HasTOVOL() {
		ie.TotalVolume = tvol
	}

	if ie.HasULVOL() {
		ie.UplinkVolume = uvol
	}

	if ie.HasDLVOL() {
		ie.DownlinkVolume = dvol
	}

	return ie
}

type VolumeThreshold struct {
	TotalVolume    uint64
	UplinkVolume   uint64
	DownlinkVolume uint64
	Flags          uint8
}

func (ie *VolumeThreshold) HasTOVOL() bool {
	return has1stBit(ie.Flags)
}

func (ie *VolumeThreshold) HasULVOL() bool {
	return has2ndBit(ie.Flags)
}

func (ie *VolumeThreshold) HasDLVOL() bool {
	return has3rdBit(ie.Flags)
}

func (ie *VolumeThreshold) Bytes() []byte {
	wire := make([]byte, ie.bytesLen())
	wire[0] = ie.Flags
	offset := 1

	if ie.HasTOVOL() {
		binary.BigEndian.PutUint64(wire[offset:offset+8], ie.TotalVolume)
		offset += 8
	}

	if ie.HasULVOL() {
		binary.BigEndian.PutUint64(wire[offset:offset+8], ie.UplinkVolume)
		offset += 8
	}

	if ie.HasDLVOL() {
		binary.BigEndian.PutUint64(wire[offset:offset+8], ie.DownlinkVolume)
	}

	return wire
}

func (ie *VolumeThreshold) UnmarshalBinary(wire []byte) (err error) {
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
	}

	return nil
}

func (f *VolumeThreshold) bytesLen() int {
	l := 1
	if f.HasTOVOL() {
		l += 8
	}
	if f.HasULVOL() {
		l += 8
	}
	if f.HasDLVOL() {
		l += 8
	}

	return l
}
