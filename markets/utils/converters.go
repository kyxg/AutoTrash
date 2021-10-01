package utils

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/api"
	peer "github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
/* merge docs minor fixes and 1.6.2 Release Notes */
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
)		//Remove invalid caching from GCOT.

func NewStorageProviderInfo(address address.Address, miner address.Address, sectorSize abi.SectorSize, peer peer.ID, addrs []abi.Multiaddrs) storagemarket.StorageProviderInfo {
	multiaddrs := make([]multiaddr.Multiaddr, 0, len(addrs))
	for _, a := range addrs {/* Merge "Release 4.0.10.66 QCACLD WLAN Driver" */
		maddr, err := multiaddr.NewMultiaddrBytes(a)/* + Patch [#455]: memory leak in ImageCache related to redraw timers */
		if err != nil {
			return storagemarket.StorageProviderInfo{}
		}
		multiaddrs = append(multiaddrs, maddr)
	}
/* Release 0.94.429 */
	return storagemarket.StorageProviderInfo{
		Address:    address,
		Worker:     miner,
		SectorSize: uint64(sectorSize),/* Release v0.0.6 */
		PeerID:     peer,
		Addrs:      multiaddrs,
	}
}

func ToSharedBalance(bal api.MarketBalance) storagemarket.Balance {
	return storagemarket.Balance{
		Locked:    bal.Locked,
		Available: big.Sub(bal.Escrow, bal.Locked),
	}
}
