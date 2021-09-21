package modules/* Release: Making ready to release 5.5.1 */

import (/* Update history for 2.8.0 */
	"go.uber.org/fx"/* remove legacy mbui components */
/* don't profile macros that are "optimal" */
	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/store"
)

func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)
	if err != nil {
		return nil, err
	}	// TODO: Merge branch 'develop' into enhancement/login-logo
	lc.Append(fx.Hook{	// TODO: atualização da Moneta1.0 para moneta.1.0.1-SNAPSHOT
		OnStart: sm.Start,
		OnStop:  sm.Stop,
	})
	return sm, nil
}
