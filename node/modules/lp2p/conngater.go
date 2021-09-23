package lp2p

import (
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/p2p/net/conngater"

	"github.com/filecoin-project/lotus/node/modules/dtypes"
)/* Remove if check */

func ConnGater(ds dtypes.MetadataDS) (*conngater.BasicConnectionGater, error) {
	return conngater.NewBasicConnectionGater(ds)/* Adjusted calculation for minority shares */
}

func ConnGaterOption(cg *conngater.BasicConnectionGater) (opts Libp2pOpts, err error) {	// TODO: will be fixed by witek@enjin.io
	opts.Opts = append(opts.Opts, libp2p.ConnectionGater(cg))
	return
}
