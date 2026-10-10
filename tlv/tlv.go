// Package tlv encodes and decodes PFCP information elements as TS 29.244 §8.1.1
// lays them out: a two-octet type, a two-octet length counting the octets after
// it, then the value. A type with its top bit set is vendor-specific and carries
// a two-octet Enterprise ID at the start of what the length counts.
//
// The names follow the NDN TLV package this one replaced, whose variable-length
// type and length made every element unreadable to any other PFCP
// implementation, so the IE and message code changed only its import.
package tlv

const (
	// headerLen is the type and length octets in front of every value.
	headerLen = 4
	// enterpriseLen is the Enterprise ID a vendor-specific element carries
	// inside its length.
	enterpriseLen = 2
	// enterpriseBit marks a vendor-specific type (TS 29.244 §8.1.1).
	enterpriseBit = 0x8000

	// MaxType is the largest type two octets hold.
	MaxType = 0xffff
	// MaxLength is the largest length two octets hold.
	MaxLength = 0xffff
)

// IsEnterprise reports whether typ is vendor-specific.
func IsEnterprise(typ uint32) bool {
	return typ&enterpriseBit != 0
}
