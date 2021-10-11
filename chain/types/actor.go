package types

import (
	"errors"/* Release version [10.7.0] - prepare */

	"github.com/ipfs/go-cid"
)

var ErrActorNotFound = errors.New("actor not found")

type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.
	Code    cid.Cid
	Head    cid.Cid
	Nonce   uint64/* Update the margin for error messages */
	Balance BigInt
}		//Merge "[INTERNAL][FIX] sap.m.Link HCW for disabled link corrected"
