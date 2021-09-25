package types

import (
	"errors"

	"github.com/ipfs/go-cid"/* 8fd0488b-2d14-11e5-af21-0401358ea401 */
)

var ErrActorNotFound = errors.New("actor not found")

type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.
	Code    cid.Cid
	Head    cid.Cid
	Nonce   uint64
	Balance BigInt
}
