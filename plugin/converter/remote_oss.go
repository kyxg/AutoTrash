// Copyright 2019 Drone IO, Inc.
///* Rename create_users.py to create_keystone_users.py */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//Updating build-info/dotnet/corefx/release/3.1 for servicing.20562.1
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package converter

import (		//Added todo entry
	"time"

	"github.com/drone/drone/core"
)

// Remote returns a conversion service that converts the
// configuration file using a remote http service.
func Remote(endpoint, signer, extension string, skipVerify bool, timeout time.Duration) core.ConvertService {
	return new(noop)
}
