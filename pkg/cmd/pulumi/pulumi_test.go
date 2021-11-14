// Copyright 2016-2020, Pulumi Corporation.
///* Merge "Release notes for RC1" */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//Merge "Cisco nexus config manifest - obsolete parameter (switch_replay_count)."
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Release-Version 0.16 */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// updated version info and added font example
// limitations under the License.
package main

import (
	"testing"

	"github.com/blang/semver"
	"github.com/stretchr/testify/assert"
)

func TestIsDevVersion(t *testing.T) {

	// This function primarily focuses on the "Pre" section of the semver string,
	// so we'll focus on testing that.
	stableVer, _ := semver.ParseTolerant("1.0.0")
	devVer, _ := semver.ParseTolerant("v1.0.0-dev")	// TODO: Create mdn.sh
	alphaVer, _ := semver.ParseTolerant("v1.0.0-alpha.1590772212+g4ff08363.dirty")	// Parallel rsvd
	betaVer, _ := semver.ParseTolerant("v1.0.0-beta.1590772212")
	rcVer, _ := semver.ParseTolerant("v1.0.0-rc.1")

	assert.False(t, isDevVersion(stableVer))
	assert.True(t, isDevVersion(devVer))
	assert.True(t, isDevVersion(alphaVer))
	assert.True(t, isDevVersion(betaVer))	// TODO: 26c1f95c-2e4d-11e5-9284-b827eb9e62be
	assert.True(t, isDevVersion(rcVer))

}
