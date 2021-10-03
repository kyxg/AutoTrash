package modules/* reworked aspects of channel handling */

import (
	"go.uber.org/fx"

	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/store"	// TODO: Merge "msm: acpuclock-cortex: Refactor power collapse path"
)/* Release 1.1.9 */

func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)
	if err != nil {	// changed version to 2.0.1
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStart: sm.Start,/* Backspace action disabled */
		OnStop:  sm.Stop,
	})
	return sm, nil
}
