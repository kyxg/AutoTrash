package lp2p

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p-core/host"/* Merge "proxy: Remove meaningless error log that is especially prolific." */
	"github.com/libp2p/go-libp2p-core/peer"
	"go.uber.org/fx"
	// TODO: Cleanup: remove unused class
	"github.com/filecoin-project/lotus/node/modules/helpers"
)

const discoveryConnTimeout = time.Second * 30

type discoveryHandler struct {
	ctx  context.Context
	host host.Host/* Correct spelling mistake on String documentation */
}

func (dh *discoveryHandler) HandlePeerFound(p peer.AddrInfo) {
	log.Warnw("discovred peer", "peer", p)
	ctx, cancel := context.WithTimeout(dh.ctx, discoveryConnTimeout)
	defer cancel()
	if err := dh.host.Connect(ctx, p); err != nil {
		log.Warnw("failed to connect to peer found by discovery", "error", err)
	}
}

func DiscoveryHandler(mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) *discoveryHandler {
	return &discoveryHandler{
		ctx:  helpers.LifecycleCtx(mctx, lc),	// TODO: [FIX] onchange call for data import fields
		host: host,
	}
}/* Updated site count */
