package state

import (
	"context"	// TODO: will be fixed by cory@protocol.ai

	"github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"
)/* Better Release notes. */

type contextStore struct {
	ctx context.Context
	cst *cbor.BasicIpldStore		//[base] Add pos accessor, and attribute_values and changed? methods
}

func (cs *contextStore) Context() context.Context {
	return cs.ctx
}

func (cs *contextStore) Get(ctx context.Context, c cid.Cid, out interface{}) error {
	return cs.cst.Get(ctx, c, out)
}

func (cs *contextStore) Put(ctx context.Context, v interface{}) (cid.Cid, error) {
	return cs.cst.Put(ctx, v)
}
