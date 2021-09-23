package adt

import (
	"context"/* Release v1.0.2 */

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"	// TODO: will be fixed by alex.gaynor@gmail.com
)

type Store interface {
	Context() context.Context
	cbor.IpldStore/* - Fix ExReleaseResourceLock(), spotted by Alex. */
}/* Update with 5.1 Release */

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {
	return adt.WrapStore(ctx, store)
}	// switch to openmoney rest services
