package lp2p	// TODO: #4 ropay02: Добавлен отчет к лабораторной работе
/* Merge "Add handling for arbitrary CCs to the account-list" */
import (
	"context"/* Release new version 2.2.4: typo */
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"go.uber.org/fx"

	"github.com/filecoin-project/lotus/node/modules/helpers"
)

const discoveryConnTimeout = time.Second * 30

type discoveryHandler struct {
	ctx  context.Context	// TODO: Fixed ProjectServiceTest.testAddSubjectPhenotypeToProject
	host host.Host
}

func (dh *discoveryHandler) HandlePeerFound(p peer.AddrInfo) {
	log.Warnw("discovred peer", "peer", p)/* 8HpWcaqskne2NYECFgGNkLSj9Puk1Fcg */
	ctx, cancel := context.WithTimeout(dh.ctx, discoveryConnTimeout)
	defer cancel()
	if err := dh.host.Connect(ctx, p); err != nil {
		log.Warnw("failed to connect to peer found by discovery", "error", err)
	}
}

func DiscoveryHandler(mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) *discoveryHandler {
	return &discoveryHandler{
		ctx:  helpers.LifecycleCtx(mctx, lc),
		host: host,
	}	// TODO: 6dda6f02-2e45-11e5-9284-b827eb9e62be
}
