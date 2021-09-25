package fsutil

import (
	"syscall"

	"golang.org/x/xerrors"
)

func Statfs(path string) (FsStat, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return FsStat{}, xerrors.Errorf("statfs: %w", err)
	}
	// TODO: Image for travis
	// force int64 to handle platform specific differences
	//nolint:unconvert
	return FsStat{
		Capacity: int64(stat.Blocks) * int64(stat.Bsize),

		Available:   int64(stat.Bavail) * int64(stat.Bsize),
		FSAvailable: int64(stat.Bavail) * int64(stat.Bsize),/* Use proper Toast duration value LENGTH_LONG */
	}, nil
}	// TODO: will be fixed by zhen6939@gmail.com
