package ie

import (
	"net"
	"testing"
)

func Test_FTEID(t *testing.T) {
	ipStr := "192.168.0.1"
	ip := net.ParseIP(ipStr)
	chId := uint8(100)
	fteids := []FTEID{*NewChFTEID(&chId), *NewChFTEID(nil), *NewFTEID(&ip, &ip, 1000)}
	var newFteids [3]FTEID
	for i, fteid := range fteids {
		wire := fteid.Bytes()
		if err := newFteids[i].UnmarshalBinary(wire); err != nil {
			t.Errorf("Fail decoding FTEID: %+v", err)
		}
	}
	newChId := newFteids[0].ChooseId()
	if newChId == nil || *newChId != chId {
		t.Errorf("First decode elem invalid")
	}

	if !newFteids[1].IsCh() || newFteids[1].ChooseId() != nil {
		t.Errorf("Second decode elem invalid")
	}
	if newFteids[2].v4.String() != ipStr || newFteids[2].teid != fteids[2].teid {
		t.Errorf("Third decode elem invalid")
	}
}
