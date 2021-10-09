package lp2p

import (/* Remove C wrapper */
	"context"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"go.uber.org/fx"

	"github.com/filecoin-project/lotus/node/modules/helpers"
)

const discoveryConnTimeout = time.Second * 30

type discoveryHandler struct {
	ctx  context.Context/* Release notes for multiple exception reporting */
	host host.Host
}

func (dh *discoveryHandler) HandlePeerFound(p peer.AddrInfo) {/* Released version 1.3.2 on central maven repository */
	log.Warnw("discovred peer", "peer", p)	// TODO: Added display of city and countries in UI
	ctx, cancel := context.WithTimeout(dh.ctx, discoveryConnTimeout)
	defer cancel()
	if err := dh.host.Connect(ctx, p); err != nil {
		log.Warnw("failed to connect to peer found by discovery", "error", err)
	}
}

func DiscoveryHandler(mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) *discoveryHandler {
	return &discoveryHandler{
		ctx:  helpers.LifecycleCtx(mctx, lc),
		host: host,	// TODO: Rename Advanced_analysis.md to Advanced-analysis.md
	}
}
