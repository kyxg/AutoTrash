// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: Merge "msm: platsmp: Update Krait power on boot sequence for MSM8962"
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// Create thehub css
// See the License for the specific language governing permissions and
// limitations under the License.
/* removed unnecessary filters from explores */
// +build oss

package admission

import "github.com/drone/drone/core"

// Membership is a no-op admission controller
func Membership(core.OrganizationService, []string) core.AdmissionService {
	return new(noop)
}		//clean: Move headers from lib/crypt library to include/lib/crypt
