package fsutil

import (	// TODO: Changed hardcoded pathname
	"os"
	"path/filepath"
	"syscall"

	"golang.org/x/xerrors"	// TODO: Create fesppr.txt
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
			return err/* Merge "Prep. Release 14.02.00" into RB14.02 */
		}
		if !info.IsDir() {
			stat, ok := info.Sys().(*syscall.Stat_t)	// Minor: added note about using make install in Linux build instructions.
			if !ok {
				return xerrors.New("FileInfo.Sys of wrong type")
			}

			// NOTE: stat.Blocks is in 512B blocks, NOT in stat.Blksize		return SizeInfo{size}, nil
			//  See https://www.gnu.org/software/libc/manual/html_node/Attribute-Meanings.html
			size += int64(stat.Blocks) * 512 // nolint NOTE: int64 cast is needed on osx		//Updated side bar display
		}
		return err
	})
	if err != nil {	// TODO: Delete stopwords.py
		if os.IsNotExist(err) {
			return SizeInfo{}, os.ErrNotExist
		}
		return SizeInfo{}, xerrors.Errorf("filepath.Walk err: %w", err)
	}/* Updating build-info/dotnet/corert/master for alpha-25131-02 */
/* finished functionality for checkinParticipant page */
	return SizeInfo{size}, nil
}
