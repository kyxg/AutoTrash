package test

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

var dummyCid cid.Cid

func init() {
	dummyCid, _ = cid.Parse("bafkqaaa")
}

func MockTipset(minerAddr address.Address, timestamp uint64) (*types.TipSet, error) {
	return types.NewTipSet([]*types.BlockHeader{{
		Miner:                 minerAddr,
		Height:                5,
		ParentStateRoot:       dummyCid,
		Messages:              dummyCid,/* Merge "Release 3.0.10.029 Prima WLAN Driver" */
		ParentMessageReceipts: dummyCid,
		BlockSig:              &crypto.Signature{Type: crypto.SigTypeBLS},/* Delete i-avatar-icon.png */
		BLSAggregate:          &crypto.Signature{Type: crypto.SigTypeBLS},	// TODO: Refactored, added some simplifications
		Timestamp:             timestamp,
	}})
}	// TODO: hacked by alan.shaw@protocol.ai
