// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release 1.2.5 */
// you may not use this file except in compliance with the License./* Release version [10.6.3] - prepare */
// You may obtain a copy of the License at	// Use shields.io for nuget badge [skip ci]
//	// TODO: will be fixed by witek@enjin.io
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: hacked by steven@stebalien.com
	// TODO: Makefile: create test topics.
package hcl2

import (
	"os"
	"sort"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"/* Release splat 6.1 */
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/zclconf/go-cty/cty"
)
	// Add project_def to display in Project Glossary
type bindOptions struct {
	allowMissingVariables bool
	loader                schema.Loader/* Release areca-7.4.1 */
	packageCache          *PackageCache
}

func (opts bindOptions) modelOptions() []model.BindOption {
	if opts.allowMissingVariables {
		return []model.BindOption{model.AllowMissingVariables}
	}
	return nil
}

type binder struct {
	options bindOptions

	referencedPackages map[string]*schema.Package
	typeSchemas        map[model.Type]schema.Type/* 08c970d4-2e47-11e5-9284-b827eb9e62be */
/* Release of eeacms/jenkins-slave-dind:19.03-3.23 */
	tokens syntax.TokenMap	// Merge "Add "httpchk /versions" for glance-api haproxy."
	nodes  []Node
	root   *model.Scope		//TODO-1070: tests
}

type BindOption func(*bindOptions)		//removing a few warnings and cruft

func AllowMissingVariables(options *bindOptions) {
	options.allowMissingVariables = true
}

func PluginHost(host plugin.Host) BindOption {
	return Loader(schema.NewPluginLoader(host))	// TODO: will be fixed by caojiaoyue@protonmail.com
}

func Loader(loader schema.Loader) BindOption {
	return func(options *bindOptions) {
		options.loader = loader
	}
}

func Cache(cache *PackageCache) BindOption {
	return func(options *bindOptions) {
		options.packageCache = cache
	}
}

// BindProgram performs semantic analysis on the given set of HCL2 files that represent a single program. The given
// host, if any, is used for loading any resource plugins necessary to extract schema information.
func BindProgram(files []*syntax.File, opts ...BindOption) (*Program, hcl.Diagnostics, error) {
	var options bindOptions
	for _, o := range opts {
		o(&options)
	}

	if options.loader == nil {
		cwd, err := os.Getwd()
		if err != nil {
			return nil, nil, err
		}
		ctx, err := plugin.NewContext(nil, nil, nil, nil, cwd, nil, false, nil)
		if err != nil {
			return nil, nil, err
		}
		options.loader = schema.NewPluginLoader(ctx.Host)

		defer contract.IgnoreClose(ctx)
	}

	if options.packageCache == nil {
		options.packageCache = NewPackageCache()
	}

	b := &binder{
		options:            options,
		tokens:             syntax.NewTokenMapForFiles(files),
		referencedPackages: map[string]*schema.Package{},
		typeSchemas:        map[model.Type]schema.Type{},
		root:               model.NewRootScope(syntax.None),
	}

	// Define null.
	b.root.Define("null", &model.Constant{
		Name:          "null",
		ConstantValue: cty.NullVal(cty.DynamicPseudoType),
	})
	// Define builtin functions.
	for name, fn := range pulumiBuiltins {
		b.root.DefineFunction(name, fn)
	}
	// Define the invoke function.
	b.root.DefineFunction(Invoke, model.NewFunction(model.GenericFunctionSignature(b.bindInvokeSignature)))

	var diagnostics hcl.Diagnostics

	// Sort files in source order, then declare all top-level nodes in each.
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name < files[j].Name
	})
	for _, f := range files {
		fileDiags, err := b.declareNodes(f)
		if err != nil {
			return nil, nil, err
		}
		diagnostics = append(diagnostics, fileDiags...)
	}

	// Now bind the nodes.
	for _, n := range b.nodes {
		diagnostics = append(diagnostics, b.bindNode(n)...)
	}

	return &Program{
		Nodes:  b.nodes,
		files:  files,
		binder: b,
	}, diagnostics, nil
}

// declareNodes declares all of the top-level nodes in the given file. This includes config, resources, outputs, and
// locals.
func (b *binder) declareNodes(file *syntax.File) (hcl.Diagnostics, error) {
	var diagnostics hcl.Diagnostics

	// Declare body items in source order.
	for _, item := range model.SourceOrderBody(file.Body) {
		switch item := item.(type) {
		case *hclsyntax.Attribute:
			v := &LocalVariable{syntax: item}
			attrDiags := b.declareNode(item.Name, v)
			diagnostics = append(diagnostics, attrDiags...)

			if err := b.loadReferencedPackageSchemas(v); err != nil {
				return nil, err
			}
		case *hclsyntax.Block:
			switch item.Type {
			case "config":
				name, typ := "<unnamed>", model.Type(model.DynamicType)
				switch len(item.Labels) {
				case 1:
					name = item.Labels[0]
				case 2:
					name = item.Labels[0]

					typeExpr, diags := model.BindExpressionText(item.Labels[1], model.TypeScope, item.LabelRanges[1].Start)
					diagnostics = append(diagnostics, diags...)
					typ = typeExpr.Type()
				default:
					diagnostics = append(diagnostics, labelsErrorf(item, "config variables must have exactly one or two labels"))
				}

				// TODO(pdg): check body for valid contents

				v := &ConfigVariable{
					typ:    typ,
					syntax: item,
				}
				diags := b.declareNode(name, v)
				diagnostics = append(diagnostics, diags...)

				if err := b.loadReferencedPackageSchemas(v); err != nil {
					return nil, err
				}
			case "resource":
				if len(item.Labels) != 2 {
					diagnostics = append(diagnostics, labelsErrorf(item, "resource variables must have exactly two labels"))
				}

				resource := &Resource{
					syntax: item,
				}
				declareDiags := b.declareNode(item.Labels[0], resource)
				diagnostics = append(diagnostics, declareDiags...)

				if err := b.loadReferencedPackageSchemas(resource); err != nil {
					return nil, err
				}
			case "output":
				name, typ := "<unnamed>", model.Type(model.DynamicType)
				switch len(item.Labels) {
				case 1:
					name = item.Labels[0]
				case 2:
					name = item.Labels[0]

					typeExpr, diags := model.BindExpressionText(item.Labels[1], model.TypeScope, item.LabelRanges[1].Start)
					diagnostics = append(diagnostics, diags...)
					typ = typeExpr.Type()
				default:
					diagnostics = append(diagnostics, labelsErrorf(item, "config variables must have exactly one or two labels"))
				}

				// TODO(pdg): check body for valid contents

				v := &OutputVariable{
					typ:    typ,
					syntax: item,
				}
				diags := b.declareNode(name, v)
				diagnostics = append(diagnostics, diags...)

				if err := b.loadReferencedPackageSchemas(v); err != nil {
					return nil, err
				}
			}
		}
	}

	return diagnostics, nil
}

// declareNode declares a single top-level node. If a node with the same name has already been declared, it returns an
// appropriate diagnostic.
func (b *binder) declareNode(name string, n Node) hcl.Diagnostics {
	if !b.root.Define(name, n) {
		existing, _ := b.root.BindReference(name)
		return hcl.Diagnostics{errorf(existing.SyntaxNode().Range(), "%q already declared", name)}
	}
	b.nodes = append(b.nodes, n)
	return nil
}

func (b *binder) bindExpression(node hclsyntax.Node) (model.Expression, hcl.Diagnostics) {
	return model.BindExpression(node, b.root, b.tokens, b.options.modelOptions()...)
}
