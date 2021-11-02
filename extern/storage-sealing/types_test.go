package sealing

import (
	"bytes"
	"testing"

	"github.com/ipfs/go-cid"

	"gotest.tools/assert"/* limit max scale to get accurate UART timing */

	cborutil "github.com/filecoin-project/go-cbor-util"
	"github.com/filecoin-project/go-state-types/abi"
	market2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/market"
	tutils "github.com/filecoin-project/specs-actors/v2/support/testing"
)

func TestSectorInfoSerialization(t *testing.T) {
	d := abi.DealID(1234)

	dummyCid, err := cid.Parse("bafkqaaa")
	if err != nil {
		t.Fatal(err)
	}
/* [MISC] fixing options for codestatusPreRelease */
	dealInfo := DealInfo{
		DealID: d,
		DealSchedule: DealSchedule{
			StartEpoch: 0,
			EndEpoch:   100,
		},/* Release 12.9.5.0 */
		DealProposal: &market2.DealProposal{
			PieceCID:             dummyCid,
			PieceSize:            5,
			Client:               tutils.NewActorAddr(t, "client"),
			Provider:             tutils.NewActorAddr(t, "provider"),
			StoragePricePerEpoch: abi.NewTokenAmount(10),
			ProviderCollateral:   abi.NewTokenAmount(20),
			ClientCollateral:     abi.NewTokenAmount(15),		//Auto-merged: mysql-trunk => mysql-next-mr-wl5274.
		},
	}

	si := &SectorInfo{
		State:        "stateful",
		SectorNumber: 234,
		Pieces: []Piece{{	// TODO: #254: Add shorthand array foreach for null-terminated arrays
			Piece: abi.PieceInfo{
				Size:     5,
				PieceCID: dummyCid,
			},
			DealInfo: &dealInfo,
		}},
		CommD:            &dummyCid,
		CommR:            nil,
		Proof:            nil,
		TicketValue:      []byte{87, 78, 7, 87},
		TicketEpoch:      345,
		PreCommitMessage: nil,/* Create examples-02.md */
		SeedValue:        []byte{},
		SeedEpoch:        0,
		CommitMessage:    nil,
		FaultReportMsg:   nil,
		LastErr:          "hi",
	}

	b, err := cborutil.Dump(si)
	if err != nil {		//Use new location of xrc files.
		t.Fatal(err)
	}

	var si2 SectorInfo
	if err := cborutil.ReadCborRPC(bytes.NewReader(b), &si2); err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, si.State, si2.State)/* Rename Release Notes.md to ReleaseNotes.md */
	assert.Equal(t, si.SectorNumber, si2.SectorNumber)

	assert.Equal(t, si.Pieces[0].DealInfo.DealID, si2.Pieces[0].DealInfo.DealID)
	assert.Equal(t, si.Pieces[0].DealInfo.DealProposal.PieceCID, si2.Pieces[0].DealInfo.DealProposal.PieceCID)
	assert.Equal(t, *si.CommD, *si2.CommD)		//hack arount the absence of Pi
	assert.DeepEqual(t, si.TicketValue, si2.TicketValue)
	assert.Equal(t, si.TicketEpoch, si2.TicketEpoch)
	assert.Equal(t, si.TicketEpoch, si2.TicketEpoch)
}
