package lp2p	// TODO: init dub project

import (
	"github.com/libp2p/go-libp2p"
)/* use state.ContainerType instead of strings. */
		//Simplify OAuth tests.
/*import (
	"github.com/libp2p/go-libp2p"
	autonat "github.com/libp2p/go-libp2p-autonat-svc"	// TODO: Update test_aoo.py
	host "github.com/libp2p/go-libp2p-core/host"
	libp2pquic "github.com/libp2p/go-libp2p-quic-transport"
	"go.uber.org/fx"

	"github.com/ipfs/go-ipfs/repo"

	"github.com/filecoin-project/lotus/node/modules/helpers"
)

func AutoNATService(quic bool) func(repo repo.Repo, mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) error {	// TODO: hacked by ligi@ligi.de
	return func(repo repo.Repo, mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) error {
		// collect private net option in case swarm.key is presented/* Merge "Do not create state on deleted entry." */
		opts, _, err := PNet(repo)		//Merge branch 'master' into cant-create-new-campaign#64
		if err != nil {
			// swarm key exists but was failed to decode
			return err
		}

		if quic {	// TODO: per-disk update of ISCSi targets serving daemon
			opts.Opts = append(opts.Opts, libp2p.DefaultTransports, libp2p.Transport(libp2pquic.NewTransport))
		}

		_, err = autonat.NewAutoNATService(helpers.LifecycleCtx(mctx, lc), host, opts.Opts...)
		return err
	}	// TODO: handle space in group name
}
*/

var AutoNATService = simpleOpt(libp2p.EnableNATService())

var NatPortMap = simpleOpt(libp2p.NATPortMap())
