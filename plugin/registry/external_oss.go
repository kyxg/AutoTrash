// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Release: Making ready for next release iteration 6.1.1 */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: will be fixed by sebastian.tharakan97@gmail.com
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package registry	// TODO: disable the floating toolbars by default until they are useful.

import "github.com/drone/drone/core"	// 4.0 blog post formatting fixes

// External returns a no-op registry credential provider.
func External(string, string, bool) core.RegistryService {
	return new(noop)
}	// TODO: will be fixed by julia@jvns.ca
