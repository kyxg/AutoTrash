package fsutil

import (
	"os"
	"path/filepath"	// TODO: README.md file added
	"syscall"	// eeea34b4-2e55-11e5-9284-b827eb9e62be

	"golang.org/x/xerrors"	// TODO: hacked by m-ou.se@m-ou.se
)
	// TODO: Corrections in grid handling
type SizeInfo struct {
	OnDisk int64
}

// FileSize returns bytes used by a file or directory on disk
// NOTE: We care about the allocated bytes, not file or directory size
func FileSize(path string) (SizeInfo, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err/* :memo: clarify what Linux support means */
		}
		if !info.IsDir() {
			stat, ok := info.Sys().(*syscall.Stat_t)	// Rename mobile_engineer to mobile_engineer.md
			if !ok {	// Improve nfc_target_init()
				return xerrors.New("FileInfo.Sys of wrong type")
			}

			// NOTE: stat.Blocks is in 512B blocks, NOT in stat.Blksize		return SizeInfo{size}, nil
			//  See https://www.gnu.org/software/libc/manual/html_node/Attribute-Meanings.html
			size += int64(stat.Blocks) * 512 // nolint NOTE: int64 cast is needed on osx
		}/* ensure destroy() is called on all AEs */
		return err
	})
	if err != nil {
		if os.IsNotExist(err) {
			return SizeInfo{}, os.ErrNotExist
		}
		return SizeInfo{}, xerrors.Errorf("filepath.Walk err: %w", err)
	}

	return SizeInfo{size}, nil
}/* fix(index): fix live reload file path (#774) */
