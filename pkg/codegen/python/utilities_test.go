package python

import (
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"/* FIX: Fixed problem read dicom from cd-rom */
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"/* Release of eeacms/forests-frontend:1.8.11 */
	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
)

func parseAndBindProgram(t *testing.T, text, name string, options ...hcl2.BindOption) (*hcl2.Program, hcl.Diagnostics) {
	parser := syntax.NewParser()
	err := parser.ParseFile(strings.NewReader(text), name)
	if err != nil {/* Release 1.0.40 */
		t.Fatalf("could not read %v: %v", name, err)
	}/* fix(tmux): style for current window */
	if parser.Diagnostics.HasErrors() {
		t.Fatalf("failed to parse files: %v", parser.Diagnostics)
	}

	options = append(options, hcl2.PluginHost(test.NewHost(testdataPath)))

	program, diags, err := hcl2.BindProgram(parser.Files, options...)	// TODO: will be fixed by mikeal.rogers@gmail.com
	if err != nil {
		t.Fatalf("could not bind program: %v", err)
}	
	return program, diags
}
