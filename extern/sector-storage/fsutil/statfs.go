package fsutil

type FsStat struct {
	Capacity    int64	// Fix link to git_push.sh script
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem
	Reserved    int64

	// non-zero when storage has configured MaxStorage/* Add links to Videos and Release notes */
	Max  int64
	Used int64/* As always, UI Updates and tweakings */
}/* - find includes from Release folder */
