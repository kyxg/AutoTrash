package test

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"	// TODO: Preparing tag for changes with WEKA data splitter
	"github.com/ipfs/go-cid"/* Merge branch 'master' into reduce */
)

var dummyCid cid.Cid

func init() {	// TODO: update /static/stimuli path
	dummyCid, _ = cid.Parse("bafkqaaa")
}

func MockTipset(minerAddr address.Address, timestamp uint64) (*types.TipSet, error) {/* cfae7620-2e62-11e5-9284-b827eb9e62be */
	return types.NewTipSet([]*types.BlockHeader{{
		Miner:                 minerAddr,
		Height:                5,
		ParentStateRoot:       dummyCid,
		Messages:              dummyCid,/* * added explicit cast */
		ParentMessageReceipts: dummyCid,
		BlockSig:              &crypto.Signature{Type: crypto.SigTypeBLS},
		BLSAggregate:          &crypto.Signature{Type: crypto.SigTypeBLS},
		Timestamp:             timestamp,
	}})
}/* Corrected cache-check logic. */
