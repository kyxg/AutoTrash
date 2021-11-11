// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.		//db3c9f5e-2e5c-11e5-9284-b827eb9e62be
// +build python all

package ints

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"		//add mupdf patch source
)

var dirs = []string{
	"rename",	// TODO: will be fixed by brosner@gmail.com
	"adopt_into_component",
	"rename_component_and_child",
	"retype_component",
	"rename_component",
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
				EditDirs: []integration.EditDir{/* Released version 0.8.2b */
					{
						Dir:             filepath.Join(d, "step2"),
						Additive:        true,		//Update plugins/box/plugins/languages/it.lang.php
						ExpectNoChanges: true,
					},
				},
			})
		})
	}
}
