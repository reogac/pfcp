package ie

import (
	"encoding/binary"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

func NewSDFFilter(fd, ttc, spi, fl string, fid uint32) *SDFFilter {
	ie := &SDFFilter{}
	if fd != "" {
		ie.FlowDescription = fd
		ie.FDLength = uint16(len([]byte(fd)))
		ie.SetFDFlag()
	}

	if ttc != "" {
		ie.ToSTrafficClass = ttc
		ie.SetTTCFlag()
	}

	if spi != "" {
		ie.SecurityParameterIndex = spi
		ie.SetSPIFlag()
	}

	if fl != "" {
		ie.FlowLabel = fl
		ie.SetFLFlag()
	}

	if fid != 0 {
		ie.SDFFilterID = fid
		ie.SetBIDFlag()
	}
	return ie
}

// TODO: implement this data structure
type SDFFilter struct {
	Flags                  uint8
	FDLength               uint16
	FlowDescription        string
	ToSTrafficClass        string // 2 octets
	SecurityParameterIndex string // 4 octets
	FlowLabel              string // 3 octets
	SDFFilterID            uint32
}

func (ie *SDFFilter) HasBID() bool {
	return has5thBit(ie.Flags)
}

func (ie *SDFFilter) SetBIDFlag() {
	ie.Flags |= 0x10
}

func (ie *SDFFilter) HasFL() bool {
	return has4thBit(ie.Flags)
}

func (ie *SDFFilter) SetFLFlag() {
	ie.Flags |= 0x08
}

func (ie *SDFFilter) HasSPI() bool {
	return has3rdBit(ie.Flags)
}

func (ie *SDFFilter) SetSPIFlag() {
	ie.Flags |= 0x04
}

func (ie *SDFFilter) HasTTC() bool {
	return has2ndBit(ie.Flags)
}

func (ie *SDFFilter) SetTTCFlag() {
	ie.Flags |= 0x02
}

func (ie *SDFFilter) HasFD() bool {
	return has1stBit(ie.Flags)
}

func (ie *SDFFilter) SetFDFlag() {
	ie.Flags |= 0x01
}

func (ie *SDFFilter) Bytes() []byte {
	l := ie.bytesLen()
	if l < 1 {
		return nil
	}
	wire := make([]byte, l)

	wire[0] = ie.Flags
	offset := 2 // 2nd octet is spare

	if ie.HasFD() {
		binary.BigEndian.PutUint16(wire[offset:offset+2], ie.FDLength)
		copy(wire[offset+2:offset+2+int(ie.FDLength)], ie.FlowDescription)
		offset += 2 + int(ie.FDLength)
	}

	if ie.HasTTC() {
		copy(wire[offset:offset+2], ie.ToSTrafficClass)
		offset += 2
	}

	if ie.HasSPI() {
		copy(wire[offset:offset+4], ie.SecurityParameterIndex)
		offset += 4
	}

	if ie.HasFL() {
		copy(wire[offset:offset+3], ie.FlowLabel)
		offset += 3
	}

	if ie.HasBID() {
		binary.BigEndian.PutUint32(wire[offset:offset+4], ie.SDFFilterID)
	}

	return wire
}
func (ie *SDFFilter) UnmarshalBinary(wire []byte) (err error) {
	l := len(wire)
	if l < 3 {
		return tlv.ErrIncomplete
	}

	ie.Flags = wire[0]
	offset := 2 // 2nd octet is spare

	if ie.HasFD() {
		if len(wire[offset:]) < 3 {
			return tlv.ErrIncomplete
		}
		ie.FDLength = binary.BigEndian.Uint16(wire[offset : offset+2])
		ie.FlowDescription = string(wire[offset+2 : offset+2+int(ie.FDLength)])
		offset += 2 + int(ie.FDLength)
	}

	if ie.HasTTC() {
		if len(wire[offset:]) < 2 {
			return tlv.ErrIncomplete
		}
		ie.ToSTrafficClass = string(wire[offset : offset+2])
		offset += 2
	}

	if ie.HasSPI() {
		if len(wire[offset:]) < 4 {
			return tlv.ErrIncomplete
		}
		ie.SecurityParameterIndex = string(wire[offset : offset+4])
		offset += 4
	}

	if ie.HasFL() {
		if len(wire[offset:]) < 3 {
			return tlv.ErrIncomplete
		}
		ie.FlowLabel = string(wire[offset : offset+3])
		offset += 3
	}

	if ie.HasBID() {
		if len(wire[offset:]) < 4 {
			return tlv.ErrIncomplete
		}
		ie.SDFFilterID = binary.BigEndian.Uint32(wire[offset : offset+4])
	}

	return nil
}

func (ie *SDFFilter) bytesLen() int {
	l := 2 // Flag + Spare

	if ie.HasFD() {
		l += 2 + int(ie.FDLength)
	}

	if ie.HasTTC() {
		l += 2
	}

	if ie.HasSPI() {
		l += 4
	}

	if ie.HasFL() {
		l += 3
	}

	if ie.HasBID() {
		l += 4
	}

	return l
}
