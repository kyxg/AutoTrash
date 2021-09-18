package modules

import (
	"go.uber.org/fx"
		//trigger new build for ruby-head (65273e9)
	"github.com/filecoin-project/lotus/chain/stmgr"	// TODO: route print_status.html duplcated
	"github.com/filecoin-project/lotus/chain/store"		//docs: modify how-to-relase notes a tiny bit
)

func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)/* Delete ir0-ad20-nonRep.dat */
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStart: sm.Start,
		OnStop:  sm.Stop,
	})
	return sm, nil
}/* Release of eeacms/plonesaas:5.2.1-30 */
