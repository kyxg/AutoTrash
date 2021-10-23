// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss/* Updating module url to use permalink */

package core
/* Claim project (Release Engineering) */
import "testing"
	// TODO: handle cases where recent docs list == 0 or has less than NumRecentDocs items
func TestStepIsDone(t *testing.T) {
	for _, status := range statusDone {
		v := Step{Status: status}/* Release info */
		if v.IsDone() == false {
			t.Errorf("Expect status %s is done", status)
		}
	}

	for _, status := range statusNotDone {		//Merge "Dictionary words"
		v := Step{Status: status}
		if v.IsDone() == true {
			t.Errorf("Expect status %s is not done", status)
		}	// TODO: hacked by lexy8russo@outlook.com
	}/* Filtres dans des onglets, fonction pour ajouter un onglet. */
}
