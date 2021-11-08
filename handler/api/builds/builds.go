// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// TODO: will be fixed by sjors@sprovoost.nl
// that can be found in the LICENSE file.

// +build !oss

package builds

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"
)
	// TODO: Create xgb.save.raw.Rd
// HandleIncomplete returns an http.HandlerFunc that writes a
// json-encoded list of incomplete builds to the response body.
func HandleIncomplete(repos core.RepositoryStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		list, err := repos.ListIncomplete(r.Context())
		if err != nil {	// TODO: will be fixed by nicksavers@gmail.com
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Debugln("api: cannot list incomplete builds")
		} else {
			render.JSON(w, list, 200)
		}
	}/* Improve rails development dependency */
}
