package ie

import (
	"bytes"
	"encoding/hex"
	"testing"
)

type valueCodec interface {
	UnmarshalBinary([]byte) error
	Bytes() []byte
}

// TestValuesMatchAnIndependentEncoder holds IE values as github.com/wmnsk/go-pfcp
// v0.0.24 encodes them, captured once rather than imported so this module takes
// no dependency for its tests. Each must decode, and encode back to the same
// octets: Apply Action, Gate Status, MBR, GBR, Outer Header Removal, the Node
// ID's type and the F-TEID's CHOOSE flags all disagreed with it before.
func TestValuesMatchAnIndependentEncoder(t *testing.T) {
	for _, c := range []struct {
		name  string
		value string
		codec func() valueCodec
	}{
		{"node id v4", "000a000001", func() valueCodec { return &NodeID{} }},
		{"node id fqdn", "020475706631076578616d706c65036f7267", func() valueCodec { return &NodeID{} }},
		{"node id v6", "0120010db8000000000000000000000001", func() valueCodec { return &NodeID{} }},
		{"f-teid v6", "020000007820010db8000000000000000000000002", func() valueCodec { return &FTEID{} }},
		{"f-teid v4v6", "03000000790a00000220010db8000000000000000000000002", func() valueCodec { return &FTEID{} }},
		{"f-teid choose v6", "06", func() valueCodec { return &FTEID{} }},
		{"ohr 2", "0101", func() valueCodec { return &OuterHeaderRemoval{} }},
		{"f-seid", "0200000000deadbeef0a000001", func() valueCodec { return &FSEID{} }},
		{"f-teid", "01000000770a000002", func() valueCodec { return &FTEID{} }},
		{"f-teid choose", "0d03", func() valueCodec { return &FTEID{} }},
		{"ue ip src", "020a2d0009", func() valueCodec { return &UEIPAddress{} }},
		{"ue ip dst", "060a2d0009", func() valueCodec { return &UEIPAddress{} }},
		{"apply action 1", "02", func() valueCodec { return &ApplyAction{} }},
		{"apply action 2", "0c02", func() valueCodec { return &ApplyAction{} }},
		{"gate open", "00", func() valueCodec { return &GateStatus{} }},
		{"gate ul closed", "04", func() valueCodec { return &GateStatus{} }},
		{"mbr", "00000003e800000007d0", func() valueCodec { return &MBR{} }},
		{"gbr", "000000012c0000000190", func() valueCodec { return &GBR{} }},
		{"ohc gtp v4", "0100000000070a000003", func() valueCodec { return &OuterHeaderCreation{} }},
		{"ohr 1", "0000", func() valueCodec { return &OuterHeaderRemoval{} }},
		{"sdf fd", "010000297065726d6974206f75742069702066726f6d2031302e302e302e302f3820746f2061737369676e6564", func() valueCodec { return &SDFFilter{} }},
		{"sdf fd+fid", "110000227065726d6974206f75742069702066726f6d20616e7920746f2061737369676e656400000007", func() valueCodec { return &SDFFilter{} }},
		{"measurement method", "03", func() valueCodec { return &MeasurementMethod{} }},
		{"measurement info", "02", func() valueCodec { return &MeasurementInformation{} }},
		{"volume threshold", "070000000000000bb800000000000003e800000000000007d0", func() valueCodec { return &VolumeThreshold{} }},
		{"volume quota", "010000000000001388", func() valueCodec { return &VolumeQuota{} }},
		{"volume measurement", "3f000000000000001e000000000000000a0000000000000014000000000000000300000000000000010000000000000002", func() valueCodec { return &VolumeMeasurement{} }},
		{"application id", "61707031", func() valueCodec { return &ApplicationID{} }},
		{"dl buffering count", "012c", func() valueCodec { return &DLBufferingSuggestedPacketCount{} }},
		{"fq-csid", "020a00000100010002", func() valueCodec { return &FQCSID{} }},
	} {
		value, _ := hex.DecodeString(c.value)
		codec := c.codec()
		if err := codec.UnmarshalBinary(value); err != nil {
			t.Errorf("%s: decode: %v", c.name, err)
			continue
		}
		if got := codec.Bytes(); !bytes.Equal(got, value) {
			t.Errorf("%s: encodes % x, want % x", c.name, got, value)
		}
	}
}
