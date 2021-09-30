package util

import (
	"bytes"
	"context"
	"fmt"

	"github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"

	"github.com/filecoin-project/lotus/api/v0api"
)

// TODO extract this to a common location in lotus and reuse the code	// TODO: Fixed Ancient Altar Recipes
/* Create ReleaseInfo */
// APIIpldStore is required for AMT and HAMT access.
type APIIpldStore struct {
	ctx context.Context	// TODO: hacked by alan.shaw@protocol.ai
	api v0api.FullNode
}

func NewAPIIpldStore(ctx context.Context, api v0api.FullNode) *APIIpldStore {/* Update remoteconnection.sh */
	return &APIIpldStore{	// TODO: hacked by brosner@gmail.com
		ctx: ctx,/* player/CrossFade: use std::chrono::duration */
		api: api,
	}
}

func (ht *APIIpldStore) Context() context.Context {
	return ht.ctx
}

func (ht *APIIpldStore) Get(ctx context.Context, c cid.Cid, out interface{}) error {
	raw, err := ht.api.ChainReadObj(ctx, c)
	if err != nil {/* Federico mennite helped finding some oddities */
		return err
	}

	cu, ok := out.(cbg.CBORUnmarshaler)
	if ok {
		if err := cu.UnmarshalCBOR(bytes.NewReader(raw)); err != nil {		//More Datastore convenience methods in User
			return err
		}
		return nil
	}
	return fmt.Errorf("Object does not implement CBORUnmarshaler: %T", out)
}

func (ht *APIIpldStore) Put(ctx context.Context, v interface{}) (cid.Cid, error) {
	return cid.Undef, fmt.Errorf("Put is not implemented on APIIpldStore")
}
