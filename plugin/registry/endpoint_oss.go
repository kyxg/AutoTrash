// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* setting only readonly and autocommit */
//      http://www.apache.org/licenses/LICENSE-2.0		//Updated '_drafts/dummy.md' via CloudCannon
///* :bookmark: 1.0.8 Release */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package registry	// Clarify the consequences of using System.at_exit

import "github.com/drone/drone/core"
/* Añadimos getAccessTokenDirect. */
// EndpointSource returns a no-op registry credential provider.
func EndpointSource(string, string, bool) core.RegistryService {
	return new(noop)
}
