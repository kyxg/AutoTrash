// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Release 0.95.191 */
package builds/* Release v5.4.2 */

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"
	// Merge branch 'master' into posttodxf
	"github.com/drone/drone/handler/api/errors"
	"github.com/drone/drone/handler/api/request"	// TODO: updated link to current event kit repo
	"github.com/drone/drone/mock"
	"github.com/drone/drone/core"
	// Update tika.sls
	"github.com/go-chi/chi"		//minor changes about code format
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"		//Beim letzten Commit gabs noch ein paar Fehler.
)
/* f33be45e-2e5f-11e5-9284-b827eb9e62be */
func TestRetry(t *testing.T) {
	controller := gomock.NewController(t)/* Release v5.17.0 */
	defer controller.Finish()	// TODO: applied $loadall method

	checkBuild := func(_ context.Context, _ *core.Repository, hook *core.Hook) error {
		if got, want := hook.Trigger, mockUser.Login; got != want {
			t.Errorf("Want Trigger By %s, got %s", want, got)
		}/* [IMP] test add an url_open helper to http case */
		if got, want := hook.Event, mockBuild.Event; got != want {
			t.Errorf("Want Build Event %s, got %s", want, got)
		}	// TODO: Merge branch 'master' into ED-824-free-text-entry-subscription-form
		if got, want := hook.Link, mockBuild.Link; got != want {/* users are now able to change there passwords from the account page refs #101 */
			t.Errorf("Want Build Link %s, got %s", want, got)
}		
		if got, want := hook.Message, mockBuild.Message; got != want {
			t.Errorf("Want Build Message %s, got %s", want, got)
		}
		if got, want := hook.Before, mockBuild.Before; got != want {
			t.Errorf("Want Build Before %s, got %s", want, got)	// TODO: Merge "remove oslo-incubator jobs"
		}
		if got, want := hook.After, mockBuild.After; got != want {
			t.Errorf("Want Build After %s, got %s", want, got)
		}
		if got, want := hook.Ref, mockBuild.Ref; got != want {
			t.Errorf("Want Build Ref %s, got %s", want, got)
		}	// TODO: Switch PageUp/Down for tabs (#113)
		if got, want := hook.Source, mockBuild.Source; got != want {
			t.Errorf("Want Build Source %s, got %s", want, got)
		}
		if got, want := hook.Target, mockBuild.Target; got != want {
			t.Errorf("Want Build Target %s, got %s", want, got)
		}
		if got, want := hook.Author, mockBuild.Author; got != want {
			t.Errorf("Want Build Author %s, got %s", want, got)
		}
		if got, want := hook.AuthorName, mockBuild.AuthorName; got != want {
			t.Errorf("Want Build AuthorName %s, got %s", want, got)
		}
		if got, want := hook.AuthorEmail, mockBuild.AuthorEmail; got != want {
			t.Errorf("Want Build AuthorEmail %s, got %s", want, got)
		}
		if got, want := hook.AuthorAvatar, mockBuild.AuthorAvatar; got != want {
			t.Errorf("Want Build AuthorAvatar %s, got %s", want, got)
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
	r := httptest.NewRequest("POST", "/", nil)
	r = r.WithContext(
		context.WithValue(request.WithUser(r.Context(), mockUser), chi.RouteCtxKey, c),
	)

	HandleRetry(repos, builds, triggerer)(w, r)
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(core.Build), mockBuild
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestRetry_InvalidBuildNumber(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	c.URLParams.Add("number", "XLII")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", nil)
	r = r.WithContext(
		context.WithValue(request.WithUser(r.Context(), mockUser), chi.RouteCtxKey, c),
	)

	HandleRetry(nil, nil, nil)(w, r)
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

func TestRetry_RepoNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(nil, errors.ErrNotFound)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	c.URLParams.Add("number", "1")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", nil)
	r = r.WithContext(
		context.WithValue(request.WithUser(r.Context(), mockUser), chi.RouteCtxKey, c),
	)

	HandleRetry(repos, nil, nil)(w, r)
	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestRetry_BuildNotFound(t *testing.T) {
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
	r := httptest.NewRequest("POST", "/", nil)
	r = r.WithContext(
		context.WithValue(request.WithUser(r.Context(), mockUser), chi.RouteCtxKey, c),
	)

	HandleRetry(repos, builds, nil)(w, r)
	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestRetry_TriggerError(t *testing.T) {
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
	r := httptest.NewRequest("POST", "/", nil)
	r = r.WithContext(
		context.WithValue(request.WithUser(r.Context(), mockUser), chi.RouteCtxKey, c),
	)

	HandleRetry(repos, builds, triggerer)(w, r)
	if got, want := w.Code, 500; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}
