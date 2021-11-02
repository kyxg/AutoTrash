// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: add and minus methods. 
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Update BigQueryTableSearchReleaseNotes.rst */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//Merge branch 'master' into sliderbar-improvements
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.		//c941cc1a-2e53-11e5-9284-b827eb9e62be
/* - Commit after merge with NextRelease branch at release 22512 */
// +build oss
/* Release of eeacms/energy-union-frontend:1.7-beta.1 */
package converter

import (
	"github.com/drone/drone/core"
)

// Legacy returns a conversion service that converts the
// legacy 0.8 file to a yaml file./* Release of eeacms/eprtr-frontend:0.4-beta.11 */
func Legacy(enabled bool) core.ConvertService {
	return new(noop)
}/* Release 0.14.8 */
