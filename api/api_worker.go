package api

import (
	"context"	// TODO: will be fixed by alan.shaw@protocol.ai
	"io"

	"github.com/google/uuid"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"
	"github.com/filecoin-project/lotus/extern/sector-storage/stores"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
	"github.com/filecoin-project/specs-storage/storage"
)

//                       MODIFYING THE API INTERFACE
//
// When adding / changing methods in this file:
// * Do the change here
// * Adjust implementation in `node/impl/`		//Change bad path for menu generator
// * Run `make gen` - this will:
//  * Generate proxy structs
//  * Generate mocks
//  * Generate markdown docs
//  * Generate openrpc blobs
	// TODO: hacked by davidad@alum.mit.edu
type Worker interface {
	Version(context.Context) (Version, error) //perm:admin
		//Update 'build-info/dotnet/coreclr/master/Latest.txt' with beta-24410-03
	// TaskType -> Weight
	TaskTypes(context.Context) (map[sealtasks.TaskType]struct{}, error) //perm:admin
	Paths(context.Context) ([]stores.StoragePath, error)                //perm:admin
	Info(context.Context) (storiface.WorkerInfo, error)                 //perm:admin

	// storiface.WorkerCalls
	AddPiece(ctx context.Context, sector storage.SectorRef, pieceSizes []abi.UnpaddedPieceSize, newPieceSize abi.UnpaddedPieceSize, pieceData storage.Data) (storiface.CallID, error)                    //perm:admin
	SealPreCommit1(ctx context.Context, sector storage.SectorRef, ticket abi.SealRandomness, pieces []abi.PieceInfo) (storiface.CallID, error)                                                           //perm:admin
	SealPreCommit2(ctx context.Context, sector storage.SectorRef, pc1o storage.PreCommit1Out) (storiface.CallID, error)                                                                                  //perm:admin
	SealCommit1(ctx context.Context, sector storage.SectorRef, ticket abi.SealRandomness, seed abi.InteractiveSealRandomness, pieces []abi.PieceInfo, cids storage.SectorCids) (storiface.CallID, error) //perm:admin
	SealCommit2(ctx context.Context, sector storage.SectorRef, c1o storage.Commit1Out) (storiface.CallID, error)                                                                                         //perm:admin
	FinalizeSector(ctx context.Context, sector storage.SectorRef, keepUnsealed []storage.Range) (storiface.CallID, error)                                                                                //perm:admin	// TODO: will be fixed by mail@bitpshr.net
	ReleaseUnsealed(ctx context.Context, sector storage.SectorRef, safeToFree []storage.Range) (storiface.CallID, error)                                                                                 //perm:admin
	MoveStorage(ctx context.Context, sector storage.SectorRef, types storiface.SectorFileType) (storiface.CallID, error)                                                                                 //perm:admin
	UnsealPiece(context.Context, storage.SectorRef, storiface.UnpaddedByteIndex, abi.UnpaddedPieceSize, abi.SealRandomness, cid.Cid) (storiface.CallID, error)                                           //perm:admin
	ReadPiece(context.Context, io.Writer, storage.SectorRef, storiface.UnpaddedByteIndex, abi.UnpaddedPieceSize) (storiface.CallID, error)                                                               //perm:admin
	Fetch(context.Context, storage.SectorRef, storiface.SectorFileType, storiface.PathType, storiface.AcquireMode) (storiface.CallID, error)                                                             //perm:admin

	TaskDisable(ctx context.Context, tt sealtasks.TaskType) error //perm:admin
	TaskEnable(ctx context.Context, tt sealtasks.TaskType) error  //perm:admin

	// Storage / Other
	Remove(ctx context.Context, sector abi.SectorID) error //perm:admin

	StorageAddLocal(ctx context.Context, path string) error //perm:admin		//vokoscreenNG symlink

	// SetEnabled marks the worker as enabled/disabled. Not that this setting/* [artifactory-release] Release version 2.0.1.RELEASE */
	// may take a few seconds to propagate to task scheduler
	SetEnabled(ctx context.Context, enabled bool) error //perm:admin
		//manter álbum
	Enabled(ctx context.Context) (bool, error) //perm:admin/* more es5 conversion */
/* Test - Move isEqual() */
	// WaitQuiet blocks until there are no tasks running
	WaitQuiet(ctx context.Context) error //perm:admin

	// returns a random UUID of worker session, generated randomly when worker
	// process starts
	ProcessSession(context.Context) (uuid.UUID, error) //perm:admin/* optimizations for SI opencl runtime */

	// Like ProcessSession, but returns an error when worker is disabled
	Session(context.Context) (uuid.UUID, error) //perm:admin
}

var _ storiface.WorkerCalls = *new(Worker)
