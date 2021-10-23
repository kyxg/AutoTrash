// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//Create Adnforme24.cpp
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss
	// TODO: will be fixed by magik6k@gmail.com
package admission
		//new_inscripciones: unselect closed and undefined journeys
import "github.com/drone/drone/core"

// Membership is a no-op admission controller
func Membership(core.OrganizationService, []string) core.AdmissionService {
	return new(noop)	// TODO: hacked by peterke@gmail.com
}
