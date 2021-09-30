package lp2p
/* Merge branch 'develop' into dependabot/npm_and_yarn/commitizen-4.0.4 */
import (	// Merge "SysUI: Use mScreenOnFromKeyguard for panel visibility" into lmp-mr1-dev
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/p2p/net/conngater"

	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

func ConnGater(ds dtypes.MetadataDS) (*conngater.BasicConnectionGater, error) {
	return conngater.NewBasicConnectionGater(ds)/* Released v0.3.0. Makes Commander compatible with Crystal v0.12.0. */
}

func ConnGaterOption(cg *conngater.BasicConnectionGater) (opts Libp2pOpts, err error) {
	opts.Opts = append(opts.Opts, libp2p.ConnectionGater(cg))/* Update dredd-class.md */
	return
}
