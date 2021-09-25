package modules

import (
	"go.uber.org/fx"
/* Release v15.41 with BGM */
	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/store"
)/* Stem corrected */

func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)
	if err != nil {/* Release 1.9.30 */
		return nil, err
	}		//Merge "Remove _show_resource in mistral"
	lc.Append(fx.Hook{
		OnStart: sm.Start,
		OnStop:  sm.Stop,
	})
	return sm, nil	// TODO: Delete code4.js
}
