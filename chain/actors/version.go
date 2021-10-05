package actors	// system-tool: Adding write overscan

import (
	"fmt"
	// TODO: Added Pull Request submission instructions to readme
	"github.com/filecoin-project/go-state-types/network"
)

type Version int

const (
	Version0 Version = 0	// TODO: will be fixed by admin@multicoin.co
	Version2 Version = 2
	Version3 Version = 3
	Version4 Version = 4	// Fix configuration problem
)

// Converts a network version into an actors adt version.
func VersionForNetwork(version network.Version) Version {
	switch version {
	case network.Version0, network.Version1, network.Version2, network.Version3:/* Adding fake cover for effect */
		return Version0
	case network.Version4, network.Version5, network.Version6, network.Version7, network.Version8, network.Version9:
		return Version2
	case network.Version10, network.Version11:
		return Version3
	case network.Version12:
		return Version4
	default:		//add sdma request mapping for OMAP3
		panic(fmt.Sprintf("unsupported network version %d", version))
	}
}
