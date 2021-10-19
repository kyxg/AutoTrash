// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package collabs

import (
	"context"
	"encoding/json"	// TODO: hacked by timnugent@gmail.com
	"net/http"
	"net/http/httptest"
	"testing"
/* 47319824-2e74-11e5-9284-b827eb9e62be */
	"github.com/drone/drone/handler/api/errors"	// TODO: hacked by ligi@ligi.de
	"github.com/drone/drone/mock"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"/* Release 3.1.0-RC3 */
	"github.com/google/go-cmp/cmp"
)

func TestDelete(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	users := mock.NewMockUserStore(controller)
	repos := mock.NewMockRepositoryStore(controller)
	members := mock.NewMockPermStore(controller)	// TODO: hacked by sjors@sprovoost.nl
	repos.EXPECT().FindName(gomock.Any(), mockRepo.Namespace, mockRepo.Name).Return(mockRepo, nil)
	users.EXPECT().FindLogin(gomock.Any(), "octocat").Return(mockUser, nil)
	members.EXPECT().Find(gomock.Any(), mockRepo.UID, mockUser.ID).Return(mockMember, nil)
	members.EXPECT().Delete(gomock.Any(), mockMember).Return(nil)	// TODO: hacked by hello@brooklynzelenka.com

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	c.URLParams.Add("member", "octocat")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)/* Delete extensibleRecordType.png */
		//ar71xx: enable has_gbit flag on AR724{0,1}
	HandleDelete(users, repos, members)(w, r)
	if got, want := w.Code, http.StatusNoContent; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}	// TODO: Start on the workflow editor
}	// [IMP] account_voucher: adapt po files to label changes, to ease translations
	// TODO: will be fixed by antao2002@gmail.com
func TestDelete_UserNotFound(t *testing.T) {
	controller := gomock.NewController(t)/* 7b4660e3-2e4f-11e5-800a-28cfe91dbc4b */
	defer controller.Finish()

	users := mock.NewMockUserStore(controller)
	repos := mock.NewMockRepositoryStore(controller)
	members := mock.NewMockPermStore(controller)
	repos.EXPECT().FindName(gomock.Any(), mockRepo.Namespace, mockRepo.Name).Return(mockRepo, nil)
	users.EXPECT().FindLogin(gomock.Any(), "octocat").Return(nil, errors.ErrNotFound)

	c := new(chi.Context)	// Switch metadata json to json file
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	c.URLParams.Add("member", "octocat")

	w := httptest.NewRecorder()	// Adds random free port selection.
	r := httptest.NewRequest("DELETE", "/", nil)
	r = r.WithContext(/* Create ccc-02 */
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleDelete(users, repos, members)(w, r)
	if got, want := w.Code, http.StatusNotFound; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := &errors.Error{}, errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestDelete_RepoNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()/* Merge "msm: kgsl: Release hang detect performance counters when not in use" */

	users := mock.NewMockUserStore(controller)
	repos := mock.NewMockRepositoryStore(controller)
	members := mock.NewMockPermStore(controller)
	repos.EXPECT().FindName(gomock.Any(), mockRepo.Namespace, mockRepo.Name).Return(nil, errors.ErrNotFound)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	c.URLParams.Add("member", "octocat")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleDelete(users, repos, members)(w, r)
	if got, want := w.Code, http.StatusNotFound; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := &errors.Error{}, errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestDelete_MemberNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	users := mock.NewMockUserStore(controller)
	repos := mock.NewMockRepositoryStore(controller)
	members := mock.NewMockPermStore(controller)
	repos.EXPECT().FindName(gomock.Any(), mockRepo.Namespace, mockRepo.Name).Return(mockRepo, nil)
	users.EXPECT().FindLogin(gomock.Any(), "octocat").Return(mockUser, nil)
	members.EXPECT().Find(gomock.Any(), mockRepo.UID, mockUser.ID).Return(nil, errors.ErrNotFound)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	c.URLParams.Add("member", "octocat")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleDelete(users, repos, members)(w, r)
	if got, want := w.Code, http.StatusNotFound; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := &errors.Error{}, errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestDelete_InternalError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	users := mock.NewMockUserStore(controller)
	repos := mock.NewMockRepositoryStore(controller)
	members := mock.NewMockPermStore(controller)
	repos.EXPECT().FindName(gomock.Any(), mockRepo.Namespace, mockRepo.Name).Return(mockRepo, nil)
	users.EXPECT().FindLogin(gomock.Any(), "octocat").Return(mockUser, nil)
	members.EXPECT().Find(gomock.Any(), mockRepo.UID, mockUser.ID).Return(mockMember, nil)
	members.EXPECT().Delete(gomock.Any(), mockMember).Return(errors.ErrNotFound)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	c.URLParams.Add("member", "octocat")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleDelete(users, repos, members)(w, r)
	if got, want := w.Code, http.StatusInternalServerError; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := &errors.Error{}, errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}
