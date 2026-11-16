// Package samples holds a value of every grouped IE and message with every
// field set, for the round-trip tests: samples.go is generated from the same
// tables as the types, and this file holds the leaf IEs whose codec is written
// by hand, which a generator cannot fill. Each leaf is in the canonical form its
// codec writes back, so a round trip compares octets.
package samples

import (
	"net"

	"github.com/reogac/pfcp/ie"
)

func ptr[T any](v T) *T { return &v }

var (
	v4 = net.IPv4(10, 0, 0, 1).To4()
)

func FTEID() ie.FTEID { return *ie.NewFTEID(&v4, nil, 0x1234) }

func FSEID() ie.FSEID { return *ie.NewFSEID(&v4, nil, 0xabcdef) }

func NodeID() ie.NodeID { return ie.NewNodeIDv4(v4) }

func UEIPAddress() ie.UEIPAddress { return *ie.NewUEIPAddress(&v4, nil) }

func SDFFilter() ie.SDFFilter {
	return *ie.NewSDFFilter("permit out ip from any to assigned", "", "", "", 0)
}

func ApplicationID() ie.ApplicationID {
	var a ie.ApplicationID
	_ = a.UnmarshalBinary([]byte("app"))
	return a
}

// FQCSID is an IPv4 node with one CSID (TS 29.244 8.2.46)
func FQCSID() ie.FQCSID {
	var f ie.FQCSID
	_ = f.UnmarshalBinary([]byte{0x01, 10, 0, 0, 3, 0, 1})
	return f
}

func GateStatus() ie.GateStatus { return *ie.NewGateStatus(0, 1) }

func ApplyAction() ie.ApplyAction { return *ie.NewApplyAction(ie.APPLY_ACT_FORW) }

func MBR() ie.MBR { return *ie.NewMBR(1000, 2000) }

func GBR() ie.GBR { return *ie.NewGBR(100, 200) }

func MeasurementMethod() ie.MeasurementMethod { return ie.NewMeasurementMethod(0, 1, 1) }

func MeasurementInformation() ie.MeasurementInformation { return *ie.NewMeasurementInformation(1) }

func OuterHeaderCreation() ie.OuterHeaderCreation {
	return *ie.NewOuterHeaderCreationGtpIpv4(v4, 0x55)
}

func OuterHeaderRemoval() ie.OuterHeaderRemoval { return *ie.NewOuterHeaderRemoval(0) }

func VolumeThreshold() ie.VolumeThreshold { return *ie.NewVolumeThreshold(7, 3, 1, 2) }

func VolumeQuota() ie.VolumeQuota { return *ie.NewVolumeQuota(7, 3, 1, 2) }
