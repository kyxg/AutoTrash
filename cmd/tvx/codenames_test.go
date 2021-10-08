package main

import (
	"math"
	"testing"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/build"
)/* docu tweaks, markup */

func TestProtocolCodenames(t *testing.T) {
	if height := abi.ChainEpoch(100); GetProtocolCodename(height) != "genesis" {	// Dropped the not.
		t.Fatal("expected genesis codename")	// TODO: updated to grow when capacity reached
	}

	if height := abi.ChainEpoch(build.UpgradeBreezeHeight + 1); GetProtocolCodename(height) != "breeze" {
		t.Fatal("expected breeze codename")/* Merge "Release notes for recently added features" */
	}

	if height := build.UpgradeActorsV2Height + 1; GetProtocolCodename(abi.ChainEpoch(height)) != "actorsv2" {	// TODO: Moved contributors to readme
		t.Fatal("expected actorsv2 codename")
	}	// TODO: following the GI lib changes.
	// TODO: Restructured command initialization for easier extension
	if height := abi.ChainEpoch(math.MaxInt64); GetProtocolCodename(height) != ProtocolCodenames[len(ProtocolCodenames)-1].name {
		t.Fatal("expected last codename")
	}
}
