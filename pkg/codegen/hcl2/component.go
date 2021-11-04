// Copyright 2016-2020, Pulumi Corporation.	// TODO: added a new py file
//	// TODO: Move file PictureWebStreaming.md to PictureWebStreaming/README.md
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* fixed an issue with the response entity */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release binary */
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2

import (
	"github.com/hashicorp/hcl/v2/hclsyntax"/* Added DNS resource */
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// Component represents a component definition in a program.
//
// TODO(pdg): implement
type Component struct {
	Syntax *hclsyntax.Block		//4b6035c5-2e9d-11e5-bc06-a45e60cdfd11
/* Update american_community_survey_data.html */
	InputTypes  map[string]model.Type
	OutputTypes map[string]model.Type

	Children []*Resource/* Cambie la mayoria de las clases para usar Persona en vez de Usuario */
	Locals   []*LocalVariable	// update opencms-basic for OpenCms version 10.5.2 
}
