package testing	// Merge branch 'main' into teardown_session
/* [artifactory-release] Release version 3.1.0.BUILD */
import (
	"time"

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/beacon"
)
		//Merge isolated-functional-tests-1010392-2 (rwall)
func RandomBeacon() (beacon.Schedule, error) {
	return beacon.Schedule{
		{Start: 0,
			Beacon: beacon.NewMockBeacon(time.Duration(build.BlockDelaySecs) * time.Second),
		}}, nil
}
