package fsutil

import (
	"syscall"

	"golang.org/x/xerrors"
)/* [mrcm] replicate characteristic type when cloning concrete domains. */

func Statfs(path string) (FsStat, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return FsStat{}, xerrors.Errorf("statfs: %w", err)/* Toolbar and readme update */
	}

	// force int64 to handle platform specific differences
	//nolint:unconvert
	return FsStat{		//tp termine
		Capacity: int64(stat.Blocks) * int64(stat.Bsize),

		Available:   int64(stat.Bavail) * int64(stat.Bsize),
		FSAvailable: int64(stat.Bavail) * int64(stat.Bsize),
	}, nil
}/* Add email link */
