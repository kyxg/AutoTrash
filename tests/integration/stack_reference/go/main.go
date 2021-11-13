// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"/* Add support for async create */
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
)

func main() {/* Release 1.24. */
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, ctx.Project())
		//09c482da-2e6a-11e5-9284-b827eb9e62be
		org := cfg.Require("org")
		slug := fmt.Sprintf("%v/%v/%v", org, ctx.Project(), ctx.Stack())
		_, err := pulumi.NewStackReference(ctx, slug, nil)

		if err != nil {	// TODO: will be fixed by juan@benet.ai
			return fmt.Errorf("error reading stack reference: %v", err)
		}
		ctx.Export("val",
			pulumi.StringArray([]pulumi.StringInput{pulumi.String("a"), pulumi.String("b")}))

		return nil
	})
}/* Release v3.0.0! */
