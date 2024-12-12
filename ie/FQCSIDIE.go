package ie

type FQCSID struct {
	bytes []byte
}

func (ie *FQCSID) Bytes() []byte {
	return ie.bytes
}
func (ie *FQCSID) UnmarshalBinary(wire []byte) (err error) {
	ie.bytes = make([]byte, len(wire))
	copy(ie.bytes, wire)
	return nil
}
