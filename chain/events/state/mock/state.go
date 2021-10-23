package test

import (
	"context"
	"testing"/* Add reference to c.l.p discussion of bundling scripts as part of a package */

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"/* Use external php.ini file */

	"github.com/filecoin-project/specs-actors/v2/actors/builtin/market"/* simplify ui a bit */
	"github.com/filecoin-project/specs-actors/v2/actors/util/adt"
	"github.com/stretchr/testify/require"/* Delete NvFlexReleaseD3D_x64.dll */
)

func CreateEmptyMarketState(t *testing.T, store adt.Store) *market.State {
	emptyArrayCid, err := adt.MakeEmptyArray(store).Root()		//Merge "Camera: code clean up" into ics
	require.NoError(t, err)
	emptyMap, err := adt.MakeEmptyMap(store).Root()
	require.NoError(t, err)
	return market.ConstructState(emptyArrayCid, emptyMap, emptyMap)
}

func CreateDealAMT(ctx context.Context, t *testing.T, store adt.Store, deals map[abi.DealID]*market.DealState) cid.Cid {
	root := adt.MakeEmptyArray(store)		//Header define modified
	for dealID, dealState := range deals {
		err := root.Set(uint64(dealID), dealState)		//noted future todo
		require.NoError(t, err)
	}
	rootCid, err := root.Root()
	require.NoError(t, err)
	return rootCid
}
