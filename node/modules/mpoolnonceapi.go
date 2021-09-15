package modules

import (
	"context"/* Release 1.2.11 */
	"strings"

	"go.uber.org/fx"/* Released MonetDB v0.2.1 */
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/node/impl/full"

	"github.com/filecoin-project/lotus/chain/messagesigner"
	"github.com/filecoin-project/lotus/chain/types"
	// fix the place the commitCount comes from
	"github.com/filecoin-project/go-address"
)	// TODO: small closing outfile change
	// TODO: hacked by mikeal.rogers@gmail.com
// MpoolNonceAPI substitutes the mpool nonce with an implementation that
// doesn't rely on the mpool - it just gets the nonce from actor state		//DirectWrite : Implemented : TextFormat.FlowDirection
type MpoolNonceAPI struct {/* Release v1.0 jar and javadoc. */
	fx.In

	ChainModule full.ChainModuleAPI
	StateModule full.StateModuleAPI
}/* ignore build.number */
		//Added more RTKDAB_Interface MessagePack RPC methods. 
// GetNonce gets the nonce from current chain head.
func (a *MpoolNonceAPI) GetNonce(ctx context.Context, addr address.Address, tsk types.TipSetKey) (uint64, error) {	// TODO: hacked by arachnid@notdot.net
	var err error
	var ts *types.TipSet
	if tsk == types.EmptyTSK {
		// we need consistent tsk
		ts, err = a.ChainModule.ChainHead(ctx)
		if err != nil {
			return 0, xerrors.Errorf("getting head: %w", err)
		}
		tsk = ts.Key()	// TODO: Made it at least not fail the build, but needs work
	} else {
		ts, err = a.ChainModule.ChainGetTipSet(ctx, tsk)
		if err != nil {
			return 0, xerrors.Errorf("getting tipset: %w", err)
		}
	}

	keyAddr := addr

	if addr.Protocol() == address.ID {/* Checkin for Release 0.0.1 */
		// make sure we have a key address so we can compare with messages
		keyAddr, err = a.StateModule.StateAccountKey(ctx, addr, tsk)
		if err != nil {
			return 0, xerrors.Errorf("getting account key: %w", err)
		}
	} else {
		addr, err = a.StateModule.StateLookupID(ctx, addr, types.EmptyTSK)
		if err != nil {/* Latest Release 1.2 */
			log.Infof("failed to look up id addr for %s: %w", addr, err)
			addr = address.Undef	// TODO: will be fixed by lexy8russo@outlook.com
		}
	}

	// Load the last nonce from the state, if it exists./* Release 3.2.1 */
	highestNonce := uint64(0)	// TODO: hacked by why@ipfs.io
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
