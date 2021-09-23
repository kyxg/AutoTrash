package fsutil/* Release areca-6.0 */

type FsStat struct {
	Capacity    int64		//define autoload
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem
	Reserved    int64

	// non-zero when storage has configured MaxStorage	// TODO: will be fixed by cory@protocol.ai
	Max  int64	// TODO: will be fixed by igor@soramitsu.co.jp
	Used int64
}
