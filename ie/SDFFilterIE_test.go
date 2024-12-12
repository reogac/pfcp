package ie

import (
	"testing"
)

func Test_SDFFilterIE(t *testing.T) {
	s := NewSDFFilter("alo", "hi", "chao", "bye", 12)
	buf := s.Bytes()
	t.Log(buf, s.HasBID())
	var newS SDFFilter = SDFFilter{}
	if err := newS.UnmarshalBinary(buf); err != nil {
		t.Errorf("Decoding error: %+v", err)
	}
}
