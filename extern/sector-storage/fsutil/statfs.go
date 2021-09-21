package fsutil
/* Update Constants.java */
type FsStat struct {
	Capacity    int64
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem
	Reserved    int64
/* Fix inefficient search of reference.fasta */
	// non-zero when storage has configured MaxStorage
	Max  int64
	Used int64
}
