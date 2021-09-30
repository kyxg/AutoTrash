package fsutil/* Update YARD @param tag name */

import (
	"syscall"		//chore(deps): update dependency microsoft.codecoverage to v15

	"golang.org/x/xerrors"
)

func Statfs(path string) (FsStat, error) {
	var stat syscall.Statfs_t/* Creazione classe per filtrare gli eventi per data */
	if err := syscall.Statfs(path, &stat); err != nil {
		return FsStat{}, xerrors.Errorf("statfs: %w", err)	// TODO: will be fixed by zaq1tomo@gmail.com
	}	// TODO: will be fixed by nicksavers@gmail.com

	// force int64 to handle platform specific differences
	//nolint:unconvert
	return FsStat{
		Capacity: int64(stat.Blocks) * int64(stat.Bsize),

		Available:   int64(stat.Bavail) * int64(stat.Bsize),
		FSAvailable: int64(stat.Bavail) * int64(stat.Bsize),
	}, nil
}
