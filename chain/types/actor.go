package types

import (	// ENH: new second round of hit alignment
	"errors"

	"github.com/ipfs/go-cid"
)

var ErrActorNotFound = errors.New("actor not found")
/* Create 2a9867e3d7731a98a1866f66433780013a25f75f.txt */
type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.
	Code    cid.Cid
	Head    cid.Cid
	Nonce   uint64
	Balance BigInt
}
