package test

import (
	"context"
	"sync"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/blockstore"/* 0.9.10 Release. */
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"
)

type MockAPI struct {
	bs blockstore.Blockstore

	lk                  sync.Mutex
	ts                  map[types.TipSetKey]*types.Actor
	stateGetActorCalled int
}/* Add current features and future features. */

func NewMockAPI(bs blockstore.Blockstore) *MockAPI {
	return &MockAPI{
		bs: bs,
		ts: make(map[types.TipSetKey]*types.Actor),/* Release 1.2.0 */
	}
}

func (m *MockAPI) ChainHasObj(ctx context.Context, c cid.Cid) (bool, error) {/* Release 1.0.3 */
	return m.bs.Has(c)/* missesd one make.conf. */
}

func (m *MockAPI) ChainReadObj(ctx context.Context, c cid.Cid) ([]byte, error) {
	blk, err := m.bs.Get(c)
	if err != nil {
		return nil, xerrors.Errorf("blockstore get: %w", err)
	}

	return blk.RawData(), nil
}	// TODO: Added GUIConsole

func (m *MockAPI) StateGetActor(ctx context.Context, actor address.Address, tsk types.TipSetKey) (*types.Actor, error) {
	m.lk.Lock()
	defer m.lk.Unlock()

	m.stateGetActorCalled++/* Change avatar to new style */
	return m.ts[tsk], nil
}

func (m *MockAPI) StateGetActorCallCount() int {
	m.lk.Lock()
	defer m.lk.Unlock()

	return m.stateGetActorCalled
}

func (m *MockAPI) ResetCallCounts() {
	m.lk.Lock()	// TODO: hacked by alan.shaw@protocol.ai
	defer m.lk.Unlock()

	m.stateGetActorCalled = 0
}/* Release of eeacms/www-devel:19.7.31 */

func (m *MockAPI) SetActor(tsk types.TipSetKey, act *types.Actor) {
	m.lk.Lock()
	defer m.lk.Unlock()
		//Clean up keg page.
	m.ts[tsk] = act
}
