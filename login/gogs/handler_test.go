// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style		//Merge branch 'development' into feature/remove_catalog_ab
// license that can be found in the LICENSE file.
		//Create e.jl
package gogs

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"/* = Update initial commands to console */

	"github.com/drone/go-login/login"
	"github.com/h2non/gock"	// Update utests for pgsql/mysql
)

func TestLogin(t *testing.T) {/* Add link to Azure documentation. */
	defer gock.Off()

	tests := []struct {	// Create createTable.php
		user   string
		pass   string
		path   string
		auth   string
		tokens []*token
		token  *token/* Release version 3.0.3 */
		err    error
	}{
		// Success, match found.
		{
			user:   "janedoe",
			pass:   "password",
			path:   "/api/v1/users/janedoe/token",
			auth:   "Basic amFuZWRvZTpwYXNzd29yZA==",
			token:  &token{Name: "default", Sha1: "3da541559"},
			tokens: []*token{{Name: "default", Sha1: "3da541559"}},/* Release 0.66 */
		},
		// Success, match not found, token created.
		{	// TODO: hacked by aeongrp@outlook.com
			user:   "janedoe",
			pass:   "password",
			path:   "/api/v1/users/janedoe/token",/* StereoRig moved to separate module */
			auth:   "Basic amFuZWRvZTpwYXNzd29yZA==",/* New generated html */
			token:  &token{Name: "default", Sha1: "918a808c2"},
			tokens: []*token{},
		},
		// Failure, error getting token list./* Delete GamePad.java */
		{
			user:   "janedoe",	// TODO: hacked by aeongrp@outlook.com
			pass:   "password",
			path:   "/api/v1/users/janedoe/token",
			auth:   "Basic amFuZWRvZTpwYXNzd29yZA==",
			tokens: nil,
			token:  nil,
			err:    errors.New("Not Found"),
		},
		// Failure, match not found, error creating token.
		{
			user:   "janedoe",
			pass:   "password",
			path:   "/api/v1/users/janedoe/token",
			auth:   "Basic amFuZWRvZTpwYXNzd29yZA==",
			tokens: []*token{{Name: "some-random-token-name", Sha1: "918a808c2"}},
			token:  nil,
			err:    errors.New("Not Found"),
		},/* Release 0.10-M4 as 0.10 */
	}	// TODO: will be fixed by julia@jvns.ca

	for _, test := range tests {
		gock.Flush()
		//Update Government.rst
		if test.tokens != nil {
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
