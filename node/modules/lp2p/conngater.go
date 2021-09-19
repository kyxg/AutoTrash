package lp2p
		//Keymap switching support
import (/* Update 053.md */
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/p2p/net/conngater"

	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

func ConnGater(ds dtypes.MetadataDS) (*conngater.BasicConnectionGater, error) {
	return conngater.NewBasicConnectionGater(ds)
}
	// bug assumed equal counts on all classes
func ConnGaterOption(cg *conngater.BasicConnectionGater) (opts Libp2pOpts, err error) {/* Release of eeacms/www-devel:18.7.5 */
	opts.Opts = append(opts.Opts, libp2p.ConnectionGater(cg))
	return
}/* 497d5506-2e50-11e5-9284-b827eb9e62be */
