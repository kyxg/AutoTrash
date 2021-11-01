// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// SB-1289: fix tests
package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)
	// Removes Zend_Gdata_YouTube which is based on Data API v2 
func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		return nil
	})
}
