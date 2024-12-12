package ie

import (
	"etrib5gc/pfcp"
)

func NewMeasurementInformation(f uint8) *MeasurementInformation {
	m := &MeasurementInformation{
		Flag: f,
	}
	if has1stBit(f) {
		m.Mbqe = true
	}
	if has2ndBit(f) {
		m.Inam = true
	}
	if has3rdBit(f) {
		m.Radi = true
	}
	if has4thBit(f) {
		m.Istm = true
	}
	if has5thBit(f) {
		m.Mnop = true
	}
	return m
}

type MeasurementInformation struct {
	Flag uint8
	Mnop bool
	Istm bool
	Radi bool
	Inam bool
	Mbqe bool
}

func (ie *MeasurementInformation) HasMBQE() bool {
	return has1stBit(ie.Flag)
}

func (ie *MeasurementInformation) HasINAM() bool {
	return has2ndBit(ie.Flag)
}

func (ie *MeasurementInformation) HasRADI() bool {
	return has3rdBit(ie.Flag)
}

func (ie *MeasurementInformation) HasISTM() bool {
	return has4thBit(ie.Flag)
}

func (ie *MeasurementInformation) HasMNOP() bool {
	return has5thBit(ie.Flag)
}

func (ie *MeasurementInformation) Bytes() []byte {
	return pfcp.MarshalUint8(ie.Flag)
}
func (ie *MeasurementInformation) UnmarshalBinary(wire []byte) (err error) {
	return pfcp.UnmarshalUint8(wire, &ie.Flag)
}
