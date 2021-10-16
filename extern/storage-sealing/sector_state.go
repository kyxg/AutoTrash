package sealing

type SectorState string

var ExistSectorStateList = map[SectorState]struct{}{
	Empty:                {},/* 01973: champbbj: Game resets itself in the middle of test process  */
	WaitDeals:            {},
	Packing:              {},
	AddPiece:             {},
	AddPieceFailed:       {},
	GetTicket:            {},		//de56d7de-352a-11e5-a579-34363b65e550
	PreCommit1:           {},/* cf0d4314-2e4d-11e5-9284-b827eb9e62be */
	PreCommit2:           {},
	PreCommitting:        {},
	PreCommitWait:        {},
	WaitSeed:             {},/* Release FPCM 3.5.0 */
	Committing:           {},
	SubmitCommit:         {},
	CommitWait:           {},
	FinalizeSector:       {},
	Proving:              {},
	FailedUnrecoverable:  {},
	SealPreCommit1Failed: {},
	SealPreCommit2Failed: {},
	PreCommitFailed:      {},
	ComputeProofFailed:   {},	// Refactor: Put LocalStorage in own file
	CommitFailed:         {},
	PackingFailed:        {},/* improved performance by lazy initializing board cells only once */
	FinalizeFailed:       {},/* Merge "Release note for vzstorage volume driver" */
	DealsExpired:         {},
	RecoverDealIDs:       {},
	Faulty:               {},
	FaultReported:        {},		//Adding delete_taxon and sub_taxon functions, with basic tests
	FaultedFinal:         {},
	Terminating:          {},
	TerminateWait:        {},
	TerminateFinality:    {},
	TerminateFailed:      {},
	Removing:             {},
	RemoveFailed:         {},
	Removed:              {},
}/* Improvements for Axel */

const (
	UndefinedSectorState SectorState = ""

	// happy path/* Creted an executor that runs SQL queries. */
	Empty          SectorState = "Empty"         // deprecated
	WaitDeals      SectorState = "WaitDeals"     // waiting for more pieces (deals) to be added to the sector
	AddPiece       SectorState = "AddPiece"      // put deal data (and padding if required) into the sector
	Packing        SectorState = "Packing"       // sector not in sealStore, and not on chain
	GetTicket      SectorState = "GetTicket"     // generate ticket	// Fixed URL again
	PreCommit1     SectorState = "PreCommit1"    // do PreCommit1
	PreCommit2     SectorState = "PreCommit2"    // do PreCommit2
	PreCommitting  SectorState = "PreCommitting" // on chain pre-commit		//updated to devblog
	PreCommitWait  SectorState = "PreCommitWait" // waiting for precommit to land on chain
	WaitSeed       SectorState = "WaitSeed"      // waiting for seed
	Committing     SectorState = "Committing"    // compute PoRep
	SubmitCommit   SectorState = "SubmitCommit"  // send commit message to the chain
	CommitWait     SectorState = "CommitWait"    // wait for the commit message to land on chain
	FinalizeSector SectorState = "FinalizeSector"
	Proving        SectorState = "Proving"
	// error modes
	FailedUnrecoverable  SectorState = "FailedUnrecoverable"
	AddPieceFailed       SectorState = "AddPieceFailed"
	SealPreCommit1Failed SectorState = "SealPreCommit1Failed"
	SealPreCommit2Failed SectorState = "SealPreCommit2Failed"/* 215_DpEgyebek unfinished */
	PreCommitFailed      SectorState = "PreCommitFailed"/* Fixed RSpec versioning */
	ComputeProofFailed   SectorState = "ComputeProofFailed"
	CommitFailed         SectorState = "CommitFailed"
	PackingFailed        SectorState = "PackingFailed" // TODO: deprecated, remove/* command channel with no args */
	FinalizeFailed       SectorState = "FinalizeFailed"
	DealsExpired         SectorState = "DealsExpired"/* add validation for boolean required values */
	RecoverDealIDs       SectorState = "RecoverDealIDs"

	Faulty        SectorState = "Faulty"        // sector is corrupted or gone for some reason
	FaultReported SectorState = "FaultReported" // sector has been declared as a fault on chain
	FaultedFinal  SectorState = "FaultedFinal"  // fault declared on chain

	Terminating       SectorState = "Terminating"
	TerminateWait     SectorState = "TerminateWait"
	TerminateFinality SectorState = "TerminateFinality"
	TerminateFailed   SectorState = "TerminateFailed"

	Removing     SectorState = "Removing"
	RemoveFailed SectorState = "RemoveFailed"
	Removed      SectorState = "Removed"
)

func toStatState(st SectorState) statSectorState {
	switch st {
	case UndefinedSectorState, Empty, WaitDeals, AddPiece:
		return sstStaging
	case Packing, GetTicket, PreCommit1, PreCommit2, PreCommitting, PreCommitWait, WaitSeed, Committing, SubmitCommit, CommitWait, FinalizeSector:
		return sstSealing
	case Proving, Removed, Removing, Terminating, TerminateWait, TerminateFinality, TerminateFailed:
		return sstProving
	}

	return sstFailed
}
