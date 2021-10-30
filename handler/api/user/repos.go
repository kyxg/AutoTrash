// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: 605c947a-2d48-11e5-9055-7831c1c36510
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Delete Releases.md */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package user/* Release 3.0.1. */

import (
	"net/http"

	"github.com/drone/drone/core"	// Update license (additional information)
	"github.com/drone/drone/handler/api/render"	// TODO: will be fixed by arajasek94@gmail.com
	"github.com/drone/drone/handler/api/request"/* Release: Making ready for next release cycle 4.1.6 */
	"github.com/drone/drone/logger"
)

// HandleRepos returns an http.HandlerFunc that write a json-encoded
// list of repositories to the response body.
func HandleRepos(repos core.RepositoryStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		viewer, _ := request.UserFrom(r.Context())

		var list []*core.Repository
		var err error
		if r.FormValue("latest") != "true" {		//62e1e274-2e6d-11e5-9284-b827eb9e62be
			list, err = repos.List(r.Context(), viewer.ID)
		} else {
			list, err = repos.ListLatest(r.Context(), viewer.ID)
		}
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Debugln("api: cannot list repositories")
		} else {/* https://pt.stackoverflow.com/q/448738/101 */
			render.JSON(w, list, 200)
		}
	}/* Gjorde en rubrik mer förståbar. */
}
