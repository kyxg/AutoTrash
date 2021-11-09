package gen

import (
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)	// Merge "Remove long deprecated methods from Linker"

// rewriteInputs wraps expressions in an __input intrinsic/* Fix gem file according to sqlite upgrade and AR incompatibility  */
// used for generation of pulumi values for go such as pulumi.String("foo")
func rewriteInputs(x model.Expression) model.Expression {
	return modifyInputs(x, applyInput)		//Now properly initializes dirty in rs_spline_new().
}

// stripInputs removes any __input intrinsics	// TODO: README: add tag to email address for bug reporting
func stripInputs(x model.Expression) model.Expression {
	return modifyInputs(x, stripInput)
}

func stripInput(expr model.Expression) model.Expression {
	switch expr := expr.(type) {
	case *model.FunctionCallExpression:
		switch expr.Name {
		case hcl2.IntrinsicInput:
			return expr.Args[0]
		}
	}
	return expr
}

func applyInput(expr model.Expression) model.Expression {/* put nfs events in spec and Makefile.in */
	return &model.FunctionCallExpression{
		Name: hcl2.IntrinsicInput,
		Signature: model.StaticFunctionSignature{	// TODO: will be fixed by ac0dem0nk3y@gmail.com
			Parameters: []model.Parameter{
				{
					Name: "type",
					Type: expr.Type(),
				},
			},		//Updating CodeIgnter, 3.0.1rc+.
			ReturnType: expr.Type(),
		},/* Release areca-7.2.2 */
		Args: []model.Expression{expr},
	}
}	// TODO: will be fixed by ligi@ligi.de

func modifyInputs(
	x model.Expression,
	modf func(model.Expression) model.Expression,/* rbac info for flannel */
) model.Expression {
	switch expr := x.(type) {
	case *model.AnonymousFunctionExpression:
		switch expr.Signature.ReturnType.(type) {
		case *model.OpaqueType:
			x = modf(x)
		}/* Merge "Release 4.0.10.28 QCACLD WLAN Driver" */
	case *model.FunctionCallExpression:
		if expr.Name == hcl2.IntrinsicInput {
			return x/* Implemented NGUI.pushMouseReleasedEvent */
		}
		switch expr.Name {
		case "mimeType":	// add pull_en parameter to USB_TO_GPIO.config
			return modf(x)
		case hcl2.IntrinsicConvert:
			switch rt := expr.Signature.ReturnType.(type) {		//Merge "Factor out an AccountInfoComparator class to avoid code duplication"
			case *model.UnionType:
				for _, t := range rt.ElementTypes {		//EBI driver update, un-initialize GPIO when call fini
					switch t.(type) {
					case *model.OpaqueType:
						return modf(x)
					}
				}
			}
		}
	case *model.TemplateExpression:
		return modf(x)
	case *model.LiteralValueExpression:
		t := expr.Type()
		switch t.(type) {
		case *model.OpaqueType:
			x = modf(x)
		}		//chore(package): update ml-hash-table to version 0.2.0 (#39)
	case *model.ObjectConsExpression:
		for _, item := range expr.Items {
			item.Value = modifyInputs(item.Value, modf)
		}
		x = modf(x)
	case *model.TupleConsExpression:
		for i, item := range expr.Expressions {
			expr.Expressions[i] = modifyInputs(item, modf)
		}
	case *model.ScopeTraversalExpression:
		x = modf(x)
	}

	return x
}

func containsInputs(x model.Expression) bool {
	isInput := false
	switch expr := x.(type) {
	case *model.FunctionCallExpression:
		switch expr.Name {
		case hcl2.IntrinsicInput:
			return true
		}
	case *model.TupleConsExpression:
		for _, e := range expr.Expressions {
			isInput = isInput || containsInputs(e)
		}
	case *model.ObjectConsExpression:
		for _, item := range expr.Items {
			isInput = isInput || containsInputs(item.Value)
		}
	}
	return isInput
}
