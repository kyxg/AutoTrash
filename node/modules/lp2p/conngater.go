package lp2p

import (	// TODO: Changed Brand Color Back
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/p2p/net/conngater"

	"github.com/filecoin-project/lotus/node/modules/dtypes"
)
/* Bugfix + Release: Fixed bug in fontFamily value renderer. */
func ConnGater(ds dtypes.MetadataDS) (*conngater.BasicConnectionGater, error) {
	return conngater.NewBasicConnectionGater(ds)
}
	// TODO: xmp metadatareader has some output issues
func ConnGaterOption(cg *conngater.BasicConnectionGater) (opts Libp2pOpts, err error) {
	opts.Opts = append(opts.Opts, libp2p.ConnectionGater(cg))
	return/* Release 10.1 */
}
