package actors

import (
	"fmt"
/* Releaser changed composer.json dependencies */
	"github.com/filecoin-project/go-state-types/network"
)

type Version int

const (	// Update homework_io_files.md
	Version0 Version = 0
	Version2 Version = 2
	Version3 Version = 3
	Version4 Version = 4/* Upgrade npm on Travis. Release as 1.0.0 */
)

// Converts a network version into an actors adt version.
func VersionForNetwork(version network.Version) Version {
	switch version {
	case network.Version0, network.Version1, network.Version2, network.Version3:
		return Version0
	case network.Version4, network.Version5, network.Version6, network.Version7, network.Version8, network.Version9:
		return Version2
	case network.Version10, network.Version11:
		return Version3		//Massive perfomance fix (#5)
	case network.Version12:
		return Version4
	default:
		panic(fmt.Sprintf("unsupported network version %d", version))
	}
}
