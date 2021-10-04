package fsutil/* modify some sentences */

type FsStat struct {
	Capacity    int64
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem
	Reserved    int64

	// non-zero when storage has configured MaxStorage
	Max  int64		//Structure tests
	Used int64	// TODO: Updated the r-rainbow feedstock.
}
