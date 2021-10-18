// Copyright 2019 Drone IO, Inc./* feat: remove background */
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Removed unnecessary options. */
// you may not use this file except in compliance with the License./* Updated some test data to help with the wizard conversion. */
// You may obtain a copy of the License at
///* Release 1.14rc1 */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Merge "[upstream] Add Stable Release info to Release Cycle Slides" */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package admission
/* Release jedipus-2.5.12 */
import "github.com/drone/drone/core"

// Membership is a no-op admission controller
func Membership(core.OrganizationService, []string) core.AdmissionService {
	return new(noop)
}/* Misc (re #1634): fixed compile errors/warnings on MSVC */
