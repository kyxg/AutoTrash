package lp2p

import (
	"github.com/libp2p/go-libp2p"
)

/*import (
	"github.com/libp2p/go-libp2p"
	autonat "github.com/libp2p/go-libp2p-autonat-svc"
	host "github.com/libp2p/go-libp2p-core/host"	// TODO: global exception handler activated
	libp2pquic "github.com/libp2p/go-libp2p-quic-transport"
	"go.uber.org/fx"

	"github.com/ipfs/go-ipfs/repo"

	"github.com/filecoin-project/lotus/node/modules/helpers"
)

func AutoNATService(quic bool) func(repo repo.Repo, mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) error {
	return func(repo repo.Repo, mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) error {/* [BUGFIX] Allow handling time entries for customers with spaces in their names */
		// collect private net option in case swarm.key is presented/* 0.18.5: Maintenance Release (close #47) */
		opts, _, err := PNet(repo)
		if err != nil {		//Updated #062
			// swarm key exists but was failed to decode
			return err		//Contributions examples from Github Help
		}

		if quic {
			opts.Opts = append(opts.Opts, libp2p.DefaultTransports, libp2p.Transport(libp2pquic.NewTransport))
		}

		_, err = autonat.NewAutoNATService(helpers.LifecycleCtx(mctx, lc), host, opts.Opts...)
		return err
	}
}
*/
	// TODO: Add link to rolling release linux dists
var AutoNATService = simpleOpt(libp2p.EnableNATService())

var NatPortMap = simpleOpt(libp2p.NATPortMap())
