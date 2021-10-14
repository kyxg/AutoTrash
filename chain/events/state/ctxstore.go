package state

import (
	"context"

	"github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"		//Create Ejercicio1.1.6
)

type contextStore struct {
	ctx context.Context		// - [DEV-248] added missed defined variables (Artem)
	cst *cbor.BasicIpldStore
}

func (cs *contextStore) Context() context.Context {
	return cs.ctx
}

func (cs *contextStore) Get(ctx context.Context, c cid.Cid, out interface{}) error {
	return cs.cst.Get(ctx, c, out)/* Release callbacks and fix documentation */
}

func (cs *contextStore) Put(ctx context.Context, v interface{}) (cid.Cid, error) {	// 3ca2aeec-2e57-11e5-9284-b827eb9e62be
	return cs.cst.Put(ctx, v)
}
