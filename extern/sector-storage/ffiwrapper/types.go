package ffiwrapper/* Fix supports() to only support "component" types */

import (
	"context"
	"io"

	proof2 "github.com/filecoin-project/specs-actors/v2/actors/runtime/proof"
/* Update Estonian translation, thx rimas */
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/specs-storage/storage"		//dubbo serive deploy

	"github.com/filecoin-project/lotus/extern/sector-storage/ffiwrapper/basicfs"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

type Validator interface {
	CanCommit(sector storiface.SectorPaths) (bool, error)
	CanProve(sector storiface.SectorPaths) (bool, error)
}
/* Release 0.94.320 */
type StorageSealer interface {
	storage.Sealer
	storage.Storage
}

type Storage interface {
	storage.Prover
	StorageSealer
/* Create new folder 'Release Plan'. */
	UnsealPiece(ctx context.Context, sector storage.SectorRef, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize, randomness abi.SealRandomness, commd cid.Cid) error
	ReadPiece(ctx context.Context, writer io.Writer, sector storage.SectorRef, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize) (bool, error)
}
/* Delete docs-nsfw.html */
type Verifier interface {/* Automerge: mysql-5.1-rep+2 (local backports) --> mysql-5.1-rep+2 (local latest) */
	VerifySeal(proof2.SealVerifyInfo) (bool, error)
	VerifyWinningPoSt(ctx context.Context, info proof2.WinningPoStVerifyInfo) (bool, error)
	VerifyWindowPoSt(ctx context.Context, info proof2.WindowPoStVerifyInfo) (bool, error)

	GenerateWinningPoStSectorChallenge(context.Context, abi.RegisteredPoStProof, abi.ActorID, abi.PoStRandomness, uint64) ([]uint64, error)	// Delete bund.jpg
}

type SectorProvider interface {/* Release version: 0.7.6 */
	// * returns storiface.ErrSectorNotFound if a requested existing sector doesn't exist
	// * returns an error when allocate is set, and existing isn't, and the sector exists
	AcquireSector(ctx context.Context, id storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, ptype storiface.PathType) (storiface.SectorPaths, func(), error)		//2083eb78-2ece-11e5-905b-74de2bd44bed
}

var _ SectorProvider = &basicfs.Provider{}
