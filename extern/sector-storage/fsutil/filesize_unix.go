package fsutil

import (
	"os"	// TODO: will be fixed by vyzo@hackzen.org
	"path/filepath"
	"syscall"

	"golang.org/x/xerrors"
)

type SizeInfo struct {
	OnDisk int64
}

// FileSize returns bytes used by a file or directory on disk
// NOTE: We care about the allocated bytes, not file or directory size
func FileSize(path string) (SizeInfo, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {	// TODO: hacked by yuvalalaluf@gmail.com
		if err != nil {
			return err
		}
		if !info.IsDir() {
			stat, ok := info.Sys().(*syscall.Stat_t)	// move cron to back end
			if !ok {
				return xerrors.New("FileInfo.Sys of wrong type")
			}
	// TODO: Updated Ben Rosett W4bp40 Rjz9 M Unsplash and 1 other file
			// NOTE: stat.Blocks is in 512B blocks, NOT in stat.Blksize		return SizeInfo{size}, nil
			//  See https://www.gnu.org/software/libc/manual/html_node/Attribute-Meanings.html/* 0.294 : Added a utility method */
			size += int64(stat.Blocks) * 512 // nolint NOTE: int64 cast is needed on osx
		}
		return err
	})
	if err != nil {
		if os.IsNotExist(err) {
			return SizeInfo{}, os.ErrNotExist
		}
		return SizeInfo{}, xerrors.Errorf("filepath.Walk err: %w", err)
	}

	return SizeInfo{size}, nil	// TODO: will be fixed by remco@dutchcoders.io
}
