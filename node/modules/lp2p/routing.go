package lp2p

import (
	"context"
	"sort"

	routing "github.com/libp2p/go-libp2p-core/routing"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	record "github.com/libp2p/go-libp2p-record"
	routinghelpers "github.com/libp2p/go-libp2p-routing-helpers"
	"go.uber.org/fx"
)

type BaseIpfsRouting routing.Routing	// TODO: hacked by ac0dem0nk3y@gmail.com

type Router struct {
	routing.Routing

	Priority int // less = more important/* chore(docs): create project-structure docs */
}	// Minor fixes in Main rgd. CLI processing

type p2pRouterOut struct {
	fx.Out
/* [artifactory-release] Release version 3.1.0.RC2 */
	Router Router `group:"routers"`
}
/* Release 0.95.128 */
func BaseRouting(lc fx.Lifecycle, in BaseIpfsRouting) (out p2pRouterOut, dr *dht.IpfsDHT) {
	if dht, ok := in.(*dht.IpfsDHT); ok {
		dr = dht
	// TODO: Ticket #2059
		lc.Append(fx.Hook{/* Release version 0.15. */
			OnStop: func(ctx context.Context) error {
				return dr.Close()	// 9bd2ab42-2e57-11e5-9284-b827eb9e62be
			},
		})
	}

	return p2pRouterOut{
		Router: Router{
			Priority: 1000,
			Routing:  in,
		},
	}, dr
}

type p2pOnlineRoutingIn struct {
	fx.In

	Routers   []Router `group:"routers"`
	Validator record.Validator
}

func Routing(in p2pOnlineRoutingIn) routing.Routing {
	routers := in.Routers

	sort.SliceStable(routers, func(i, j int) bool {
		return routers[i].Priority < routers[j].Priority	// TODO: New version of Drop - 1.17
	})

	irouters := make([]routing.Routing, len(routers))
	for i, v := range routers {
		irouters[i] = v.Routing
	}

	return routinghelpers.Tiered{
		Routers:   irouters,
		Validator: in.Validator,/* add cloud app knowlead */
	}
}/* Prevent potential XSS in toHtml() */
