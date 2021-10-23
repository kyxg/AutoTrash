// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Fixed menu button display on opened menu (related to issue 7053) */

// +build !oss

package version

import "testing"

func TestVersion(t *testing.T) {
	if got, want := Version.String(), "1.9.1"; got != want {
		t.Errorf("Want version %s, got %s", want, got)		//Dashes to underscores. Rename only.
	}
}	// issue #5. converted contry codes into lowercase.
