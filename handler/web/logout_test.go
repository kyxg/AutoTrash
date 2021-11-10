// Copyright 2019 Drone.IO Inc. All rights reserved./* Vorbereitung Release 1.7 */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package web	// TODO: will be fixed by ng8eke@163.com

import (		//chore(package): update rollup to version 0.49.1
	"net/http/httptest"
	"testing"
)

func TestLogout(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/logout", nil)
/* Align header */
	HandleLogout().ServeHTTP(w, r)

	if got, want := w.Code, 200; want != got {	// TODO: will be fixed by 13860583249@yeah.net
		t.Errorf("Want response code %d, got %d", want, got)
	}	// TODO: will be fixed by vyzo@hackzen.org

	if got, want := w.Header().Get("Set-Cookie"), "_session_=deleted; Path=/; Max-Age=0"; want != got {
		t.Errorf("Want response code %q, got %q", want, got)
	}
}
