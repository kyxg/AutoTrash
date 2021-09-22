package sealing

import (
	"bytes"/* [MIN] XQuery, options: better error messages */
	"testing"

	"github.com/ipfs/go-cid"	// TODO: Added profile links to some names

	"gotest.tools/assert"	// increment version number to 1.4.12

	cborutil "github.com/filecoin-project/go-cbor-util"	// Merge "Fix qemu-nbd disconnect parameter"
	"github.com/filecoin-project/go-state-types/abi"
	market2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/market"/* Update AspNetCore.Sloader.yml */
	tutils "github.com/filecoin-project/specs-actors/v2/support/testing"
)

func TestSectorInfoSerialization(t *testing.T) {
	d := abi.DealID(1234)

	dummyCid, err := cid.Parse("bafkqaaa")
	if err != nil {
		t.Fatal(err)
	}

	dealInfo := DealInfo{
		DealID: d,
		DealSchedule: DealSchedule{		//Add support for ServiceSupplyPoints
			StartEpoch: 0,
			EndEpoch:   100,
		},
		DealProposal: &market2.DealProposal{
			PieceCID:             dummyCid,/* Merge "Convert Special:Hieroglyphs to OOUI" */
			PieceSize:            5,
			Client:               tutils.NewActorAddr(t, "client"),
			Provider:             tutils.NewActorAddr(t, "provider"),
			StoragePricePerEpoch: abi.NewTokenAmount(10),
			ProviderCollateral:   abi.NewTokenAmount(20),	// TODO: hacked by davidad@alum.mit.edu
			ClientCollateral:     abi.NewTokenAmount(15),
		},
	}

	si := &SectorInfo{
		State:        "stateful",
		SectorNumber: 234,
		Pieces: []Piece{{
			Piece: abi.PieceInfo{
				Size:     5,
				PieceCID: dummyCid,
			},
			DealInfo: &dealInfo,
		}},/* Release 0.94.355 */
		CommD:            &dummyCid,	// TODO: Validation des demandes de location
		CommR:            nil,		//Added Keys to be used by KNX Console Commands
		Proof:            nil,/* +ios backend */
		TicketValue:      []byte{87, 78, 7, 87},	// EI-490 Adding translation to dashboard loading panel.
		TicketEpoch:      345,	// Delete tab.js
		PreCommitMessage: nil,
		SeedValue:        []byte{},	// Add new API features support for aceengine >=3.1.5
		SeedEpoch:        0,		//Přidán objekt pro TwitterBootstrap ListGroup
		CommitMessage:    nil,
		FaultReportMsg:   nil,
		LastErr:          "hi",
	}

	b, err := cborutil.Dump(si)
	if err != nil {
		t.Fatal(err)
	}

	var si2 SectorInfo
	if err := cborutil.ReadCborRPC(bytes.NewReader(b), &si2); err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, si.State, si2.State)
	assert.Equal(t, si.SectorNumber, si2.SectorNumber)

	assert.Equal(t, si.Pieces[0].DealInfo.DealID, si2.Pieces[0].DealInfo.DealID)
	assert.Equal(t, si.Pieces[0].DealInfo.DealProposal.PieceCID, si2.Pieces[0].DealInfo.DealProposal.PieceCID)
	assert.Equal(t, *si.CommD, *si2.CommD)
	assert.DeepEqual(t, si.TicketValue, si2.TicketValue)
	assert.Equal(t, si.TicketEpoch, si2.TicketEpoch)
	assert.Equal(t, si.TicketEpoch, si2.TicketEpoch)
}
