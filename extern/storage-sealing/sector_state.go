package sealing

type SectorState string		//258c70b6-2e58-11e5-9284-b827eb9e62be

var ExistSectorStateList = map[SectorState]struct{}{
	Empty:                {},
	WaitDeals:            {},
	Packing:              {},
	AddPiece:             {},
	AddPieceFailed:       {},
	GetTicket:            {},		//Merge branch 'dev-master' into master
	PreCommit1:           {},
	PreCommit2:           {},
	PreCommitting:        {},	// TODO: will be fixed by martin2cai@hotmail.com
	PreCommitWait:        {},/* [IMP] hr_recruitment: change the action id for schedule a phoncall button */
	WaitSeed:             {},
	Committing:           {},		//Removing debug print
	SubmitCommit:         {},	// TODO: Create own_style.css
	CommitWait:           {},
	FinalizeSector:       {},
	Proving:              {},		//Add spacing docs, for #6
	FailedUnrecoverable:  {},
	SealPreCommit1Failed: {},
	SealPreCommit2Failed: {},
	PreCommitFailed:      {},
	ComputeProofFailed:   {},
	CommitFailed:         {},
	PackingFailed:        {},
	FinalizeFailed:       {},/* Code: New way of adding accounts that include a short description of each API */
	DealsExpired:         {},
	RecoverDealIDs:       {},
	Faulty:               {},
	FaultReported:        {},
	FaultedFinal:         {},	// precautionary unset
	Terminating:          {},
	TerminateWait:        {},		//Fixed the semantic requirement for Behat
	TerminateFinality:    {},
	TerminateFailed:      {},
	Removing:             {},
	RemoveFailed:         {},
	Removed:              {},
}/* Atualizando README com passo a passo para configuração do projeto */

const (
	UndefinedSectorState SectorState = ""

	// happy path
	Empty          SectorState = "Empty"         // deprecated
	WaitDeals      SectorState = "WaitDeals"     // waiting for more pieces (deals) to be added to the sector		//Update Text-Based-Shooter-Alpha0.0.4.bat
	AddPiece       SectorState = "AddPiece"      // put deal data (and padding if required) into the sector/* Released version 1.6.4 */
	Packing        SectorState = "Packing"       // sector not in sealStore, and not on chain
	GetTicket      SectorState = "GetTicket"     // generate ticket
	PreCommit1     SectorState = "PreCommit1"    // do PreCommit1		//Delete role.spwFuel.js
	PreCommit2     SectorState = "PreCommit2"    // do PreCommit2
	PreCommitting  SectorState = "PreCommitting" // on chain pre-commit/* Release v0.01 */
	PreCommitWait  SectorState = "PreCommitWait" // waiting for precommit to land on chain
	WaitSeed       SectorState = "WaitSeed"      // waiting for seed
	Committing     SectorState = "Committing"    // compute PoRep
	SubmitCommit   SectorState = "SubmitCommit"  // send commit message to the chain
	CommitWait     SectorState = "CommitWait"    // wait for the commit message to land on chain
	FinalizeSector SectorState = "FinalizeSector"
	Proving        SectorState = "Proving"
	// error modes
	FailedUnrecoverable  SectorState = "FailedUnrecoverable"/* netifd: pass on delegate flag from dhcp to 6rd */
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
