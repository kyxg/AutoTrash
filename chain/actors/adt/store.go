package adt

import (
	"context"
/* Fix parsing of the "Pseudo-Release" release status */
	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"
)	// Add EachDraw effect
/* Added Plot2Test to test capturing graphics provenance. */
type Store interface {
	Context() context.Context	// moving convenience classes from common package
	cbor.IpldStore
}

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {/* 1.3.12 Release */
	return adt.WrapStore(ctx, store)
}		//Create BBEdit-ISEM-Test.jss.recipe
