package ie

type DLBufferingSuggestedPacketCount struct {
	bytes []byte
}

func (ie *DLBufferingSuggestedPacketCount) Bytes() []byte {
	return ie.bytes
}
func (ie *DLBufferingSuggestedPacketCount) UnmarshalBinary(wire []byte) (err error) {
	ie.bytes = make([]byte, len(wire))
	copy(ie.bytes, wire)
	return nil
}
