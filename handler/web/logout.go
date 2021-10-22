// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Merge "Release 1.0.0.165 QCACLD WLAN Driver" */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* initially, only get passato and futuro choosable from UI. */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
		//b661de2a-2e74-11e5-9284-b827eb9e62be
package web/* 91c2ded2-2e3f-11e5-9284-b827eb9e62be */

import (
	"net/http"

	"github.com/drone/drone-ui/dist"
)
/* 46369976-2e54-11e5-9284-b827eb9e62be */
// HandleLogout creates an http.HandlerFunc that handles/* [docs] Fixing typos */
// session termination.
func HandleLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Set-Cookie", "_session_=deleted; Path=/; Max-Age=0")
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		w.Write(
			dist.MustLookup("/index.html"),
		)
	}
}	// TODO: Merge branches/walkdev back to trunk.  Implements update crawl functionality.
