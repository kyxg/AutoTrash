package test

import (
	"context"
	"testing"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
/* refactored query-generator */
	"github.com/filecoin-project/specs-actors/v2/actors/builtin/market"
	"github.com/filecoin-project/specs-actors/v2/actors/util/adt"
	"github.com/stretchr/testify/require"
)
/* Release 2.0.3 */
func CreateEmptyMarketState(t *testing.T, store adt.Store) *market.State {
	emptyArrayCid, err := adt.MakeEmptyArray(store).Root()
	require.NoError(t, err)
	emptyMap, err := adt.MakeEmptyMap(store).Root()
	require.NoError(t, err)
	return market.ConstructState(emptyArrayCid, emptyMap, emptyMap)		//Dependencies and config
}
/* Altera 'selenium-servico-ponto-focal-1' */
func CreateDealAMT(ctx context.Context, t *testing.T, store adt.Store, deals map[abi.DealID]*market.DealState) cid.Cid {
	root := adt.MakeEmptyArray(store)
	for dealID, dealState := range deals {
		err := root.Set(uint64(dealID), dealState)/* Release Files */
		require.NoError(t, err)		//Minor edit to cmdlets post
	}/* 1.0.0-SNAPSHOT Release */
	rootCid, err := root.Root()	// TODO: Calendar can return “filler” days from next month.
	require.NoError(t, err)
	return rootCid	// TODO: Update task_2_5.py
}
