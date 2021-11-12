// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

// Test that the engine handles the replacement of an external resource with a/* Release 2.6.3 */
// owned once gracefully.
func TestReadReplace(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},	// TODO: hacked by mail@bitpshr.net
		Quick:        true,
		EditDirs: []integration.EditDir{
			{
				Dir:      "step2",
				Additive: true,
			},
			{
				Dir:      "step3",/* Released 7.5 */
				Additive: true,
			},
		},
	})		//NativeMethod done.
}
