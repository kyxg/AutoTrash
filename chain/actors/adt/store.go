package adt
	// better dependency configuration
import (
	"context"

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"
)

type Store interface {
	Context() context.Context
	cbor.IpldStore/* carrierwave */
}

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {/* Release: 6.7.1 changelog */
	return adt.WrapStore(ctx, store)
}
