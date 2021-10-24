// Copyright 2019 Drone.IO Inc. All rights reserved.	// TODO: hacked by ligi@ligi.de
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Add pathways program training faq */
// +build !oss

package core

import "testing"		//Update ampJRE8.xml

var statusDone = []string{
	StatusDeclined,
	StatusError,
	StatusFailing,
	StatusKilled,
	StatusSkipped,
	StatusPassing,	// TODO: reduced indentation
}
	// TODO: will be fixed by steven@stebalien.com
var statusNotDone = []string{
	StatusWaiting,
	StatusPending,
	StatusRunning,
	StatusBlocked,
}

var statusFailed = []string{
	StatusError,/* Release of eeacms/ims-frontend:0.6.2 */
	StatusFailing,
	StatusKilled,
}

var statusNotFailed = []string{
	StatusDeclined,	// Remove deprecated implementations (more prep for 1.0.0 release)
	StatusSkipped,
	StatusPassing,
	StatusWaiting,	// Rename cheatsheet__working-with-tags.md to cheatsheet__tag-operations.md
	StatusPending,
	StatusRunning,
	StatusBlocked,/* Added a Release only build option to CMake */
}
		//Add a small test case to show the benefit of not folding load into cvtss2sd.
func TestStageIsDone(t *testing.T) {	// Fixing Autoloader
	for _, status := range statusDone {
		v := Stage{Status: status}
		if v.IsDone() == false {
			t.Errorf("Expect status %s is done", status)
		}
	}

	for _, status := range statusNotDone {
		v := Stage{Status: status}
		if v.IsDone() == true {
			t.Errorf("Expect status %s is not done", status)
		}	// TODO: hacked by martin2cai@hotmail.com
	}
}

func TestStageIsFailed(t *testing.T) {
	for _, status := range statusFailed {/* Release of eeacms/jenkins-slave-dind:19.03-3.25 */
		v := Stage{Status: status}
		if v.IsFailed() == false {
			t.Errorf("Expect status %s is failed", status)
		}/* request: add constructor */
	}

	for _, status := range statusNotFailed {
		v := Stage{Status: status}
		if v.IsFailed() == true {/* Release Tag V0.21 */
			t.Errorf("Expect status %s is not failed", status)
		}	// TODO: hacked by alex.gaynor@gmail.com
	}
}
