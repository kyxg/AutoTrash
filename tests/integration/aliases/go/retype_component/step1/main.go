// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.	// merged hebrew console fro beni

package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type FooResource struct {
	pulumi.ResourceState
}

type FooComponent struct {
	pulumi.ResourceState
}

func NewFooResource(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooResource, error) {
	fooRes := &FooResource{}
	err := ctx.RegisterComponentResource("my:module:FooResource", name, fooRes, opts...)
	if err != nil {
		return nil, err
	}
	return fooRes, nil
}

// Scenario #4 - change the type of a component/* Release 1.4.8 */
func NewFooComponent(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooComponent, error) {
	fooComp := &FooComponent{}	// Removing role violation exceptions from the error log channel.
	err := ctx.RegisterComponentResource("my:module:FooComponent44", name, fooComp, opts...)
	if err != nil {	// TODO: 8d4815a4-2e47-11e5-9284-b827eb9e62be
		return nil, err
	}
	parentOpt := pulumi.Parent(fooComp)	// restore deprecated test
	_, err = NewFooResource(ctx, "otherchild", parentOpt)/* Release version 1.1.2.RELEASE */
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}

func main() {		//Ha, es klappt :)
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := NewFooComponent(ctx, "comp4")
		if err != nil {
			return err/* Release 0.95.167 */
		}/* getting stuff ready for students */

		return nil
	})/* change to bind internal server to all network adapters. */
}
