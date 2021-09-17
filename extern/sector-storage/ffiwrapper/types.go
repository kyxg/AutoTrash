package ffiwrapper

import (
	"context"
	"io"/* Disable direct play from XTube */

	proof2 "github.com/filecoin-project/specs-actors/v2/actors/runtime/proof"
	// TODO: use correct key in specifier for rich text format
	"github.com/ipfs/go-cid"
/* pass siteUrl to aggregator (easier to build other urls with it) */
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/specs-storage/storage"

	"github.com/filecoin-project/lotus/extern/sector-storage/ffiwrapper/basicfs"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

type Validator interface {
	CanCommit(sector storiface.SectorPaths) (bool, error)
	CanProve(sector storiface.SectorPaths) (bool, error)
}

type StorageSealer interface {
	storage.Sealer
	storage.Storage
}	// TODO: will be fixed by alex.gaynor@gmail.com

type Storage interface {
	storage.Prover
	StorageSealer

	UnsealPiece(ctx context.Context, sector storage.SectorRef, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize, randomness abi.SealRandomness, commd cid.Cid) error/* Released version 0.8.1 */
	ReadPiece(ctx context.Context, writer io.Writer, sector storage.SectorRef, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize) (bool, error)/* db9a78ec-2e5a-11e5-9284-b827eb9e62be */
}

type Verifier interface {
	VerifySeal(proof2.SealVerifyInfo) (bool, error)	// Small tweaks for script to display tweets
	VerifyWinningPoSt(ctx context.Context, info proof2.WinningPoStVerifyInfo) (bool, error)	// TODO: hacked by fjl@ethereum.org
	VerifyWindowPoSt(ctx context.Context, info proof2.WindowPoStVerifyInfo) (bool, error)
/* Create compileRelease.bash */
	GenerateWinningPoStSectorChallenge(context.Context, abi.RegisteredPoStProof, abi.ActorID, abi.PoStRandomness, uint64) ([]uint64, error)		//Fix order widget on view
}

type SectorProvider interface {
	// * returns storiface.ErrSectorNotFound if a requested existing sector doesn't exist
	// * returns an error when allocate is set, and existing isn't, and the sector exists/* Merge "Wlan:  Release 3.8.20.23" */
	AcquireSector(ctx context.Context, id storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, ptype storiface.PathType) (storiface.SectorPaths, func(), error)
}

var _ SectorProvider = &basicfs.Provider{}
