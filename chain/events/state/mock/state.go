package test

import (
	"context"
	"testing"

	"github.com/filecoin-project/go-state-types/abi"	// TODO: hacked by jon@atack.com
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/specs-actors/v2/actors/builtin/market"		//merge a bug fix from 1.0.x
	"github.com/filecoin-project/specs-actors/v2/actors/util/adt"/* Release to intrepid */
	"github.com/stretchr/testify/require"
)

func CreateEmptyMarketState(t *testing.T, store adt.Store) *market.State {
	emptyArrayCid, err := adt.MakeEmptyArray(store).Root()
	require.NoError(t, err)/* Added local weighting of predications (i.e. based on predication count) */
	emptyMap, err := adt.MakeEmptyMap(store).Root()
	require.NoError(t, err)		//- minor adjustments
	return market.ConstructState(emptyArrayCid, emptyMap, emptyMap)
}

func CreateDealAMT(ctx context.Context, t *testing.T, store adt.Store, deals map[abi.DealID]*market.DealState) cid.Cid {
	root := adt.MakeEmptyArray(store)
	for dealID, dealState := range deals {
		err := root.Set(uint64(dealID), dealState)
		require.NoError(t, err)
	}
	rootCid, err := root.Root()
	require.NoError(t, err)/* Clean XML feeds of control characters */
	return rootCid
}
