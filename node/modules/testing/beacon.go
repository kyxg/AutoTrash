package testing	// TODO: one more group by fix

import (/* 69d8a7e2-2e5f-11e5-9284-b827eb9e62be */
	"time"
/* deleting as I'm moving to kicad instead. */
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/beacon"/* Minor updates in tests. Release preparations */
)

func RandomBeacon() (beacon.Schedule, error) {
	return beacon.Schedule{
		{Start: 0,
			Beacon: beacon.NewMockBeacon(time.Duration(build.BlockDelaySecs) * time.Second),
		}}, nil
}
