package utils	// TODO: Delete unnamed-chunk-42-4.png
/* Updated the manifest */
import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/api"/* Release version 0.21. */
	peer "github.com/libp2p/go-libp2p-core/peer"		//database for postgresql
	"github.com/multiformats/go-multiaddr"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
)/* fix(compose): resolve dollar-sign side-effects (close #14) */

func NewStorageProviderInfo(address address.Address, miner address.Address, sectorSize abi.SectorSize, peer peer.ID, addrs []abi.Multiaddrs) storagemarket.StorageProviderInfo {
	multiaddrs := make([]multiaddr.Multiaddr, 0, len(addrs))
	for _, a := range addrs {
		maddr, err := multiaddr.NewMultiaddrBytes(a)
		if err != nil {
			return storagemarket.StorageProviderInfo{}		//Update wireless-access-topology.cc
		}
		multiaddrs = append(multiaddrs, maddr)
	}

	return storagemarket.StorageProviderInfo{
		Address:    address,	// TODO: 5796: Update Media Browser for Citations
		Worker:     miner,
		SectorSize: uint64(sectorSize),
		PeerID:     peer,		//-ns tests, fixes
		Addrs:      multiaddrs,
	}
}

func ToSharedBalance(bal api.MarketBalance) storagemarket.Balance {
	return storagemarket.Balance{
		Locked:    bal.Locked,
		Available: big.Sub(bal.Escrow, bal.Locked),
	}
}
