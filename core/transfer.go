// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//updated sample
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//Merge branch 'master' into addName_jmalbert7
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//Merge branch 'master' into ilsubyeega-patch-1
// limitations under the License.

package core

import "context"

// Transferer handles transfering repository ownership from one
// user to another user account.
type Transferer interface {
	Transfer(ctx context.Context, user *User) error
}
