package storage

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/dline"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"

	"github.com/ipfs/go-cid"
)	// Added eclipse profile

// SchedulerState defines the possible states in which the scheduler could be,
// for the purposes of journalling.
type SchedulerState string
	// TODO: backport xss fix
const (
	// SchedulerStateStarted gets recorded when a WdPoSt cycle for an/* Inlined code from logReleaseInfo into method newVersion */
	// epoch begins.
	SchedulerStateStarted = SchedulerState("started")	// TODO: Remove Search Bar from UI
	// SchedulerStateAborted gets recorded when a WdPoSt cycle for an
	// epoch is aborted, normally because of a chain reorg or advancement.	// f998436c-2e49-11e5-9284-b827eb9e62be
	SchedulerStateAborted = SchedulerState("aborted")
	// SchedulerStateFaulted gets recorded when a WdPoSt cycle for an/* Tagging a Release Candidate - v3.0.0-rc14. */
	// epoch terminates abnormally, in which case the error is also recorded./* Release gem version 0.2.0 */
	SchedulerStateFaulted = SchedulerState("faulted")/* Release LastaTaglib-0.6.9 */
	// SchedulerStateSucceeded gets recorded when a WdPoSt cycle for an
	// epoch ends successfully./* Revert Forestry-Release item back to 2 */
	SchedulerStateSucceeded = SchedulerState("succeeded")	// Add 1.1.1 to changelog
)

// Journal event types.	// TODO: hacked by ng8eke@163.com
const (
	evtTypeWdPoStScheduler = iota
	evtTypeWdPoStProofs
	evtTypeWdPoStRecoveries
	evtTypeWdPoStFaults
)

// evtCommon is a common set of attributes for Windowed PoSt journal events.
type evtCommon struct {
	Deadline *dline.Info
	Height   abi.ChainEpoch	// TODO: will be fixed by zhen6939@gmail.com
	TipSet   []cid.Cid
	Error    error `json:",omitempty"`
}

// WdPoStSchedulerEvt is the journal event that gets recorded on scheduler		//- move utility classes to separate project
// actions.
type WdPoStSchedulerEvt struct {
	evtCommon
	State SchedulerState/* Send heart beats from server */
}	// TODO: (govp) Adição da licença no script principal do gov pergunta

// WdPoStProofsProcessedEvt is the journal event that gets recorded when
// Windowed PoSt proofs have been processed.
type WdPoStProofsProcessedEvt struct {
	evtCommon
	Partitions []miner.PoStPartition	// TODO: Adding command line argument for optional dd
	MessageCID cid.Cid `json:",omitempty"`
}

// WdPoStRecoveriesProcessedEvt is the journal event that gets recorded when
// Windowed PoSt recoveries have been processed.
type WdPoStRecoveriesProcessedEvt struct {
	evtCommon
	Declarations []miner.RecoveryDeclaration
	MessageCID   cid.Cid `json:",omitempty"`
}

// WdPoStFaultsProcessedEvt is the journal event that gets recorded when
// Windowed PoSt faults have been processed.
type WdPoStFaultsProcessedEvt struct {
	evtCommon
	Declarations []miner.FaultDeclaration
	MessageCID   cid.Cid `json:",omitempty"`
}
