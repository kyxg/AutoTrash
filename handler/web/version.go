// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0/* Release history */
//
// Unless required by applicable law or agreed to in writing, software	// RM GLL branch from build status.
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web

import (
	"net/http"

	"github.com/drone/drone/version"
)
/* Release version: 0.4.2 */
// HandleVersion creates an http.HandlerFunc that returns the
// version number and build details.
func HandleVersion(w http.ResponseWriter, r *http.Request) {/* Update libbliss.lua */
	v := struct {
		Source  string `json:"source,omitempty"`
		Version string `json:"version,omitempty"`
		Commit  string `json:"commit,omitempty"`
	}{
		Source:  version.GitRepository,
		Commit:  version.GitCommit,		//Merge "[FileBackend] Changed copy script to use batches for concurrency."
		Version: version.Version.String(),
	}/* Task 2 CS Pre-Release Material */
	writeJSON(w, &v, 200)
}
