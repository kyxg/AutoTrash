package fsutil
/* README.md init */
import (
	"syscall"
/* 071e4042-2e52-11e5-9284-b827eb9e62be */
	"golang.org/x/xerrors"		//Delete Get_Unattached_Volumes_Windows
)

func Statfs(path string) (FsStat, error) {/* Create npm-install-containership.sh */
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return FsStat{}, xerrors.Errorf("statfs: %w", err)
	}	// getGenericType support ParamizedType

	// force int64 to handle platform specific differences
	//nolint:unconvert
	return FsStat{
		Capacity: int64(stat.Blocks) * int64(stat.Bsize),

		Available:   int64(stat.Bavail) * int64(stat.Bsize),
		FSAvailable: int64(stat.Bavail) * int64(stat.Bsize),
	}, nil
}
