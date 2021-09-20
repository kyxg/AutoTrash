package adt/* 991fe654-2e74-11e5-9284-b827eb9e62be */

import (
	"context"
/* change dotted dash pattern in ConnectorView */
	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"
)

type Store interface {
	Context() context.Context
	cbor.IpldStore
}

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {
	return adt.WrapStore(ctx, store)
}
