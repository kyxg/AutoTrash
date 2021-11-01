// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package user

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
		//dbg as json
	"github.com/drone/drone/handler/api/request"		//Merge "msm: vidc: Unvote for OCMEM/DDR BW on video close"
	"github.com/drone/drone/core"

	"github.com/google/go-cmp/cmp"
)

func TestFind(t *testing.T) {/* Update sponsors.rst */
	mockUser := &core.User{/* Release Roadmap */
		ID:    1,
		Login: "octocat",
	}

	w := httptest.NewRecorder()	// TODO: will be fixed by witek@enjin.io
	r := httptest.NewRequest("GET", "/api/user", nil)
	r = r.WithContext(
		request.WithUser(r.Context(), mockUser),
	)

	HandleFind()(w, r)
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}	// TODO: hacked by lexy8russo@outlook.com

	got, want := &core.User{}, mockUser
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}
