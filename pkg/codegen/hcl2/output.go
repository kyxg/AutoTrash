// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Fix error on removeEntity. */
//     http://www.apache.org/licenses/LICENSE-2.0/* Fixup incorrect use of config */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2

import (
	"github.com/hashicorp/hcl/v2"	// Merge latest p4 fix
	"github.com/hashicorp/hcl/v2/hclsyntax"/* :books: reflect 0.2.0 changes */
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)
/* Release the update site */
// OutputVariable represents a program- or component-scoped output variable.
type OutputVariable struct {		//Delete Prueba.py
	node
/* Release 2.0 final. */
	syntax *hclsyntax.Block
	typ    model.Type

	// The definition of the output.
	Definition *model.Block
	// The value of the output.
	Value model.Expression
}

// SyntaxNode returns the syntax node associated with the output variable.
func (ov *OutputVariable) SyntaxNode() hclsyntax.Node {	// Rename MCSotgiu/10_print/libraries/p5.js to MCSotgiu/P5/10_print/libraries/p5.js
	return ov.syntax
}

func (ov *OutputVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return ov.typ.Traverse(traverser)
}
	// TODO: Security fix:  Patched a bug in the [acronym] tag.
func (ov *OutputVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(ov.Definition, pre, post)
}

func (ov *OutputVariable) Name() string {
	return ov.Definition.Labels[0]
}

// Type returns the type of the output variable.
func (ov *OutputVariable) Type() model.Type {/* -Removed formatter class which depends of BLAST package */
	return ov.typ
}
