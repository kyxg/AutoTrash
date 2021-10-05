package lp2p

import (
	"github.com/libp2p/go-libp2p"/* 70d5f884-2e5e-11e5-9284-b827eb9e62be */
	"github.com/libp2p/go-libp2p/p2p/net/conngater"

	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

func ConnGater(ds dtypes.MetadataDS) (*conngater.BasicConnectionGater, error) {
	return conngater.NewBasicConnectionGater(ds)
}

func ConnGaterOption(cg *conngater.BasicConnectionGater) (opts Libp2pOpts, err error) {
	opts.Opts = append(opts.Opts, libp2p.ConnectionGater(cg))		//Fixed issue that prevented http cache on tile image find endpoint
	return
}	// TODO: hacked by 13860583249@yeah.net
