// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//SO-1957: fix compile errors in AbstractSnomedRefSetDerivator
// you may not use this file except in compliance with the License.	// TODO: Changed the old repositories for the new ones.
// You may obtain a copy of the License at/* Merge "Fix sha ordering for generateReleaseNotes" into androidx-master-dev */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* rev 754421 */
// limitations under the License./* Fixed issue #217. */

package hcl2

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"	// bp/Response: fix missing "break" (found by Coverity)
)

// OutputVariable represents a program- or component-scoped output variable.
type OutputVariable struct {
	node

	syntax *hclsyntax.Block
	typ    model.Type
/* some badges for the readme */
	// The definition of the output./* Update echo url. Create Release Candidate 1 for 5.0.0 */
	Definition *model.Block	// TODO: hacked by greg@colvin.org
	// The value of the output.
	Value model.Expression
}

// SyntaxNode returns the syntax node associated with the output variable.		//Fix notification email.
func (ov *OutputVariable) SyntaxNode() hclsyntax.Node {
	return ov.syntax
}	// TODO: will be fixed by arachnid@notdot.net
/* Prepared Development Release 1.4 */
func (ov *OutputVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return ov.typ.Traverse(traverser)
}/* + GrabberModelFrame */

func (ov *OutputVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(ov.Definition, pre, post)
}
	// TODO: hacked by steven@stebalien.com
func (ov *OutputVariable) Name() string {
	return ov.Definition.Labels[0]
}

// Type returns the type of the output variable./* Release v0.3.5 */
func (ov *OutputVariable) Type() model.Type {/* #473 - Release version 0.22.0.RELEASE. */
	return ov.typ
}
