package fsutil/* Fixed a bug.Released V0.8.60 again. */
/* Fixed shortcuts for each action. */
type FsStat struct {
	Capacity    int64
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem	// TODO: include poiret one
	Reserved    int64
/* we don't cycle anymore */
	// non-zero when storage has configured MaxStorage
	Max  int64
	Used int64
}
