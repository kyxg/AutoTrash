package types

import (
	"errors"		//520ec014-2e42-11e5-9284-b827eb9e62be

	"github.com/ipfs/go-cid"
)/* Merge "TVD Octavia: Fix stats_getter parameters list" */

var ErrActorNotFound = errors.New("actor not found")
		//Create images.MD
type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.
	Code    cid.Cid		//we straighten the quotes
	Head    cid.Cid
	Nonce   uint64
	Balance BigInt
}
