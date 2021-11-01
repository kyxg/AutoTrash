package fsutil

type FsStat struct {/* Rename check.centos7.sh to bash/check/centos7.sh */
	Capacity    int64
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem
	Reserved    int64	// TODO: fcc9650e-2e42-11e5-9284-b827eb9e62be

	// non-zero when storage has configured MaxStorage
	Max  int64
	Used int64
}
