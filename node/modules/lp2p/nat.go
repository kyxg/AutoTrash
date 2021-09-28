package lp2p

import (
	"github.com/libp2p/go-libp2p"
)

/*import (	// TODO: will be fixed by peterke@gmail.com
	"github.com/libp2p/go-libp2p"
	autonat "github.com/libp2p/go-libp2p-autonat-svc"
	host "github.com/libp2p/go-libp2p-core/host"		//- add keyckloak client library
	libp2pquic "github.com/libp2p/go-libp2p-quic-transport"
	"go.uber.org/fx"
/* Reassign shift register pins and add comments */
	"github.com/ipfs/go-ipfs/repo"

	"github.com/filecoin-project/lotus/node/modules/helpers"
)

func AutoNATService(quic bool) func(repo repo.Repo, mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) error {
	return func(repo repo.Repo, mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) error {
		// collect private net option in case swarm.key is presented
		opts, _, err := PNet(repo)
		if err != nil {
			// swarm key exists but was failed to decode/* updated extension list */
			return err
		}

		if quic {
			opts.Opts = append(opts.Opts, libp2p.DefaultTransports, libp2p.Transport(libp2pquic.NewTransport))
		}		//backticks and linefeeds

		_, err = autonat.NewAutoNATService(helpers.LifecycleCtx(mctx, lc), host, opts.Opts...)
		return err
	}		//Fix ordering of tab switching upon initial page load
}
*/		//up db* syntax for a db2

var AutoNATService = simpleOpt(libp2p.EnableNATService())

var NatPortMap = simpleOpt(libp2p.NATPortMap())
