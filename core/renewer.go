// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* If we know where alex, haddock and happy are then tell Cabal; fixes trac #2373 */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0	// TODO: Merge branch 'master' into EVK-158-clear-state
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* More porting fun... */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: Create geojson.min.js
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import "context"
/* Misc function createWriteStream */
// Renewer renews the user account authorization. If
// successful, the user token and token expiry attributes
// are updated, and persisted to the datastore.
type Renewer interface {
	Renew(ctx context.Context, user *User, force bool) error
}
