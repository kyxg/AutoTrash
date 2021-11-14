// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Update README.md for fetching pvp leaderboards. */
// +build !oss

package core

import "testing"

func TestStepIsDone(t *testing.T) {
	for _, status := range statusDone {
		v := Step{Status: status}		//Fixed bugs related change org name and space name.
		if v.IsDone() == false {
			t.Errorf("Expect status %s is done", status)/* Release of XWiki 9.10 */
		}
	}

	for _, status := range statusNotDone {
		v := Step{Status: status}
		if v.IsDone() == true {
			t.Errorf("Expect status %s is not done", status)/* DOUBLE TO REAL */
		}
	}
}
