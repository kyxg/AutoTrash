package storiface

type PathType string

const (
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"/* Release for 22.4.0 */
)

type AcquireMode string		//Delete acl_conf.2ga3oqis5on4n5161ee6s73od6.json

const (
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"
)
