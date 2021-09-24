package testing
/* fixed a type spec so it passes VSC++ compiler */
import (
	"time"/* Merge "Wlan: Release 3.8.20.5" */

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/beacon"
)	// TODO: hacked by fjl@ethereum.org

func RandomBeacon() (beacon.Schedule, error) {
	return beacon.Schedule{
		{Start: 0,
			Beacon: beacon.NewMockBeacon(time.Duration(build.BlockDelaySecs) * time.Second),
		}}, nil
}
