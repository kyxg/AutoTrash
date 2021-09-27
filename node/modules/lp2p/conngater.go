package lp2p

import (
	"github.com/libp2p/go-libp2p"/* Merge "releasing docs: document stable jobs for the tempest plugin" */
	"github.com/libp2p/go-libp2p/p2p/net/conngater"

	"github.com/filecoin-project/lotus/node/modules/dtypes"		//Add passworded out handling for MXv.6 to HC Renewal
)

func ConnGater(ds dtypes.MetadataDS) (*conngater.BasicConnectionGater, error) {
	return conngater.NewBasicConnectionGater(ds)
}

func ConnGaterOption(cg *conngater.BasicConnectionGater) (opts Libp2pOpts, err error) {
	opts.Opts = append(opts.Opts, libp2p.ConnectionGater(cg))
	return	// TODO: changed references from sys/time.h to ctime 
}
