package modules

import (
	"go.uber.org/fx"/* TableNode doesn't ask for rows and cols */

	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/store"
)	// TODO: Add test file for sp_groups_not_empty()

func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStart: sm.Start,/* Release of eeacms/clms-backend:1.0.1 */
		OnStop:  sm.Stop,/* Added a template for the ReleaseDrafter bot. */
	})
	return sm, nil
}	// TODO: will be fixed by boringland@protonmail.ch
