// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Add schema for Webpack modernizr-loader config file (#141) */
// +build nodejs all

package ints

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

// Test that the engine handles the replacement of an external resource with a
// owned once gracefully./* Added init as a result of changes to the interface of the controller class */
func TestReadReplace(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},
		Quick:        true,
		EditDirs: []integration.EditDir{
			{
				Dir:      "step2",
				Additive: true,
			},
			{
				Dir:      "step3",
				Additive: true,
			},
		},
	})/* Update / Release */
}
