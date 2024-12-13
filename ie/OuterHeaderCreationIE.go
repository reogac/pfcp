package ie

import (
	"encoding/binary"
	"github.com/reogac/pfcp"
	"fmt"
	"net"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

const (
	OuterHeaderCreationGtpUUdpIpv4 uint8 = 1
	OuterHeaderCreationGtpUUdpIpv6 uint8 = 1 << 1
	OuterHeaderCreationUdpIpv4     uint8 = 1 << 2
	OuterHeaderCreationUdpIpv6     uint8 = 1 << 3
)

type OuterHeaderCreation struct {
	desc uint8
	teid uint32
	ip   net.IP //must be the right size respecting to the description
	port uint16
}

func NewOuterHeaderCreationGtpIpv4(ip net.IP, teid uint32) *OuterHeaderCreation {
	ie := &OuterHeaderCreation{
		desc: OuterHeaderCreationGtpUUdpIpv4,
		teid: teid,
	}
	ie.ip = ip.To4()
	return ie
}

func NewOuterHeaderCreationGtpIpv6(ip net.IP, teid uint32) *OuterHeaderCreation {
	ie := &OuterHeaderCreation{
		desc: OuterHeaderCreationGtpUUdpIpv6,
		teid: teid,
	}
	ie.ip = ip.To16()
	return ie
}

func NewOuterHeaderCreationUdpIpv4(ip net.IP, port uint16) *OuterHeaderCreation {
	ie := &OuterHeaderCreation{
		desc: OuterHeaderCreationUdpIpv4,
		port: port,
	}
	ie.ip = ip.To4()
	return ie
}

func NewOuterHeaderCreationUdpIpv6(ip net.IP, port uint16) *OuterHeaderCreation {
	ie := &OuterHeaderCreation{
		desc: OuterHeaderCreationUdpIpv6,
		port: port,
	}
	ie.ip = ip.To16()
	return ie
}

func (ie *OuterHeaderCreation) IsGtp() bool {
	return ie.desc == OuterHeaderCreationGtpUUdpIpv4 ||
		ie.desc == OuterHeaderCreationGtpUUdpIpv6
}

func (ie *OuterHeaderCreation) IP() net.IP {
	return ie.ip
}

func (ie *OuterHeaderCreation) Desc() uint16 {
	return uint16(ie.desc)
}

func (ie *OuterHeaderCreation) Teid() uint32 {
	return ie.teid
}

func (ie *OuterHeaderCreation) Port() uint16 {
	return ie.port
}

func (ie *OuterHeaderCreation) Bytes() (wire []byte) {
	//default is GTP IPv4
	isGtp := true
	isV4 := true
	switch ie.desc {
	case OuterHeaderCreationGtpUUdpIpv4:
		isGtp, isV4 = true, true
	case OuterHeaderCreationGtpUUdpIpv6:
		isGtp, isV4 = true, false
	case OuterHeaderCreationUdpIpv4:
		isGtp, isV4 = false, true
	case OuterHeaderCreationUdpIpv6:
		isGtp, isV4 = false, false
	}

	//NOTE: Desc is encode in two bytes, right (least significant) is spared
	wire = []byte{ie.desc, 0}
	if isGtp {
		wire = append(wire, pfcp.MarshalUint32(ie.teid)...)
	}
	//NOTE: make sure to create right IP addr size, event for the case where
	//ie.ip is empty
	var ipBytes []byte
	if isV4 {
		ipBytes = make([]byte, net.IPv4len)
	} else {
		ipBytes = make([]byte, net.IPv6len)
	}
	copy(ipBytes, ie.ip)
	wire = append(wire, ipBytes...)
	if !isGtp { //udp
		wire = append(wire, pfcp.MarshalUint16(ie.port)...)
	}
	return
}
func (ie *OuterHeaderCreation) UnmarshalBinary(wire []byte) error {
	wireLen := len(wire)
	if wireLen < 2 {
		return tlv.ErrIncomplete
	}

	ie.desc = wire[0]
	isGtp := true
	isV4 := true
	switch ie.desc {
	case OuterHeaderCreationGtpUUdpIpv4:
		isGtp, isV4 = true, true
	case OuterHeaderCreationGtpUUdpIpv6:
		isGtp, isV4 = true, false
	case OuterHeaderCreationUdpIpv4:
		isGtp, isV4 = false, true
	case OuterHeaderCreationUdpIpv6:
		isGtp, isV4 = false, false
	default:
		return fmt.Errorf("Unknown outer header creation description")
	}
	offset := 2
	if isGtp {
		if offset+4 > wireLen {
			return tlv.ErrIncomplete
		}
		ie.teid = binary.BigEndian.Uint32(wire[offset : offset+4])
		offset += 4
	}
	ipLen := net.IPv6len
	if isV4 {
		ipLen = net.IPv4len
	}
	if offset+ipLen > wireLen {
		return tlv.ErrIncomplete
	}
	ie.ip = net.IP(make([]byte, ipLen))
	copy(ie.ip[:], wire[offset:offset+ipLen])
	offset += ipLen

	if !isGtp { //UDP
		if offset+2 > wireLen {
			return tlv.ErrIncomplete
		}
		ie.port = binary.BigEndian.Uint16(wire[offset : offset+2])
		offset += 2
	}

	//check there is no junk
	if offset < wireLen {
		return tlv.ErrTail
	}
	return nil
}
