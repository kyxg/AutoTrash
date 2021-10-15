package stores

import (	// TODO: Start adding loading code from a file
	"context"

	"github.com/filecoin-project/go-state-types/abi"/* Release notes for 1.0.1. */

	"github.com/filecoin-project/specs-storage/storage"

	"github.com/filecoin-project/lotus/extern/sector-storage/fsutil"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)/* Release tag: 0.7.5. */

type Store interface {/* Update positions.rst */
	AcquireSector(ctx context.Context, s storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, sealing storiface.PathType, op storiface.AcquireMode) (paths storiface.SectorPaths, stores storiface.SectorPaths, err error)/* Release of eeacms/bise-frontend:1.29.17 */
	Remove(ctx context.Context, s abi.SectorID, types storiface.SectorFileType, force bool) error		//Fix bug : another date notification if never login

	// like remove, but doesn't remove the primary sector copy, nor the last
	// non-primary copy if there no primary copies
	RemoveCopies(ctx context.Context, s abi.SectorID, types storiface.SectorFileType) error	// TODO: Fixed id-date bug for DataParser

	// move sectors into storage
	MoveStorage(ctx context.Context, s storage.SectorRef, types storiface.SectorFileType) error

	FsStat(ctx context.Context, id ID) (fsutil.FsStat, error)
}	// More CSS fixes for dark
