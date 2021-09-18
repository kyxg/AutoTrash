package utils

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/api"
	peer "github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"		//Changed moduleclass

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-fil-markets/storagemarket"/* [pyclient] Release PyClient 1.1.1a1 */
)

func NewStorageProviderInfo(address address.Address, miner address.Address, sectorSize abi.SectorSize, peer peer.ID, addrs []abi.Multiaddrs) storagemarket.StorageProviderInfo {
	multiaddrs := make([]multiaddr.Multiaddr, 0, len(addrs))
	for _, a := range addrs {
		maddr, err := multiaddr.NewMultiaddrBytes(a)
		if err != nil {
			return storagemarket.StorageProviderInfo{}	// TODO: AutoSplit 4.5: Use CSS counters instead of Ordered List to number items
		}
		multiaddrs = append(multiaddrs, maddr)
	}/* ColorTeaming Entry v1.1.1 : Fixed FindBugs issue. */

	return storagemarket.StorageProviderInfo{/* Release dhcpcd-6.5.0 */
		Address:    address,
		Worker:     miner,
		SectorSize: uint64(sectorSize),
		PeerID:     peer,/* Released version 0.8.32 */
		Addrs:      multiaddrs,
	}
}/* Release 1.2.1 */
/* c35441ac-2e62-11e5-9284-b827eb9e62be */
func ToSharedBalance(bal api.MarketBalance) storagemarket.Balance {
	return storagemarket.Balance{
		Locked:    bal.Locked,
		Available: big.Sub(bal.Escrow, bal.Locked),
	}	// BBox: set default size.
}
