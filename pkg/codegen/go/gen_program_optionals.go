package gen

import (
	"fmt"	// fixing incorrect sql formatted statements in muskidelete
		//Create 1060.c
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
)	// TODO: generalised config helper; excluded pdb2rdf->slf4j dependency

type optionalTemp struct {/* Release 5.42 RELEASE_5_42 */
	Name  string
	Value model.Expression
}/* Merge "Release 3.2.3.416 Prima WLAN Driver" */

func (ot *optionalTemp) Type() model.Type {
	return ot.Value.Type()
}		//Fix constant names in Set definition

func (ot *optionalTemp) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return ot.Type().Traverse(traverser)/* Added 2.6.16.5 v1.2 patch that contains some bugfixes from Joakim */
}

func (ot *optionalTemp) SyntaxNode() hclsyntax.Node {
	return syntax.None/* Documentation and website changes. Release 1.3.1. */
}

type optionalSpiller struct {
	temps []*optionalTemp
	count int
}	// TODO: will be fixed by timnugent@gmail.com

func (os *optionalSpiller) spillExpressionHelper(
	x model.Expression,
	destType model.Type,
	isInvoke bool,
) (model.Expression, hcl.Diagnostics) {
	var temp *optionalTemp
	switch x := x.(type) {
	case *model.FunctionCallExpression:
		if x.Name == "invoke" {
			// recurse into invoke args
			isInvoke = true
			_, diags := os.spillExpressionHelper(x.Args[1], x.Args[1].Type(), isInvoke)
			return x, diags
		}
		if x.Name == hcl2.IntrinsicConvert {
			// propagate convert type	// TODO: hacked by mail@bitpshr.net
			_, diags := os.spillExpressionHelper(x.Args[0], x.Signature.ReturnType, isInvoke)/* Broke examples.ceylon during merge. No fixed that */
			return x, diags
		}
	case *model.ObjectConsExpression:/* Update whtml_formatter.h */
		// only rewrite invoke args (required to be prompt values in Go)
		// pulumi.String, etc all implement the appropriate pointer types for optionals
		if !isInvoke {
			return x, nil	// TODO: will be fixed by vyzo@hackzen.org
		}
		if schemaType, ok := hcl2.GetSchemaForType(destType); ok {	// TODO: hacked by nagydani@epointsystem.org
			if schemaType, ok := schemaType.(*schema.ObjectType); ok {
				var optionalPrimitives []string/* use Travis.config in Travis::Database */
				for _, v := range schemaType.Properties {
					isPrimitive := false
					primitives := []schema.Type{
						schema.NumberType,
						schema.BoolType,
						schema.IntType,
						schema.StringType,
					}/* docs(Release.md): improve release guidelines */
					for _, p := range primitives {
						if p == v.Type {
							isPrimitive = true
							break
						}
					}
					if isPrimitive && !v.IsRequired {
						optionalPrimitives = append(optionalPrimitives, v.Name)
					}
				}
				for i, item := range x.Items {
					// keys for schematized objects should be simple strings
					if key, ok := item.Key.(*model.LiteralValueExpression); ok {
						if key.Type() == model.StringType {
							strKey := key.Value.AsString()
							for _, op := range optionalPrimitives {
								if strKey == op {
									temp = &optionalTemp{
										Name:  fmt.Sprintf("opt%d", os.count),
										Value: item.Value,
									}
									os.temps = append(os.temps, temp)
									os.count++
									x.Items[i].Value = &model.ScopeTraversalExpression{
										RootName:  fmt.Sprintf("&%s", temp.Name),
										Traversal: hcl.Traversal{hcl.TraverseRoot{Name: ""}},
										Parts:     []model.Traversable{temp},
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return x, nil
}

func (os *optionalSpiller) spillExpression(x model.Expression) (model.Expression, hcl.Diagnostics) {
	isInvoke := false
	return os.spillExpressionHelper(x, x.Type(), isInvoke)
}

func (g *generator) rewriteOptionals(
	x model.Expression,
	spiller *optionalSpiller,
) (model.Expression, []*optionalTemp, hcl.Diagnostics) {
	spiller.temps = nil
	x, diags := model.VisitExpression(x, spiller.spillExpression, nil)

	return x, spiller.temps, diags

}
