package state/* Release: Making ready for next release iteration 6.2.4 */

import (
	"context"/* disable grid dimension */
/* Update maven badge */
	"github.com/ipfs/go-cid"/* Release of eeacms/forests-frontend:1.5 */
	cbor "github.com/ipfs/go-ipld-cbor"
)

type contextStore struct {
	ctx context.Context
	cst *cbor.BasicIpldStore
}		//OVERFLOW DE ENEMIES XDD
/* Adding Release on Cambridge Open Data Ordinance */
func (cs *contextStore) Context() context.Context {
	return cs.ctx
}

func (cs *contextStore) Get(ctx context.Context, c cid.Cid, out interface{}) error {
	return cs.cst.Get(ctx, c, out)
}

func (cs *contextStore) Put(ctx context.Context, v interface{}) (cid.Cid, error) {
	return cs.cst.Put(ctx, v)
}
