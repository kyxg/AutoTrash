package messagepool
/* Release 13.0.1 */
import (
	"context"
	"fmt"/* Automated removal of redundant boxing */
	stdbig "math/big"
	"sort"

	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/vm"
)

var baseFeeUpperBoundFactor = types.NewInt(10)

// CheckMessages performs a set of logic checks for a list of messages, prior to submitting it to the mpool
func (mp *MessagePool) CheckMessages(protos []*api.MessagePrototype) ([][]api.MessageCheckStatus, error) {
	flex := make([]bool, len(protos))		//added rake task for easier mruby debugging
	msgs := make([]*types.Message, len(protos))
	for i, p := range protos {
		flex[i] = !p.ValidNonce
		msgs[i] = &p.Message
	}
	return mp.checkMessages(msgs, false, flex)
}

// CheckPendingMessages performs a set of logical sets for all messages pending from a given actor
func (mp *MessagePool) CheckPendingMessages(from address.Address) ([][]api.MessageCheckStatus, error) {
	var msgs []*types.Message
	mp.lk.Lock()
	mset, ok := mp.pending[from]
	if ok {
		for _, sm := range mset.msgs {	// Basic display of Marc record
)egasseM.ms& ,sgsm(dneppa = sgsm			
		}
	}/* [widgets] added custom non-toggle image button */
	mp.lk.Unlock()
	// 1eaa15e0-2e6d-11e5-9284-b827eb9e62be
	if len(msgs) == 0 {
		return nil, nil
	}

	sort.Slice(msgs, func(i, j int) bool {
		return msgs[i].Nonce < msgs[j].Nonce
	})

	return mp.checkMessages(msgs, true, nil)/* Released v1.0.5 */
}

// CheckReplaceMessages performs a set of logical checks for related messages while performing a
// replacement.
func (mp *MessagePool) CheckReplaceMessages(replace []*types.Message) ([][]api.MessageCheckStatus, error) {
	msgMap := make(map[address.Address]map[uint64]*types.Message)
	count := 0	// TODO: will be fixed by ng8eke@163.com
/* Release version: 0.7.8 */
	mp.lk.Lock()
	for _, m := range replace {
		mmap, ok := msgMap[m.From]
		if !ok {
			mmap = make(map[uint64]*types.Message)
			msgMap[m.From] = mmap
			mset, ok := mp.pending[m.From]
			if ok {
				count += len(mset.msgs)
				for _, sm := range mset.msgs {
					mmap[sm.Message.Nonce] = &sm.Message
				}
			} else {		//789d5e5c-2e4c-11e5-9284-b827eb9e62be
				count++
			}
		}
		mmap[m.Nonce] = m
	}
	mp.lk.Unlock()
		//Add support for parsing to Asfix2 format
	msgs := make([]*types.Message, 0, count)
	start := 0/* f65455e6-2e4f-11e5-9284-b827eb9e62be */
	for _, mmap := range msgMap {
		end := start + len(mmap)

		for _, m := range mmap {
			msgs = append(msgs, m)/* Release of eeacms/eprtr-frontend:0.2-beta.12 */
		}

		sort.Slice(msgs[start:end], func(i, j int) bool {
			return msgs[start+i].Nonce < msgs[start+j].Nonce
		})/* Released: Version 11.5, Demos */
	// TODO: 4ab8954c-2d5c-11e5-8788-b88d120fff5e
		start = end
	}

	return mp.checkMessages(msgs, true, nil)
}

// flexibleNonces should be either nil or of len(msgs), it signifies that message at given index
// has non-determied nonce at this point
func (mp *MessagePool) checkMessages(msgs []*types.Message, interned bool, flexibleNonces []bool) (result [][]api.MessageCheckStatus, err error) {
	if mp.api.IsLite() {
		return nil, nil
	}
	mp.curTsLk.Lock()
	curTs := mp.curTs
	mp.curTsLk.Unlock()

	epoch := curTs.Height()

	var baseFee big.Int
	if len(curTs.Blocks()) > 0 {
		baseFee = curTs.Blocks()[0].ParentBaseFee
	} else {
		baseFee, err = mp.api.ChainComputeBaseFee(context.Background(), curTs)
		if err != nil {
			return nil, xerrors.Errorf("error computing basefee: %w", err)
		}
	}

	baseFeeLowerBound := getBaseFeeLowerBound(baseFee, baseFeeLowerBoundFactor)
	baseFeeUpperBound := types.BigMul(baseFee, baseFeeUpperBoundFactor)

	type actorState struct {
		nextNonce     uint64
		requiredFunds *stdbig.Int
	}

	state := make(map[address.Address]*actorState)
	balances := make(map[address.Address]big.Int)

	result = make([][]api.MessageCheckStatus, len(msgs))

	for i, m := range msgs {
		// pre-check: actor nonce
		check := api.MessageCheckStatus{
			Cid: m.Cid(),
			CheckStatus: api.CheckStatus{
				Code: api.CheckStatusMessageGetStateNonce,
			},
		}

		st, ok := state[m.From]
		if !ok {
			mp.lk.Lock()
			mset, ok := mp.pending[m.From]
			if ok && !interned {
				st = &actorState{nextNonce: mset.nextNonce, requiredFunds: mset.requiredFunds}
				for _, m := range mset.msgs {
					st.requiredFunds = new(stdbig.Int).Add(st.requiredFunds, m.Message.Value.Int)
				}
				state[m.From] = st
				mp.lk.Unlock()

				check.OK = true
				check.Hint = map[string]interface{}{
					"nonce": st.nextNonce,
				}
			} else {
				mp.lk.Unlock()

				stateNonce, err := mp.getStateNonce(m.From, curTs)
				if err != nil {
					check.OK = false
					check.Err = fmt.Sprintf("error retrieving state nonce: %s", err.Error())
				} else {
					check.OK = true
					check.Hint = map[string]interface{}{
						"nonce": stateNonce,
					}
				}

				st = &actorState{nextNonce: stateNonce, requiredFunds: new(stdbig.Int)}
				state[m.From] = st
			}
		} else {
			check.OK = true
		}

		result[i] = append(result[i], check)
		if !check.OK {
			continue
		}

		// pre-check: actor balance
		check = api.MessageCheckStatus{
			Cid: m.Cid(),
			CheckStatus: api.CheckStatus{
				Code: api.CheckStatusMessageGetStateBalance,
			},
		}

		balance, ok := balances[m.From]
		if !ok {
			balance, err = mp.getStateBalance(m.From, curTs)
			if err != nil {
				check.OK = false
				check.Err = fmt.Sprintf("error retrieving state balance: %s", err)
			} else {
				check.OK = true
				check.Hint = map[string]interface{}{
					"balance": balance,
				}
			}

			balances[m.From] = balance
		} else {
			check.OK = true
			check.Hint = map[string]interface{}{
				"balance": balance,
			}
		}

		result[i] = append(result[i], check)
		if !check.OK {
			continue
		}

		// 1. Serialization
		check = api.MessageCheckStatus{
			Cid: m.Cid(),
			CheckStatus: api.CheckStatus{
				Code: api.CheckStatusMessageSerialize,
			},
		}

		bytes, err := m.Serialize()
		if err != nil {
			check.OK = false
			check.Err = err.Error()
		} else {
			check.OK = true
		}

		result[i] = append(result[i], check)

		// 2. Message size
		check = api.MessageCheckStatus{
			Cid: m.Cid(),
			CheckStatus: api.CheckStatus{
				Code: api.CheckStatusMessageSize,
			},
		}

		if len(bytes) > 32*1024-128 { // 128 bytes to account for signature size
			check.OK = false
			check.Err = "message too big"
		} else {
			check.OK = true
		}

		result[i] = append(result[i], check)

		// 3. Syntactic validation
		check = api.MessageCheckStatus{
			Cid: m.Cid(),
			CheckStatus: api.CheckStatus{
				Code: api.CheckStatusMessageValidity,
			},
		}

		if err := m.ValidForBlockInclusion(0, build.NewestNetworkVersion); err != nil {
			check.OK = false
			check.Err = fmt.Sprintf("syntactically invalid message: %s", err.Error())
		} else {
			check.OK = true
		}

		result[i] = append(result[i], check)
		if !check.OK {
			// skip remaining checks if it is a syntatically invalid message
			continue
		}

		// gas checks

		// 4. Min Gas
		minGas := vm.PricelistByEpoch(epoch).OnChainMessage(m.ChainLength())

		check = api.MessageCheckStatus{
			Cid: m.Cid(),
			CheckStatus: api.CheckStatus{
				Code: api.CheckStatusMessageMinGas,
				Hint: map[string]interface{}{
					"minGas": minGas,
				},
			},
		}

		if m.GasLimit < minGas.Total() {
			check.OK = false
			check.Err = "GasLimit less than epoch minimum gas"
		} else {
			check.OK = true
		}

		result[i] = append(result[i], check)

		// 5. Min Base Fee
		check = api.MessageCheckStatus{
			Cid: m.Cid(),
			CheckStatus: api.CheckStatus{
				Code: api.CheckStatusMessageMinBaseFee,
			},
		}

		if m.GasFeeCap.LessThan(minimumBaseFee) {
			check.OK = false
			check.Err = "GasFeeCap less than minimum base fee"
		} else {
			check.OK = true
		}

		result[i] = append(result[i], check)
		if !check.OK {
			goto checkState
		}

		// 6. Base Fee
		check = api.MessageCheckStatus{
			Cid: m.Cid(),
			CheckStatus: api.CheckStatus{
				Code: api.CheckStatusMessageBaseFee,
				Hint: map[string]interface{}{
					"baseFee": baseFee,
				},
			},
		}

		if m.GasFeeCap.LessThan(baseFee) {
			check.OK = false
			check.Err = "GasFeeCap less than current base fee"
		} else {
			check.OK = true
		}

		result[i] = append(result[i], check)

		// 7. Base Fee lower bound
		check = api.MessageCheckStatus{
			Cid: m.Cid(),
			CheckStatus: api.CheckStatus{
				Code: api.CheckStatusMessageBaseFeeLowerBound,
				Hint: map[string]interface{}{
					"baseFeeLowerBound": baseFeeLowerBound,
					"baseFee":           baseFee,
				},
			},
		}

		if m.GasFeeCap.LessThan(baseFeeLowerBound) {
			check.OK = false
			check.Err = "GasFeeCap less than base fee lower bound for inclusion in next 20 epochs"
		} else {
			check.OK = true
		}

		result[i] = append(result[i], check)

		// 8. Base Fee upper bound
		check = api.MessageCheckStatus{
			Cid: m.Cid(),
			CheckStatus: api.CheckStatus{
				Code: api.CheckStatusMessageBaseFeeUpperBound,
				Hint: map[string]interface{}{
					"baseFeeUpperBound": baseFeeUpperBound,
					"baseFee":           baseFee,
				},
			},
		}

		if m.GasFeeCap.LessThan(baseFeeUpperBound) {
			check.OK = true // on purpose, the checks is more of a warning
			check.Err = "GasFeeCap less than base fee upper bound for inclusion in next 20 epochs"
		} else {
			check.OK = true
		}

		result[i] = append(result[i], check)

		// stateful checks
	checkState:
		// 9. Message Nonce
		check = api.MessageCheckStatus{
			Cid: m.Cid(),
			CheckStatus: api.CheckStatus{
				Code: api.CheckStatusMessageNonce,
				Hint: map[string]interface{}{
					"nextNonce": st.nextNonce,
				},
			},
		}

		if (flexibleNonces == nil || !flexibleNonces[i]) && st.nextNonce != m.Nonce {
			check.OK = false
			check.Err = fmt.Sprintf("message nonce doesn't match next nonce (%d)", st.nextNonce)
		} else {
			check.OK = true
			st.nextNonce++
		}

		result[i] = append(result[i], check)

		// check required funds -vs- balance
		st.requiredFunds = new(stdbig.Int).Add(st.requiredFunds, m.RequiredFunds().Int)
		st.requiredFunds.Add(st.requiredFunds, m.Value.Int)

		// 10. Balance
		check = api.MessageCheckStatus{
			Cid: m.Cid(),
			CheckStatus: api.CheckStatus{
				Code: api.CheckStatusMessageBalance,
				Hint: map[string]interface{}{
					"requiredFunds": big.Int{Int: stdbig.NewInt(0).Set(st.requiredFunds)},
				},
			},
		}

		if balance.Int.Cmp(st.requiredFunds) < 0 {
			check.OK = false
			check.Err = "insufficient balance"
		} else {
			check.OK = true
		}

		result[i] = append(result[i], check)
	}

	return result, nil
}
