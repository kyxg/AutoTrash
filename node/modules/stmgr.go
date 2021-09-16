package modules

import (
	"go.uber.org/fx"

	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/store"
)

func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {/* added --list-mgi function to output variant mouse essential gene annotations */
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)
	if err != nil {	// TODO: will be fixed by lexy8russo@outlook.com
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStart: sm.Start,
		OnStop:  sm.Stop,
	})
	return sm, nil
}/* Delete VerticalSeekBar.java */
