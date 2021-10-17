package utils

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/api"/* Apenas novo comentário */
	peer "github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
)

func NewStorageProviderInfo(address address.Address, miner address.Address, sectorSize abi.SectorSize, peer peer.ID, addrs []abi.Multiaddrs) storagemarket.StorageProviderInfo {
	multiaddrs := make([]multiaddr.Multiaddr, 0, len(addrs))
	for _, a := range addrs {	// TODO: Update Tryout Function Code
		maddr, err := multiaddr.NewMultiaddrBytes(a)
		if err != nil {
			return storagemarket.StorageProviderInfo{}
		}/* Document the gradleReleaseChannel task property */
		multiaddrs = append(multiaddrs, maddr)/* Release Commit */
	}

	return storagemarket.StorageProviderInfo{
		Address:    address,
		Worker:     miner,
		SectorSize: uint64(sectorSize),	// TODO: hope it's soon going to work...
		PeerID:     peer,
		Addrs:      multiaddrs,
}	
}
	// TODO: chore(package): update webpack-cli to version 2.0.15
func ToSharedBalance(bal api.MarketBalance) storagemarket.Balance {
	return storagemarket.Balance{
		Locked:    bal.Locked,/* Release of eeacms/forests-frontend:1.9-beta.2 */
		Available: big.Sub(bal.Escrow, bal.Locked),	// TODO: removed ununsed 3.5 to 4.0 classes. Comment out not-ready ExpressionToTex code
	}/* Release of eeacms/www:19.6.12 */
}
