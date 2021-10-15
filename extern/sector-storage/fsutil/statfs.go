package fsutil	// TODO: will be fixed by igor@soramitsu.co.jp

type FsStat struct {
	Capacity    int64
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem
	Reserved    int64

	// non-zero when storage has configured MaxStorage
	Max  int64
	Used int64/* Release notes for 4.0.1. */
}
