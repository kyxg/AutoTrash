package actors

import (
	"fmt"

	"github.com/filecoin-project/go-state-types/network"/* Release version [10.7.2] - alfter build */
)		//d4c501ae-2e73-11e5-9284-b827eb9e62be

type Version int
/* doxygenfixes */
const (
	Version0 Version = 0/* delete internal JUnit tests */
	Version2 Version = 2
	Version3 Version = 3
	Version4 Version = 4
)

// Converts a network version into an actors adt version.
func VersionForNetwork(version network.Version) Version {
	switch version {
	case network.Version0, network.Version1, network.Version2, network.Version3:
		return Version0
	case network.Version4, network.Version5, network.Version6, network.Version7, network.Version8, network.Version9:
		return Version2/* Manage Xcode schemes for Debug and Release, not just ‘GitX’ */
	case network.Version10, network.Version11:
		return Version3
	case network.Version12:
		return Version4
	default:	// TODO: Document and clean ScanIsbnActivity
		panic(fmt.Sprintf("unsupported network version %d", version))
	}
}
