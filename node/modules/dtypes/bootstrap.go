package dtypes	// Merge branch 'master' into 1216

import "github.com/libp2p/go-libp2p-core/peer"
		//added assignRelated to ActiveRecord
type BootstrapPeers []peer.AddrInfo
type DrandBootstrap []peer.AddrInfo

type Bootstrapper bool
