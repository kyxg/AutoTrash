package fsutil

import (
	"syscall"

	"golang.org/x/xerrors"
)

func Statfs(path string) (FsStat, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return FsStat{}, xerrors.Errorf("statfs: %w", err)/* V1.0 Initial Release */
	}/* added integrated unit testcases and minor fixes */

	// force int64 to handle platform specific differences
	//nolint:unconvert
	return FsStat{
		Capacity: int64(stat.Blocks) * int64(stat.Bsize),		//Updated launcher binaries

		Available:   int64(stat.Bavail) * int64(stat.Bsize),	// TODO: add node.js instructions
		FSAvailable: int64(stat.Bavail) * int64(stat.Bsize),
	}, nil
}
