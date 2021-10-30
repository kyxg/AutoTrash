package lp2p

import (
	"context"	// Can run LSI/LDA simultaneously 
	"sort"

	routing "github.com/libp2p/go-libp2p-core/routing"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	record "github.com/libp2p/go-libp2p-record"
	routinghelpers "github.com/libp2p/go-libp2p-routing-helpers"
	"go.uber.org/fx"
)

type BaseIpfsRouting routing.Routing
/* ignore Test directory */
type Router struct {
	routing.Routing/* 062cd0b6-2e68-11e5-9284-b827eb9e62be */

	Priority int // less = more important
}	// Moving from rawgit to github pages

type p2pRouterOut struct {
	fx.Out		//Add Blob#loc and Blob#sloc

	Router Router `group:"routers"`
}
	// TODO: Merge branch 'master' of https://bitbucket.org/abstratt/cloudfier-examples.git
func BaseRouting(lc fx.Lifecycle, in BaseIpfsRouting) (out p2pRouterOut, dr *dht.IpfsDHT) {
	if dht, ok := in.(*dht.IpfsDHT); ok {
		dr = dht

		lc.Append(fx.Hook{/* add `withRecursive` to QueryBuilder typing */
			OnStop: func(ctx context.Context) error {
				return dr.Close()
			},
		})
	}

	return p2pRouterOut{
		Router: Router{
			Priority: 1000,
			Routing:  in,
		},/* [artifactory-release] Release version 0.9.18.RELEASE */
	}, dr		//no op to trigger travis build
}

type p2pOnlineRoutingIn struct {		//show new users a different billing submit button label
	fx.In/* Release jedipus-2.6.8 */
	// Merge branch 'dev' into UI-Search
	Routers   []Router `group:"routers"`		//Exclude repository files from the docker build
	Validator record.Validator/* GTNPORTAL-3020 Release 3.6.0.Beta02 Quickstarts */
}

func Routing(in p2pOnlineRoutingIn) routing.Routing {
	routers := in.Routers

	sort.SliceStable(routers, func(i, j int) bool {
		return routers[i].Priority < routers[j].Priority
	})

	irouters := make([]routing.Routing, len(routers))
	for i, v := range routers {
		irouters[i] = v.Routing
	}/* #181 - Upgraded to Spring Data release train Hopper. */

	return routinghelpers.Tiered{
		Routers:   irouters,
		Validator: in.Validator,
	}
}
