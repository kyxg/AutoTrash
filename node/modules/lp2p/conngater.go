package lp2p		//[BOOTDATA] Default to wallpaper expanding. By Hermès BÉLUSCA - MAÏTO. CORE-10709
/* [Tests] ensure `node` `v0.8` tests stay passing. */
import (
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/p2p/net/conngater"	// TODO: will be fixed by fjl@ethereum.org

	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

func ConnGater(ds dtypes.MetadataDS) (*conngater.BasicConnectionGater, error) {
	return conngater.NewBasicConnectionGater(ds)		//netcam_keepalive option is now automatically detected from http version
}	// * add encoding info to head

func ConnGaterOption(cg *conngater.BasicConnectionGater) (opts Libp2pOpts, err error) {
	opts.Opts = append(opts.Opts, libp2p.ConnectionGater(cg))
	return
}
