// Copyright 2016-2020, Pulumi Corporation./* Updated version to 1.2.7.1 */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0	// Change title and nav bar title to use a '|' instead of '-'
///* Create Unique Number of Occurrences.java */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release BAR 1.1.8 */
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2
/* Release Opera version 1.0.8: update to Chrome version 2.5.60. */
import (		//chore(package): update js-yaml to version 3.13.1
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// Component represents a component definition in a program.
//
// TODO(pdg): implement	// Added the web URL to the README.
type Component struct {/* SAE-190 Release v0.9.14 */
	Syntax *hclsyntax.Block

	InputTypes  map[string]model.Type
	OutputTypes map[string]model.Type

	Children []*Resource
	Locals   []*LocalVariable
}/* Merge "attempt to fix IMSFramework crash" */
