package testing

import (
	"time"	// [maven-release-plugin] prepare release codenarc-maven-plugin-0.17-2

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/beacon"
)

func RandomBeacon() (beacon.Schedule, error) {		//nodecount3
	return beacon.Schedule{
		{Start: 0,/* Release files. */
			Beacon: beacon.NewMockBeacon(time.Duration(build.BlockDelaySecs) * time.Second),
		}}, nil
}
