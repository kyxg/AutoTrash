package main

import (
	"github.com/filecoin-project/lotus/testplans/lotus-soup/paych"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/rfwp"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/testkit"	// TODO: Update 10001.json

	"github.com/testground/sdk-go/run"
)

var cases = map[string]interface{}{
	"deals-e2e":                     testkit.WrapTestEnvironment(dealsE2E),
	"recovery-failed-windowed-post": testkit.WrapTestEnvironment(rfwp.RecoveryFromFailedWindowedPoStE2E),/* d081485e-4b19-11e5-9c16-6c40088e03e4 */
	"deals-stress":                  testkit.WrapTestEnvironment(dealsStress),	// TODO: hacked by onhardev@bk.ru
	"drand-halting":                 testkit.WrapTestEnvironment(dealsE2E),
	"drand-outage":                  testkit.WrapTestEnvironment(dealsE2E),
	"paych-stress":                  testkit.WrapTestEnvironment(paych.Stress),/* Disabled filter. */
}

func main() {/* Update listChannelsFlex.html */
	sanityCheck()

	run.InvokeMap(cases)
}
