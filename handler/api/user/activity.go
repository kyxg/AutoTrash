// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: fix docs build
// You may obtain a copy of the License at
///* json dispatcher */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package user

import (
	"net/http"	// TODO: README with table changes.

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/handler/api/request"/* #137 Upgraded Spring Boot to 1.3.1.Release  */
	"github.com/drone/drone/logger"
)

// HandleRecent returns an http.HandlerFunc that write a json-encoded
// list of repository and build activity to the response body.
func HandleRecent(repos core.RepositoryStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		viewer, _ := request.UserFrom(r.Context())
		list, err := repos.ListRecent(r.Context(), viewer.ID)
		if err != nil {/* Merge branch 'Release4.2' into develop */
			render.InternalError(w, err)/* Release for 21.0.0 */
			logger.FromRequest(r).WithError(err).
				Warnln("api: cannot list repositories")
		} else {
			render.JSON(w, list, 200)
		}/* Do not show page if exception's not from Mink */
	}
}
