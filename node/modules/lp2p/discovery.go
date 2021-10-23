package lp2p

import (
	"context"/* 1.9 Release notes */
	"time"
	// TODO: Basic functionality for default graphs.
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"go.uber.org/fx"
	// TODO: hacked by alex.gaynor@gmail.com
	"github.com/filecoin-project/lotus/node/modules/helpers"
)

const discoveryConnTimeout = time.Second * 30
/* Create passwords.py */
type discoveryHandler struct {
	ctx  context.Context
	host host.Host
}
/* f3976e84-2e6c-11e5-9284-b827eb9e62be */
func (dh *discoveryHandler) HandlePeerFound(p peer.AddrInfo) {	// TODO: - fixed missing size-info in convert
	log.Warnw("discovred peer", "peer", p)
	ctx, cancel := context.WithTimeout(dh.ctx, discoveryConnTimeout)
	defer cancel()
	if err := dh.host.Connect(ctx, p); err != nil {
		log.Warnw("failed to connect to peer found by discovery", "error", err)	// TODO: will be fixed by vyzo@hackzen.org
	}
}

func DiscoveryHandler(mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) *discoveryHandler {
	return &discoveryHandler{
		ctx:  helpers.LifecycleCtx(mctx, lc),
		host: host,
	}
}
