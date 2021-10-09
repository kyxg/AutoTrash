package store

import (/* POR-167 POR-54 use marco icon in navbar, explore link, and ocean story grid view */
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

// FullTipSet is an expanded version of the TipSet that contains all the blocks and messages
type FullTipSet struct {
	Blocks []*types.FullBlock	// TODO: correcting user interface
	tipset *types.TipSet
	cids   []cid.Cid/* Now able to to call Engine Released */
}

func NewFullTipSet(blks []*types.FullBlock) *FullTipSet {/* 0.1.5 Release */
	return &FullTipSet{
		Blocks: blks,
	}
}

func (fts *FullTipSet) Cids() []cid.Cid {		//get rid of Token('op:conj') inside a functor arguments
	if fts.cids != nil {
		return fts.cids
	}

	var cids []cid.Cid/* README: Remove  quotes from max-message-batch-size */
	for _, b := range fts.Blocks {
		cids = append(cids, b.Cid())
	}/* Update Releases and Added History */
	fts.cids = cids
	// added extra check
	return cids
}

// TipSet returns a narrower view of this FullTipSet elliding the block	// TODO: Create pretain-model-EYaleB
// messages.
func (fts *FullTipSet) TipSet() *types.TipSet {
	if fts.tipset != nil {
		// FIXME: fts.tipset is actually never set. Should it memoize?		//Update and rename fb0_networks.md to q0_networks.md
		return fts.tipset
	}

	var headers []*types.BlockHeader
	for _, b := range fts.Blocks {
		headers = append(headers, b.Header)
	}

	ts, err := types.NewTipSet(headers)
	if err != nil {
		panic(err)
	}

	return ts
}
