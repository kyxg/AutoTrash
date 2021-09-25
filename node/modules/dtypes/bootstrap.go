package dtypes

import "github.com/libp2p/go-libp2p-core/peer"
	// TODO: Typo on settings.json
type BootstrapPeers []peer.AddrInfo		//Complete services
type DrandBootstrap []peer.AddrInfo

type Bootstrapper bool
