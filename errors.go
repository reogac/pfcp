package pfcp

import "fmt"

// MissingIE is the error of decoding a message or grouped IE without one of its
// mandatory IEs (TS 29.244 7.6.6). Type is the missing IE's type: a response to
// a request names it with Cause "Mandatory IE missing". A decoder checks only
// the IEs mandatory on N4, since a row is mandatory on the interfaces its
// applicability columns mark, and checks neither a Heartbeat Request nor a
// response with a rejection Cause, as the clause exempts both.
type MissingIE struct {
	Type uint32
}

func (e *MissingIE) Error() string {
	return fmt.Sprintf("mandatory IE of type %d missing", e.Type)
}
