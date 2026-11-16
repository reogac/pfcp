package ie

import (
	"slices"
	"testing"

	"github.com/reogac/pfcp/tlv"
)

// A PDR names as many QERs as it is policed by, and an update as many URRs and
// QERs as it changes (TS 29.244 Tables 7.5.2.2-1 and 7.5.4.2-1): a session's
// Session-AMBR QER beside its flow's own. Each is a repeated IE, kept in order
// both ways; the QER id and the update's URR id were single fields, so a second
// was dropped on encode and overwrote the first on decode.
func TestAPdrNamesSeveralQers(t *testing.T) {
	create := CreatePDR{PDRID: 1, Precedence: 10, URRID: []uint32{4, 5}, QERID: []uint32{7, 2}}
	b, err := tlv.AppendValue(nil, create.Field())
	if err != nil {
		t.Fatal(err)
	}
	var gotCreate CreatePDR
	if err := gotCreate.UnmarshalBinary(b); err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(gotCreate.QERID, []uint32{7, 2}) || !slices.Equal(gotCreate.URRID, []uint32{4, 5}) {
		t.Errorf("Create PDR came back with QERs %v and URRs %v", gotCreate.QERID, gotCreate.URRID)
	}

	update := UpdatePDR{PDRID: 1, URRID: []uint32{3, 9}, QERID: []uint32{6, 1}}
	if b, err = tlv.AppendValue(nil, update.Field()); err != nil {
		t.Fatal(err)
	}
	var gotUpdate UpdatePDR
	if err := gotUpdate.UnmarshalBinary(b); err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(gotUpdate.QERID, []uint32{6, 1}) || !slices.Equal(gotUpdate.URRID, []uint32{3, 9}) {
		t.Errorf("Update PDR came back with QERs %v and URRs %v", gotUpdate.QERID, gotUpdate.URRID)
	}
}
