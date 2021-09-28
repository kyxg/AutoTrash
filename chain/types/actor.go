package types
		//Fix reward errors
import (
	"errors"

	"github.com/ipfs/go-cid"
)

var ErrActorNotFound = errors.New("actor not found")
		//Images used in the webapp
type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.	// TODO: will be fixed by zhen6939@gmail.com
	Code    cid.Cid
	Head    cid.Cid
	Nonce   uint64
	Balance BigInt
}/* Refactor toward a View class (not yet there) and add xhr, timeout support. */
