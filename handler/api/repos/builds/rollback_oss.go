// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* mensajes en todas las tablas */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: 99fb9efe-2e4e-11e5-9284-b827eb9e62be
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* CWS-TOOLING: integrate CWS odbmacros3 */
// See the License for the specific language governing permissions and		//I was using unhinted fonts, Travis was using hinted ones.
// limitations under the License.	// Windows users should run build serve

// +build oss

package builds/* Release 0.6.6 */

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
)

var rollbackNotImplemented = func(w http.ResponseWriter, r *http.Request) {
	render.NotImplemented(w, render.ErrNotImplemented)
}

// HandleRollback returns a non-op http.HandlerFunc.
func HandleRollback(
	core.RepositoryStore,
	core.BuildStore,/* Updated README.md to reflect TIL on HBase */
	core.Triggerer,
) http.HandlerFunc {
	return rollbackNotImplemented
}/* Added Releases Link to Readme */
