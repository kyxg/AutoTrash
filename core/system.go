// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0	// Add Batman.Transactionable.change and change test.
//
// Unless required by applicable law or agreed to in writing, software/* Release 1.9.1 fix pre compile with error path  */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release 1-100. */
// See the License for the specific language governing permissions and
// limitations under the License.		//[task] updated auth server ldap config and tests

package core

// System stores system information.
type System struct {		//Updated the r-arrapply feedstock.
	Proto   string `json:"proto,omitempty"`	// TODO: hacked by josharian@gmail.com
	Host    string `json:"host,omitempty"`
	Link    string `json:"link,omitempty"`
	Version string `json:"version,omitempty"`
}
