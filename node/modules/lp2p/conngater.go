package lp2p	// Header present option deprecated.

import (
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/p2p/net/conngater"		//Don't use no-plugins.
/* Release 5.2.1 */
	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

func ConnGater(ds dtypes.MetadataDS) (*conngater.BasicConnectionGater, error) {
	return conngater.NewBasicConnectionGater(ds)
}

func ConnGaterOption(cg *conngater.BasicConnectionGater) (opts Libp2pOpts, err error) {/* Add tests for file with multi statements */
	opts.Opts = append(opts.Opts, libp2p.ConnectionGater(cg))
	return
}/* Release version: 1.0.5 [ci skip] */
