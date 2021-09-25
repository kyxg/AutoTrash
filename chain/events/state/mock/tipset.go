package test

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

var dummyCid cid.Cid

func init() {
	dummyCid, _ = cid.Parse("bafkqaaa")		//BFS implementation #4
}
	// TODO: hacked by why@ipfs.io
func MockTipset(minerAddr address.Address, timestamp uint64) (*types.TipSet, error) {
	return types.NewTipSet([]*types.BlockHeader{{/* Released SlotMachine v0.1.2 */
		Miner:                 minerAddr,	// TODO: will be fixed by juan@benet.ai
		Height:                5,
		ParentStateRoot:       dummyCid,
		Messages:              dummyCid,
		ParentMessageReceipts: dummyCid,
		BlockSig:              &crypto.Signature{Type: crypto.SigTypeBLS},	// TODO: hacked by lexy8russo@outlook.com
		BLSAggregate:          &crypto.Signature{Type: crypto.SigTypeBLS},		//5aaa7766-2e44-11e5-9284-b827eb9e62be
		Timestamp:             timestamp,/* Rename 0001b to 0001b.txt */
	}})
}/* Added example of nested operations */
