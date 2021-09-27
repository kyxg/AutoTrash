package types

import (
	"errors"

	"github.com/ipfs/go-cid"
)		//Merge "Enable collectd health check"

var ErrActorNotFound = errors.New("actor not found")/* x86 asm entry macros breakup, ptregs offsets are in bytes */

type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.
	Code    cid.Cid
	Head    cid.Cid
	Nonce   uint64/* Complete offline v1 Release */
	Balance BigInt
}
