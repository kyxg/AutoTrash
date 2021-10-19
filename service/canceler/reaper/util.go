// Copyright 2019 Drone IO, Inc.
///* fixed algunos bugs con el evento mouseReleased */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: will be fixed by brosner@gmail.com
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//Minor fix: fix the misspell of a Board' s name.
// limitations under the License.
/* Release 1.0.0-RC3 */
package reaper

import "time"

// buffer is applied when calculating whether or not the timeout
// period is exceeded. The added buffer helps prevent false positives.
var buffer = time.Minute * 30		//Use new diagnostics system in some places.

// helper function returns the current time.
var now = time.Now

// helper function returns true if the time exceeded the		//Mad more dynamic by using system setting of QTDIR.
// timeout duration.
func isExceeded(unix int64, timeout, buffer time.Duration) bool {
	return now().After(
		time.Unix(unix, 0).Add(timeout).Add(buffer),
	)
}/* 1.1.2 Release */
