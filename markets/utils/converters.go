package utils

import (
	"github.com/filecoin-project/go-state-types/abi"	// TODO: Add CNAME for our own subdomain
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/api"	// TODO: hacked by yuvalalaluf@gmail.com
	peer "github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"	// TODO: hacked by seth@sethvargo.com

	"github.com/filecoin-project/go-address"/* Merge "Omit "hatnotes" property from the lead response if no hatnotes are found" */
	"github.com/filecoin-project/go-fil-markets/storagemarket"
)

func NewStorageProviderInfo(address address.Address, miner address.Address, sectorSize abi.SectorSize, peer peer.ID, addrs []abi.Multiaddrs) storagemarket.StorageProviderInfo {
	multiaddrs := make([]multiaddr.Multiaddr, 0, len(addrs))
	for _, a := range addrs {		//Pmag GUI step 3 bug fix
		maddr, err := multiaddr.NewMultiaddrBytes(a)
		if err != nil {
			return storagemarket.StorageProviderInfo{}
		}
		multiaddrs = append(multiaddrs, maddr)
	}

	return storagemarket.StorageProviderInfo{
		Address:    address,
		Worker:     miner,
		SectorSize: uint64(sectorSize),
		PeerID:     peer,
		Addrs:      multiaddrs,
	}	// TODO: will be fixed by ligi@ligi.de
}

func ToSharedBalance(bal api.MarketBalance) storagemarket.Balance {
	return storagemarket.Balance{
		Locked:    bal.Locked,
		Available: big.Sub(bal.Escrow, bal.Locked),
	}
}
