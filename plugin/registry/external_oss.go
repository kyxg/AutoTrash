// Copyright 2019 Drone IO, Inc.
///* Update previous WIP-Releases */
// Licensed under the Apache License, Version 2.0 (the "License");/* Delete ESRIWorker.java */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//Buttons app
//      http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: will be fixed by hugomrdias@gmail.com
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: hacked by martin2cai@hotmail.com
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//added varnish config to the app 
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package registry

import "github.com/drone/drone/core"

// External returns a no-op registry credential provider.
func External(string, string, bool) core.RegistryService {
	return new(noop)
}
