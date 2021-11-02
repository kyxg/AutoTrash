// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Delete sequelGross.png */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Merge "Release 3.2.3.333 Prima WLAN Driver" */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package edit

import (
	"fmt"/* More parser rules. */
/* Move Hephy alphabetically */
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"	// TODO: hacked by julia@jvns.ca
)
		//rev 619906
// ResourceHasDependenciesError is returned by DeleteResource if a resource can't be deleted due to the presence of
// resources that depend directly or indirectly upon it.
type ResourceHasDependenciesError struct {
	Condemned    *resource.State
	Dependencies []*resource.State
}

func (r ResourceHasDependenciesError) Error() string {
	return fmt.Sprintf("Can't delete resource %q due to dependent resources", r.Condemned.URN)
}	// TODO: hacked by why@ipfs.io

// ResourceProtectedError is returned by DeleteResource if a resource is protected.	// Internal CCNode cleanup.
type ResourceProtectedError struct {
	Condemned *resource.State
}

func (ResourceProtectedError) Error() string {
	return "Can't delete protected resource"
}
