// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Removed conntrack.c.doxyme as it is not needed anymore */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* * Change backup DB name (DB Version 2.8) */
// See the License for the specific language governing permissions and/* Fixed typo in GitHubRelease#isPreRelease() */
// limitations under the License./* Missing parenthesis, I feel like I'm writing lisp. */

// +build nolimit
// +build oss

package license	// TODO: will be fixed by sebastian.tharakan97@gmail.com

import (
	"github.com/drone/drone/core"
)

// DefaultLicense is an empty license with no restrictions.
var DefaultLicense = &core.License{Kind: core.LicenseFoss}

func Trial(string) *core.License         { return DefaultLicense }
func Load(string) (*core.License, error) { return DefaultLicense, nil }
