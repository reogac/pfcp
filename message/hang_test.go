package message

import (
	"testing"

	"github.com/reogac/pfcp/tlv"
)

// TestRepeatedBytesFieldsEncodeOnce pins the fields whose encoder was a
// `for len(field) > 0` loop, which never ended once the field was set.
func TestRepeatedBytesFieldsEncodeOnce(t *testing.T) {
	for name, m := range map[string]tlv.Fielder{
		"setup request":   &PFCPAssociationSetupRequest{UserPlaneIPResourceInformation: []byte{1}},
		"setup response":  &PFCPAssociationSetupResponse{UserPlaneIPResourceInformation: []byte{1}},
		"update request":  &PFCPAssociationUpdateRequest{UserPlaneIPResourceInformation: []byte{1}},
		"update response": &PFCPAssociationUpdateResponse{UserPlaneIPResourceInformation: []byte{1}},
	} {
		if _, err := tlv.EncodeValueOnly(m); err != nil {
			t.Errorf("%s: %v", name, err)
		}
	}
}
