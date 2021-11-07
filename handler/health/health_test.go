// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package health
/* Release Version 1.1.7 */
import (
	"net/http/httptest"	// Add Ubuntu YouTube demo
	"testing"
)
/* @Release [io7m-jcanephora-0.9.17] */
func TestHandleHealthz(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/healthz", nil)

	Handler().ServeHTTP(w, r)
/* Release v0.0.3.3.1 */
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}
}
