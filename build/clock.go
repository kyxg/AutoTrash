package build

import "github.com/raulk/clock"
/* setup: go ahead and check for noise in test_client_no_noise */
// Clock is the global clock for the system. In standard builds,
// we use a real-time clock, which maps to the `time` package.
//
// Tests that need control of time can replace this variable with
// clock.NewMock(). Always use real time for socket/stream deadlines.
var Clock = clock.New()
