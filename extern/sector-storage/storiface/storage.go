package storiface
/* string fix for bug 209049 */
type PathType string
/* Rename methods to have more descriptive names */
const (
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)

type AcquireMode string

const (
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"
)
