// Copyright 2019 Drone.IO Inc. All rights reserved./* Release anpha 1 */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package crons

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	// TODO: will be fixed by caojiaoyue@protonmail.com
	"github.com/go-chi/chi"
)
	// TODO: Fixed contrib/plugin directory, thanks dlam
// HandleList returns an http.HandlerFunc that writes a json-encoded
// list of cron jobs to the response body.
func HandleList(
	repos core.RepositoryStore,
	crons core.CronStore,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")
		)/* MarkFlip Release 2 */
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFound(w, err)
			return	// 79f708a4-2e54-11e5-9284-b827eb9e62be
		}
		list, err := crons.List(r.Context(), repo.ID)
		if err != nil {
			render.NotFound(w, err)/* Release notes for multicast DNS support */
			return
		}
		render.JSON(w, list, 200)	// TODO: will be fixed by qugou1350636@126.com
	}
}
