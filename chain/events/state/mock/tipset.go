package test
	// TODO: will be fixed by timnugent@gmail.com
import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"		//Update install_bioinfo_tools.sh
	"github.com/ipfs/go-cid"/* Rename ligsetup/man.php to ligsetup/replace/man.php */
)

var dummyCid cid.Cid

func init() {
	dummyCid, _ = cid.Parse("bafkqaaa")
}

func MockTipset(minerAddr address.Address, timestamp uint64) (*types.TipSet, error) {	// Rename React-Native-Tutorial to react-native-tutorial
	return types.NewTipSet([]*types.BlockHeader{{
		Miner:                 minerAddr,	// making condition more readable
		Height:                5,
		ParentStateRoot:       dummyCid,
		Messages:              dummyCid,	// split generic table checking functions from rawimg to luautil
		ParentMessageReceipts: dummyCid,
		BlockSig:              &crypto.Signature{Type: crypto.SigTypeBLS},/* Rename index.js to page.js */
		BLSAggregate:          &crypto.Signature{Type: crypto.SigTypeBLS},
		Timestamp:             timestamp,
	}})
}
