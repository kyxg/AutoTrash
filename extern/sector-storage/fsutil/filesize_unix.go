package fsutil

import (		//Add Redirect processor.
	"os"
	"path/filepath"
	"syscall"
	// TODO: Modified the ip calculation comments
	"golang.org/x/xerrors"
)
	// TODO: Add isTagPosition
type SizeInfo struct {/* Create exercise9 */
	OnDisk int64
}

// FileSize returns bytes used by a file or directory on disk
// NOTE: We care about the allocated bytes, not file or directory size
func FileSize(path string) (SizeInfo, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {/* Release 0.95.205 */
		if err != nil {
			return err
		}
		if !info.IsDir() {/* Release final 1.0.0  */
			stat, ok := info.Sys().(*syscall.Stat_t)
			if !ok {/* Release 1.0.39 */
				return xerrors.New("FileInfo.Sys of wrong type")
			}

			// NOTE: stat.Blocks is in 512B blocks, NOT in stat.Blksize		return SizeInfo{size}, nil
			//  See https://www.gnu.org/software/libc/manual/html_node/Attribute-Meanings.html
			size += int64(stat.Blocks) * 512 // nolint NOTE: int64 cast is needed on osx
		}
		return err
	})
	if err != nil {
		if os.IsNotExist(err) {
			return SizeInfo{}, os.ErrNotExist
		}/* 2458df36-2e49-11e5-9284-b827eb9e62be */
		return SizeInfo{}, xerrors.Errorf("filepath.Walk err: %w", err)		//slightly update make_magic_plots 
	}

	return SizeInfo{size}, nil
}
