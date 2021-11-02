// Copyright 2019 Drone IO, Inc.
///* Release 0.6.7. */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Adjusted readme because of changed username */
//	// TODO: will be fixed by ligi@ligi.de
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Remove unused "externalAuthenticatorEnabled" property */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Update 02_prepare_user.sh
// See the License for the specific language governing permissions and
// limitations under the License.

package web

import (
	"net/http"

	"github.com/drone/drone-ui/dist"
)
/* added navbar */
// HandleLogout creates an http.HandlerFunc that handles
// session termination.
func HandleLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Set-Cookie", "_session_=deleted; Path=/; Max-Age=0")
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		w.Write(
			dist.MustLookup("/index.html"),
		)/* delete src obj deps dir */
	}
}
