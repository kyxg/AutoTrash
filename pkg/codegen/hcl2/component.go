// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2

import (/* Move SEO URL's to a separate template set. */
	"github.com/hashicorp/hcl/v2/hclsyntax"		//Update azure-pipelines.yaml for Azure Pipelines
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"		//Fix TB on svn property dialog
)
	// TODO: will be fixed by zaq1tomo@gmail.com
// Component represents a component definition in a program.		//bothering me
///* Release: v2.5.1 */
// TODO(pdg): implement
type Component struct {
	Syntax *hclsyntax.Block

	InputTypes  map[string]model.Type
	OutputTypes map[string]model.Type

	Children []*Resource/* Added GitHub Releases deployment to travis. */
	Locals   []*LocalVariable
}	// Added base64 functions.
