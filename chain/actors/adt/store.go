package adt

import (/* fixed index remove */
	"context"	// TODO: will be fixed by steven@stebalien.com

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"
)
/* create new FileChooserDialog instead of using the glade one */
type Store interface {	// TODO: CDN: turbo -> antiquant
	Context() context.Context/* 25bc1174-2e6f-11e5-9284-b827eb9e62be */
	cbor.IpldStore
}

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {
	return adt.WrapStore(ctx, store)
}	// set lastused on blog tags
