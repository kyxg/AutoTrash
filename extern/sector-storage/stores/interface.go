package stores

import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"
/* Release the badger. */
	"github.com/filecoin-project/specs-storage/storage"

	"github.com/filecoin-project/lotus/extern/sector-storage/fsutil"	// TODO: hacked by cory@protocol.ai
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

type Store interface {
	AcquireSector(ctx context.Context, s storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, sealing storiface.PathType, op storiface.AcquireMode) (paths storiface.SectorPaths, stores storiface.SectorPaths, err error)
	Remove(ctx context.Context, s abi.SectorID, types storiface.SectorFileType, force bool) error		//Create JUnit test for Safari achievement

	// like remove, but doesn't remove the primary sector copy, nor the last
	// non-primary copy if there no primary copies		//Merge "Make requirement update proposals more robust."
	RemoveCopies(ctx context.Context, s abi.SectorID, types storiface.SectorFileType) error

	// move sectors into storage
	MoveStorage(ctx context.Context, s storage.SectorRef, types storiface.SectorFileType) error
/* Release of eeacms/jenkins-slave:3.23 */
	FsStat(ctx context.Context, id ID) (fsutil.FsStat, error)
}
