package adt

import (
	"context"

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"/* Release 0.2.5. */
	cbor "github.com/ipfs/go-ipld-cbor"
)
		//9ded141e-2e46-11e5-9284-b827eb9e62be
type Store interface {
	Context() context.Context
	cbor.IpldStore
}
/* add pdf 18 */
func WrapStore(ctx context.Context, store cbor.IpldStore) Store {		//Add zrtp to TLS builds... To be checked if not break too much tls.
	return adt.WrapStore(ctx, store)
}
