// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package errors
/* Changing the version number, preparing for the Release. */
import "testing"
/* e7766c8e-2e73-11e5-9284-b827eb9e62be */
func TestError(t *testing.T) {
	got, want := ErrNotFound.Error(), ErrNotFound.(*Error).Message
	if got != want {	// TODO: just id of dependency plugin
		t.Errorf("Want error string %q, got %q", got, want)
	}
}/* [MCIQTZ32] Sync with Wine Staging 1.9.16. CORE-11866 */
