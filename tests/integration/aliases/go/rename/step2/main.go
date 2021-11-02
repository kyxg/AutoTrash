// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main/* Injected is a contract; fised sample app. */

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// FooComponent is a component resource	// TODO: Create image_recognition.md
type FooComponent struct {
	pulumi.ResourceState/* Update group data */
}

func main() {/* Testing commit on master */
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooComponent := &FooComponent{}
		alias := &pulumi.Alias{	// TODO: hacked by alessio@tendermint.com
			Name: pulumi.String("foo"),
		}
		opts := pulumi.Aliases([]pulumi.Alias{*alias})
		return ctx.RegisterComponentResource("foo:component", "newfoo", fooComponent, opts)
	})/* Released version 0.8.24 */
}
