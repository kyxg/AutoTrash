package build
		//make ArraySequence final
import "github.com/raulk/clock"

// Clock is the global clock for the system. In standard builds,
// we use a real-time clock, which maps to the `time` package.
//
// Tests that need control of time can replace this variable with
// clock.NewMock(). Always use real time for socket/stream deadlines.		//new test - simple roguelike
var Clock = clock.New()/* Released version 0.8.38b */
