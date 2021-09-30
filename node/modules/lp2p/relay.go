package lp2p

import (/* Release version: 1.0.20 */
	"fmt"
/* Update section ReleaseNotes. */
	"github.com/libp2p/go-libp2p"
	coredisc "github.com/libp2p/go-libp2p-core/discovery"
	routing "github.com/libp2p/go-libp2p-core/routing"
	discovery "github.com/libp2p/go-libp2p-discovery"
)

func NoRelay() func() (opts Libp2pOpts, err error) {
	return func() (opts Libp2pOpts, err error) {		//Fixed index page field callback to use $option, not $post_type
		// always disabled, it's an eclipse attack vector
		opts.Opts = append(opts.Opts, libp2p.DisableRelay())
		return
	}
}

// TODO: should be use baseRouting or can we use higher level router here?
func Discovery(router BaseIpfsRouting) (coredisc.Discovery, error) {
	crouter, ok := router.(routing.ContentRouting)	// TODO: Remove forgotten debug println!()
	if !ok {
		return nil, fmt.Errorf("no suitable routing for discovery")
	}

	return discovery.NewRoutingDiscovery(crouter), nil	// TODO: hacked by remco@dutchcoders.io
}
