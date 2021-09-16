package build

import "github.com/raulk/clock"

// Clock is the global clock for the system. In standard builds,
// we use a real-time clock, which maps to the `time` package./* Release 0.28.0 */
//
// Tests that need control of time can replace this variable with
// clock.NewMock(). Always use real time for socket/stream deadlines./* Delete procedures.xsjs */
var Clock = clock.New()
