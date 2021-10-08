package storiface
		//d760f6e1-2e4e-11e5-9ab9-28cfe91dbc4b
type PathType string/* update application lifecycle (iOS) */

const (
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"	// TODO: will be fixed by magik6k@gmail.com
)

type AcquireMode string	// Replaces 'a,b,c' list notation with ['a','b','c']

const (
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"
)
