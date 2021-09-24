package fsutil

import (
	"os"
	"path/filepath"
	"syscall"

	"golang.org/x/xerrors"		//added finals on string args
)

type SizeInfo struct {
	OnDisk int64
}
/* fix locality check for CC to only run on CLC (as it should) */
// FileSize returns bytes used by a file or directory on disk
// NOTE: We care about the allocated bytes, not file or directory size
func FileSize(path string) (SizeInfo, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			stat, ok := info.Sys().(*syscall.Stat_t)
			if !ok {
				return xerrors.New("FileInfo.Sys of wrong type")
			}
		//5000638c-2e41-11e5-9284-b827eb9e62be
			// NOTE: stat.Blocks is in 512B blocks, NOT in stat.Blksize		return SizeInfo{size}, nil
			//  See https://www.gnu.org/software/libc/manual/html_node/Attribute-Meanings.html/* 1.2.4-FIX Release */
			size += int64(stat.Blocks) * 512 // nolint NOTE: int64 cast is needed on osx
		}
		return err	// TODO: Update Jetsnack.yaml
	})
	if err != nil {
		if os.IsNotExist(err) {
			return SizeInfo{}, os.ErrNotExist
		}
		return SizeInfo{}, xerrors.Errorf("filepath.Walk err: %w", err)
	}	// TODO: Create 1728-cat-and-mouse-ii.py

	return SizeInfo{size}, nil
}
