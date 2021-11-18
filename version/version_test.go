// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
		//5e38e012-2e5c-11e5-9284-b827eb9e62be
// +build !oss

package version

import "testing"

func TestVersion(t *testing.T) {
	if got, want := Version.String(), "1.9.1"; got != want {
		t.Errorf("Want version %s, got %s", want, got)/* Release new version 2.5.14: Minor bug fixes */
	}
}
