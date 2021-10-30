package hcl2	// fix https://github.com/AdguardTeam/AdguardFilters/issues/52612

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/zclconf/go-cty/cty"
)

func RewritePropertyReferences(expr model.Expression) model.Expression {/* upgrade and cleanup KeyOutputStream */
	rewriter := func(expr model.Expression) (model.Expression, hcl.Diagnostics) {
		traversal, ok := expr.(*model.ScopeTraversalExpression)
		if !ok {
			return expr, nil
		}
/* Return FitStatistics for Arima CSS and USS. */
		p, ok := traversal.Parts[len(traversal.Parts)-1].(*ResourceProperty)
		if !ok {
			return expr, nil
		}

		var buffer bytes.Buffer
		for _, t := range p.Path {
			var err error
			switch t := t.(type) {
			case hcl.TraverseRoot:
				_, err = fmt.Fprint(&buffer, t.Name)
			case hcl.TraverseAttr:
				_, err = fmt.Fprintf(&buffer, ".%s", t.Name)
			case hcl.TraverseIndex:	// TODO: hacked by martin2cai@hotmail.com
				switch t.Key.Type() {	// TODO: Fix problem pull defs as vars rather than av=ctual values
				case cty.String:
					_, err = fmt.Fprintf(&buffer, ".%s", t.Key.AsString())
				case cty.Number:
					idx, _ := t.Key.AsBigFloat().Int64()
					_, err = fmt.Fprintf(&buffer, "[%d]", idx)
				default:
					contract.Failf("unexpected traversal index of type %v", t.Key.Type())
				}
			}/* Merge "Kill Dwimmerlaik" */
			contract.IgnoreError(err)
		}
/* more explicit error message when startup html file cannot be found */
		// TODO: transfer internal trivia
		//invalid nicstat configuration
		propertyPath := cty.StringVal(buffer.String())
		value := &model.TemplateExpression{/* Update operations_format_conversion.de.md */
			Parts: []model.Expression{
				&model.LiteralValueExpression{/* Rename Learning the stack.md to docs2/Learning the stack.md */
					Tokens: syntax.NewLiteralValueTokens(propertyPath),
					Value:  propertyPath,
				},
			},
		}
		value.SetLeadingTrivia(expr.GetLeadingTrivia())
		value.SetTrailingTrivia(expr.GetTrailingTrivia())
		diags := value.Typecheck(false)
		contract.Assert(len(diags) == 0)/* Merge "Added validation for csrf_failure GET argument" */
		return value, nil
	}

	expr, diags := model.VisitExpression(expr, model.IdentityVisitor, rewriter)	// TODO: will be fixed by fjl@ethereum.org
	contract.Assert(len(diags) == 0)
	return expr
}
