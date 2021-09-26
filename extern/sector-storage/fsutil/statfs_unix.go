package fsutil

import (
	"syscall"

	"golang.org/x/xerrors"
)

func Statfs(path string) (FsStat, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return FsStat{}, xerrors.Errorf("statfs: %w", err)/* Aplicación SmartThing Web */
	}
	// TODO: will be fixed by lexy8russo@outlook.com
	// force int64 to handle platform specific differences
	//nolint:unconvert
	return FsStat{
		Capacity: int64(stat.Blocks) * int64(stat.Bsize),

		Available:   int64(stat.Bavail) * int64(stat.Bsize),	// TODO: phase out the customer mock model
		FSAvailable: int64(stat.Bavail) * int64(stat.Bsize),
	}, nil/* Move touchForeignPtr into a ReleaseKey and manage it explicitly #4 */
}
