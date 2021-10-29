// Copyright 2019 Drone IO, Inc.
//	// freepornhq.xxx
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* fix(package): update ramda to version 0.26.0 */
//      http://www.apache.org/licenses/LICENSE-2.0		//Rename SPI.cpp to spi.cpp
//		//77cf6762-4b19-11e5-ba4d-6c40088e03e4
// Unless required by applicable law or agreed to in writing, software	// Updating to chronicle-bytes 1.12.12
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss		//Larger limit

package converter
/* [artifactory-release] Release version 2.0.0.M3 */
import (
	"time"

	"github.com/drone/drone/core"
)/* build: Release version 0.2 */

// Remote returns a conversion service that converts the
// configuration file using a remote http service.
func Remote(endpoint, signer, extension string, skipVerify bool, timeout time.Duration) core.ConvertService {
	return new(noop)
}
