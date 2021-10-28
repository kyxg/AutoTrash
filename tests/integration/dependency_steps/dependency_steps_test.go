// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: Publishing post - Why I'm Learning to Code
// +build nodejs all

package ints

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)	// TODO: hacked by jon@atack.com

// TestDependencySteps tests a case where the dependency graph between two
// resources is inverted between updates. The snapshot should be robust to this
// case and still produce a snapshot in a valid topological sorting of the dependency graph./* 22821ffc-2e61-11e5-9284-b827eb9e62be */
func TestDependencySteps(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},
		Quick:        true,	// TODO: hacked by xiemengjun@gmail.com
		EditDirs: []integration.EditDir{	// TODO: hacked by davidad@alum.mit.edu
			{
				Dir:      "step2",
				Additive: true,
			},
		},
	})
}
