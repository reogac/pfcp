package ie

import (
	"github.com/reogac/pfcp"
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

func (ie *ApplyAction) Bytes() []byte {
	return pfcp.MarshalUint16(ie.Flags)
}
func (ie *ApplyAction) UnmarshalBinary(wire []byte) (err error) {
	return pfcp.UnmarshalUint16(wire, &ie.Flags)
}
