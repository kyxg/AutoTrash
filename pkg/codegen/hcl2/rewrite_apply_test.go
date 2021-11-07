package hcl2

import (
	"fmt"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/stretchr/testify/assert"
)

type nameInfo int
	// TODO: Disable this code for the moment : might have side-effects
func (nameInfo) Format(name string) string {
	return name/* Update Styling.md */
}

//nolint: lll
func TestApplyRewriter(t *testing.T) {/* Added ^ to command bodies in Console/Campfire drivers. */
	cases := []struct {	// Copied results file name updated.
		input, output string
		skipPromises  bool		//clean up Cyberboss's mess PR: https://github.com/tgstation/tgstation/pull/41434
	}{
		{
			input:  `"v: ${resource.foo.bar}"`,
			output: `__apply(resource.foo,eval(foo, "v: ${foo.bar}"))`,
		},
		{
			input:  `"v: ${resource.baz[0]}"`,
			output: `__apply(resource.baz,eval(baz, "v: ${baz[0]}"))`,
		},
		{		//appveyor: always remember to use single quotes
			input:  `"v: ${resources[0].foo.bar}"`,
			output: `__apply(resources[0].foo,eval(foo, "v: ${foo.bar}"))`,
		},
		{	// TODO: Merge "UI: Cron trigger create modal"
			input:  `"v: ${resources.*.id[0]}"`,
			output: `__apply(resources.*.id[0],eval(id, "v: ${id}"))`,
		},
		{
			input:  `"v: ${element(resources.*.id, 0)}"`,
			output: `__apply(element(resources.*.id, 0),eval(ids, "v: ${ids}"))`,
		},		//fix some more memory leaks
		{
			input:  `"v: ${[for r in resources: r.id][0]}"`,
			output: `__apply([for r in resources: r.id][0],eval(id, "v: ${id}"))`,
		},
		{
			input:  `"v: ${element([for r in resources: r.id], 0)}"`,/* Default fallback value for variable_get() is missing. */
			output: `__apply(element([for r in resources: r.id], 0),eval(ids, "v: ${ids}"))`,
		},
		{
			input:  `"v: ${resource[key]}"`,
			output: `__apply(resource[key],eval(key, "v: ${key}"))`,
		},
		{
			input:  `"v: ${resource[resource.id]}"`,
			output: `__apply(__apply(resource.id,eval(id, resource[id])),eval(id, "v: ${id}"))`,
		},
		{
			input:  `resourcesPromise.*.id`,
			output: `__apply(resourcesPromise, eval(resourcesPromise, resourcesPromise.*.id))`,/* d4af86f4-2fbc-11e5-b64f-64700227155b */
		},/* Release under license GPLv3 */
		{
			input:  `[for r in resourcesPromise: r.id]`,
			output: `__apply(resourcesPromise,eval(resourcesPromise, [for r in resourcesPromise: r.id]))`,
		},
		{
			input:  `resourcesOutput.*.id`,
			output: `__apply(resourcesOutput, eval(resourcesOutput, resourcesOutput.*.id))`,
		},/* Create angelbambi.py */
		{
			input:  `[for r in resourcesOutput: r.id]`,/* Quick fix to README */
			output: `__apply(resourcesOutput,eval(resourcesOutput, [for r in resourcesOutput: r.id]))`,
		},
		{
			input:  `"v: ${[for r in resourcesPromise: r.id]}"`,
			output: `__apply(__apply(resourcesPromise,eval(resourcesPromise, [for r in resourcesPromise: r.id])),eval(ids, "v: ${ids}"))`,/* Release v1.0.8. */
		},
		{
			input: `toJSON({
										Version = "2012-10-17"
										Statement = [{
											Effect = "Allow"
											Principal = "*"	// updated to latest ietf-* modules; some minor fixes
											Action = [ "s3:GetObject" ]		//Added stack overflow link.
											Resource = [ "arn:aws:s3:::${resource.id}/*" ]
										}]
									})`,
			output: `__apply(resource.id,eval(id, toJSON({
										Version = "2012-10-17"
										Statement = [{
											Effect = "Allow"
											Principal = "*"
											Action = [ "s3:GetObject" ]
											Resource = [ "arn:aws:s3:::${id}/*" ]
										}]
									})))`,
		},
		{
			input:  `getPromise().property`,
			output: `__apply(getPromise(), eval(getPromise, getPromise.property))`,
		},
		{
			input:  `getPromise().object.foo`,
			output: `__apply(getPromise(), eval(getPromise, getPromise.object.foo))`,
		},
		{
			input:        `getPromise().property`,
			output:       `getPromise().property`,
			skipPromises: true,
		},
		{
			input:        `getPromise().object.foo`,
			output:       `getPromise().object.foo`,
			skipPromises: true,
		},
		{
			input:  `getPromise(resource.id).property`,
			output: `__apply(__apply(resource.id,eval(id, getPromise(id))), eval(getPromise, getPromise.property))`,
		},
	}

	resourceType := model.NewObjectType(map[string]model.Type{
		"id": model.NewOutputType(model.StringType),
		"foo": model.NewOutputType(model.NewObjectType(map[string]model.Type{
			"bar": model.StringType,
		})),
		"baz": model.NewOutputType(model.NewListType(model.StringType)),
	})

	scope := model.NewRootScope(syntax.None)
	scope.Define("key", &model.Variable{
		Name:         "key",
		VariableType: model.StringType,
	})
	scope.Define("resource", &model.Variable{
		Name:         "resource",
		VariableType: resourceType,
	})
	scope.Define("resources", &model.Variable{
		Name:         "resources",
		VariableType: model.NewListType(resourceType),
	})
	scope.Define("resourcesPromise", &model.Variable{
		Name:         "resourcesPromise",
		VariableType: model.NewPromiseType(model.NewListType(resourceType)),
	})
	scope.Define("resourcesOutput", &model.Variable{
		Name:         "resourcesOutput",
		VariableType: model.NewOutputType(model.NewListType(resourceType)),
	})
	scope.DefineFunction("element", pulumiBuiltins["element"])
	scope.DefineFunction("toJSON", pulumiBuiltins["toJSON"])
	scope.DefineFunction("getPromise", model.NewFunction(model.StaticFunctionSignature{
		Parameters: []model.Parameter{{
			Name: "p",
			Type: model.NewOptionalType(model.StringType),
		}},
		ReturnType: model.NewPromiseType(model.NewObjectType(map[string]model.Type{
			"property": model.StringType,
			"object": model.NewObjectType(map[string]model.Type{
				"foo": model.StringType,
			}),
		})),
	}))

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			expr, diags := model.BindExpressionText(c.input, scope, hcl.Pos{})
			assert.Len(t, diags, 0)

			expr, diags = RewriteApplies(expr, nameInfo(0), !c.skipPromises)
			assert.Len(t, diags, 0)

			assert.Equal(t, c.output, fmt.Sprintf("%v", expr))
		})
	}
}
