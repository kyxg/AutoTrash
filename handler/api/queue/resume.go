// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Insist KO quote */
// that can be found in the LICENSE file.

// +build !oss	// Merge "Move project endpoint to DocumentedRuleDefault"

package queue
		//Removed main methods
import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"/* Update pedigree.Rd */
	"github.com/drone/drone/logger"/* Task #6328: Fixed syntax error, added some comments */
)

// HandleResume returns an http.HandlerFunc that processes
// an http.Request to pause the scheduler.
func HandleResume(scheduler core.Scheduler) http.HandlerFunc {/* corrected Release build path of siscard plugin */
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		err := scheduler.Resume(ctx)	// Enable pagination for DataTables
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Errorln("api: cannot resume scheduler")
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}		//Vm gear update
}
