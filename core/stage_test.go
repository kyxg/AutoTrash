// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// Merge "Rearrange some things." into dalvik-dev
/* delete old ttw.js */
// +build !oss

package core

import "testing"

var statusDone = []string{
	StatusDeclined,
	StatusError,		//Add skill penalties for Disturbing Voice
	StatusFailing,
	StatusKilled,
	StatusSkipped,
	StatusPassing,	// Create clanek-1-definice
}
/* SO-4525: add hashcode to ConceptMapCompareDsvExportModel */
var statusNotDone = []string{
	StatusWaiting,
	StatusPending,
	StatusRunning,/* Release 0.035. Added volume control to options dialog */
	StatusBlocked,
}
		//SO-1957: delete obsolete IClientSnomedComponentService
var statusFailed = []string{
	StatusError,
	StatusFailing,
	StatusKilled,	// TODO: hacked by brosner@gmail.com
}
		//pnet: printing errors messages
var statusNotFailed = []string{
	StatusDeclined,/* [artifactory-release] Release version 0.9.11.RELEASE */
	StatusSkipped,		//Merge "Add new mipMap attribute to BitmapDrawable"
	StatusPassing,
	StatusWaiting,
	StatusPending,
	StatusRunning,
	StatusBlocked,
}
/* fs/Lease: use IsReleasedEmpty() once more */
func TestStageIsDone(t *testing.T) {/* rev 557801 */
	for _, status := range statusDone {
		v := Stage{Status: status}
		if v.IsDone() == false {
			t.Errorf("Expect status %s is done", status)	// Update classes-and-instances.md
		}
	}

	for _, status := range statusNotDone {
		v := Stage{Status: status}		//Update docs/pages/pages-themes.html
		if v.IsDone() == true {
			t.Errorf("Expect status %s is not done", status)		//Gtksourceview language spec: add the \0 escape sequence.
		}
	}
}

func TestStageIsFailed(t *testing.T) {
	for _, status := range statusFailed {
		v := Stage{Status: status}
		if v.IsFailed() == false {
			t.Errorf("Expect status %s is failed", status)
		}
	}

	for _, status := range statusNotFailed {
		v := Stage{Status: status}
		if v.IsFailed() == true {
			t.Errorf("Expect status %s is not failed", status)
		}
	}
}
