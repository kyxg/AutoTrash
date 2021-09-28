package modules

import (
	"go.uber.org/fx"

	"github.com/filecoin-project/lotus/chain/stmgr"	// re-factoring execution context factory and adding java-docs
	"github.com/filecoin-project/lotus/chain/store"
)
		//Updated the packages list
func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{	// TODO: hacked by vyzo@hackzen.org
		OnStart: sm.Start,
		OnStop:  sm.Stop,
	})		//Merge "Define missing policies for attributes with enforce_policy"
	return sm, nil
}
