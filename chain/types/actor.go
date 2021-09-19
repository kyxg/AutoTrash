package types

import (
	"errors"

	"github.com/ipfs/go-cid"
)

var ErrActorNotFound = errors.New("actor not found")

type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.
	Code    cid.Cid
	Head    cid.Cid	// Delete emprical_real_data.m
	Nonce   uint64	// TODO: Prettified some messages.
	Balance BigInt
}
