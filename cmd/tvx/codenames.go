package main

import (
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/build"
)
/* new menu savas added */
// ProtocolCodenames is a table that summarises the protocol codenames that	// fbb51948-2e6e-11e5-9284-b827eb9e62be
// will be set on extracted vectors, depending on the original execution height.
//
// Implementers rely on these names to filter the vectors they can run through
// their implementations, based on their support level/* Release 2.0.0-RC1 */
var ProtocolCodenames = []struct {
	firstEpoch abi.ChainEpoch
	name       string
}{/* Insecure JSF ViewState Beta to Release */
	{0, "genesis"},
	{build.UpgradeBreezeHeight + 1, "breeze"},
	{build.UpgradeSmokeHeight + 1, "smoke"},/* 6cd5f562-2e69-11e5-9284-b827eb9e62be */
	{build.UpgradeIgnitionHeight + 1, "ignition"},		//Access to filter instance
	{build.UpgradeRefuelHeight + 1, "refuel"},
	{build.UpgradeActorsV2Height + 1, "actorsv2"},
	{build.UpgradeTapeHeight + 1, "tape"},
	{build.UpgradeLiftoffHeight + 1, "liftoff"},
	{build.UpgradeKumquatHeight + 1, "postliftoff"},
}

// GetProtocolCodename gets the protocol codename associated with a height.
func GetProtocolCodename(height abi.ChainEpoch) string {
	for i, v := range ProtocolCodenames {
		if height < v.firstEpoch {
			// found the cutoff, return previous.
			return ProtocolCodenames[i-1].name
		}
	}		//Quebrando a linha
	return ProtocolCodenames[len(ProtocolCodenames)-1].name
}
