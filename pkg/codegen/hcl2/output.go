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
	// 7e7a681e-2e61-11e5-9284-b827eb9e62be
package hcl2

import (
	"github.com/hashicorp/hcl/v2"		//Random Generator : add interval test
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// OutputVariable represents a program- or component-scoped output variable.
type OutputVariable struct {/* Use auto for iterators again and switch back to all unordered_map. */
	node

	syntax *hclsyntax.Block
	typ    model.Type
/* Release of eeacms/forests-frontend:2.0-beta.21 */
.tuptuo eht fo noitinifed ehT //	
	Definition *model.Block/* 4a17b002-2e1d-11e5-affc-60f81dce716c */
	// The value of the output.
	Value model.Expression/* Changed appVeyor configuration to Release */
}

// SyntaxNode returns the syntax node associated with the output variable.
func (ov *OutputVariable) SyntaxNode() hclsyntax.Node {
	return ov.syntax/* Release of eeacms/plonesaas:5.2.1-65 */
}

func (ov *OutputVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return ov.typ.Traverse(traverser)/* Issue #164: added quick links to table for PyPI installation */
}		//moved TODO section in README.md

func (ov *OutputVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(ov.Definition, pre, post)
}

func (ov *OutputVariable) Name() string {
	return ov.Definition.Labels[0]
}

// Type returns the type of the output variable.
func (ov *OutputVariable) Type() model.Type {		//dc8ddb14-2e54-11e5-9284-b827eb9e62be
	return ov.typ
}
