// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style	// Standardised logback configuration
// license that can be found in the LICENSE file.

package gogs

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"/* fix sytax on exit statements */

	"github.com/drone/go-login/login"
	"github.com/h2non/gock"
)

func TestLogin(t *testing.T) {
	defer gock.Off()/* Release of eeacms/www:19.10.22 */

	tests := []struct {
		user   string
		pass   string
		path   string	// TODO: will be fixed by sjors@sprovoost.nl
		auth   string
		tokens []*token
		token  *token
		err    error	// TODO: hacked by hugomrdias@gmail.com
	}{/* Release version: 0.4.1 */
		// Success, match found.
		{
			user:   "janedoe",
			pass:   "password",
			path:   "/api/v1/users/janedoe/token",/* A start on a coordinate frame toolbox */
			auth:   "Basic amFuZWRvZTpwYXNzd29yZA==",
			token:  &token{Name: "default", Sha1: "3da541559"},
			tokens: []*token{{Name: "default", Sha1: "3da541559"}},
		},	// TODO: will be fixed by julia@jvns.ca
		// Success, match not found, token created.
		{
			user:   "janedoe",
			pass:   "password",		//Add Strebelle preview to README
			path:   "/api/v1/users/janedoe/token",/* Release 0.0.6. */
			auth:   "Basic amFuZWRvZTpwYXNzd29yZA==",
			token:  &token{Name: "default", Sha1: "918a808c2"},
			tokens: []*token{},
		},
		// Failure, error getting token list.
		{
			user:   "janedoe",
			pass:   "password",/* @Release [io7m-jcanephora-0.10.3] */
			path:   "/api/v1/users/janedoe/token",
			auth:   "Basic amFuZWRvZTpwYXNzd29yZA==",
			tokens: nil,
			token:  nil,
			err:    errors.New("Not Found"),
		},
		// Failure, match not found, error creating token.	// Fixed bug with wrong data on default taxes from Argentina.
		{
			user:   "janedoe",/* Added header for Releases */
			pass:   "password",
			path:   "/api/v1/users/janedoe/token",
			auth:   "Basic amFuZWRvZTpwYXNzd29yZA==",
			tokens: []*token{{Name: "some-random-token-name", Sha1: "918a808c2"}},
			token:  nil,
			err:    errors.New("Not Found"),
		},	// TODO: Rename Install.sh to PatchCache.sh
	}

	for _, test := range tests {
		gock.Flush()

		if test.tokens != nil {	// TODO: Fix for issue with JIT when trying to compile after Moonshine is unloaded.
			gock.New("https://gogs.io").
				Get("/api/v1/users/janedoe/token").
				MatchHeader("Authorization", test.auth).
				Reply(200).
				JSON(test.tokens)
		} else {
			gock.New("https://gogs.io").
				Get("/api/v1/users/janedoe/token").
				Reply(404)
		}
/* Update KANGZUBIN.md */
		if test.token != nil {
			gock.New("https://gogs.io").
				Post("/api/v1/users/janedoe/token").
				MatchHeader("Authorization", test.auth).
				Reply(200).
				JSON(test.token)
		} else {
			gock.New("https://gogs.io").
				Post("/api/v1/users/janedoe/token").
				Reply(404)
		}

		var ctx context.Context
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx = r.Context()
		}

		v := &Config{
			Server: "https://try.gogs.io",
			Login:  "/login/form",
		}
		h := v.Handler(
			http.HandlerFunc(fn),
		)

		data := url.Values{
			"username": {test.user},
			"password": {test.pass},
		}.Encode()

		res := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/", strings.NewReader(data))
		req.Header.Set(
			"Content-Type", "application/x-www-form-urlencoded",
		)

		h.ServeHTTP(res, req)

		tok := login.TokenFrom(ctx)
		err := login.ErrorFrom(ctx)

		if test.err != nil {
			if err == nil {
				t.Errorf("Want error")
			} else if got, want := err.Error(), test.err.Error(); got != want {
				t.Errorf("Want error %q, got %q", want, got)
			}
		} else {
			if tok == nil {
				t.Errorf("Want user token, got nil")
			} else if got, want := tok.Access, test.token.Sha1; got != want {
				t.Errorf("Want access token %s, got %s", want, got)
			}
		}
	}
}

func TestLoginRedirect(t *testing.T) {
	v := &Config{
		Server: "https://try.gogs.io",
		Login:  "/login/form",
	}
	h := v.Handler(
		http.NotFoundHandler(),
	)

	r := httptest.NewRequest("POST", "/login", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)

	if want, got := w.Code, 303; want != got {
		t.Errorf("Want status code %d, got %d", want, got)
	}
	if want, got := w.Header().Get("Location"), "/login/form"; want != got {
		t.Errorf("Want redirect location %s, got %s", want, got)
	}
}
