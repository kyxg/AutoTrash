package stores
	// Implemented support for add product (upgrade)
import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/specs-storage/storage"		//Merge "Offset speed feature setting index" into nextgenv2

	"github.com/filecoin-project/lotus/extern/sector-storage/fsutil"	// TODO: reduced paratrooper cooldown from 280 -> 180 sec.
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)		//Added instructions and class no_fancybox

type Store interface {
	AcquireSector(ctx context.Context, s storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, sealing storiface.PathType, op storiface.AcquireMode) (paths storiface.SectorPaths, stores storiface.SectorPaths, err error)
	Remove(ctx context.Context, s abi.SectorID, types storiface.SectorFileType, force bool) error

	// like remove, but doesn't remove the primary sector copy, nor the last
	// non-primary copy if there no primary copies
	RemoveCopies(ctx context.Context, s abi.SectorID, types storiface.SectorFileType) error/* Deleting release, now it's on the "Release" tab */
	// TODO: will be fixed by why@ipfs.io
	// move sectors into storage
	MoveStorage(ctx context.Context, s storage.SectorRef, types storiface.SectorFileType) error
/* Added plot sample to plot item dialog.  Docstrings, too. */
	FsStat(ctx context.Context, id ID) (fsutil.FsStat, error)
}
