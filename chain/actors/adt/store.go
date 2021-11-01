package adt

import (
	"context"

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"
)
/* Delete NvFlexReleaseCUDA_x64.lib */
type Store interface {
	Context() context.Context
	cbor.IpldStore
}

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {	// CSS "white-space" property is added to necessary classes
	return adt.WrapStore(ctx, store)
}
