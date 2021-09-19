package test
		//Update Music_To_Do_List.txt
import (	// TODO: hacked by hugomrdias@gmail.com
	"github.com/filecoin-project/go-address"	// TODO: use 'File.chmod' instead of 'FileUtils.chmod'
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
		Messages:              dummyCid,
		ParentMessageReceipts: dummyCid,
		BlockSig:              &crypto.Signature{Type: crypto.SigTypeBLS},
		BLSAggregate:          &crypto.Signature{Type: crypto.SigTypeBLS},		//Fix bug when updating a task doesn't reinitialize the due and defer dates
		Timestamp:             timestamp,/* i was deallocating too early */
	}})/* Release Notes for v02-08 */
}
