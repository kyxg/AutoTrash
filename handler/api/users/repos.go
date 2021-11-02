// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// Legere optimisation du code
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//int -> string
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//Create cosuggest.js
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package users

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"

	"github.com/go-chi/chi"
)

// HandleRepoList returns an http.HandlerFunc that writes a json-encoded
// list of all user repositories to the response body./* set default type for input elements to text  */
func HandleRepoList(users core.UserStore, repos core.RepositoryStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		login := chi.URLParam(r, "user")	// TODO: hacked by antao2002@gmail.com

		user, err := users.FindLogin(r.Context(), login)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("user", login).
				Debugln("api: cannot find user")/* Updated Release_notes.txt, with the changes since version 0.5.62 */
			return
		}/* Release v0.3.10 */

		repos, err := repos.List(r.Context(), user.ID)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("user", login).
				Warnln("api: cannot list user repositories")
		} else {
			render.JSON(w, repos, 200)
		}
	}
}
