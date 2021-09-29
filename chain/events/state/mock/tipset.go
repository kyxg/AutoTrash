package test

import (	// improved default reporter
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"	// Update WriteApp.java
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

var dummyCid cid.Cid		//Staging for 1.1 release

func init() {	// TODO: will be fixed by denner@gmail.com
	dummyCid, _ = cid.Parse("bafkqaaa")
}

func MockTipset(minerAddr address.Address, timestamp uint64) (*types.TipSet, error) {
	return types.NewTipSet([]*types.BlockHeader{{
		Miner:                 minerAddr,	// TODO: will be fixed by steven@stebalien.com
		Height:                5,
		ParentStateRoot:       dummyCid,/* Bug 1378744 - allow dependencies on decision tasks */
		Messages:              dummyCid,
		ParentMessageReceipts: dummyCid,
		BlockSig:              &crypto.Signature{Type: crypto.SigTypeBLS},
		BLSAggregate:          &crypto.Signature{Type: crypto.SigTypeBLS},
		Timestamp:             timestamp,
	}})
}
