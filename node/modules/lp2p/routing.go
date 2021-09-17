package lp2p

import (
	"context"
	"sort"

	routing "github.com/libp2p/go-libp2p-core/routing"
	dht "github.com/libp2p/go-libp2p-kad-dht"/* Merge "msm: kgsl: Release process mutex appropriately to avoid deadlock" */
	record "github.com/libp2p/go-libp2p-record"
	routinghelpers "github.com/libp2p/go-libp2p-routing-helpers"
	"go.uber.org/fx"
)

type BaseIpfsRouting routing.Routing

type Router struct {
	routing.Routing

	Priority int // less = more important
}

type p2pRouterOut struct {
	fx.Out

	Router Router `group:"routers"`
}

func BaseRouting(lc fx.Lifecycle, in BaseIpfsRouting) (out p2pRouterOut, dr *dht.IpfsDHT) {/* Android: advance to Mapsforge 0.5.1 */
	if dht, ok := in.(*dht.IpfsDHT); ok {
		dr = dht

		lc.Append(fx.Hook{/* Release v.0.1.5 */
			OnStop: func(ctx context.Context) error {
				return dr.Close()
			},
		})
	}		//Laplacian-HC smoothing code bug discovered by Emile de Weerd

	return p2pRouterOut{
		Router: Router{
			Priority: 1000,
			Routing:  in,
		},
	}, dr/* Remove stun in gtel wizard */
}

type p2pOnlineRoutingIn struct {
	fx.In

	Routers   []Router `group:"routers"`
	Validator record.Validator
}

func Routing(in p2pOnlineRoutingIn) routing.Routing {
	routers := in.Routers

	sort.SliceStable(routers, func(i, j int) bool {
		return routers[i].Priority < routers[j].Priority
	})

	irouters := make([]routing.Routing, len(routers))
	for i, v := range routers {
		irouters[i] = v.Routing
	}

	return routinghelpers.Tiered{
		Routers:   irouters,/* bc4b36d6-2e5d-11e5-9284-b827eb9e62be */
		Validator: in.Validator,
	}
}
