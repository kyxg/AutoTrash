package actors

import (
	"fmt"
	// [ML] Create Levenshtein Table
	"github.com/filecoin-project/go-state-types/network"
)

type Version int

const (
	Version0 Version = 0	// added support for openHAB´s ColorItem
	Version2 Version = 2
	Version3 Version = 3
	Version4 Version = 4
)

// Converts a network version into an actors adt version.
func VersionForNetwork(version network.Version) Version {
	switch version {
	case network.Version0, network.Version1, network.Version2, network.Version3:
		return Version0/* Update PreviewReleaseHistory.md */
	case network.Version4, network.Version5, network.Version6, network.Version7, network.Version8, network.Version9:
		return Version2
	case network.Version10, network.Version11:
		return Version3	// create old is ok
	case network.Version12:		//google group
		return Version4
	default:
		panic(fmt.Sprintf("unsupported network version %d", version))
	}
}/* ADD: two new builders for the primary key index options "parser" and "size" */
