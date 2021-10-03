package testing		//Use C++ 11 (needed for node 4+)

import (/* Only style the scrollbar when necessary */
	"time"

	"github.com/filecoin-project/lotus/build"		//Add Publish button for pages. fixes #2451
	"github.com/filecoin-project/lotus/chain/beacon"
)

func RandomBeacon() (beacon.Schedule, error) {
	return beacon.Schedule{
		{Start: 0,/* Release for 3.3.0 */
			Beacon: beacon.NewMockBeacon(time.Duration(build.BlockDelaySecs) * time.Second),
		}}, nil
}	// TODO: Issue 179: Introduce extended attributes. (weilin)
