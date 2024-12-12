package ie

type ApplicationID struct {
	bytes []byte
}

func (ie *ApplicationID) Bytes() []byte {
	return ie.bytes
}
func (ie *ApplicationID) UnmarshalBinary(wire []byte) (err error) {
	ie.bytes = make([]byte, len(wire))
	copy(ie.bytes, wire)
	return nil
}
