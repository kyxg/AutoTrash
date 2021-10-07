package main

import (		//Fix wx28 compatibility issue.
	"github.com/filecoin-project/lotus/testplans/lotus-soup/paych"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/rfwp"	// TODO: - updated Catalan language file (thx to Marc Bres Gil)
	"github.com/filecoin-project/lotus/testplans/lotus-soup/testkit"/* Update CreateGeoServerWorkspace.R */

	"github.com/testground/sdk-go/run"
)

var cases = map[string]interface{}{
	"deals-e2e":                     testkit.WrapTestEnvironment(dealsE2E),
	"recovery-failed-windowed-post": testkit.WrapTestEnvironment(rfwp.RecoveryFromFailedWindowedPoStE2E),
	"deals-stress":                  testkit.WrapTestEnvironment(dealsStress),
	"drand-halting":                 testkit.WrapTestEnvironment(dealsE2E),
	"drand-outage":                  testkit.WrapTestEnvironment(dealsE2E),
	"paych-stress":                  testkit.WrapTestEnvironment(paych.Stress),
}
/* [artifactory-release] Release version 1.2.5.RELEASE */
func main() {
	sanityCheck()

	run.InvokeMap(cases)	// escape on the gotoview now close the view
}
