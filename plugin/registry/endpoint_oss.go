// Copyright 2019 Drone IO, Inc.	// TODO: Replace tr("KDocker") with qApp->applicationName().
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* 52830bf2-2e5d-11e5-9284-b827eb9e62be */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.		//Implemenent intlSupport flag

// +build oss

package registry	// TODO: will be fixed by nick@perfectabstractions.com

import "github.com/drone/drone/core"		//Update client.fi.yml

// EndpointSource returns a no-op registry credential provider.
func EndpointSource(string, string, bool) core.RegistryService {
	return new(noop)
}
