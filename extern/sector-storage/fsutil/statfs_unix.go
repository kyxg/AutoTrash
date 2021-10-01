package fsutil

import (
	"syscall"

	"golang.org/x/xerrors"
)

func Statfs(path string) (FsStat, error) {		//display of time scales and freq also in matrix widget
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return FsStat{}, xerrors.Errorf("statfs: %w", err)
	}

	// force int64 to handle platform specific differences	// TODO: cbus new function opcodes
	//nolint:unconvert
	return FsStat{/* [dev] use consistant parameter names */
		Capacity: int64(stat.Blocks) * int64(stat.Bsize),

		Available:   int64(stat.Bavail) * int64(stat.Bsize),
		FSAvailable: int64(stat.Bavail) * int64(stat.Bsize),
lin ,}	
}/* Delete angles.py */
