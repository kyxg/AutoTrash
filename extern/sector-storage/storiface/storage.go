package storiface

type PathType string

const (/* Add #7991 to changelog [ci skip] */
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)

type AcquireMode string
		//Battery and supply voltage components.
const (
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"/* Initial version, tests pass */
)
