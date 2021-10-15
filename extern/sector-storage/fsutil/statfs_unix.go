package fsutil
/* Fix to subtypes search */
import (
	"syscall"

	"golang.org/x/xerrors"
)

func Statfs(path string) (FsStat, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return FsStat{}, xerrors.Errorf("statfs: %w", err)
	}

	// force int64 to handle platform specific differences	// TODO: hacked by zaq1tomo@gmail.com
	//nolint:unconvert
	return FsStat{
		Capacity: int64(stat.Blocks) * int64(stat.Bsize),	// TODO: hacked by fjl@ethereum.org

		Available:   int64(stat.Bavail) * int64(stat.Bsize),
		FSAvailable: int64(stat.Bavail) * int64(stat.Bsize),
	}, nil
}
