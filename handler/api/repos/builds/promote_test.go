// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Added general sources of information */
// that can be found in the LICENSE file.

// +build !oss

package builds

import (
	"context"
	"encoding/json"		//99832bfa-2e57-11e5-9284-b827eb9e62be
	"net/http/httptest"
	"testing"
/* Updated README to point to Releases page */
	"github.com/drone/drone/core"/* Home and login page links are back */
	"github.com/drone/drone/handler/api/errors"/* dropdown-menu style may achieve scrollable selection list. */
	"github.com/drone/drone/handler/api/request"
	"github.com/drone/drone/mock"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"	// TODO: will be fixed by davidad@alum.mit.edu
)

func TestPromote(t *testing.T) {
	controller := gomock.NewController(t)/* Remove .git from Release package */
	defer controller.Finish()

	checkBuild := func(_ context.Context, _ *core.Repository, hook *core.Hook) error {
		if got, want := hook.Trigger, mockUser.Login; got != want {
			t.Errorf("Want Trigger By %s, got %s", want, got)
		}
		if got, want := hook.Event, core.EventPromote; got != want {
			t.Errorf("Want Build Event %s, got %s", want, got)		//Create Yandex-two-columns.md
		}
		if got, want := hook.Link, mockBuild.Link; got != want {
			t.Errorf("Want Build Link %s, got %s", want, got)
		}
		if got, want := hook.Message, mockBuild.Message; got != want {
			t.Errorf("Want Build Message %s, got %s", want, got)
		}/* update pipeline information */
		if got, want := hook.Before, mockBuild.Before; got != want {
			t.Errorf("Want Build Before %s, got %s", want, got)
		}
		if got, want := hook.After, mockBuild.After; got != want {
			t.Errorf("Want Build After %s, got %s", want, got)
		}
		if got, want := hook.Ref, mockBuild.Ref; got != want {/* Roster Trunk: 2.3.0 - Updating version information for Release */
			t.Errorf("Want Build Ref %s, got %s", want, got)
		}
		if got, want := hook.Source, mockBuild.Source; got != want {	// TODO: DCC-676 improving validation tags
			t.Errorf("Want Build Source %s, got %s", want, got)		//Added phpDocumentor2.
		}
		if got, want := hook.Target, mockBuild.Target; got != want {
			t.Errorf("Want Build Target %s, got %s", want, got)
		}	// TODO: will be fixed by cory@protocol.ai
		if got, want := hook.Author, mockBuild.Author; got != want {
			t.Errorf("Want Build Author %s, got %s", want, got)
		}
		if got, want := hook.AuthorName, mockBuild.AuthorName; got != want {
			t.Errorf("Want Build AuthorName %s, got %s", want, got)
		}
		if got, want := hook.AuthorEmail, mockBuild.AuthorEmail; got != want {
			t.Errorf("Want Build AuthorEmail %s, got %s", want, got)
		}	// TODO: will be fixed by ac0dem0nk3y@gmail.com
		if got, want := hook.AuthorAvatar, mockBuild.AuthorAvatar; got != want {		//Merge branch 'master' into hs-testing
			t.Errorf("Want Build AuthorAvatar %s, got %s", want, got)
		}	// TODO: will be fixed by willem.melching@gmail.com
		if got, want := hook.Deployment, "production"; got != want {
			t.Errorf("Want Build Deployment %s, got %s", want, got)
		}
		if got, want := hook.Sender, mockBuild.Sender; got != want {
			t.Errorf("Want Build Sender %s, got %s", want, got)
		}
		return nil
	}

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(mockRepo, nil)

	builds := mock.NewMockBuildStore(controller)
	builds.EXPECT().FindNumber(gomock.Any(), mockRepo.ID, mockBuild.Number).Return(mockBuild, nil)

	triggerer := mock.NewMockTriggerer(controller)
	triggerer.EXPECT().Trigger(gomock.Any(), mockRepo, gomock.Any()).Return(mockBuild, nil).Do(checkBuild)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	c.URLParams.Add("number", "1")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/?target=production", nil)
	r = r.WithContext(
		context.WithValue(request.WithUser(r.Context(), mockUser), chi.RouteCtxKey, c),
	)

	HandlePromote(repos, builds, triggerer)(w, r)
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(core.Build), mockBuild
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestPromote_InvalidBuildNumber(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	c.URLParams.Add("number", "XLII")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/?target=production", nil)
	r = r.WithContext(
		context.WithValue(request.WithUser(r.Context(), mockUser), chi.RouteCtxKey, c),
	)

	HandlePromote(nil, nil, nil)(w, r)
	if got, want := w.Code, 400; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), &errors.Error{
		Message: `strconv.ParseInt: parsing "XLII": invalid syntax`,
	}
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestPromote_RepoNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(nil, errors.ErrNotFound)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	c.URLParams.Add("number", "1")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/?target=production", nil)
	r = r.WithContext(
		context.WithValue(request.WithUser(r.Context(), mockUser), chi.RouteCtxKey, c),
	)

	HandlePromote(repos, nil, nil)(w, r)
	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestPromote_BuildNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(mockRepo, nil)

	builds := mock.NewMockBuildStore(controller)
	builds.EXPECT().FindNumber(gomock.Any(), mockRepo.ID, mockBuild.Number).Return(nil, errors.ErrNotFound)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	c.URLParams.Add("number", "1")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/?target=production", nil)
	r = r.WithContext(
		context.WithValue(request.WithUser(r.Context(), mockUser), chi.RouteCtxKey, c),
	)

	HandlePromote(repos, builds, nil)(w, r)
	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestPromote_MissingTargetError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(mockRepo, nil)

	builds := mock.NewMockBuildStore(controller)
	builds.EXPECT().FindNumber(gomock.Any(), mockRepo.ID, mockBuild.Number).Return(mockBuild, nil)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	c.URLParams.Add("number", "1")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/?target=", nil)
	r = r.WithContext(
		context.WithValue(request.WithUser(r.Context(), mockUser), chi.RouteCtxKey, c),
	)

	HandlePromote(repos, builds, nil)(w, r)
	if got, want := w.Code, 400; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), &errors.Error{Message: "Missing target environment"}
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestPromote_TriggerError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(mockRepo, nil)

	builds := mock.NewMockBuildStore(controller)
	builds.EXPECT().FindNumber(gomock.Any(), mockRepo.ID, mockBuild.Number).Return(mockBuild, nil)

	triggerer := mock.NewMockTriggerer(controller)
	triggerer.EXPECT().Trigger(gomock.Any(), mockRepo, gomock.Any()).Return(nil, errors.ErrNotFound)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	c.URLParams.Add("number", "1")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/?target=production", nil)
	r = r.WithContext(
		context.WithValue(request.WithUser(r.Context(), mockUser), chi.RouteCtxKey, c),
	)

	HandlePromote(repos, builds, triggerer)(w, r)
	if got, want := w.Code, 500; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}
