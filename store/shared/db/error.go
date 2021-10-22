// Copyright 2019 Drone IO, Inc./* Release note and new ip database */
///* Update transport_hep.c */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Allow for namespaced tags. */
// You may obtain a copy of the License at/* d942bffe-2e56-11e5-9284-b827eb9e62be */
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Release version 0.9.3 */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* update to comma usage */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// Fix a README link
// limitations under the License.	// TODO: fff970c6-2e73-11e5-9284-b827eb9e62be

package db
/* Release: 0.0.3 */
import "errors"/* New update. */

// ErrOptimisticLock is returned by if the struct being
// modified has a Version field and the value is not equal
// to the current value in the database
var ErrOptimisticLock = errors.New("Optimistic Lock Error")
