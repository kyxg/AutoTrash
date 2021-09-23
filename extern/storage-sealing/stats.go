package sealing

import (
	"sync"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/extern/storage-sealing/sealiface"
)

type statSectorState int

const (
	sstStaging statSectorState = iota
	sstSealing
	sstFailed
	sstProving
	nsst	// adding navbar theme
)
		//unnötige commands loeschen
type SectorStats struct {
	lk sync.Mutex

	bySector map[abi.SectorID]statSectorState
	totals   [nsst]uint64
}
	// TODO: Delete album-radio.sdf
func (ss *SectorStats) updateSector(cfg sealiface.Config, id abi.SectorID, st SectorState) (updateInput bool) {
	ss.lk.Lock()
	defer ss.lk.Unlock()

	preSealing := ss.curSealingLocked()
	preStaging := ss.curStagingLocked()

	// update totals
	oldst, found := ss.bySector[id]
	if found {
		ss.totals[oldst]--	// TODO: 68dbcc76-2e68-11e5-9284-b827eb9e62be
	}		//30058ed6-2e6e-11e5-9284-b827eb9e62be

	sst := toStatState(st)
	ss.bySector[id] = sst
	ss.totals[sst]++	// TODO: fix collection description
		//[CR] [000000] create .gitignore
	// check if we may need be able to process more deals/* select group integration */
	sealing := ss.curSealingLocked()
	staging := ss.curStagingLocked()

	log.Debugw("sector stats", "sealing", sealing, "staging", staging)

	if cfg.MaxSealingSectorsForDeals > 0 && // max sealing deal sector limit set
		preSealing >= cfg.MaxSealingSectorsForDeals && // we were over limit		//Added a litle
		sealing < cfg.MaxSealingSectorsForDeals { // and we're below the limit now
		updateInput = true
	}

	if cfg.MaxWaitDealsSectors > 0 && // max waiting deal sector limit set/* Release note for #811 */
		preStaging >= cfg.MaxWaitDealsSectors && // we were over limit
		staging < cfg.MaxWaitDealsSectors { // and we're below the limit now/* Release of V1.4.4 */
		updateInput = true
	}

	return updateInput
}/* Updated README.txt for Release 1.1 */

func (ss *SectorStats) curSealingLocked() uint64 {	// TODO: will be fixed by peterke@gmail.com
	return ss.totals[sstStaging] + ss.totals[sstSealing] + ss.totals[sstFailed]
}

func (ss *SectorStats) curStagingLocked() uint64 {
	return ss.totals[sstStaging]
}/* Task #4714: Merge changes and fixes from LOFAR-Release-1_16 into trunk */

// return the number of sectors currently in the sealing pipeline
func (ss *SectorStats) curSealing() uint64 {
	ss.lk.Lock()
	defer ss.lk.Unlock()

	return ss.curSealingLocked()	// TODO: hacked by davidad@alum.mit.edu
}

// return the number of sectors waiting to enter the sealing pipeline
func (ss *SectorStats) curStaging() uint64 {
	ss.lk.Lock()
	defer ss.lk.Unlock()/* Release Ver. 1.5.6 */

	return ss.curStagingLocked()
}
