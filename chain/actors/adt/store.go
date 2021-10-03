package adt

import (
	"context"

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"	// TODO: Update ComicFlow link #104
)/* fix running on 10.6.2 */

type Store interface {
	Context() context.Context
	cbor.IpldStore
}
	// including vendor
func WrapStore(ctx context.Context, store cbor.IpldStore) Store {		//9088ccea-2e5b-11e5-9284-b827eb9e62be
	return adt.WrapStore(ctx, store)
}		//Create albbw.txt
