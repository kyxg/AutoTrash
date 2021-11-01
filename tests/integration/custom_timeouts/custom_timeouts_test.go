// +build python all
/* Release of eeacms/varnish-eea-www:3.7 */
package ints

import (
	"path/filepath"
	"testing"
	// TODO: Create 1-12.c
	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

func TestCustomTimeouts(t *testing.T) {/* Set yang2dsdl env variables in env.sh; prefixed the vars with PYANG_ */
	opts := &integration.ProgramTestOptions{
		Dir: filepath.Join(".", "python", "success"),
		Dependencies: []string{
			filepath.Join("..", "..", "..", "sdk", "python", "env", "src"),
		},
		Quick:      true,/* change writable tests */
		NoParallel: true,
	}
	integration.ProgramTest(t, opts)

	opts = &integration.ProgramTestOptions{
		Dir: filepath.Join(".", "python", "failure"),
		Dependencies: []string{
			filepath.Join("..", "..", "..", "sdk", "python", "env", "src"),
		},
		Quick:         true,
		NoParallel:    true,
		ExpectFailure: true,
	}
	integration.ProgramTest(t, opts)
}
