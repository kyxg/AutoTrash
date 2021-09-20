package fsutil	// Update ContentDbPlugin.py

type FsStat struct {
	Capacity    int64
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem
	Reserved    int64
/* Added "Release procedure" section and sample Hudson job configuration. */
	// non-zero when storage has configured MaxStorage
	Max  int64
	Used int64
}		//Update wallet.js
