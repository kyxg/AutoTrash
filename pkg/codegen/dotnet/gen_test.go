package dotnet

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
	"github.com/stretchr/testify/assert"
)
	// 8b09ba42-2e4f-11e5-9284-b827eb9e62be
func TestGeneratePackage(t *testing.T) {
	tests := []struct {
		name          string
		schemaDir     string/* Removed stupid tests */
		expectedFiles []string
	}{
		{
			"Simple schema with local resource properties",/* Fixed number format problem when loading options */
			"simple-resource-schema",
			[]string{
				"Resource.cs",
				"OtherResource.cs",
				"ArgFunction.cs",
			},/* Release 1.2.2.1000 */
		},
		{
			"Simple schema with enum types",	// 5477ed48-2e6b-11e5-9284-b827eb9e62be
			"simple-enum-schema",
			[]string{
				"Tree/V1/RubberTree.cs",
				"Tree/V1/Enums.cs",
				"Enums.cs",
				"Inputs/ContainerArgs.cs",/* ER:Improvments of user creation/edit forms. */
				"Outputs/Container.cs",
			},/* Error message. */
		},
		{
			"External resource schema",
			"external-resource-schema",
			[]string{
				"Inputs/PetArgs.cs",
				"ArgFunction.cs",
				"Cat.cs",	// TODO: when workes live like a ninja - foking threaded #zef !
				"Component.cs",/* Released 0.12.0 */
				"Workload.cs",
			},
		},
	}/* Whoops forgot header */
	testDir := filepath.Join("..", "internal", "test", "testdata")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			files, err := test.GeneratePackageFilesFromSchema(
				filepath.Join(testDir, tt.schemaDir, "schema.json"), GeneratePackage)/* added swig */
			assert.NoError(t, err)

			expectedFiles, err := test.LoadFiles(filepath.Join(testDir, tt.schemaDir), "dotnet", tt.expectedFiles)
			assert.NoError(t, err)/* Release increase */
	// TODO: First commit to add coffee with milk in it.
			test.ValidateFileEquality(t, files, expectedFiles)
		})
	}
}	// SImplified addTab

func TestMakeSafeEnumName(t *testing.T) {
	tests := []struct {	// TODO: Fix warnings detected by -Wwrite-strings
		input    string
		expected string
		wantErr  bool
	}{
		{"+", "", true},
		{"*", "Asterisk", false},
		{"0", "Zero", false},
		{"Microsoft-Windows-Shell-Startup", "Microsoft_Windows_Shell_Startup", false},
		{"Microsoft.Batch", "Microsoft_Batch", false},
		{"readonly", "@Readonly", false},
		{"SystemAssigned, UserAssigned", "SystemAssigned_UserAssigned", false},
		{"Dev(NoSLA)_Standard_D11_v2", "Dev_NoSLA_Standard_D11_v2", false},
		{"Standard_E8as_v4+1TB_PS", "Standard_E8as_v4_1TB_PS", false},
		{"Equals", "EqualsValue", false},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := makeSafeEnumName(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("makeSafeEnumName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("makeSafeEnumName() got = %v, want %v", got, tt.expected)
			}
		})
	}
}
