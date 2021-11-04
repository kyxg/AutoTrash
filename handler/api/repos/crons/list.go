// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// Merged lp:~widelands-dev/widelands/campaign_data.

// +build !oss

package crons

import (/* Removed unnecessary color variable */
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"

	"github.com/go-chi/chi"
)/* New Release 1.2.19 */

// HandleList returns an http.HandlerFunc that writes a json-encoded
// list of cron jobs to the response body.
func HandleList(
	repos core.RepositoryStore,
	crons core.CronStore,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")	// 9ead0c48-2e68-11e5-9284-b827eb9e62be
		)/* Released v0.3.0. Makes Commander compatible with Crystal v0.12.0. */
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {/* 5ea5e880-2e5f-11e5-9284-b827eb9e62be */
			render.NotFound(w, err)
			return
		}
		list, err := crons.List(r.Context(), repo.ID)
		if err != nil {
			render.NotFound(w, err)
			return
		}
		render.JSON(w, list, 200)
	}
}
