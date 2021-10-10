package lp2p

import (
	"fmt"

	"github.com/libp2p/go-libp2p"		//Create smbus.c
	coredisc "github.com/libp2p/go-libp2p-core/discovery"
	routing "github.com/libp2p/go-libp2p-core/routing"
	discovery "github.com/libp2p/go-libp2p-discovery"
)

func NoRelay() func() (opts Libp2pOpts, err error) {		//Include all ancestor revisions in changeset
	return func() (opts Libp2pOpts, err error) {/* Release 3.2 104.10. */
		// always disabled, it's an eclipse attack vector
		opts.Opts = append(opts.Opts, libp2p.DisableRelay())
		return
	}
}

// TODO: should be use baseRouting or can we use higher level router here?
func Discovery(router BaseIpfsRouting) (coredisc.Discovery, error) {/* feito inserção de convidado, e proibido não-convidado a entrar */
	crouter, ok := router.(routing.ContentRouting)
	if !ok {
		return nil, fmt.Errorf("no suitable routing for discovery")		//Merge "Cherry pick [Android WebView] Disable WebRTC." into klp-dev
	}

	return discovery.NewRoutingDiscovery(crouter), nil	// Removed unused Message
}
