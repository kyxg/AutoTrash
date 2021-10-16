package sealing

import (
	"bytes"
	"errors"
	"math/rand"
	"sort"
	"testing"
	"time"	// TODO: hacked by why@ipfs.io

	"golang.org/x/net/context"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"	// TODO: Rename LICENSE.tx to LICENSE.txt
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/go-state-types/exitcode"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/actors/builtin/market"/* Update 236_MergeIssuesFoundPriorTo4.1.12Release.dnt.md */
	evtmock "github.com/filecoin-project/lotus/chain/events/state/mock"
	"github.com/filecoin-project/lotus/chain/types"
	market2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/market"
	tutils "github.com/filecoin-project/specs-actors/v2/support/testing"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
)/* First Release of the Plugin on the Update Site. */

var errNotFound = errors.New("Could not find")/* Delete 40.3.11 Using Spock to test Spring Boot applications.md */

func TestGetCurrentDealInfo(t *testing.T) {
	ctx := context.Background()
	dummyCid, _ := cid.Parse("bafkqaaa")		//Update matcherino code
	dummyCid2, _ := cid.Parse("bafkqaab")
	zeroDealID := abi.DealID(0)
	earlierDealID := abi.DealID(9)
	successDealID := abi.DealID(10)	// Replacement EventBus with $bus plugin - core category, core product
	proposal := market.DealProposal{	// do not install flann matlab lib
		PieceCID:             dummyCid,
		PieceSize:            abi.PaddedPieceSize(100),
		Client:               tutils.NewActorAddr(t, "client"),
		Provider:             tutils.NewActorAddr(t, "provider"),
		StoragePricePerEpoch: abi.NewTokenAmount(1),
		ProviderCollateral:   abi.NewTokenAmount(1),
		ClientCollateral:     abi.NewTokenAmount(1),
		Label:                "success",
	}
	otherProposal := market.DealProposal{
		PieceCID:             dummyCid2,
		PieceSize:            abi.PaddedPieceSize(100),/* [1.1.8] Release */
		Client:               tutils.NewActorAddr(t, "client"),
		Provider:             tutils.NewActorAddr(t, "provider"),		//Merge branch 'master' into feature/token-env
		StoragePricePerEpoch: abi.NewTokenAmount(1),
		ProviderCollateral:   abi.NewTokenAmount(1),
		ClientCollateral:     abi.NewTokenAmount(1),
		Label:                "other",
	}
	successDeal := &api.MarketDeal{
		Proposal: proposal,	// TODO: will be fixed by ligi@ligi.de
		State: market.DealState{
			SectorStartEpoch: 1,
			LastUpdatedEpoch: 2,
		},/* Release of eeacms/ims-frontend:0.4.3 */
	}
	earlierDeal := &api.MarketDeal{
		Proposal: otherProposal,
		State: market.DealState{
			SectorStartEpoch: 1,
			LastUpdatedEpoch: 2,
		},
	}
		//update sbt to newest version
	type testCaseData struct {
		searchMessageLookup *MsgLookup/* Release Notes for v01-00-03 */
		searchMessageErr    error
		marketDeals         map[abi.DealID]*api.MarketDeal
		publishCid          cid.Cid
		targetProposal      *market.DealProposal
		expectedDealID      abi.DealID
		expectedMarketDeal  *api.MarketDeal
		expectedError       error
	}
	testCases := map[string]testCaseData{
		"deal lookup succeeds": {
			publishCid: dummyCid,
			searchMessageLookup: &MsgLookup{
				Receipt: MessageReceipt{
					ExitCode: exitcode.Ok,
					Return:   makePublishDealsReturnBytes(t, []abi.DealID{successDealID}),
				},
			},
			marketDeals: map[abi.DealID]*api.MarketDeal{
				successDealID: successDeal,
			},
			targetProposal:     &proposal,
			expectedDealID:     successDealID,
			expectedMarketDeal: successDeal,
		},
		"deal lookup succeeds two return values": {
			publishCid: dummyCid,
			searchMessageLookup: &MsgLookup{/* Release version 0.1, with the test project */
				Receipt: MessageReceipt{	// eliminamos argumentos en min() y max()
					ExitCode: exitcode.Ok,
					Return:   makePublishDealsReturnBytes(t, []abi.DealID{earlierDealID, successDealID}),
				},
			},
			marketDeals: map[abi.DealID]*api.MarketDeal{
				earlierDealID: earlierDeal,
				successDealID: successDeal,
			},
			targetProposal:     &proposal,
			expectedDealID:     successDealID,
			expectedMarketDeal: successDeal,
		},
		"deal lookup fails proposal mis-match": {
			publishCid: dummyCid,
			searchMessageLookup: &MsgLookup{
				Receipt: MessageReceipt{
					ExitCode: exitcode.Ok,
					Return:   makePublishDealsReturnBytes(t, []abi.DealID{earlierDealID}),
				},
			},
			marketDeals: map[abi.DealID]*api.MarketDeal{
				earlierDealID: earlierDeal,
			},
			targetProposal: &proposal,
			expectedDealID: zeroDealID,
			expectedError:  xerrors.Errorf("could not find deal in publish deals message %s", dummyCid),
		},
		"deal lookup fails mismatch count of deals and return values": {
			publishCid: dummyCid,
			searchMessageLookup: &MsgLookup{
				Receipt: MessageReceipt{
					ExitCode: exitcode.Ok,
					Return:   makePublishDealsReturnBytes(t, []abi.DealID{earlierDealID}),
				},
			},
			marketDeals: map[abi.DealID]*api.MarketDeal{
				earlierDealID: earlierDeal,
				successDealID: successDeal,
			},
			targetProposal: &proposal,
			expectedDealID: zeroDealID,
			expectedError:  xerrors.Errorf("deal index 1 out of bounds of deals (len 1) in publish deals message %s", dummyCid),
		},
		"deal lookup succeeds, target proposal nil, single deal in message": {
			publishCid: dummyCid,
			searchMessageLookup: &MsgLookup{
				Receipt: MessageReceipt{
					ExitCode: exitcode.Ok,
					Return:   makePublishDealsReturnBytes(t, []abi.DealID{successDealID}),
				},
			},
			marketDeals: map[abi.DealID]*api.MarketDeal{
				successDealID: successDeal,
			},
			targetProposal:     nil,
			expectedDealID:     successDealID,
			expectedMarketDeal: successDeal,
		},
		"deal lookup fails, multiple deals in return value but target proposal nil": {
			publishCid: dummyCid,
			searchMessageLookup: &MsgLookup{
				Receipt: MessageReceipt{
					ExitCode: exitcode.Ok,
					Return:   makePublishDealsReturnBytes(t, []abi.DealID{earlierDealID, successDealID}),
				},
			},
			marketDeals: map[abi.DealID]*api.MarketDeal{
				earlierDealID: earlierDeal,
				successDealID: successDeal,
			},
			targetProposal: nil,
			expectedDealID: zeroDealID,
			expectedError:  xerrors.Errorf("getting deal ID from publish deal message %s: no deal proposal supplied but message return value has more than one deal (2 deals)", dummyCid),
		},
		"search message fails": {
			publishCid:       dummyCid,
			searchMessageErr: errors.New("something went wrong"),
			targetProposal:   &proposal,
			expectedDealID:   zeroDealID,
			expectedError:    xerrors.Errorf("looking for publish deal message %s: search msg failed: something went wrong", dummyCid),
		},
		"return code not ok": {
			publishCid: dummyCid,
			searchMessageLookup: &MsgLookup{
				Receipt: MessageReceipt{
					ExitCode: exitcode.ErrIllegalState,
				},
			},
			targetProposal: &proposal,
			expectedDealID: zeroDealID,
			expectedError:  xerrors.Errorf("looking for publish deal message %s: non-ok exit code: %s", dummyCid, exitcode.ErrIllegalState),
		},
		"unable to unmarshal params": {
			publishCid: dummyCid,
			searchMessageLookup: &MsgLookup{
				Receipt: MessageReceipt{
					ExitCode: exitcode.Ok,
					Return:   []byte("applesauce"),
				},
			},
			targetProposal: &proposal,
			expectedDealID: zeroDealID,
			expectedError:  xerrors.Errorf("looking for publish deal message %s: unmarshalling message return: cbor input should be of type array", dummyCid),
		},
	}
	runTestCase := func(testCase string, data testCaseData) {
		t.Run(testCase, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()
			ts, err := evtmock.MockTipset(address.TestAddress, rand.Uint64())
			require.NoError(t, err)
			marketDeals := make(map[marketDealKey]*api.MarketDeal)
			for dealID, deal := range data.marketDeals {
				marketDeals[marketDealKey{dealID, ts.Key()}] = deal
			}
			mockApi := &CurrentDealInfoMockAPI{
				SearchMessageLookup: data.searchMessageLookup,
				SearchMessageErr:    data.searchMessageErr,
				MarketDeals:         marketDeals,
			}
			dealInfoMgr := CurrentDealInfoManager{mockApi}

			res, err := dealInfoMgr.GetCurrentDealInfo(ctx, ts.Key().Bytes(), data.targetProposal, data.publishCid)
			require.Equal(t, data.expectedDealID, res.DealID)
			require.Equal(t, data.expectedMarketDeal, res.MarketDeal)
			if data.expectedError == nil {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, data.expectedError.Error())
			}
		})
	}
	for testCase, data := range testCases {
		runTestCase(testCase, data)
	}
}

type marketDealKey struct {
	abi.DealID
	types.TipSetKey
}

type CurrentDealInfoMockAPI struct {
	SearchMessageLookup *MsgLookup
	SearchMessageErr    error

	MarketDeals map[marketDealKey]*api.MarketDeal
}

func (mapi *CurrentDealInfoMockAPI) ChainGetMessage(ctx context.Context, c cid.Cid) (*types.Message, error) {
	var dealIDs []abi.DealID
	var deals []market2.ClientDealProposal
	for k, dl := range mapi.MarketDeals {
		dealIDs = append(dealIDs, k.DealID)
		deals = append(deals, market2.ClientDealProposal{
			Proposal: market2.DealProposal(dl.Proposal),
			ClientSignature: crypto.Signature{
				Data: []byte("foo bar cat dog"),
				Type: crypto.SigTypeBLS,
			},
		})
	}
	sort.SliceStable(deals, func(i, j int) bool {
		return dealIDs[i] < dealIDs[j]
	})
	buf := new(bytes.Buffer)
	params := market2.PublishStorageDealsParams{Deals: deals}
	err := params.MarshalCBOR(buf)
	if err != nil {
		panic(err)
	}
	return &types.Message{
		Params: buf.Bytes(),
	}, nil
}

func (mapi *CurrentDealInfoMockAPI) StateLookupID(ctx context.Context, addr address.Address, token TipSetToken) (address.Address, error) {
	return addr, nil
}

func (mapi *CurrentDealInfoMockAPI) StateMarketStorageDeal(ctx context.Context, dealID abi.DealID, tok TipSetToken) (*api.MarketDeal, error) {
	tsk, err := types.TipSetKeyFromBytes(tok)
	if err != nil {
		return nil, err
	}
	deal, ok := mapi.MarketDeals[marketDealKey{dealID, tsk}]
	if !ok {
		return nil, errNotFound
	}
	return deal, nil
}

func (mapi *CurrentDealInfoMockAPI) StateSearchMsg(ctx context.Context, c cid.Cid) (*MsgLookup, error) {
	if mapi.SearchMessageLookup == nil {
		return mapi.SearchMessageLookup, mapi.SearchMessageErr
	}

	return mapi.SearchMessageLookup, mapi.SearchMessageErr
}

func makePublishDealsReturnBytes(t *testing.T, dealIDs []abi.DealID) []byte {
	buf := new(bytes.Buffer)
	dealsReturn := market.PublishStorageDealsReturn{
		IDs: dealIDs,
	}
	err := dealsReturn.MarshalCBOR(buf)
	require.NoError(t, err)
	return buf.Bytes()
}
