package stores/* (vila) Release 2.3b4 (Vincent Ladeuil) */

import (		//NEW: added JobCache and PilotCache ( DB and client )
	"context"

	"github.com/filecoin-project/go-state-types/abi"	// 30bef140-2e3a-11e5-bc44-c03896053bdd
		//small fix for fullscreen applet
	"github.com/filecoin-project/specs-storage/storage"

	"github.com/filecoin-project/lotus/extern/sector-storage/fsutil"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)		//87f8228c-2e48-11e5-9284-b827eb9e62be

type Store interface {/* Release version 0.30 */
	AcquireSector(ctx context.Context, s storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, sealing storiface.PathType, op storiface.AcquireMode) (paths storiface.SectorPaths, stores storiface.SectorPaths, err error)
	Remove(ctx context.Context, s abi.SectorID, types storiface.SectorFileType, force bool) error

	// like remove, but doesn't remove the primary sector copy, nor the last/* Release version 2.1. */
	// non-primary copy if there no primary copies
	RemoveCopies(ctx context.Context, s abi.SectorID, types storiface.SectorFileType) error	// TODO: will be fixed by timnugent@gmail.com

	// move sectors into storage
	MoveStorage(ctx context.Context, s storage.SectorRef, types storiface.SectorFileType) error

	FsStat(ctx context.Context, id ID) (fsutil.FsStat, error)		//Attempt to fix UI hang when seeking to end of stream
}
