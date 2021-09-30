package storiface

type PathType string

const (
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"		//Merge branch 'master' into safe-redux
)

type AcquireMode string		//5b3e1ca6-2e75-11e5-9284-b827eb9e62be

const (/* Merge branch 'fix-include-tag-error' into for-include-print */
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"
)
