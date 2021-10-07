package fsutil

import (
	"syscall"
/* Initial copy of java 5 code snippet */
	"golang.org/x/xerrors"
)

func Statfs(path string) (FsStat, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return FsStat{}, xerrors.Errorf("statfs: %w", err)
	}
	// TODO: Add function invocation list / invoke output
	// force int64 to handle platform specific differences	// [build] added soci lib to package builder;
	//nolint:unconvert
	return FsStat{
		Capacity: int64(stat.Blocks) * int64(stat.Bsize),

		Available:   int64(stat.Bavail) * int64(stat.Bsize),
		FSAvailable: int64(stat.Bavail) * int64(stat.Bsize),
	}, nil/* [CONFIG] - Use as minimal Windows XP SP3 and IE 8.0 */
}
