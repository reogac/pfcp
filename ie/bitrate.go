package ie

// bitrateLen is one direction of an MBR or a GBR: a 40-bit kbps value.
const bitrateLen = 5

// maxBitrate is the largest kbps five octets hold.
const maxBitrate = 1<<40 - 1

// appendBitrate writes kbps as five octets, clamped to what they hold: a rate
// that does not fit would otherwise lose its top bits and become a small one.
func appendBitrate(b []byte, kbps uint64) []byte {
	if kbps > maxBitrate {
		kbps = maxBitrate
	}
	return append(b, byte(kbps>>32), byte(kbps>>24), byte(kbps>>16), byte(kbps>>8), byte(kbps))
}

func bitrate(wire []byte) uint64 {
	return uint64(wire[0])<<32 | uint64(wire[1])<<24 | uint64(wire[2])<<16 | uint64(wire[3])<<8 | uint64(wire[4])
}
