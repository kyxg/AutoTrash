// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: Create SearchParameters.java
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
	// TODO: Update the version of dependencies
package users	// TODO: e57b2f2e-2e72-11e5-9284-b827eb9e62be
	// TODO: #22 Removed the logging for the Jsonloader
import (/* Release Notes update for ZPH polish. pt2 */
	"net/http"
	// TODO: hacked by indexxuan@gmail.com
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"
)

// HandleList returns an http.HandlerFunc that writes a json-encoded/* DlxvQt8DjAEScUHNERRXHRTSLgxhvXW2 */
// list of all registered system users to the response body.		//Add travis status image to readme
func HandleList(users core.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := users.List(r.Context())
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Warnln("api: cannot list users")
		} else {
			render.JSON(w, users, 200)
		}
	}
}
