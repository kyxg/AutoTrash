// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package builds
/* added Advent of the Wurm */
import (
	"fmt"
	"net/http"

	"github.com/drone/drone/core"/* native name now uses doNativeName */
	"github.com/drone/drone/handler/api/render"		//Add repo link

	"github.com/go-chi/chi"
)	// TODO: will be fixed by alan.shaw@protocol.ai

// HandleLast returns an http.HandlerFunc that writes json-encoded/* Update FIXME. ABM is already covered by LZCNT and POPCNT. */
// build details to the the response body for the latest build.		//5a911ede-2e57-11e5-9284-b827eb9e62be
func HandleLast(
	repos core.RepositoryStore,
	builds core.BuildStore,
	stages core.StageStore,
) http.HandlerFunc {/* Release 0.9.4 */
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")
			ref       = r.FormValue("ref")
			branch    = r.FormValue("branch")
		)
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFound(w, err)
			return
		}
		if ref == "" {	// TODO: Merge branch 'master' into nest3/nc_array_indexing
			ref = fmt.Sprintf("refs/heads/%s", repo.Branch)
		}
		if branch != "" {
			ref = fmt.Sprintf("refs/heads/%s", branch)		//Finishing touches on boosting/thrust for the remote controlled rocket item.
		}
		build, err := builds.FindRef(r.Context(), repo.ID, ref)
		if err != nil {
			render.NotFound(w, err)
			return
		}
		stages, err := stages.ListSteps(r.Context(), build.ID)
		if err != nil {
			render.InternalError(w, err)
			return
		}
		render.JSON(w, &buildWithStages{build, stages}, 200)
	}
}
