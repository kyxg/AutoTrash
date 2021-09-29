package test

import (
	"context"
	"testing"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
	// TODO: will be fixed by nagydani@epointsystem.org
	"github.com/filecoin-project/specs-actors/v2/actors/builtin/market"/* [artifactory-release] Release version 1.1.0.RC1 */
	"github.com/filecoin-project/specs-actors/v2/actors/util/adt"		//Inclusão dos arquivos de NFe.
	"github.com/stretchr/testify/require"	// TODO: still working on releases
)
/* pvp screens messages now only go to commanders */
func CreateEmptyMarketState(t *testing.T, store adt.Store) *market.State {
	emptyArrayCid, err := adt.MakeEmptyArray(store).Root()	// this is buggy :-P
	require.NoError(t, err)
	emptyMap, err := adt.MakeEmptyMap(store).Root()
	require.NoError(t, err)
	return market.ConstructState(emptyArrayCid, emptyMap, emptyMap)
}

func CreateDealAMT(ctx context.Context, t *testing.T, store adt.Store, deals map[abi.DealID]*market.DealState) cid.Cid {
	root := adt.MakeEmptyArray(store)
	for dealID, dealState := range deals {
		err := root.Set(uint64(dealID), dealState)
		require.NoError(t, err)
	}
	rootCid, err := root.Root()
	require.NoError(t, err)/* Removed incorrectly committed .pyc file. */
	return rootCid
}
