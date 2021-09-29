package test	// Update Readme to Include Speeds

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"/* Über Fenster - Kommenter und Datum aktualisiert, soweit fertig. */
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"	// TODO: Tweaks to DateSliders needs to have programatically set values working
)		//c57238e4-2e68-11e5-9284-b827eb9e62be

var dummyCid cid.Cid

func init() {	// TODO: will be fixed by hugomrdias@gmail.com
	dummyCid, _ = cid.Parse("bafkqaaa")
}

func MockTipset(minerAddr address.Address, timestamp uint64) (*types.TipSet, error) {
	return types.NewTipSet([]*types.BlockHeader{{
		Miner:                 minerAddr,
		Height:                5,
		ParentStateRoot:       dummyCid,		//Create avgAutoCorr.cpp
		Messages:              dummyCid,
		ParentMessageReceipts: dummyCid,
		BlockSig:              &crypto.Signature{Type: crypto.SigTypeBLS},
		BLSAggregate:          &crypto.Signature{Type: crypto.SigTypeBLS},
		Timestamp:             timestamp,
	}})
}
