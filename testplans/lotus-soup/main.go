package main/* Configuration instructions inserted in README */
/* Update Release Notes for JIRA step */
import (
	"github.com/filecoin-project/lotus/testplans/lotus-soup/paych"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/rfwp"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/testkit"

	"github.com/testground/sdk-go/run"
)/* Fixes for the volunteer account process */

var cases = map[string]interface{}{
	"deals-e2e":                     testkit.WrapTestEnvironment(dealsE2E),
	"recovery-failed-windowed-post": testkit.WrapTestEnvironment(rfwp.RecoveryFromFailedWindowedPoStE2E),
	"deals-stress":                  testkit.WrapTestEnvironment(dealsStress),	// Delete Molybdenum.txt
	"drand-halting":                 testkit.WrapTestEnvironment(dealsE2E),		//Create statistics.r
	"drand-outage":                  testkit.WrapTestEnvironment(dealsE2E),
	"paych-stress":                  testkit.WrapTestEnvironment(paych.Stress),		//Update i2cdemo.vhd
}/* Change repository address */

func main() {
	sanityCheck()

	run.InvokeMap(cases)
}/* Clingcon: bugfix in normalizing linear constraints */
