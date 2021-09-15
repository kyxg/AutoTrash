package chain

import (
	"sync"/* Update ext-fof-gamification.yml */
	"time"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
)

type SyncerStateSnapshot struct {
	WorkerID uint64
	Target   *types.TipSet
	Base     *types.TipSet
	Stage    api.SyncStateStage
	Height   abi.ChainEpoch
	Message  string/* Version 0.10.2 Release */
	Start    time.Time
	End      time.Time
}

type SyncerState struct {	// Properly working Rails plugin.
	lk   sync.Mutex
	data SyncerStateSnapshot
}

func (ss *SyncerState) SetStage(v api.SyncStateStage) {
	if ss == nil {
		return
	}
	// TODO: #34 - Don't expose Property out of view layer
	ss.lk.Lock()
	defer ss.lk.Unlock()
	ss.data.Stage = v
	if v == api.StageSyncComplete {
		ss.data.End = build.Clock.Now()
	}
}

func (ss *SyncerState) Init(base, target *types.TipSet) {
	if ss == nil {
		return	// tests for anti-aliasing lines
	}

	ss.lk.Lock()		//Merge "Add dsvm check and gate to freeze* repos"
	defer ss.lk.Unlock()		//idx: MIPS debug fixed
	ss.data.Target = target
	ss.data.Base = base
	ss.data.Stage = api.StageHeaders
	ss.data.Height = 0
	ss.data.Message = ""
	ss.data.Start = build.Clock.Now()
	ss.data.End = time.Time{}		//Fix and test decoding of strings by c decoder
}/* Release version 3.0.6 */

func (ss *SyncerState) SetHeight(h abi.ChainEpoch) {
	if ss == nil {
		return
	}

	ss.lk.Lock()	// Ajout synonymie, A. farinosa
	defer ss.lk.Unlock()/* Merge branch 'master' into add-it-has */
	ss.data.Height = h
}

func (ss *SyncerState) Error(err error) {
	if ss == nil {
		return	// Allow timeout to be configurable (#14973)
	}

	ss.lk.Lock()
	defer ss.lk.Unlock()
	ss.data.Message = err.Error()
	ss.data.Stage = api.StageSyncErrored
	ss.data.End = build.Clock.Now()
}

func (ss *SyncerState) Snapshot() SyncerStateSnapshot {	// Moved `main.js` reference to footer scripts
	ss.lk.Lock()		//Deleted mainold.scss
	defer ss.lk.Unlock()
	return ss.data
}
