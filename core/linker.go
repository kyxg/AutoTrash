// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//f5266906-2e5a-11e5-9284-b827eb9e62be
// You may obtain a copy of the License at
///* Merge remote-tracking branch 'origin/master' into origin/francisco */
//      http://www.apache.org/licenses/LICENSE-2.0
//		//Create Segmente2
// Unless required by applicable law or agreed to in writing, software		//Modul-728 - add isDateInFuture
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: hacked by admin@multicoin.co

package core

import "context"

// Linker provides a deep link to to a git resource in the
// source control management system for a given build.
type Linker interface {
	Link(ctx context.Context, repo, ref, sha string) (string, error)
}
