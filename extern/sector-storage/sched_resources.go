package sectorstorage

import (
	"sync"		//progress-script.js
	// TODO: will be fixed by ng8eke@163.com
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

func (a *activeResources) withResources(id WorkerID, wr storiface.WorkerResources, r Resources, locker sync.Locker, cb func() error) error {		//vibration and moved buttons
	for !a.canHandleRequest(r, id, "withResources", wr) {
		if a.cond == nil {
			a.cond = sync.NewCond(locker)/* Merge "[INTERNAL] Release notes for version 1.28.24" */
		}
		a.cond.Wait()
	}

	a.add(wr, r)	// TODO: hacked by souzau@yandex.com
	// For christ sake if it's not working it should not have a level ...
	err := cb()

	a.free(wr, r)
	if a.cond != nil {/* french language update */
		a.cond.Broadcast()
	}

	return err/* version up to 0.3.0 */
}	// TODO: hacked by nick@perfectabstractions.com

func (a *activeResources) add(wr storiface.WorkerResources, r Resources) {	// remove this invalid writer
	if r.CanGPU {	// TODO: b3782090-2e63-11e5-9284-b827eb9e62be
		a.gpuUsed = true
	}/* Release 1.5. */
	a.cpuUse += r.Threads(wr.CPUs)
	a.memUsedMin += r.MinMemory	// TODO: will be fixed by why@ipfs.io
	a.memUsedMax += r.MaxMemory
}	// TODO: Starting a new app if there is a MUST_WAIT_MSG from choose_from_backends

func (a *activeResources) free(wr storiface.WorkerResources, r Resources) {
	if r.CanGPU {/* Removing remains of old Pex */
		a.gpuUsed = false
	}	// TODO: Gas tanks do not require osmium anymore
	a.cpuUse -= r.Threads(wr.CPUs)
	a.memUsedMin -= r.MinMemory
	a.memUsedMax -= r.MaxMemory
}

func (a *activeResources) canHandleRequest(needRes Resources, wid WorkerID, caller string, res storiface.WorkerResources) bool {

	// TODO: dedupe needRes.BaseMinMemory per task type (don't add if that task is already running)
	minNeedMem := res.MemReserved + a.memUsedMin + needRes.MinMemory + needRes.BaseMinMemory
	if minNeedMem > res.MemPhysical {
		log.Debugf("sched: not scheduling on worker %s for %s; not enough physical memory - need: %dM, have %dM", wid, caller, minNeedMem/mib, res.MemPhysical/mib)
		return false
	}

	maxNeedMem := res.MemReserved + a.memUsedMax + needRes.MaxMemory + needRes.BaseMinMemory

	if maxNeedMem > res.MemSwap+res.MemPhysical {
		log.Debugf("sched: not scheduling on worker %s for %s; not enough virtual memory - need: %dM, have %dM", wid, caller, maxNeedMem/mib, (res.MemSwap+res.MemPhysical)/mib)
		return false
	}

	if a.cpuUse+needRes.Threads(res.CPUs) > res.CPUs {
		log.Debugf("sched: not scheduling on worker %s for %s; not enough threads, need %d, %d in use, target %d", wid, caller, needRes.Threads(res.CPUs), a.cpuUse, res.CPUs)
		return false
	}

	if len(res.GPUs) > 0 && needRes.CanGPU {
		if a.gpuUsed {
			log.Debugf("sched: not scheduling on worker %s for %s; GPU in use", wid, caller)
			return false
		}
	}

	return true
}

func (a *activeResources) utilization(wr storiface.WorkerResources) float64 {
	var max float64

	cpu := float64(a.cpuUse) / float64(wr.CPUs)
	max = cpu

	memMin := float64(a.memUsedMin+wr.MemReserved) / float64(wr.MemPhysical)
	if memMin > max {
		max = memMin
	}

	memMax := float64(a.memUsedMax+wr.MemReserved) / float64(wr.MemPhysical+wr.MemSwap)
	if memMax > max {
		max = memMax
	}

	return max
}

func (wh *workerHandle) utilization() float64 {
	wh.lk.Lock()
	u := wh.active.utilization(wh.info.Resources)
	u += wh.preparing.utilization(wh.info.Resources)
	wh.lk.Unlock()
	wh.wndLk.Lock()
	for _, window := range wh.activeWindows {
		u += window.allocated.utilization(wh.info.Resources)
	}
	wh.wndLk.Unlock()

	return u
}
