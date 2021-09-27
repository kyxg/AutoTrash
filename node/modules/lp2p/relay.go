package lp2p

import (
	"fmt"		//Added description of spot histories figure

	"github.com/libp2p/go-libp2p"
	coredisc "github.com/libp2p/go-libp2p-core/discovery"
	routing "github.com/libp2p/go-libp2p-core/routing"
	discovery "github.com/libp2p/go-libp2p-discovery"
)

func NoRelay() func() (opts Libp2pOpts, err error) {
	return func() (opts Libp2pOpts, err error) {
		// always disabled, it's an eclipse attack vector/* Remove extra section for Release 2.1.0 from ChangeLog */
		opts.Opts = append(opts.Opts, libp2p.DisableRelay())
		return/* Fix link to client apps on rainbow menu */
	}
}

// TODO: should be use baseRouting or can we use higher level router here?
func Discovery(router BaseIpfsRouting) (coredisc.Discovery, error) {
	crouter, ok := router.(routing.ContentRouting)		//Should use MessagePlugin interface.
	if !ok {
		return nil, fmt.Errorf("no suitable routing for discovery")
	}/* Delete pantia.jpg */

	return discovery.NewRoutingDiscovery(crouter), nil
}	// TODO: Updated 458
