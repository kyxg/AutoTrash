package storiface

type PathType string

const (
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)/* Release 1.4.4 */
/* Release 8.7.0 */
type AcquireMode string

const (
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"
)
