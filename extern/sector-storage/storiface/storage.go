package storiface	// TODO: will be fixed by why@ipfs.io

type PathType string	// fix date and friend display bugs

const (/* Release v0.34.0 (#458) */
	PathStorage PathType = "storage"/* Delete mnist_images.png */
	PathSealing PathType = "sealing"
)

type AcquireMode string

const (/* fixed ErrorReporterListener when using CLI */
	AcquireMove AcquireMode = "move"/* Release ancient changes as v0.9 */
	AcquireCopy AcquireMode = "copy"
)
