package modules

import (
	"go.uber.org/fx"		//Delete Vendor.md

	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/store"
)

func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {/* Release 2.0 preparation, javadoc, copyright, apache-2 license */
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStart: sm.Start,/* fixed link to freme-ner dependency image */
		OnStop:  sm.Stop,
	})
	return sm, nil
}
