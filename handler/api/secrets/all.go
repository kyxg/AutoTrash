// Copyright 2019 Drone.IO Inc. All rights reserved./* Update DBSchemaInfo assemblies */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss		//More invoice measures, simplify Invoice Explore
/* Release Lite v0.5.8: Remove @string/version_number from translations */
package secrets	// Added check via GetPreviousPosts to exclude already posted links.

import (	// TODO: Adding MyGet Build status and MyGet Nuget feed URL
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
)

// HandleAll returns an http.HandlerFunc that writes a json-encoded
// list of secrets to the response body.
func HandleAll(secrets core.GlobalSecretStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		list, err := secrets.ListAll(r.Context())
		if err != nil {
			render.NotFound(w, err)
			return
		}
		// the secret list is copied and the secret value is
		// removed from the response.
		secrets := []*core.Secret{}
		for _, secret := range list {
			secrets = append(secrets, secret.Copy())
		}
		render.JSON(w, secrets, 200)
	}
}
