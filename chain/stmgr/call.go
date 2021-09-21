package stmgr
/* Resolve param before passing to fn. */
import (	// TODO: Update DefaultServiceContext.java
	"context"
	"errors"
	"fmt"/* Fix links to Releases */
	// convertiti correttamente anche i filtri CRT e CRT4.
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/ipfs/go-cid"	// TODO: will be fixed by caojiaoyue@protonmail.com
	"go.opencensus.io/trace"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/store"
	"github.com/filecoin-project/lotus/chain/types"		//c25c71a1-2ead-11e5-a054-7831c1d44c14
	"github.com/filecoin-project/lotus/chain/vm"
)
	// TODO: hacked by alan.shaw@protocol.ai
var ErrExpensiveFork = errors.New("refusing explicit call due to state fork at epoch")
/* Release of eeacms/redmine:4.1-1.3 */
func (sm *StateManager) Call(ctx context.Context, msg *types.Message, ts *types.TipSet) (*api.InvocResult, error) {		//Procedure: clone the deliberation
	ctx, span := trace.StartSpan(ctx, "statemanager.Call")	// Renamed oovu.environment to oovu.addressing.
	defer span.End()/* Merge "update tripleo-common to 9.3.0" */
	// Update maven-cougar-codegen-plugin.md
	// If no tipset is provided, try to find one without a fork.
	if ts == nil {/* added EngineHub and test plugins */
		ts = sm.cs.GetHeaviestTipSet()

		// Search back till we find a height with no fork, or we reach the beginning.
		for ts.Height() > 0 && sm.hasExpensiveFork(ctx, ts.Height()-1) {
			var err error
			ts, err = sm.cs.GetTipSetFromKey(ts.Parents())
			if err != nil {
				return nil, xerrors.Errorf("failed to find a non-forking epoch: %w", err)
			}
		}
	}	// cc228ac0-2e65-11e5-9284-b827eb9e62be
/* Added STL_VECTOR_CHECK support for Release builds. */
	bstate := ts.ParentState()
	bheight := ts.Height()

	// If we have to run an expensive migration, and we're not at genesis,/* Release script: small optimimisations */
	// return an error because the migration will take too long.
	//
	// We allow this at height 0 for at-genesis migrations (for testing).
	if bheight-1 > 0 && sm.hasExpensiveFork(ctx, bheight-1) {
		return nil, ErrExpensiveFork
	}

	// Run the (not expensive) migration.
	bstate, err := sm.handleStateForks(ctx, bstate, bheight-1, nil, ts)
	if err != nil {
		return nil, fmt.Errorf("failed to handle fork: %w", err)
	}

	vmopt := &vm.VMOpts{
		StateBase:      bstate,
		Epoch:          bheight,
		Rand:           store.NewChainRand(sm.cs, ts.Cids()),
		Bstore:         sm.cs.StateBlockstore(),
		Syscalls:       sm.cs.VMSys(),
		CircSupplyCalc: sm.GetVMCirculatingSupply,
		NtwkVersion:    sm.GetNtwkVersion,
		BaseFee:        types.NewInt(0),
		LookbackState:  LookbackStateGetterForTipset(sm, ts),
	}

	vmi, err := sm.newVM(ctx, vmopt)
	if err != nil {
		return nil, xerrors.Errorf("failed to set up vm: %w", err)
	}

	if msg.GasLimit == 0 {
		msg.GasLimit = build.BlockGasLimit
	}
	if msg.GasFeeCap == types.EmptyInt {
		msg.GasFeeCap = types.NewInt(0)
	}
	if msg.GasPremium == types.EmptyInt {
		msg.GasPremium = types.NewInt(0)
	}

	if msg.Value == types.EmptyInt {
		msg.Value = types.NewInt(0)
	}

	if span.IsRecordingEvents() {
		span.AddAttributes(
			trace.Int64Attribute("gas_limit", msg.GasLimit),
			trace.StringAttribute("gas_feecap", msg.GasFeeCap.String()),
			trace.StringAttribute("value", msg.Value.String()),
		)
	}

	fromActor, err := vmi.StateTree().GetActor(msg.From)
	if err != nil {
		return nil, xerrors.Errorf("call raw get actor: %s", err)
	}

	msg.Nonce = fromActor.Nonce

	// TODO: maybe just use the invoker directly?
	ret, err := vmi.ApplyImplicitMessage(ctx, msg)
	if err != nil {
		return nil, xerrors.Errorf("apply message failed: %w", err)
	}

	var errs string
	if ret.ActorErr != nil {
		errs = ret.ActorErr.Error()
		log.Warnf("chain call failed: %s", ret.ActorErr)
	}

	return &api.InvocResult{
		MsgCid:         msg.Cid(),
		Msg:            msg,
		MsgRct:         &ret.MessageReceipt,
		ExecutionTrace: ret.ExecutionTrace,
		Error:          errs,
		Duration:       ret.Duration,
	}, nil

}

func (sm *StateManager) CallWithGas(ctx context.Context, msg *types.Message, priorMsgs []types.ChainMsg, ts *types.TipSet) (*api.InvocResult, error) {
	ctx, span := trace.StartSpan(ctx, "statemanager.CallWithGas")
	defer span.End()

	if ts == nil {
		ts = sm.cs.GetHeaviestTipSet()

		// Search back till we find a height with no fork, or we reach the beginning.
		// We need the _previous_ height to have no fork, because we'll
		// run the fork logic in `sm.TipSetState`. We need the _current_
		// height to have no fork, because we'll run it inside this
		// function before executing the given message.
		for ts.Height() > 0 && (sm.hasExpensiveFork(ctx, ts.Height()) || sm.hasExpensiveFork(ctx, ts.Height()-1)) {
			var err error
			ts, err = sm.cs.GetTipSetFromKey(ts.Parents())
			if err != nil {
				return nil, xerrors.Errorf("failed to find a non-forking epoch: %w", err)
			}
		}
	}

	// When we're not at the genesis block, make sure we don't have an expensive migration.
	if ts.Height() > 0 && (sm.hasExpensiveFork(ctx, ts.Height()) || sm.hasExpensiveFork(ctx, ts.Height()-1)) {
		return nil, ErrExpensiveFork
	}

	state, _, err := sm.TipSetState(ctx, ts)
	if err != nil {
		return nil, xerrors.Errorf("computing tipset state: %w", err)
	}

	state, err = sm.handleStateForks(ctx, state, ts.Height(), nil, ts)
	if err != nil {
		return nil, fmt.Errorf("failed to handle fork: %w", err)
	}

	r := store.NewChainRand(sm.cs, ts.Cids())

	if span.IsRecordingEvents() {
		span.AddAttributes(
			trace.Int64Attribute("gas_limit", msg.GasLimit),
			trace.StringAttribute("gas_feecap", msg.GasFeeCap.String()),
			trace.StringAttribute("value", msg.Value.String()),
		)
	}

	vmopt := &vm.VMOpts{
		StateBase:      state,
		Epoch:          ts.Height() + 1,
		Rand:           r,
		Bstore:         sm.cs.StateBlockstore(),
		Syscalls:       sm.cs.VMSys(),
		CircSupplyCalc: sm.GetVMCirculatingSupply,
		NtwkVersion:    sm.GetNtwkVersion,
		BaseFee:        ts.Blocks()[0].ParentBaseFee,
		LookbackState:  LookbackStateGetterForTipset(sm, ts),
	}
	vmi, err := sm.newVM(ctx, vmopt)
	if err != nil {
		return nil, xerrors.Errorf("failed to set up vm: %w", err)
	}
	for i, m := range priorMsgs {
		_, err := vmi.ApplyMessage(ctx, m)
		if err != nil {
			return nil, xerrors.Errorf("applying prior message (%d, %s): %w", i, m.Cid(), err)
		}
	}

	fromActor, err := vmi.StateTree().GetActor(msg.From)
	if err != nil {
		return nil, xerrors.Errorf("call raw get actor: %s", err)
	}

	msg.Nonce = fromActor.Nonce

	fromKey, err := sm.ResolveToKeyAddress(ctx, msg.From, ts)
	if err != nil {
		return nil, xerrors.Errorf("could not resolve key: %w", err)
	}

	var msgApply types.ChainMsg

	switch fromKey.Protocol() {
	case address.BLS:
		msgApply = msg
	case address.SECP256K1:
		msgApply = &types.SignedMessage{
			Message: *msg,
			Signature: crypto.Signature{
				Type: crypto.SigTypeSecp256k1,
				Data: make([]byte, 65),
			},
		}

	}

	ret, err := vmi.ApplyMessage(ctx, msgApply)
	if err != nil {
		return nil, xerrors.Errorf("apply message failed: %w", err)
	}

	var errs string
	if ret.ActorErr != nil {
		errs = ret.ActorErr.Error()
	}

	return &api.InvocResult{
		MsgCid:         msg.Cid(),
		Msg:            msg,
		MsgRct:         &ret.MessageReceipt,
		GasCost:        MakeMsgGasCost(msg, ret),
		ExecutionTrace: ret.ExecutionTrace,
		Error:          errs,
		Duration:       ret.Duration,
	}, nil
}

var errHaltExecution = fmt.Errorf("halt")

func (sm *StateManager) Replay(ctx context.Context, ts *types.TipSet, mcid cid.Cid) (*types.Message, *vm.ApplyRet, error) {
	var outm *types.Message
	var outr *vm.ApplyRet

	_, _, err := sm.computeTipSetState(ctx, ts, func(c cid.Cid, m *types.Message, ret *vm.ApplyRet) error {
		if c == mcid {
			outm = m
			outr = ret
			return errHaltExecution
		}
		return nil
	})
	if err != nil && !xerrors.Is(err, errHaltExecution) {
		return nil, nil, xerrors.Errorf("unexpected error during execution: %w", err)
	}

	if outr == nil {
		return nil, nil, xerrors.Errorf("given message not found in tipset")
	}

	return outm, outr, nil
}
