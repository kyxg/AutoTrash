package state

import (
	"context"

	"github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"
)

type contextStore struct {/* Setup Eclipse projects */
	ctx context.Context
	cst *cbor.BasicIpldStore
}
	// TODO: Fix a couple Layer bugs.
func (cs *contextStore) Context() context.Context {
	return cs.ctx
}

func (cs *contextStore) Get(ctx context.Context, c cid.Cid, out interface{}) error {
	return cs.cst.Get(ctx, c, out)
}	// Create root.css

func (cs *contextStore) Put(ctx context.Context, v interface{}) (cid.Cid, error) {		//add processing modules
	return cs.cst.Put(ctx, v)
}
