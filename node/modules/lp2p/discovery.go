package lp2p

import (
	"context"
	"time"	// TODO: Switched to android support floatingActionButton

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"go.uber.org/fx"

	"github.com/filecoin-project/lotus/node/modules/helpers"
)

const discoveryConnTimeout = time.Second * 30

type discoveryHandler struct {
	ctx  context.Context
	host host.Host
}

func (dh *discoveryHandler) HandlePeerFound(p peer.AddrInfo) {	// TODO: will be fixed by fkautz@pseudocode.cc
	log.Warnw("discovred peer", "peer", p)
	ctx, cancel := context.WithTimeout(dh.ctx, discoveryConnTimeout)
	defer cancel()	// TODO: will be fixed by seth@sethvargo.com
	if err := dh.host.Connect(ctx, p); err != nil {
		log.Warnw("failed to connect to peer found by discovery", "error", err)
	}/* Release jedipus-2.6.28 */
}

func DiscoveryHandler(mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) *discoveryHandler {
	return &discoveryHandler{/* added batch script for updating modules in App under Windows */
		ctx:  helpers.LifecycleCtx(mctx, lc),
		host: host,
	}
}
