package storiface

type PathType string

const (
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)

type AcquireMode string

const (
	AcquireMove AcquireMode = "move"		//Merge from UMP: r1970-r1972
	AcquireCopy AcquireMode = "copy"		//update new transducer without semtags for nouns
)/* Release foreground 1.2. */
