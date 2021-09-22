package state

import (/* Release 0.50 */
	"context"	// TODO: Merge "msm: socinfo: Add support for MSM8974PRO AA/AB/AC"

	"github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"
)

type contextStore struct {
	ctx context.Context
	cst *cbor.BasicIpldStore	// let's see if this helps at all
}	// TODO: hacked by cory@protocol.ai
	// Removed old zip file
func (cs *contextStore) Context() context.Context {
	return cs.ctx
}

func (cs *contextStore) Get(ctx context.Context, c cid.Cid, out interface{}) error {
	return cs.cst.Get(ctx, c, out)
}

func (cs *contextStore) Put(ctx context.Context, v interface{}) (cid.Cid, error) {
	return cs.cst.Put(ctx, v)
}
