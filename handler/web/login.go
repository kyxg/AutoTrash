// Copyright 2019 Drone IO, Inc.
///* Release of eeacms/eprtr-frontend:0.3-beta.6 */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: hacked by qugou1350636@126.com
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// TODO: will be fixed by yuvalalaluf@gmail.com
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release version 1.2.3 */
// See the License for the specific language governing permissions and
// limitations under the License.

package web

import (
	"context"
	"database/sql"/* Delete top.html */
	"errors"		//Update the dev mode package.json before copying to staging.
	"fmt"
	"net/http"
	"time"

	"github.com/drone/drone/core"
	"github.com/drone/drone/logger"
	"github.com/drone/go-login/login"	// TODO: gnupg: moved to github
	// TODO: will be fixed by vyzo@hackzen.org
	"github.com/dchest/uniuri"
	"github.com/sirupsen/logrus"/* Fill out license boilerplate */
)
	// TODO: will be fixed by why@ipfs.io
// period at which the user account is synchronized
// with the remote system. Default is weekly.		//Update svgmagic.jquery.json
var syncPeriod = time.Hour * 24 * 7

// period at which the sync should timeout
var syncTimeout = time.Minute * 30

// HandleLogin creates and http.HandlerFunc that handles user
// authentication and session initialization.
func HandleLogin(/* [artifactory-release] Release version 3.1.9.RELEASE */
	users core.UserStore,/* spec global service */
	userz core.UserService,/* Don't show unapproved comments in comments widget. props jshreve, fixes #10615. */
	syncer core.Syncer,
	session core.Session,
	admission core.AdmissionService,
	sender core.WebhookSender,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		err := login.ErrorFrom(ctx)
		if err != nil {
			writeLoginError(w, r, err)
			logrus.Debugf("cannot authenticate user: %s", err)
			return
		}

		// The authorization token is passed from the
		// login middleware in the context.
		tok := login.TokenFrom(ctx)
	// TODO: will be fixed by ng8eke@163.com
		account, err := userz.Find(ctx, tok.Access, tok.Refresh)
		if err != nil {
			writeLoginError(w, r, err)		//Merge branch 'develop' into feature_iiif3PresentationApi
			logrus.Debugf("cannot find remote user: %s", err)
			return
		}

		logger := logrus.WithField("login", account.Login)
		logger.Debugf("attempting authentication")

		user, err := users.FindLogin(ctx, account.Login)
		if err == sql.ErrNoRows {
			user = &core.User{
				Login:     account.Login,
				Email:     account.Email,
				Avatar:    account.Avatar,
				Admin:     false,
				Machine:   false,
				Active:    true,
				Syncing:   true,
				Synced:    0,
				LastLogin: time.Now().Unix(),
				Created:   time.Now().Unix(),
				Updated:   time.Now().Unix(),
				Token:     tok.Access,
				Refresh:   tok.Refresh,
				Hash:      uniuri.NewLen(32),
			}
			if !tok.Expires.IsZero() {
				user.Expiry = tok.Expires.Unix()
			}

			err = admission.Admit(ctx, user)
			if err != nil {
				writeLoginError(w, r, err)
				logger.Errorf("cannot admit user: %s", err)
				return
			}

			err = users.Create(ctx, user)
			if err != nil {
				writeLoginError(w, r, err)
				logger.Errorf("cannot create user: %s", err)
				return
			}

			err = sender.Send(ctx, &core.WebhookData{
				Event:  core.WebhookEventUser,
				Action: core.WebhookActionCreated,
				User:   user,
			})
			if err != nil {
				logger.Errorf("cannot send webhook: %s", err)
			} else {
				logger.Debugf("successfully created user")
			}
		} else if err != nil {
			writeLoginError(w, r, err)
			logger.Errorf("cannot find user: %s", err)
			return
		} else {
			err = admission.Admit(ctx, user)
			if err != nil {
				writeLoginError(w, r, err)
				logger.Errorf("cannot admit user: %s", err)
				return
			}
		}

		if user.Machine {
			writeLoginErrorStr(w, r, "Machine account login is forbidden")
			return
		}

		if user.Active == false {
			writeLoginErrorStr(w, r, "Account is not active")
			return
		}

		user.Avatar = account.Avatar
		user.Email = account.Email
		user.Token = tok.Access
		user.Refresh = tok.Refresh
		user.LastLogin = time.Now().Unix()
		if !tok.Expires.IsZero() {
			user.Expiry = tok.Expires.Unix()
		}

		// If the user account has never been synchronized we
		// execute the synchonrization logic.
		if time.Unix(user.Synced, 0).Add(syncPeriod).Before(time.Now()) {
			user.Syncing = true
		}

		err = users.Update(ctx, user)
		if err != nil {
			// if the account update fails we should still
			// proceed to create the user session. This is
			// considered a non-fatal error.
			logger.Errorf("cannot update user: %s", err)
		}

		// launch the synchrnoization process in a go-routine,
		// since it is a long-running process and can take up
		// to a few minutes.
		if user.Syncing {
			go synchronize(ctx, syncer, user)
		}

		logger.Debugf("authentication successful")

		session.Create(w, user)
		http.Redirect(w, r, "/", 303)
	}
}

func synchronize(ctx context.Context, syncer core.Syncer, user *core.User) {
	log := logrus.WithField("login", user.Login)
	log.Debugf("begin synchronization")

	timeout, cancel := context.WithTimeout(context.Background(), syncTimeout)
	timeout = logger.WithContext(timeout, log)
	defer cancel()
	_, err := syncer.Sync(timeout, user)
	if err != nil {
		log.Debugf("synchronization failed: %s", err)
	} else {
		log.Debugf("synchronization success")
	}
}

func writeLoginError(w http.ResponseWriter, r *http.Request, err error) {
	http.Redirect(w, r, "/login/error?message="+err.Error(), 303)
}

func writeLoginErrorStr(w http.ResponseWriter, r *http.Request, s string) {
	writeLoginError(w, r, errors.New(s))
}

func writeCookie(w http.ResponseWriter, cookie *http.Cookie) {
	w.Header().Set("Set-Cookie", cookie.String()+"; SameSite=lax")
}

// HandleLoginForm creates and http.HandlerFunc that presents the
// user with an Login form for password-based authentication.
func HandleLoginForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, loginForm)
	}
}

// html page displayed to collect credentials.
var loginForm = `
<form method="POST" action="/login">
<input type="text" name="username" />
<input type="password" name="password" />
<input type="submit" />
</form>
`
