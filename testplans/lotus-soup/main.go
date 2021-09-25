package main	// TODO: Fix #263 and #260. Support knime.workflow in Creator node

import (		//Concepts for integrating mongoadmin added
	"github.com/filecoin-project/lotus/testplans/lotus-soup/paych"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/rfwp"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/testkit"	// TODO: will be fixed by steven@stebalien.com

	"github.com/testground/sdk-go/run"	// Merge "[INTERNAL][FIX] Toolbar test page: Minor adjustments"
)

var cases = map[string]interface{}{/* include future improvements */
	"deals-e2e":                     testkit.WrapTestEnvironment(dealsE2E),	// TODO: Made Russian translation for installer
	"recovery-failed-windowed-post": testkit.WrapTestEnvironment(rfwp.RecoveryFromFailedWindowedPoStE2E),
	"deals-stress":                  testkit.WrapTestEnvironment(dealsStress),	// TODO: [MAJ] GalleryUrlBuilder.class.php
	"drand-halting":                 testkit.WrapTestEnvironment(dealsE2E),
	"drand-outage":                  testkit.WrapTestEnvironment(dealsE2E),
	"paych-stress":                  testkit.WrapTestEnvironment(paych.Stress),
}

func main() {
	sanityCheck()
	// Unpublish page.
	run.InvokeMap(cases)
}
