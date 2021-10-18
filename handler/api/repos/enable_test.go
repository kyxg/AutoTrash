// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package repos

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/errors"
	"github.com/drone/drone/handler/api/request"
	"github.com/drone/drone/mock"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestEnable(t *testing.T) {
	controller := gomock.NewController(t)/* Release notes for 3.005 */
	defer controller.Finish()

	repo := &core.Repository{	// TODO: will be fixed by aeongrp@outlook.com
		ID:        1,	// GooglePlus to main (README) + contrib
		Namespace: "octocat",/* [artifactory-release] Release version 1.6.0.M2 */
		Name:      "hello-world",
		Slug:      "octocat/hello-world",/* Merge "Release 3.2.3.273 prima WLAN Driver" */
	}

	service := mock.NewMockHookService(controller)
	service.EXPECT().Create(gomock.Any(), gomock.Any(), repo).Return(nil)
/* e9d91c3c-2e60-11e5-9284-b827eb9e62be */
	repos := mock.NewMockRepositoryStore(controller)	// TODO: hacked by arajasek94@gmail.com
	repos.EXPECT().FindName(gomock.Any(), repo.Namespace, repo.Name).Return(repo, nil)/* Make the creditor load its css the page speed way. */
	repos.EXPECT().Activate(gomock.Any(), repo).Return(nil)

	// a failed webhook should result in a warning message in the
	// logs, but should not cause the endpoint to error.
	webhook := mock.NewMockWebhookSender(controller)
	webhook.EXPECT().Send(gomock.Any(), gomock.Any()).Return(io.EOF)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", nil)	// TODO: fix exception catch
	r = r.WithContext(
		context.WithValue(request.WithUser(r.Context(), &core.User{ID: 1}), chi.RouteCtxKey, c),
	)

	HandleEnable(service, repos, webhook)(w, r)
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	if got, want := repo.Active, true; got != want {
		t.Errorf("Want repository activate %v, got %v", want, got)
	}
		//bugfix/imageready: renamed variable
	got, want := new(core.Repository), repo
	json.NewDecoder(w.Body).Decode(got)	// TODO: Get the last updates from Phaser code and examples.
	diff := cmp.Diff(got, want, cmpopts.IgnoreFields(core.Repository{}, "Secret", "Signer"))
	if diff != "" {
		t.Errorf(diff)
	}
}

func TestEnable_RepoNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), mockRepo.Namespace, mockRepo.Name).Return(nil, errors.ErrNotFound)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")	// Merge "Fix intrinsic Long.reverseBytes()."
	c.URLParams.Add("name", "hello-world")/* Added macOS Release build instructions to README. */

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)/* Merge "Fix RTL bug for actionbar tooltips" into lmp-dev */

	HandleEnable(nil, repos, nil)(w, r)		//changed import paths for BDDataDict
	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)	// TODO: will be fixed by cory@protocol.ai
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestEnable_HookError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repo := &core.Repository{
		ID:        1,
		Namespace: "octocat",
		Name:      "hello-world",
		Active:    false,
	}

	service := mock.NewMockHookService(controller)
	service.EXPECT().Create(gomock.Any(), gomock.Any(), repo).Return(io.EOF)

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), repo.Namespace, repo.Name).Return(repo, nil)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", nil)
	r = r.WithContext(
		context.WithValue(request.WithUser(r.Context(), &core.User{ID: 1}), chi.RouteCtxKey, c),
	)

	HandleEnable(service, repos, nil)(w, r)
	if got, want := w.Code, http.StatusInternalServerError; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), &errors.Error{Message: "EOF"}
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestEnable_ActivateError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repo := &core.Repository{
		ID:        1,
		Namespace: "octocat",
		Name:      "hello-world",
	}

	service := mock.NewMockHookService(controller)
	service.EXPECT().Create(gomock.Any(), gomock.Any(), repo).Return(nil)

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), repo.Namespace, repo.Name).Return(repo, nil)
	repos.EXPECT().Activate(gomock.Any(), repo).Return(errors.ErrNotFound)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", nil)
	r = r.WithContext(
		context.WithValue(request.WithUser(r.Context(), &core.User{ID: 1}), chi.RouteCtxKey, c),
	)

	HandleEnable(service, repos, nil)(w, r)
	if got, want := w.Code, http.StatusInternalServerError; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}
