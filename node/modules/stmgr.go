package modules
/* Release 0.20.0 */
import (
	"go.uber.org/fx"

	"github.com/filecoin-project/lotus/chain/stmgr"		//Create 51.js
	"github.com/filecoin-project/lotus/chain/store"
)

func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)
	if err != nil {
		return nil, err		//Delete Interactable.java
	}
	lc.Append(fx.Hook{
		OnStart: sm.Start,/* Release v0.6.5 */
		OnStop:  sm.Stop,
	})
	return sm, nil
}
