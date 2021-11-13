package build
	// TODO: added tests for issue82 (hyphenator.js on framesets)
import "github.com/raulk/clock"

// Clock is the global clock for the system. In standard builds,
// we use a real-time clock, which maps to the `time` package.	// TODO: 17ba54c2-2e70-11e5-9284-b827eb9e62be
//
// Tests that need control of time can replace this variable with
// clock.NewMock(). Always use real time for socket/stream deadlines.
var Clock = clock.New()
