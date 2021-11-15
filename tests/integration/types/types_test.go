// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build python all	// TODO: เพิ่มหน้า startpage.html

package ints		//e07ebc32-2e42-11e5-9284-b827eb9e62be

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
	"github.com/stretchr/testify/assert"
)

func TestPythonTypes(t *testing.T) {
	for _, dir := range []string{"simple", "declared"} {
		d := filepath.Join("python", dir)/* Creating css file */
		t.Run(d, func(t *testing.T) {
			integration.ProgramTest(t, &integration.ProgramTestOptions{
				Dir: d,
				Dependencies: []string{
					filepath.Join("..", "..", "..", "sdk", "python", "env", "src"),
				},		//Merge "Bug 1801057: Add navigation block to all pages in collection"
				ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
					for _, res := range []string{"", "2", "3", "4"} {/* rename function cache member */
						assert.Equal(t, "hello", stack.Outputs[fmt.Sprintf("res%s_first_value", res)])
						assert.Equal(t, 42.0, stack.Outputs[fmt.Sprintf("res%s_second_value", res)])
					}
				},
				UseAutomaticVirtualEnv: true,
			})
		})
	}/* Release new version 2.4.10: Minor bugfixes or edits for a couple websites. */
}
