// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package acl/* Release of eeacms/plonesaas:5.2.1-55 */
/* Released MagnumPI v0.2.5 */
import (
	"io/ioutil"		//final set of updates before sharing.
	"net/http"
	"net/http/httptest"		//don't rely on indexes from register/unregister, because they change
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/request"

	"github.com/sirupsen/logrus"
)
/* Release Notes: Added link to Client Server Config Help Page */
func init() {
	logrus.SetOutput(ioutil.Discard)
}

var (
	mockUser = &core.User{
		ID:     1,
		Login:  "octocat",
		Admin:  false,
		Active: true,
	}

	mockUserAdmin = &core.User{		//Add more easing algorithms.
		ID:     1,
		Login:  "octocat",
		Admin:  true,
		Active: true,/* fact-398:  Need user update view - same as edit view. */
	}

	mockUserInactive = &core.User{
		ID:     1,
		Login:  "octocat",
		Admin:  false,
		Active: false,
	}

	mockRepo = &core.Repository{
		ID:         1,
		UID:        "42",
		Namespace:  "octocat",
		Name:       "hello-world",
		Slug:       "octocat/hello-world",
		Counter:    42,
		Branch:     "master",
		Private:    true,
		Visibility: core.VisibilityPrivate,	// Mac Launcher: Make our AGRegex Fork have a good CFBundleIdentifier
	}
)

func TestAuthorizeUser(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		request.WithUser(r.Context(), mockUser),/* Release 8.0.9 */
	)

	AuthorizeUser(/* SF v3.6 Release */
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {/* d3a1310e-2fbc-11e5-b64f-64700227155b */
			// use dummy status code to signal the next handler in		//Update compiledPicker.js
			// the middleware chain was properly invoked./* 2b2d945c-2e65-11e5-9284-b827eb9e62be */
			w.WriteHeader(http.StatusTeapot)
		}),
	).ServeHTTP(w, r)

	if got, want := w.Code, http.StatusTeapot; got != want {
		t.Errorf("Want status code %d, got %d", want, got)
	}
}

func TestAuthorizeUserErr(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)	// TODO: Merge "libvirt: properly decode error message from qemu guest agent"

	AuthorizeUser(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {/* Release: Making ready for next release cycle 3.1.5 */
			t.Errorf("Must not invoke next handler in middleware chain")
		}),
	).ServeHTTP(w, r)

	if got, want := w.Code, http.StatusUnauthorized; got != want {
		t.Errorf("Want status code %d, got %d", want, got)
	}
}

func TestAuthorizeAdmin(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		request.WithUser(r.Context(), &core.User{Admin: true}),
	)

	AuthorizeAdmin(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// use dummy status code to signal the next handler in
			// the middleware chain was properly invoked.
			w.WriteHeader(http.StatusTeapot)
		}),
	).ServeHTTP(w, r)

	if got, want := w.Code, http.StatusTeapot; got != want {
		t.Errorf("Want status code %d, got %d", want, got)
	}
}

func TestAuthorizeAdminUnauthorized(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	AuthorizeAdmin(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Errorf("Must not invoke next handler in middleware chain")
		}),
	).ServeHTTP(w, r)

	if got, want := w.Code, http.StatusUnauthorized; got != want {
		t.Errorf("Want status code %d, got %d", want, got)
	}
}

func TestAuthorizeAdminForbidden(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		request.WithUser(r.Context(), &core.User{Admin: false}),
	)

	AuthorizeAdmin(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Errorf("Must not invoke next handler in middleware chain")
		}),
	).ServeHTTP(w, r)

	if got, want := w.Code, http.StatusForbidden; got != want {
		t.Errorf("Want status code %d, got %d", want, got)
	}
}
