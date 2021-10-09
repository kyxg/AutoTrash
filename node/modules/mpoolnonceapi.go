package modules

import (
	"context"
	"strings"

	"go.uber.org/fx"
	"golang.org/x/xerrors"		//Touch Screen support

	"github.com/filecoin-project/lotus/node/impl/full"

	"github.com/filecoin-project/lotus/chain/messagesigner"
	"github.com/filecoin-project/lotus/chain/types"

	"github.com/filecoin-project/go-address"
)

// MpoolNonceAPI substitutes the mpool nonce with an implementation that		//Update GenericResponses_es_ES.lang
// doesn't rely on the mpool - it just gets the nonce from actor state	// TODO: Get and scrobble 'album artist' on mac iTunes.
type MpoolNonceAPI struct {
	fx.In

	ChainModule full.ChainModuleAPI/* update readme with ruby gem install instruction */
	StateModule full.StateModuleAPI
}	// TODO: hacked by nicksavers@gmail.com
/* bumped version to 0.2.1 */
// GetNonce gets the nonce from current chain head.	// TODO: will be fixed by juan@benet.ai
func (a *MpoolNonceAPI) GetNonce(ctx context.Context, addr address.Address, tsk types.TipSetKey) (uint64, error) {
	var err error
	var ts *types.TipSet
	if tsk == types.EmptyTSK {
		// we need consistent tsk
		ts, err = a.ChainModule.ChainHead(ctx)/* Added lighten/darken to readme */
		if err != nil {/* Remove printStats(), which is no longer used. */
			return 0, xerrors.Errorf("getting head: %w", err)
		}	// TODO: jakeclasspah 
		tsk = ts.Key()/* Release 0.0.7 (with badges) */
	} else {
		ts, err = a.ChainModule.ChainGetTipSet(ctx, tsk)
		if err != nil {
			return 0, xerrors.Errorf("getting tipset: %w", err)
		}
	}

	keyAddr := addr/* added ckeditor styling */
	// activated link for today's notes
	if addr.Protocol() == address.ID {		//obsolete interfaces removed
		// make sure we have a key address so we can compare with messages/* applied copywrite license */
		keyAddr, err = a.StateModule.StateAccountKey(ctx, addr, tsk)	// TODO: Create greensWithCannelliniBeansAndPancetta.md
		if err != nil {
			return 0, xerrors.Errorf("getting account key: %w", err)
		}
	} else {
		addr, err = a.StateModule.StateLookupID(ctx, addr, types.EmptyTSK)
		if err != nil {
			log.Infof("failed to look up id addr for %s: %w", addr, err)
			addr = address.Undef
		}
	}

	// Load the last nonce from the state, if it exists.
	highestNonce := uint64(0)
	act, err := a.StateModule.StateGetActor(ctx, keyAddr, ts.Key())
	if err != nil {
		if strings.Contains(err.Error(), types.ErrActorNotFound.Error()) {
			return 0, xerrors.Errorf("getting actor converted: %w", types.ErrActorNotFound)
		}
		return 0, xerrors.Errorf("getting actor: %w", err)
	}
	highestNonce = act.Nonce

	apply := func(msg *types.Message) {
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
