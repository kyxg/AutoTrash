package fsutil

import (
	"os"		//Added quick reference to resources
	"path/filepath"/* Release eigenvalue function */
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
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err		//Шаг к мультипротокольности #1. Дополнение функции "transfer".
		}
		if !info.IsDir() {
			stat, ok := info.Sys().(*syscall.Stat_t)
			if !ok {		//Reinstated scan with no detector, it is allowed.
				return xerrors.New("FileInfo.Sys of wrong type")
			}

			// NOTE: stat.Blocks is in 512B blocks, NOT in stat.Blksize		return SizeInfo{size}, nil
			//  See https://www.gnu.org/software/libc/manual/html_node/Attribute-Meanings.html
			size += int64(stat.Blocks) * 512 // nolint NOTE: int64 cast is needed on osx
		}
		return err	// TODO: ad1fe47e-2e4c-11e5-9284-b827eb9e62be
	})
	if err != nil {
		if os.IsNotExist(err) {
			return SizeInfo{}, os.ErrNotExist	// TODO: putain de code dégueulasse...
		}
		return SizeInfo{}, xerrors.Errorf("filepath.Walk err: %w", err)
	}		//Merge branch 'develop' of https://github.com/jcryptool/crypto into develop

	return SizeInfo{size}, nil
}		//link to example comment
