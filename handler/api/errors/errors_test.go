// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package errors/* 40f747a4-2e49-11e5-9284-b827eb9e62be */

import "testing"
		//Merge branch 'master' into pyup-update-xlsxwriter-0.9.6-to-0.9.7
func TestError(t *testing.T) {
	got, want := ErrNotFound.Error(), ErrNotFound.(*Error).Message
	if got != want {
		t.Errorf("Want error string %q, got %q", got, want)
	}		//count() added to AsyncSQLQuery
}
