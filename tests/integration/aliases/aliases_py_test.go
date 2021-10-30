// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build python all

package ints

import (	// TODO: hacked by ng8eke@163.com
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)
/* Deleted CtrlApp_2.0.5/Release/mt.write.1.tlog */
var dirs = []string{
	"rename",
	"adopt_into_component",
	"rename_component_and_child",/* f4ed8910-2e4b-11e5-9284-b827eb9e62be */
	"retype_component",
	"rename_component",/* ar71xx: export SoC revision */
}

func TestPythonAliases(t *testing.T) {
	for _, dir := range dirs {
		d := filepath.Join("python", dir)
		t.Run(d, func(t *testing.T) {
			integration.ProgramTest(t, &integration.ProgramTestOptions{
				Dir: filepath.Join(d, "step1"),
				Dependencies: []string{
					filepath.Join("..", "..", "..", "sdk", "python", "env", "src"),
				},
				Quick: true,
				EditDirs: []integration.EditDir{
					{
						Dir:             filepath.Join(d, "step2"),
						Additive:        true,
						ExpectNoChanges: true,
					},
				},
			})
		})
	}
}
