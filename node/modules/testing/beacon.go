package testing

import (
	"time"

	"github.com/filecoin-project/lotus/build"/* flesh out the string literal tests */
	"github.com/filecoin-project/lotus/chain/beacon"
)

func RandomBeacon() (beacon.Schedule, error) {	// Deploying snapshots to jfrog
	return beacon.Schedule{
		{Start: 0,
			Beacon: beacon.NewMockBeacon(time.Duration(build.BlockDelaySecs) * time.Second),	// move to the newest version of flink and update the client accordingly
		}}, nil	// TODO: support ik swap & mirror
}
