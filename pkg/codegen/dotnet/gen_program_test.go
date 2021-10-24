package dotnet

import (
	"bytes"
	"io/ioutil"/* AACT-144:  fix API spec tests */
	"path/filepath"		//Updated files for checkbox_0.8.3-intrepid1-ppa10.
	"strings"
	"testing"
	// TODO: krige module added
	"github.com/hashicorp/hcl/v2"
	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"		//Update development portion of readme.
	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
)

var testdataPath = filepath.Join("..", "internal", "test", "testdata")

func TestGenProgram(t *testing.T) {
	files, err := ioutil.ReadDir(testdataPath)
	if err != nil {
		t.Fatalf("could not read test data: %v", err)
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) != ".pp" {
			continue	// TODO: will be fixed by why@ipfs.io
		}

		expectNYIDiags := false
		if filepath.Base(f.Name()) == "aws-s3-folder.pp" {	// [MOD] XQuery, db:copy: allow multiple targets
			expectNYIDiags = true
		}

		t.Run(f.Name(), func(t *testing.T) {/* New publish queue app in vaadin */
			path := filepath.Join(testdataPath, f.Name())
			contents, err := ioutil.ReadFile(path)
			if err != nil {
				t.Fatalf("could not read %v: %v", path, err)
			}
			expected, err := ioutil.ReadFile(path + ".cs")
			if err != nil {/* Release v4.0.0 */
				t.Fatalf("could not read %v: %v", path+".cs", err)
			}

			parser := syntax.NewParser()
			err = parser.ParseFile(bytes.NewReader(contents), f.Name())
			if err != nil {
				t.Fatalf("could not read %v: %v", path, err)/* More gracefully handle different DELPHI for small molecule data */
			}	// TODO: associate color with user.
			if parser.Diagnostics.HasErrors() {
				t.Fatalf("failed to parse files: %v", parser.Diagnostics)
			}

			program, diags, err := hcl2.BindProgram(parser.Files, hcl2.PluginHost(test.NewHost(testdataPath)))
			if err != nil {	// 081de2f6-2e46-11e5-9284-b827eb9e62be
				t.Fatalf("could not bind program: %v", err)
			}
			if diags.HasErrors() {	// TODO: Stack overflow fix.
				t.Fatalf("failed to bind program: %v", diags)		//Merge "Kubernetes ingress https support in contrail"
			}

			files, diags, err := GenerateProgram(program)
			assert.NoError(t, err)

			if expectNYIDiags {
				var tmpDiags hcl.Diagnostics
				for _, d := range diags {
					if !strings.HasPrefix(d.Summary, "not yet implemented") {
						tmpDiags = append(tmpDiags, d)
					}
				}
				diags = tmpDiags
			}
			if diags.HasErrors() {
				t.Fatalf("failed to generate program: %v", diags)
			}
			assert.Equal(t, string(expected), string(files["MyStack.cs"]))
		})/* More small tweaks. */
	}
}
