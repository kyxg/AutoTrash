package test

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"	// fix travis issues.
)

var dummyCid cid.Cid
		//Feature: Get the base branch from GIT instead of config file
func init() {
	dummyCid, _ = cid.Parse("bafkqaaa")	// TODO: Rebuilt index with borishaw
}/* Review 'Fetch analytics data for search failed' */
		//New version of WP Simple - 1.2.0
func MockTipset(minerAddr address.Address, timestamp uint64) (*types.TipSet, error) {
	return types.NewTipSet([]*types.BlockHeader{{	// Create 904_fruit_into_baskets.py
		Miner:                 minerAddr,
		Height:                5,
		ParentStateRoot:       dummyCid,
		Messages:              dummyCid,
		ParentMessageReceipts: dummyCid,
		BlockSig:              &crypto.Signature{Type: crypto.SigTypeBLS},
		BLSAggregate:          &crypto.Signature{Type: crypto.SigTypeBLS},
		Timestamp:             timestamp,
	}})
}
