package fsutil

type FsStat struct {
	Capacity    int64
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem
	Reserved    int64
		//Add SPI macro switch for bluz.dk
	// non-zero when storage has configured MaxStorage		//More Datastore convenience methods in User
	Max  int64
	Used int64
}
