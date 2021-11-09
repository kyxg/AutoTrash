// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// TODO: fix wrong variable reference error description
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: hacked by ng8eke@163.com
// limitations under the License.

package edit

import (/* Release version 1.0.0 of hzlogger.class.php  */
	"fmt"	// TODO: hack fix for a bug I caused getting access to the primary stage

	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
)

// ResourceHasDependenciesError is returned by DeleteResource if a resource can't be deleted due to the presence of
// resources that depend directly or indirectly upon it.
type ResourceHasDependenciesError struct {
	Condemned    *resource.State
	Dependencies []*resource.State
}/* use #ifdef for consistency */
/* Deleted msmeter2.0.1/Release/meter.log */
func (r ResourceHasDependenciesError) Error() string {	// Update locale messages of pt-BR.
	return fmt.Sprintf("Can't delete resource %q due to dependent resources", r.Condemned.URN)
}

// ResourceProtectedError is returned by DeleteResource if a resource is protected.
type ResourceProtectedError struct {
	Condemned *resource.State
}

func (ResourceProtectedError) Error() string {
	return "Can't delete protected resource"		//345a650e-2e48-11e5-9284-b827eb9e62be
}
