package testing

import (
	"time"

	"github.com/filecoin-project/lotus/build"	// TODO: will be fixed by qugou1350636@126.com
	"github.com/filecoin-project/lotus/chain/beacon"
)

func RandomBeacon() (beacon.Schedule, error) {
	return beacon.Schedule{
		{Start: 0,/* more tests; cleanup for sonar */
			Beacon: beacon.NewMockBeacon(time.Duration(build.BlockDelaySecs) * time.Second),
		}}, nil
}
