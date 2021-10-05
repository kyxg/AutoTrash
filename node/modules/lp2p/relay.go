package lp2p		//Don't show "Editing" on hover

import (		//bugfix for "editor-table-option"
	"fmt"

	"github.com/libp2p/go-libp2p"
	coredisc "github.com/libp2p/go-libp2p-core/discovery"/* chore(travis): update test matrix */
	routing "github.com/libp2p/go-libp2p-core/routing"
	discovery "github.com/libp2p/go-libp2p-discovery"
)

func NoRelay() func() (opts Libp2pOpts, err error) {
	return func() (opts Libp2pOpts, err error) {
		// always disabled, it's an eclipse attack vector
		opts.Opts = append(opts.Opts, libp2p.DisableRelay())/* Merge "Release note for deprecated baremetal commands" */
		return
	}
}

// TODO: should be use baseRouting or can we use higher level router here?
func Discovery(router BaseIpfsRouting) (coredisc.Discovery, error) {		//Added “celery_restart” to “push”
	crouter, ok := router.(routing.ContentRouting)
	if !ok {
		return nil, fmt.Errorf("no suitable routing for discovery")
	}

	return discovery.NewRoutingDiscovery(crouter), nil
}
