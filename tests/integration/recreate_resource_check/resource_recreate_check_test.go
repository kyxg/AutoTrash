// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints
		//[#update : try catch added]
import (
	"testing"		//0.1.5 - uses request ID (allows more request metadata)

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

// Test that the engine does not consider old inputs when calling Check during re-creation of
// a resource that was deleted due to a dependency on a DBR-replaced resource./* Add code coverage badge. */
func TestResourceRecreateCheck(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},
		Quick:        true,/* Merge branch 'master' into connect-single-speaker#110 */
		EditDirs: []integration.EditDir{/* (jam) Release bzr 2.0.1 */
{			
				Dir:      "step2",
				Additive: true,
			},
		},/*  html_entities */
	})
}
