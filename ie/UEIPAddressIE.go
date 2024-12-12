package ie

import (
	"net"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

type UEIPAddress struct {
	v4, v6  net.IP
	V6Flags *uint8
	Sd      bool
}

func NewUEIPAddress(v4, v6 *net.IP) *UEIPAddress {
	ie := &UEIPAddress{
		Sd: false,
	}
	if v4 != nil {
		ie.v4 = v4.To4()
	}

	if v6 != nil {
		ie.v6 = v6.To16()
	}
	return ie
}

func (ie *UEIPAddress) V4() net.IP {
	return ie.v4
}
func (ie *UEIPAddress) V6() net.IP {
	return ie.v6
}

func (ie *UEIPAddress) Bytes() (wire []byte) {
	hasV4 := len(ie.v4) > 0
	hasV6 := len(ie.v6) > 0
	flags := uint8(0)
	if hasV6 {
		flags |= 1
	}
	if hasV4 {
		flags |= 1 << 1
	}
	if ie.Sd {
		flags |= 1 << 2
	}
	if ie.V6Flags != nil {
		flags |= 1 << 3
	}
	wire = []byte{flags}

	if hasV4 {
		wire = append(wire, ie.v4...)
	}

	if hasV6 {
		wire = append(wire, ie.v6...)
		if ie.V6Flags != nil {
			wire = append(wire, *ie.V6Flags)
		}
	}
	return
}
func (ie *UEIPAddress) UnmarshalBinary(wire []byte) (err error) {
	wireLen := len(wire)
	if wireLen < 1 {
		return tlv.ErrIncomplete
	}
	hasV6 := wire[0]&1 == 1
	hasV4 := ((wire[0] & (1 << 1)) >> 1) == 1
	ie.Sd = ((wire[0] & (1 << 2)) >> 2) == 1
	v6Flags := ((wire[0] & (1 << 3)) >> 3) == 1

	offset := 1
	if hasV4 {
		if offset+net.IPv4len > wireLen {
			return tlv.ErrIncomplete
		}
		ie.v4 = net.IP(make([]byte, net.IPv4len))
		copy(ie.v4[:], wire[offset:offset+net.IPv4len])
		offset += net.IPv4len
	}
	if hasV6 {
		if offset+net.IPv6len > wireLen {
			return tlv.ErrIncomplete
		}
		ie.v6 = net.IP(make([]byte, net.IPv6len))
		copy(ie.v6[:], wire[offset:offset+net.IPv6len])
		offset += net.IPv6len
		if v6Flags {
			if offset+1 > wireLen {
				return tlv.ErrIncomplete
			}
			ie.V6Flags = new(uint8)
			*ie.V6Flags = wire[offset]
			offset++
		}
	}
	if offset > wireLen {
		return tlv.ErrTail
	}
	return nil
}
