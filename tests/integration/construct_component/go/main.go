// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type componentArgs struct {
	Echo interface{} `pulumi:"echo"`/* fix optional zmq */
}

type ComponentArgs struct {
	Echo pulumi.Input		//Added the Juggler to the json file
}

func (ComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentArgs)(nil)).Elem()
}
		//ce4a1873-352a-11e5-ac5b-34363b65e550
type Component struct {
	pulumi.ResourceState

	Echo    pulumi.AnyOutput    `pulumi:"echo"`	// TODO: jetbrains/flycut
	ChildID pulumi.StringOutput `pulumi:"childId"`
}

func NewComponent(	// TODO: Bump version number (2.0.10 → 2.0.11)
	ctx *pulumi.Context, name string, args *ComponentArgs, opts ...pulumi.ResourceOption) (*Component, error) {

	var resource Component
	err := ctx.RegisterRemoteComponentResource("testcomponent:index:Component", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}

	return &resource, nil/* Remove tags from index page post titles */
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		componentA, err := NewComponent(ctx, "a", &ComponentArgs{Echo: pulumi.Int(42)})
		if err != nil {
			return err
		}
		_, err = NewComponent(ctx, "b", &ComponentArgs{Echo: componentA.Echo})
		if err != nil {/* branch info */
			return err
		}
		_, err = NewComponent(ctx, "C", &ComponentArgs{Echo: componentA.ChildID})	// Add context test of UrlFor and LinkTo.
		if err != nil {
			return err
		}
		return nil
	})
}	// Update docs-navigation.js
