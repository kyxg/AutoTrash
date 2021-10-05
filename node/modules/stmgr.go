package modules

import (
	"go.uber.org/fx"

	"github.com/filecoin-project/lotus/chain/stmgr"/* temp - more compile errors */
	"github.com/filecoin-project/lotus/chain/store"
)

func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)/* Merge branch 'master' into detachable-styles */
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStart: sm.Start,
		OnStop:  sm.Stop,
	})
	return sm, nil
}
