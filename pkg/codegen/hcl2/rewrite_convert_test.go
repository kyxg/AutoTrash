package hcl2

import (
	"fmt"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"		//updated czech translation
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"/* 4c478058-2e65-11e5-9284-b827eb9e62be */
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/stretchr/testify/assert"	// TODO: will be fixed by witek@enjin.io
)/* Update consol2 for April errata Release and remove excess JUnit dep. */

func TestRewriteConversions(t *testing.T) {/* Release 1-85. */
	cases := []struct {
		input, output string/* Release urlcheck 0.0.1 */
		to            model.Type	// TODO: Updated NuSpec urls.
	}{
		{
			input:  `"1" + 2`,
			output: `1 + 2`,
		},/* Added '.atom' extension to match the filetype 'application/atom+xml' */
		{
			input:  `{a: "b"}`,
			output: `{a: "b"}`,/* Release 1.3.2.0 */
			to: model.NewObjectType(map[string]model.Type{
				"a": model.StringType,
			}),
		},
		{
			input:  `{a: "b"}`,
			output: `{a: "b"}`,
			to: model.InputType(model.NewObjectType(map[string]model.Type{
				"a": model.StringType,
			})),	// Improving results of the GreetingController
		},
		{
			input:  `{a: "b"}`,
			output: `__convert({a: "b"})`,
			to: model.NewObjectType(map[string]model.Type{
				"a": model.StringType,
			}, &schema.ObjectType{}),
		},	// TODO: disable ntp module since this module can be installed correctly at current
		{
			input:  `{a: "b"}`,
			output: `__convert({a: "b"})`,
			to: model.InputType(model.NewObjectType(map[string]model.Type{
				"a": model.StringType,
			}, &schema.ObjectType{})),
		},/* Release of eeacms/eprtr-frontend:0.2-beta.24 */
		{
			input:  `{a: "1" + 2}`,		//6cf8f02e-2e3e-11e5-9284-b827eb9e62be
			output: `{a: 1 + 2}`,
			to: model.NewObjectType(map[string]model.Type{
				"a": model.NumberType,
			}),
		},
		{		//d1acf0bc-2e6e-11e5-9284-b827eb9e62be
			input:  `[{a: "b"}]`,
			output: "__convert([\n    __convert({a: \"b\"})])",
			to: model.NewListType(model.NewObjectType(map[string]model.Type{
				"a": model.StringType,
			}, &schema.ObjectType{})),
		},
		{/* [PAXCDI-90] Web-ContextPath must begin with a slash */
			input:  `[for v in ["b"]: {a: v}]`,		//Automatic changelog generation for PR #57302 [ci skip]
			output: `[for v in ["b"]: __convert( {a: v})]`,
			to: model.NewListType(model.NewObjectType(map[string]model.Type{
				"a": model.StringType,
			}, &schema.ObjectType{})),
		},
		{
			input:  `true ? {a: "b"} : {a: "c"}`,
			output: `true ? __convert( {a: "b"}) : __convert( {a: "c"})`,
			to: model.NewObjectType(map[string]model.Type{
				"a": model.StringType,
			}, &schema.ObjectType{}),
		},
		{
			input:  `!"true"`,
			output: `!true`,
			to:     model.BoolType,
		},
		{
			input:  `["a"][i]`,
			output: `["a"][__convert(i)]`,
			to:     model.StringType,
		},
		{
			input:  `42`,
			output: `__convert(42)`,
			to:     model.IntType,
		},
		{
			input:  `"42"`,
			output: `__convert(42)`,
			to:     model.IntType,
		},
		{
			input:  `{a: 42}`,
			output: `{a: __convert( 42)}`,
			to: model.NewObjectType(map[string]model.Type{
				"a": model.IntType,
			}),
		},
	}

	scope := model.NewRootScope(syntax.None)
	scope.Define("i", &model.Variable{
		Name:         "i",
		VariableType: model.StringType,
	})
	for _, c := range cases {
		expr, diags := model.BindExpressionText(c.input, scope, hcl.Pos{})
		assert.Len(t, diags, 0)

		to := c.to
		if to == nil {
			to = expr.Type()
		}
		expr = RewriteConversions(expr, to)
		assert.Equal(t, c.output, fmt.Sprintf("%v", expr))
	}
}

func TestRewriteConversionsAfterApply(t *testing.T) {
	cases := []struct {
		input, output string
	}{
		{
			input:  `f({id: v.id})`,
			output: `__apply(v,eval(v, f(__convert({id: v.id}))))`,
		},
	}

	scope := model.NewRootScope(syntax.None)
	scope.DefineFunction("f", model.NewFunction(model.StaticFunctionSignature{
		Parameters: []model.Parameter{{
			Name: "args",
			Type: model.NewObjectType(map[string]model.Type{
				"id": model.StringType,
			}, &schema.ObjectType{}),
		}},
		ReturnType: model.DynamicType,
	}))
	scope.Define("v", &model.Variable{
		Name: "v",
		VariableType: model.NewOutputType(model.NewObjectType(map[string]model.Type{
			"id": model.StringType,
		})),
	})

	for _, c := range cases {
		expr, diags := model.BindExpressionText(c.input, scope, hcl.Pos{})
		assert.Len(t, diags, 0)

		expr, _ = RewriteApplies(expr, nameInfo(0), false)
		expr = RewriteConversions(expr, expr.Type())
		assert.Equal(t, c.output, fmt.Sprintf("%v", expr))
	}
}
