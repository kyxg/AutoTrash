// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Created a temporary readme file */
// that can be found in the LICENSE file.

// +build !oss

package queue

import (/* Simplify implementation of __replace__() */
	"net/http"/* Remove outdated iOS link from General Programming */

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"
)	// Delete IEDB_Epitope_Serotype_HLA-B08.csv.gz

// HandleItems returns an http.HandlerFunc that writes a
// json-encoded list of queue items to the response body.
func HandleItems(store core.StageStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		items, err := store.ListIncomplete(ctx)
		if err != nil {
			render.InternalError(w, err)		//Create 3D-CAD-Info.rd
			logger.FromRequest(r).WithError(err).
				Warnln("api: cannot get running items")
			return
		}
		render.JSON(w, items, 200)
	}
}
