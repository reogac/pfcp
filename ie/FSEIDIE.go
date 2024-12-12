package ie

import (
	"encoding/binary"
	"net"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

type FSEID struct {
	seid uint64
	v4   net.IP
	v6   net.IP
}

func NewFSEID(v4, v6 *net.IP, seid uint64) (ie *FSEID) {
	ie = &FSEID{
		seid: seid,
	}
	if v4 != nil {
		ie.v4 = v4.To4() //always convert to Ipv4
	}
	if v6 != nil {
		ie.v6 = v6.To16() //always convert to Ipv6
	}
	return
}

func (ie *FSEID) Seid() uint64 {
	return ie.seid
}

func (ie *FSEID) IpV4() net.IP {
	return ie.v4
}

func (ie *FSEID) IpV6() net.IP {
	return ie.v6
}

func (ie *FSEID) Bytes() (wire []byte) {
	hasV4 := len(ie.v4) > 0
	hasV6 := len(ie.v6) > 0

	flags := uint8(0)
	if hasV4 {
		flags |= 1 << 1 // 0000 0010
	}
	if hasV6 {
		flags |= 1 // 0000 0011
	}
	wire = make([]byte, 9)
	wire[0] = flags
	binary.BigEndian.PutUint64(wire[1:], ie.seid)

	if hasV4 {
		wire = append(wire, ie.v4...)
	}
	if hasV6 {
		wire = append(wire, ie.v6...)
	}
	return
}

func (ie *FSEID) UnmarshalBinary(wire []byte) (err error) {
	wireLen := len(wire)

	//flags + seid = 9 bytes
	if wireLen < 9 {
		return tlv.ErrIncomplete
	}

	ie.seid = binary.BigEndian.Uint64(wire[1:9])
	hasV6 := (wire[0] & 1) == 1
	hasV4 := ((wire[0] & 2) >> 1) == 1

	offset := 9
	if hasV4 {
		if offset+net.IPv4len > wireLen {
			return tlv.ErrIncomplete
		}
		ie.v4 = net.IP(wire[offset : offset+net.IPv4len])
		offset += net.IPv4len
	}
	if hasV6 {
		if offset+net.IPv6len > wireLen {
			return tlv.ErrIncomplete
		}
		ie.v6 = net.IP(wire[offset : offset+net.IPv6len])
		offset += net.IPv6len
	}
	//check there is no junk
	if offset < wireLen {
		return tlv.ErrTail
	}
	return nil
}
