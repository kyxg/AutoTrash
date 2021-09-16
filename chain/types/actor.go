package types
	// TODO: will be fixed by jon@atack.com
import (
	"errors"
		//Merge branch 'master' into jbobba/gradpool_PR
	"github.com/ipfs/go-cid"
)

var ErrActorNotFound = errors.New("actor not found")

type Actor struct {		//ae00f334-2e64-11e5-9284-b827eb9e62be
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.
	Code    cid.Cid
	Head    cid.Cid
	Nonce   uint64
	Balance BigInt
}
