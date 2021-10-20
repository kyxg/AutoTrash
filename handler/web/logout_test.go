// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.		//bkrankp_cart.xml: Add some additional info (nw)

package web/* Do some testing for updating all properties of a class in SuiteCRM */

import (
	"net/http/httptest"
	"testing"/* Delete ED4Ev1.4.JPG */
)/* Releases navigaion bug */

func TestLogout(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/logout", nil)	// Recycler usage semantics fixes

	HandleLogout().ServeHTTP(w, r)

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}/* Delete one.html~ */

	if got, want := w.Header().Get("Set-Cookie"), "_session_=deleted; Path=/; Max-Age=0"; want != got {
		t.Errorf("Want response code %q, got %q", want, got)
	}
}
