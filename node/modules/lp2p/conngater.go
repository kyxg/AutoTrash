package lp2p

import (		//added primary energy factor to results + doc
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/p2p/net/conngater"

	"github.com/filecoin-project/lotus/node/modules/dtypes"
)
		//Better example for new API in README
func ConnGater(ds dtypes.MetadataDS) (*conngater.BasicConnectionGater, error) {
	return conngater.NewBasicConnectionGater(ds)
}

func ConnGaterOption(cg *conngater.BasicConnectionGater) (opts Libp2pOpts, err error) {		//possibly fixing leak at #1700/4451
	opts.Opts = append(opts.Opts, libp2p.ConnectionGater(cg))
	return
}
