package ie

import (
	"fmt"
	"net"
	"strings"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

const (
	NodeIdTypeIpv4Address uint8 = iota
	NodeIdTypeIpv6Address
	NodeIdTypeFqdn
)

type NodeID struct {
	nodeType uint8
	bytes    []byte
}

func NewNodeIDv4(ip net.IP) (ie NodeID) {
	ie.nodeType = NodeIdTypeIpv4Address
	ie.bytes = make([]byte, net.IPv4len)
	copy(ie.bytes, ip.To4())
	return
}

func NewNodeIDv6(ip net.IP) (ie NodeID) {
	ie.nodeType = NodeIdTypeIpv6Address
	ie.bytes = make([]byte, net.IPv6len)
	copy(ie.bytes, ip.To16())
	return
}

func (ie *NodeID) SetFqdnStr(fqdn string) (err error) {
	ie.nodeType = NodeIdTypeFqdn
	ie.bytes, err = FqdnBytes(fqdn, false)
	return
}

func (ie NodeID) IP() net.IP {
	if ie.nodeType == NodeIdTypeIpv4Address {
		return net.IP(ie.bytes).To4()
	} else if ie.nodeType == NodeIdTypeIpv6Address {
		return net.IP(ie.bytes).To16()
	}
	return net.IP([]byte{})
}

func (ie *NodeID) Fqdn() string {
	if ie.nodeType == NodeIdTypeFqdn {
		if fqdn, err := FqdnString(ie.bytes); err == nil {
			return fqdn
		}
	}
	return ""
}

func (ie *NodeID) GetType() uint8 {
	return ie.nodeType
}

func (ie *NodeID) Bytes() (wire []byte) {
	switch ie.nodeType {
	case NodeIdTypeIpv4Address:
		wire = make([]byte, 1+net.IPv4len)
	case NodeIdTypeIpv6Address:
		wire = make([]byte, 1+net.IPv6len)
	default:
		wire = make([]byte, 1+len(ie.bytes))
	}
	copy(wire[1:], ie.bytes)
	wire[0] = ie.nodeType
	return
}

func (ie *NodeID) UnmarshalBinary(wire []byte) (err error) {
	if len(wire) < 1 {
		return tlv.ErrIncomplete
	}
	switch wire[0] {
	case NodeIdTypeIpv4Address:
		if len(wire) < 1+net.IPv4len {
			return tlv.ErrIncomplete
		} else if len(wire) > 1+net.IPv4len {
			return tlv.ErrTail
		}
		ie.bytes = make([]byte, net.IPv4len)
		copy(ie.bytes, wire[1:])

	case NodeIdTypeIpv6Address:
		if len(wire) < 1+net.IPv6len {
			return tlv.ErrIncomplete
		} else if len(wire) > 1+net.IPv6len {
			return tlv.ErrTail
		}
		ie.bytes = make([]byte, net.IPv6len)
		copy(ie.bytes, wire[1:])

	case NodeIdTypeFqdn:
		ie.bytes = make([]byte, len(wire)-1)
		copy(ie.bytes, wire[1:])
		if _, err := FqdnString(ie.bytes); err != nil {
			return fmt.Errorf("Fail to decode fqdn: %+v", err)
		}
	default:
		return fmt.Errorf("Unknown NodeID type")
	}
	return nil
}

func FqdnBytes(fqdn string, isDnn bool) (wire []byte, err error) {
	parts := strings.Split(fqdn, ".")

	maxPart := 63
	maxLen := 255
	if isDnn {
		maxPart = 62
		maxLen = 100
	}

	for _, part := range parts {
		// In RFC 1035 max length is 63, but in TS 23.003
		// including length octet
		if len(part) > maxPart {
			err = fmt.Errorf("FQDN limit the label to %d octets or less", maxPart)
			return
		}
		wire = append(wire, uint8(len(part)))
		wire = append(wire, []byte(part)...)
	}

	if len(wire) > maxLen {
		err = fmt.Errorf("FQDN should less then %d octet", maxLen)
	}
	return
}

func FqdnString(wire []byte) (fqdn string, err error) {
	parts := []string{}
	offset := 0
	wireLen := len(wire)
	for offset < wireLen {
		partLen := int(wire[offset])
		offset++
		if offset+partLen > wireLen {
			err = tlv.ErrIncomplete
			return
		}
		parts = append(parts, string(wire[offset:offset+partLen]))
		offset += partLen
	}
	fqdn = strings.Join(parts, ".")
	return
}
