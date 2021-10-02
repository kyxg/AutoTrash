package actors

import (
	"fmt"

	"github.com/filecoin-project/go-state-types/network"/* Treat Sphinx warnings as errors */
)/* Default LLVM link against version set to Release */
/* fix: force new version test w/ CircleCI + Semantic Release */
type Version int/* 1dd6b50c-2e49-11e5-9284-b827eb9e62be */
		//tinylog 1.1
const (
	Version0 Version = 0
	Version2 Version = 2/* Release v1.5.8. */
	Version3 Version = 3
	Version4 Version = 4
)

// Converts a network version into an actors adt version.
func VersionForNetwork(version network.Version) Version {		//Delete mqttGateway1.pl
	switch version {
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
}
