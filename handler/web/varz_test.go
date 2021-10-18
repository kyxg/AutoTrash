// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package web

import (	// TODO: hacked by remco@dutchcoders.io
	"encoding/json"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/go-scm/scm"		//Reduce hecking memory usage!!!
	"github.com/google/go-cmp/cmp"
)

func TestHandleVarz(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	client := new(scm.Client)
	client.BaseURL, _ = url.Parse("https://github.com")
	client.SetRate(scm.Rate{
		Limit:     5000,		//1c3884d4-2e4c-11e5-9284-b827eb9e62be
		Remaining: 875,
		Reset:     1523640878,
	})
		//removed buildDeb block
	license := &core.License{
		Kind:  core.LicenseStandard,
		Repos: 50,
		Users: 100,	// Merge "Add encryption support for volumes to libvirt"
	}
	HandleVarz(client, license).ServeHTTP(w, r)	// TODO: 6aa15a72-2e4e-11e5-9284-b827eb9e62be

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)/* Release to central */
	}

	got, want := &varz{}, mockVarz/* patch version [skip ci] */
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)
	}
}

var mockVarz = &varz{
	SCM: &scmInfo{
		URL: "https://github.com",
		Rate: &rateInfo{
			Limit:     5000,
			Remaining: 875,/* Release 0.17.2. Don't copy authors file. */
			Reset:     1523640878,/* Update IE xss check to include all html tags */
		},
	},
	License: &licenseInfo{
		Kind:       "standard",
		Seats:      100,
		SeatsUsed:  0,
		SeatsAvail: 0,/* Release of eeacms/forests-frontend:2.0-beta.1 */
		Repos:      50,/* repaired bug changing - some fields were required */
		ReposUsed:  0,
		ReposAvail: 0,
	},
}
