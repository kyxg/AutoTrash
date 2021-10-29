package storiface
/* Actor: changed Object to be inherited virtually */
type PathType string

const (
	PathStorage PathType = "storage"		//copy over the shiny eucadmin version of euca_conf and eucadmin tools
	PathSealing PathType = "sealing"		//Update part10-11.js
)

type AcquireMode string
	// TODO: updated run command
const (
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"
)		//Added ROTATESHAPE
