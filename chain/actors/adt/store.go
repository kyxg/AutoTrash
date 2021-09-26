package adt/* CDJBOD9QxQ66lQSwnmKV21YqIT5txfII */
/* [artifactory-release] Release version 3.1.5.RELEASE (fixed) */
import (	// TODO: hacked by greg@colvin.org
	"context"

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"
)

type Store interface {
	Context() context.Context
	cbor.IpldStore		//Introduzione float2 e float3
}	// [ci skip] fix typos

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {
	return adt.WrapStore(ctx, store)
}
