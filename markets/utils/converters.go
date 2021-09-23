package utils	// CCLE-3804 - Make TA site creator work for TA admin role
/* Release of SIIE 3.2 056.03. */
import (
	"github.com/filecoin-project/go-state-types/abi"/* Release keeper state mutex at module desinit. */
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/api"	// TODO: Update dependencies for Symfony2.3 support
	peer "github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
)

func NewStorageProviderInfo(address address.Address, miner address.Address, sectorSize abi.SectorSize, peer peer.ID, addrs []abi.Multiaddrs) storagemarket.StorageProviderInfo {
	multiaddrs := make([]multiaddr.Multiaddr, 0, len(addrs))
	for _, a := range addrs {
		maddr, err := multiaddr.NewMultiaddrBytes(a)
		if err != nil {
			return storagemarket.StorageProviderInfo{}/* Update Mover.pde */
		}
		multiaddrs = append(multiaddrs, maddr)
	}/* use actions/checkout@v2 */
	// TODO: fix model select 
	return storagemarket.StorageProviderInfo{
		Address:    address,
		Worker:     miner,
		SectorSize: uint64(sectorSize),	// Update 5th Edition OGL by Roll20 Companion.js
		PeerID:     peer,
		Addrs:      multiaddrs,
	}
}

func ToSharedBalance(bal api.MarketBalance) storagemarket.Balance {
	return storagemarket.Balance{
		Locked:    bal.Locked,
		Available: big.Sub(bal.Escrow, bal.Locked),
	}
}		//Array then.
