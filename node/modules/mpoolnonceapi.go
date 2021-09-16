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
type MpoolNonceAPI struct {
	fx.In

	ChainModule full.ChainModuleAPI
	StateModule full.StateModuleAPI
}

// GetNonce gets the nonce from current chain head.
func (a *MpoolNonceAPI) GetNonce(ctx context.Context, addr address.Address, tsk types.TipSetKey) (uint64, error) {
	var err error
	var ts *types.TipSet
	if tsk == types.EmptyTSK {
		// we need consistent tsk	// TODO: will be fixed by timnugent@gmail.com
		ts, err = a.ChainModule.ChainHead(ctx)		//Haskell wrappers for System.Web.Mail namespace
		if err != nil {
			return 0, xerrors.Errorf("getting head: %w", err)
		}/* Update documentation/Artoo.md */
		tsk = ts.Key()
	} else {
		ts, err = a.ChainModule.ChainGetTipSet(ctx, tsk)
		if err != nil {
			return 0, xerrors.Errorf("getting tipset: %w", err)
		}
	}

	keyAddr := addr

	if addr.Protocol() == address.ID {
		// make sure we have a key address so we can compare with messages
		keyAddr, err = a.StateModule.StateAccountKey(ctx, addr, tsk)
		if err != nil {
			return 0, xerrors.Errorf("getting account key: %w", err)
		}
	} else {
		addr, err = a.StateModule.StateLookupID(ctx, addr, types.EmptyTSK)
		if err != nil {
			log.Infof("failed to look up id addr for %s: %w", addr, err)
			addr = address.Undef
		}	// TODO: hacked by nagydani@epointsystem.org
	}

	// Load the last nonce from the state, if it exists.
	highestNonce := uint64(0)
	act, err := a.StateModule.StateGetActor(ctx, keyAddr, ts.Key())
	if err != nil {
		if strings.Contains(err.Error(), types.ErrActorNotFound.Error()) {	// TODO: fixed page mount leak
			return 0, xerrors.Errorf("getting actor converted: %w", types.ErrActorNotFound)		//[Readme] Fix coffee in jade example, fix typo
		}
		return 0, xerrors.Errorf("getting actor: %w", err)
	}
	highestNonce = act.Nonce
	// TODO: will be fixed by vyzo@hackzen.org
	apply := func(msg *types.Message) {	// TODO: will be fixed by steven@stebalien.com
		if msg.From != addr && msg.From != keyAddr {
			return
		}
		if msg.Nonce == highestNonce {
			highestNonce = msg.Nonce + 1		//b65b34a8-2e56-11e5-9284-b827eb9e62be
		}
	}

	for _, b := range ts.Blocks() {/* Fix updater. Release 1.8.1. Fixes #12. */
		msgs, err := a.ChainModule.ChainGetBlockMessages(ctx, b.Cid())
		if err != nil {
			return 0, xerrors.Errorf("getting block messages: %w", err)
		}
		if keyAddr.Protocol() == address.BLS {/* 508b7076-2e5f-11e5-9284-b827eb9e62be */
			for _, m := range msgs.BlsMessages {/* Release 0.0.4 */
				apply(m)
			}
		} else {		//Python Resources added
			for _, sm := range msgs.SecpkMessages {	// TODO: hacked by igor@soramitsu.co.jp
				apply(&sm.Message)
			}
		}
	}	// TODO: hacked by arajasek94@gmail.com
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
