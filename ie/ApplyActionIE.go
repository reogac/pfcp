package ie

import (
	"github.com/reogac/pfcp/tlv"
)

// Apply Action IE bits definition
//
//	in code of go-pfcp: DDPN, BDPN, and EDRT are set to "0"
//	in code of smf-free5gc: only FORW is enabled
const (
	APPLY_ACT_DROP = 1 << iota
	APPLY_ACT_FORW
	APPLY_ACT_BUFF
	APPLY_ACT_NOCP
	APPLY_ACT_DUPL
	APPLY_ACT_IPMA
	APPLY_ACT_IPMD
	APPLY_ACT_DFRT
	APPLY_ACT_EDRT
	APPLY_ACT_BDPN
	APPLY_ACT_DDPN
	APPLY_ACT_FSSM
	APPLY_ACT_MBSU
)

func NewApplyAction(f uint16) *ApplyAction {
	return &ApplyAction{
		Flags: f,
	}
}

type ApplyAction struct {
	Flags uint16
}

func (a *ApplyAction) DROP() bool {
	return a.Flags&APPLY_ACT_DROP != 0
}

func (a *ApplyAction) FORW() bool {
	return a.Flags&APPLY_ACT_FORW != 0
}

func (a *ApplyAction) BUFF() bool {
	return a.Flags&APPLY_ACT_BUFF != 0
}

func (a *ApplyAction) NOCP() bool {
	return a.Flags&APPLY_ACT_NOCP != 0
}

func (a *ApplyAction) DUPL() bool {
	return a.Flags&APPLY_ACT_DUPL != 0
}

func (a *ApplyAction) IPMA() bool {
	return a.Flags&APPLY_ACT_IPMA != 0
}

func (a *ApplyAction) IPMD() bool {
	return a.Flags&APPLY_ACT_IPMD != 0
}

func (a *ApplyAction) DFRT() bool {
	return a.Flags&APPLY_ACT_DFRT != 0
}

func (a *ApplyAction) EDRT() bool {
	return a.Flags&APPLY_ACT_EDRT != 0
}

func (a *ApplyAction) BDPN() bool {
	return a.Flags&APPLY_ACT_BDPN != 0
}

func (a *ApplyAction) DDPN() bool {
	return a.Flags&APPLY_ACT_DDPN != 0
}

func (a *ApplyAction) FSSM() bool {
	return a.Flags&APPLY_ACT_FSSM != 0
}

func (a *ApplyAction) MBSU() bool {
	return a.Flags&APPLY_ACT_MBSU != 0
}

// Bytes writes Flags in the spec's octet order (TS 29.244 §8.2.26): the low
// byte is octet 5 (DROP to DFRT) and the high byte octet 6 (EDRT to MBSU), so
// the value is not a big-endian uint16. Octet 6 is left out while it is zero,
// which is the one-octet form a peer from before it existed expects.
func (ie *ApplyAction) Bytes() []byte {
	if ie.Flags>>8 == 0 {
		return []byte{byte(ie.Flags)}
	}
	return []byte{byte(ie.Flags), byte(ie.Flags >> 8)}
}

// UnmarshalBinary takes octet 5 alone or with octet 6, and ignores any later
// octet a newer release adds.
func (ie *ApplyAction) UnmarshalBinary(wire []byte) (err error) {
	switch len(wire) {
	case 0:
		return tlv.ErrIncomplete
	case 1:
		ie.Flags = uint16(wire[0])
	default:
		ie.Flags = uint16(wire[0]) | uint16(wire[1])<<8
	}
	return nil
}
