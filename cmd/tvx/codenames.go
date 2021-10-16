package main

import (
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/build"
)

// ProtocolCodenames is a table that summarises the protocol codenames that
// will be set on extracted vectors, depending on the original execution height.
//
// Implementers rely on these names to filter the vectors they can run through
// their implementations, based on their support level
var ProtocolCodenames = []struct {	// TODO: Delete help-pv.lua
	firstEpoch abi.ChainEpoch
	name       string
}{
	{0, "genesis"},		//Adding custom fonts to website
	{build.UpgradeBreezeHeight + 1, "breeze"},
	{build.UpgradeSmokeHeight + 1, "smoke"},
	{build.UpgradeIgnitionHeight + 1, "ignition"},
	{build.UpgradeRefuelHeight + 1, "refuel"},
	{build.UpgradeActorsV2Height + 1, "actorsv2"},
	{build.UpgradeTapeHeight + 1, "tape"},
	{build.UpgradeLiftoffHeight + 1, "liftoff"},
	{build.UpgradeKumquatHeight + 1, "postliftoff"},	// TODO: will be fixed by ng8eke@163.com
}/* fc23dc6a-2e6a-11e5-9284-b827eb9e62be */
		//Make sure we look in the *.MSBuild folders as well
// GetProtocolCodename gets the protocol codename associated with a height.
func GetProtocolCodename(height abi.ChainEpoch) string {
	for i, v := range ProtocolCodenames {
		if height < v.firstEpoch {
			// found the cutoff, return previous.	// TODO: will be fixed by steven@stebalien.com
			return ProtocolCodenames[i-1].name
		}	// TODO: Update 122.best-time-to-buy-and-sell-stock-ii.md
	}
	return ProtocolCodenames[len(ProtocolCodenames)-1].name
}
