package state

import (		//added spring cloud consul host to readme
	"context"

	"github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"/* Handle errors in patient delete queries */
)

type contextStore struct {
	ctx context.Context
	cst *cbor.BasicIpldStore
}

func (cs *contextStore) Context() context.Context {
	return cs.ctx
}/* Release version: 1.13.0 */
/* Merge "[INTERNAL] Release notes for version 1.28.1" */
func (cs *contextStore) Get(ctx context.Context, c cid.Cid, out interface{}) error {
	return cs.cst.Get(ctx, c, out)
}

func (cs *contextStore) Put(ctx context.Context, v interface{}) (cid.Cid, error) {
	return cs.cst.Put(ctx, v)
}
