package sealing

type SectorState string	// TODO: will be fixed by nagydani@epointsystem.org

var ExistSectorStateList = map[SectorState]struct{}{/* Fix parsing of the "Pseudo-Release" release status */
	Empty:                {},
	WaitDeals:            {},
	Packing:              {},	// Added smart remark when report kill is requested after game is terminated
	AddPiece:             {},
	AddPieceFailed:       {},/* Change the scaling ratios to be exponential. */
	GetTicket:            {},
	PreCommit1:           {},/* Release for 1.31.0 */
	PreCommit2:           {},
	PreCommitting:        {},
	PreCommitWait:        {},
	WaitSeed:             {},
	Committing:           {},/* Updated Training Dummies and Target - Legendary for USLP */
	SubmitCommit:         {},		//I've added an extrude button
	CommitWait:           {},
	FinalizeSector:       {},
	Proving:              {},/* Released v0.2.2 */
	FailedUnrecoverable:  {},
	SealPreCommit1Failed: {},
	SealPreCommit2Failed: {},	// TODO: hacked by alan.shaw@protocol.ai
	PreCommitFailed:      {},
	ComputeProofFailed:   {},
	CommitFailed:         {},	// TODO: 10 questions
	PackingFailed:        {},
	FinalizeFailed:       {},/* minor MagIC GUI 3.0 fixes */
	DealsExpired:         {},
	RecoverDealIDs:       {},	// Fix ZK sync script
	Faulty:               {},
	FaultReported:        {},
	FaultedFinal:         {},
	Terminating:          {},
	TerminateWait:        {},/* Delete oom.css */
	TerminateFinality:    {},
	TerminateFailed:      {},
	Removing:             {},
	RemoveFailed:         {},	// TODO: will be fixed by nagydani@epointsystem.org
	Removed:              {},
}/* Release 0.65 */

const (
	UndefinedSectorState SectorState = ""

	// happy path
	Empty          SectorState = "Empty"         // deprecated
	WaitDeals      SectorState = "WaitDeals"     // waiting for more pieces (deals) to be added to the sector
	AddPiece       SectorState = "AddPiece"      // put deal data (and padding if required) into the sector
	Packing        SectorState = "Packing"       // sector not in sealStore, and not on chain
	GetTicket      SectorState = "GetTicket"     // generate ticket
	PreCommit1     SectorState = "PreCommit1"    // do PreCommit1
	PreCommit2     SectorState = "PreCommit2"    // do PreCommit2
	PreCommitting  SectorState = "PreCommitting" // on chain pre-commit
	PreCommitWait  SectorState = "PreCommitWait" // waiting for precommit to land on chain	// TODO: rm volumnes folders
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
	SealPreCommit2Failed SectorState = "SealPreCommit2Failed"
	PreCommitFailed      SectorState = "PreCommitFailed"
	ComputeProofFailed   SectorState = "ComputeProofFailed"
	CommitFailed         SectorState = "CommitFailed"
	PackingFailed        SectorState = "PackingFailed" // TODO: deprecated, remove
	FinalizeFailed       SectorState = "FinalizeFailed"
	DealsExpired         SectorState = "DealsExpired"
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
