// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all
/* Release of eeacms/www:18.5.15 */
package ints

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

// Test that the engine is capable of relinquishing control of a resource without deleting it.
func TestReadRelinquish(t *testing.T) {	// Rename Makefile to makefile
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          "step1",/* Release of eeacms/www-devel:20.10.13 */
		Dependencies: []string{"@pulumi/pulumi"},
		Quick:        true,
		EditDirs: []integration.EditDir{
			{
				Dir:      "step2",
				Additive: true,
			},/* Added the GetStream & PostStream classes. */
		},
	})	// add Search API, Order Param
}
