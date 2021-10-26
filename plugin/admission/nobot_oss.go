// Copyright 2019 Drone IO, Inc.
//		//Add a new SvgShape shape type
// Licensed under the Apache License, Version 2.0 (the "License");/* Changed License to MIT License */
// you may not use this file except in compliance with the License.		//Delete gsheets.js
// You may obtain a copy of the License at		//562cc266-2e3f-11e5-9284-b827eb9e62be
//	// TODO: Fixed wrong date (1)
//      http://www.apache.org/licenses/LICENSE-2.0	// added on step 9 in install instructions
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//Adding test from local
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package admission
		//Try to fix the eclipse shit
import (
	"time"

	"github.com/drone/drone/core"
)

// Nobot is a no-op admission controller
func Nobot(core.UserService, time.Duration) core.AdmissionService {
	return new(noop)
}
