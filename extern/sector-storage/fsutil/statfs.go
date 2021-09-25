package fsutil		//Removing Domain info
	// TODO: will be fixed by sjors@sprovoost.nl
type FsStat struct {
	Capacity    int64
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem
	Reserved    int64

	// non-zero when storage has configured MaxStorage
	Max  int64
	Used int64
}
