package ie

import "testing"

// Bytes used to build the encoding and return nil, so every usage report
// carried an empty Volume Measurement. Every combination of present values
// must come back as it went in.
func TestVolumeMeasurementRoundTrips(t *testing.T) {
	for flags := uint8(1); flags <= 0x3f; flags++ {
		in := VolumeMeasurement{Flags: flags, TotalVolume: 1, UplinkVolume: 2, DownlinkVolume: 3,
			TotalNumberOfPackets: 4, UplinkNumberOfPackets: 5, DownlinkNumberOfPackets: 6}
		wire := in.Bytes()
		if len(wire) != in.bytesLen() {
			t.Fatalf("flags %#x: %d bytes, want %d", flags, len(wire), in.bytesLen())
		}
		var out VolumeMeasurement
		if err := out.UnmarshalBinary(wire); err != nil {
			t.Fatalf("flags %#x: %v", flags, err)
		}
		want := VolumeMeasurement{Flags: flags}
		for i, p := range []*uint64{&want.TotalVolume, &want.UplinkVolume, &want.DownlinkVolume,
			&want.TotalNumberOfPackets, &want.UplinkNumberOfPackets, &want.DownlinkNumberOfPackets} {
			if flags&(1<<i) != 0 {
				*p = uint64(i + 1)
			}
		}
		if out != want {
			t.Errorf("flags %#x: %+v, want %+v", flags, out, want)
		}
	}
}
