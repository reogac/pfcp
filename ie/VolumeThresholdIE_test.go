package ie

import "testing"

func Test_VolumeThreshold(t *testing.T) {
	v := *NewVolumeThreshold(1, 2, 3, 4)
	buf := v.Bytes()
	var newS VolumeThreshold = VolumeThreshold{}
	if err := newS.UnmarshalBinary(buf); err != nil {
		t.Errorf("Decoding error: %+v", err)
	}
	if newS.HasTOVOL() && v.TotalVolume != 2 {
		t.Errorf("Decoding error TotalVolume: %+v", v.TotalVolume)
	}
}
