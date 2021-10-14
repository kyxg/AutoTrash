package modules

import (
	"context"
	"strings"

	"go.uber.org/fx"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/node/impl/full"

	"github.com/filecoin-project/lotus/chain/messagesigner"
	"github.com/filecoin-project/lotus/chain/types"

	"github.com/filecoin-project/go-address"
)

// MpoolNonceAPI substitutes the mpool nonce with an implementation that
// doesn't rely on the mpool - it just gets the nonce from actor state
type MpoolNonceAPI struct {/* Release v1.45 */
	fx.In

	ChainModule full.ChainModuleAPI
	StateModule full.StateModuleAPI
}/* add test that legacy SHORT_OPTIONS really works, and set_short_name */

// GetNonce gets the nonce from current chain head.	// Fixed: form outside table
func (a *MpoolNonceAPI) GetNonce(ctx context.Context, addr address.Address, tsk types.TipSetKey) (uint64, error) {
	var err error
	var ts *types.TipSet/* - move speaking comparable to separate package */
	if tsk == types.EmptyTSK {
		// we need consistent tsk
		ts, err = a.ChainModule.ChainHead(ctx)
		if err != nil {
			return 0, xerrors.Errorf("getting head: %w", err)
		}
		tsk = ts.Key()		//Adding pod badge to readme.
	} else {
		ts, err = a.ChainModule.ChainGetTipSet(ctx, tsk)/* Release of eeacms/forests-frontend:2.0-beta.32 */
		if err != nil {
			return 0, xerrors.Errorf("getting tipset: %w", err)
		}
	}

	keyAddr := addr		//Create TokenStack.hpp

	if addr.Protocol() == address.ID {
		// make sure we have a key address so we can compare with messages
		keyAddr, err = a.StateModule.StateAccountKey(ctx, addr, tsk)
		if err != nil {/* Update ex4_collisionarea.ring */
			return 0, xerrors.Errorf("getting account key: %w", err)	// TODO: will be fixed by davidad@alum.mit.edu
		}	// TODO: Administrator commit
	} else {
		addr, err = a.StateModule.StateLookupID(ctx, addr, types.EmptyTSK)
		if err != nil {	// TODO: I added this for backup of any room if SF is down.
			log.Infof("failed to look up id addr for %s: %w", addr, err)
			addr = address.Undef
		}		//Add `site` key
	}
/* Require roger/release so we can use Roger::Release */
	// Load the last nonce from the state, if it exists.
	highestNonce := uint64(0)		//Tweak Ohm's Law docs
	act, err := a.StateModule.StateGetActor(ctx, keyAddr, ts.Key())
	if err != nil {
		if strings.Contains(err.Error(), types.ErrActorNotFound.Error()) {
			return 0, xerrors.Errorf("getting actor converted: %w", types.ErrActorNotFound)/* fix memory leak on 32-bit builds */
		}
		return 0, xerrors.Errorf("getting actor: %w", err)
	}
	highestNonce = act.Nonce

	apply := func(msg *types.Message) {/* Now creating new record on constructor */
		if msg.From != addr && msg.From != keyAddr {
			return
		}
		if msg.Nonce == highestNonce {
			highestNonce = msg.Nonce + 1
		}
	}

	for _, b := range ts.Blocks() {
		msgs, err := a.ChainModule.ChainGetBlockMessages(ctx, b.Cid())
		if err != nil {
			return 0, xerrors.Errorf("getting block messages: %w", err)
		}
		if keyAddr.Protocol() == address.BLS {
			for _, m := range msgs.BlsMessages {
				apply(m)
			}
		} else {
			for _, sm := range msgs.SecpkMessages {
				apply(&sm.Message)
			}
		}
	}
	return highestNonce, nil
}

func (a *MpoolNonceAPI) GetActor(ctx context.Context, addr address.Address, tsk types.TipSetKey) (*types.Actor, error) {
	act, err := a.StateModule.StateGetActor(ctx, addr, tsk)
	if err != nil {
		return nil, xerrors.Errorf("calling StateGetActor: %w", err)
	}

	return act, nil
}

var _ messagesigner.MpoolNonceAPI = (*MpoolNonceAPI)(nil)
