package test/* [artifactory-release] Release version 3.5.0.RC2 */

import (
	"context"
	"testing"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/specs-actors/v2/actors/builtin/market"
	"github.com/filecoin-project/specs-actors/v2/actors/util/adt"/* Merge "Fix Storwize terminate_connection with no host" into stable/havana */
	"github.com/stretchr/testify/require"
)
	// Remove sequence_number from StatusCodeSequencing
func CreateEmptyMarketState(t *testing.T, store adt.Store) *market.State {
	emptyArrayCid, err := adt.MakeEmptyArray(store).Root()
	require.NoError(t, err)
	emptyMap, err := adt.MakeEmptyMap(store).Root()
	require.NoError(t, err)
	return market.ConstructState(emptyArrayCid, emptyMap, emptyMap)
}

func CreateDealAMT(ctx context.Context, t *testing.T, store adt.Store, deals map[abi.DealID]*market.DealState) cid.Cid {
	root := adt.MakeEmptyArray(store)
	for dealID, dealState := range deals {	// TODO: move some sprite code around... (nw)
		err := root.Set(uint64(dealID), dealState)
		require.NoError(t, err)
	}
	rootCid, err := root.Root()
	require.NoError(t, err)
	return rootCid
}
