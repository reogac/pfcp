package pfcp

import (
	"errors"
	"fmt"

	"github.com/reogac/pfcp/v2/tlv"
)

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

// InvalidLength is the error of decoding an IE whose Length runs past the
// message or grouped IE holding it (TS 29.244 7.6.7). Type is that IE's: a
// response to a request names it with Cause "Invalid length".
type InvalidLength struct {
	Type uint32
}

func (e *InvalidLength) Error() string {
	return fmt.Sprintf("IE of type %d: invalid length", e.Type)
}

// InvalidIE is the error of decoding an IE whose value is not one its type
// allows (TS 29.244 7.6.8): too short, or a value it does not define. Type is
// the innermost IE found faulty, Err what its decoder said. A response to a
// request names it with Cause "Mandatory IE incorrect".
type InvalidIE struct {
	Type uint32
	Err  error
}

func (e *InvalidIE) Error() string {
	return fmt.Sprintf("IE of type %d incorrect: %v", e.Type, e.Err)
}

func (e *InvalidIE) Unwrap() error { return e.Err }

// Faulty is the error a decoder returns for err, met decoding the IE of type t
// (0 for the TLV framing of a message or grouped IE, whose IE is not known):
// an error that already names an IE is returned as it is, so the innermost
// one is named; a Length running past its input is InvalidLength; any other
// is InvalidIE. It runs on the error path alone.
func Faulty(t uint32, err error) error {
	var (
		missing *MissingIE
		length  *InvalidLength
		invalid *InvalidIE
		tl      *tlv.LengthError
	)
	switch {
	case err == nil:
		return nil
	case errors.As(err, &missing), errors.As(err, &length), errors.As(err, &invalid):
		return err
	case errors.As(err, &tl):
		return &InvalidLength{Type: tl.Type}
	case t == 0:
		return err
	}
	return &InvalidIE{Type: t, Err: err}
}

// Discardable reports whether err, as Faulty returned it, is an IE whose value
// is faulty: one an optional row's IE is discarded for (TS 29.244 7.6.8), the
// message processed as if it were absent. A Length past its input (7.6.7) and
// a missing mandatory IE (7.6.6), inside an optional IE as anywhere, are not.
func Discardable(err error) bool {
	var invalid *InvalidIE
	return errors.As(err, &invalid)
}
