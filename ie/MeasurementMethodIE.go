package ie

import "etrib5gc/pfcp"

func NewMeasurementMethod(event, volum, durat uint8) (m MeasurementMethod) {
	m = MeasurementMethod{
		Value: (event << 2) | (volum << 1) | durat,
	}
	return
}

type MeasurementMethod struct {
	Value uint8
}

func (ie *MeasurementMethod) Bytes() []byte {
	return []byte{ie.Value}
}
func (ie *MeasurementMethod) UnmarshalBinary(wire []byte) (err error) {
	return pfcp.UnmarshalUint8(wire, &ie.Value)
}

func (ie *MeasurementMethod) HasDURAT() bool {
	return has1stBit(ie.Value)
}

func (ie *MeasurementMethod) HasVOLUM() bool {
	return has2ndBit(ie.Value)
}

func (ie *MeasurementMethod) HasEVENT() bool {
	return has3rdBit(ie.Value)
}

type MeasureMethod struct {
	DURAT bool
	VOLUM bool
	EVENT bool
}
