package ie

import "testing"

func Test_MeasurementInformation(t *testing.T) {
	v := *NewMeasurementInformation(1)
	buf := v.Bytes()
	var newS MeasurementInformation = MeasurementInformation{}
	if err := newS.UnmarshalBinary(buf); err != nil {
		t.Errorf("Decoding error: %+v", err)
	}
	if newS.HasMBQE() && !v.Mbqe {
		t.Errorf("Decoding error TotalVolume: %+v", v.Mbqe)
	}
}
