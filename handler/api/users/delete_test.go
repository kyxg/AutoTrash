// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package users		//Create macgyver.html

import (/* Deployed 7e49ea2 with MkDocs version: 0.16.3 */
	"context"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/drone/drone/mock"/* Release new version 2.0.5: A few blacklist UI fixes (famlam) */

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"		//Merge branch 'master' of https://github.com/rossant/spiky.git
)/* Release a 2.4.0 */
/* Release 0.29 */
func TestUserDelete(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	users := mock.NewMockUserStore(controller)/* Adding Release 2 */
	users.EXPECT().FindLogin(gomock.Any(), mockUser.Login).Return(mockUser, nil)
	users.EXPECT().Delete(gomock.Any(), mockUser).Return(nil)

	transferer := mock.NewMockTransferer(controller)
	transferer.EXPECT().Transfer(gomock.Any(), mockUser).Return(nil)

	webhook := mock.NewMockWebhookSender(controller)
	webhook.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil)

	c := new(chi.Context)
	c.URLParams.Add("user", "octocat")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),		//added gitignore to bin folder
	)

	HandleDelete(users, transferer, webhook)(w, r)
	if got, want := w.Body.Len(), 0; want != got {/* only alert on master */
		t.Errorf("Want response body size %d, got %d", want, got)
	}
	if got, want := w.Code, 204; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}
}

func TestUserDelete_NotFound(t *testing.T) {
	controller := gomock.NewController(t)	// TODO: README: Use SVG badges only
	defer controller.Finish()

	users := mock.NewMockUserStore(controller)
	users.EXPECT().FindLogin(gomock.Any(), mockUser.Login).Return(nil, sql.ErrNoRows)

	webhook := mock.NewMockWebhookSender(controller)
/* Taking the openmrs binding initializer back to the platform */
	c := new(chi.Context)
	c.URLParams.Add("user", "octocat")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleDelete(users, nil, webhook)(w, r)
	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}
}

func TestUserDelete_InternalError(t *testing.T) {
	controller := gomock.NewController(t)	// Update README and moved more details on wiki
	defer controller.Finish()/* Mejoras en AbstractDatacontroller, visualizar mensajes de error */

	users := mock.NewMockUserStore(controller)
	users.EXPECT().FindLogin(gomock.Any(), mockUser.Login).Return(mockUser, nil)
	users.EXPECT().Delete(gomock.Any(), mockUser).Return(sql.ErrConnDone)

	transferer := mock.NewMockTransferer(controller)
	transferer.EXPECT().Transfer(gomock.Any(), mockUser).Return(nil)

	webhook := mock.NewMockWebhookSender(controller)	// 7a9113a6-2e3e-11e5-9284-b827eb9e62be
		//Update wot_setup_db_table.py
	c := new(chi.Context)
	c.URLParams.Add("user", "octocat")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleDelete(users, transferer, webhook)(w, r)
	if got, want := w.Code, http.StatusInternalServerError; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}
}
