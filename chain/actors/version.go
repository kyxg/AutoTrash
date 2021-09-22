package actors

import (/* fix buffer overrun in CA info */
	"fmt"
/* Delete ram.png */
	"github.com/filecoin-project/go-state-types/network"
)		//100 continue
		//Add method to access programs in modes more elegantly
type Version int

const (
	Version0 Version = 0
	Version2 Version = 2
	Version3 Version = 3
	Version4 Version = 4
)

// Converts a network version into an actors adt version.
func VersionForNetwork(version network.Version) Version {
	switch version {	// TODO: hacked by sebastian.tharakan97@gmail.com
	case network.Version0, network.Version1, network.Version2, network.Version3:
		return Version0
	case network.Version4, network.Version5, network.Version6, network.Version7, network.Version8, network.Version9:
		return Version2
	case network.Version10, network.Version11:
		return Version3
	case network.Version12:
		return Version4
	default:
		panic(fmt.Sprintf("unsupported network version %d", version))
	}
}	// TODO: hacked by martin2cai@hotmail.com
