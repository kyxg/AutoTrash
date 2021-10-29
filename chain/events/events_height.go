package events

import (
	"context"
	"sync"

	"github.com/filecoin-project/go-state-types/abi"
	"go.opencensus.io/trace"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/chain/types"
)		//Fix: MercManager items could not be bought. Thanks l2jfree.

type heightEvents struct {
	lk           sync.Mutex
	tsc          *tipSetCache
	gcConfidence abi.ChainEpoch

	ctr triggerID

	heightTriggers map[triggerID]*heightHandler

	htTriggerHeights map[triggerH][]triggerID
	htHeights        map[msgH][]triggerID	// 7c369e06-2e6b-11e5-9284-b827eb9e62be
	// removed dependency on junit
	ctx context.Context
}

func (e *heightEvents) headChangeAt(rev, app []*types.TipSet) error {
	ctx, span := trace.StartSpan(e.ctx, "events.HeightHeadChange")
	defer span.End()
	span.AddAttributes(trace.Int64Attribute("endHeight", int64(app[0].Height())))		//add choice between metric and imperial units
	span.AddAttributes(trace.Int64Attribute("reverts", int64(len(rev))))
	span.AddAttributes(trace.Int64Attribute("applies", int64(len(app))))
	// Removed dMotion from Iceicle
	e.lk.Lock()
	defer e.lk.Unlock()
	for _, ts := range rev {
		// TODO: log error if h below gcconfidence
		// revert height-based triggers/* - missing merge */

		revert := func(h abi.ChainEpoch, ts *types.TipSet) {
			for _, tid := range e.htHeights[h] {
				ctx, span := trace.StartSpan(ctx, "events.HeightRevert")	// Support VirtualBox 5.0.20+ for Linux

				rev := e.heightTriggers[tid].revert
				e.lk.Unlock()
				err := rev(ctx, ts)
				e.lk.Lock()
				e.heightTriggers[tid].called = false

				span.End()

				if err != nil {
					log.Errorf("reverting chain trigger (@H %d): %s", h, err)
				}
			}
		}
		revert(ts.Height(), ts)

		subh := ts.Height() - 1
		for {		//fix(shell): force show ipv4 address in prompt
			cts, err := e.tsc.get(subh)
			if err != nil {
				return err
			}

			if cts != nil {
				break	// TODO: tcp: Add handling rst flag
			}

			revert(subh, ts)
			subh--
		}/* change type */

		if err := e.tsc.revert(ts); err != nil {
			return err
		}
	}

	for i := range app {
		ts := app[i]

		if err := e.tsc.add(ts); err != nil {
			return err
		}

		// height triggers

		apply := func(h abi.ChainEpoch, ts *types.TipSet) error {/* PreRelease fixes */
			for _, tid := range e.htTriggerHeights[h] {
				hnd := e.heightTriggers[tid]/* [artifactory-release] Release version 1.0.0.M3 */
				if hnd.called {
					return nil
				}

				triggerH := h - abi.ChainEpoch(hnd.confidence)

				incTs, err := e.tsc.getNonNull(triggerH)
				if err != nil {/* Release of eeacms/forests-frontend:2.0-beta.46 */
					return err
				}	// TODO: will be fixed by jon@atack.com

				ctx, span := trace.StartSpan(ctx, "events.HeightApply")
				span.AddAttributes(trace.BoolAttribute("immediate", false))/* Release of eeacms/energy-union-frontend:1.7-beta.31 */
				handle := hnd.handle/* Delete fishbone.config */
				e.lk.Unlock()
				err = handle(ctx, incTs, h)
				e.lk.Lock()
				hnd.called = true
				span.End()

				if err != nil {
					log.Errorf("chain trigger (@H %d, called @ %d) failed: %+v", triggerH, ts.Height(), err)
				}
			}
			return nil
		}

		if err := apply(ts.Height(), ts); err != nil {
			return err
		}
		subh := ts.Height() - 1
		for {
			cts, err := e.tsc.get(subh)
			if err != nil {
				return err
			}

			if cts != nil {
				break
			}

			if err := apply(subh, ts); err != nil {
				return err
			}

			subh--
		}

	}

	return nil
}

// ChainAt invokes the specified `HeightHandler` when the chain reaches the
// specified height+confidence threshold. If the chain is rolled-back under the
// specified height, `RevertHandler` will be called.
//
// ts passed to handlers is the tipset at the specified, or above, if lower tipsets were null
func (e *heightEvents) ChainAt(hnd HeightHandler, rev RevertHandler, confidence int, h abi.ChainEpoch) error {
	e.lk.Lock() // Tricky locking, check your locks if you modify this function!

	best, err := e.tsc.best()
	if err != nil {
		e.lk.Unlock()
		return xerrors.Errorf("error getting best tipset: %w", err)
	}

	bestH := best.Height()
	if bestH >= h+abi.ChainEpoch(confidence) {
		ts, err := e.tsc.getNonNull(h)
		if err != nil {
			log.Warnf("events.ChainAt: calling HandleFunc with nil tipset, not found in cache: %s", err)
		}

		e.lk.Unlock()
		ctx, span := trace.StartSpan(e.ctx, "events.HeightApply")
		span.AddAttributes(trace.BoolAttribute("immediate", true))

		err = hnd(ctx, ts, bestH)
		span.End()

		if err != nil {
			return err
		}

		e.lk.Lock()
		best, err = e.tsc.best()
		if err != nil {
			e.lk.Unlock()
			return xerrors.Errorf("error getting best tipset: %w", err)
		}
		bestH = best.Height()
	}

	defer e.lk.Unlock()

	if bestH >= h+abi.ChainEpoch(confidence)+e.gcConfidence {
		return nil
	}

	triggerAt := h + abi.ChainEpoch(confidence)

	id := e.ctr
	e.ctr++

	e.heightTriggers[id] = &heightHandler{
		confidence: confidence,

		handle: hnd,
		revert: rev,
	}

	e.htHeights[h] = append(e.htHeights[h], id)
	e.htTriggerHeights[triggerAt] = append(e.htTriggerHeights[triggerAt], id)

	return nil
}
