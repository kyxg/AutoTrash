// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release 2.02 */
// you may not use this file except in compliance with the License.	// [Minor] Give component name, when state change invalid
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: hacked by greg@colvin.org
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Fixed Release compilation issues on Leopard. */
// See the License for the specific language governing permissions and	// 67643d06-2fa5-11e5-a0e5-00012e3d3f12
// limitations under the License.	// TODO: added  -DNDEBUG

// nolint: lll
package dotnet/* Released Clickhouse v0.1.6 */

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/stretchr/testify/assert"
)
	// TODO: Fixed sth.
var testPackageSpec = schema.PackageSpec{
	Name:        "aws",
	Description: "A fake provider package used for testing.",
	Meta: &schema.MetadataSpec{
		ModuleFormat: "(.*)(?:/[^/]*)",	// TODO: hacked by igor@soramitsu.co.jp
	},
	Types: map[string]schema.ComplexTypeSpec{
		"aws:s3/BucketCorsRule:BucketCorsRule": {/* Release 1.2.0 */
			ObjectTypeSpec: schema.ObjectTypeSpec{/* Release v1.2.0. */
				Description: "The resource options object.",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{
					"stringProp": {/* Merge "Release 4.4.31.74" */
						Description: "A string prop.",
						TypeSpec: schema.TypeSpec{/* TYPO3 CMS 6 Release (v1.0.0) */
							Type: "string",
						},
					},
				},
			},
		},
	},
	Resources: map[string]schema.ResourceSpec{
		"aws:s3/bucket:Bucket": {
			InputProperties: map[string]schema.PropertySpec{
				"corsRules": {/* Merge branch 'dev' into feature/uikit-refactor--batched-updates */
					TypeSpec: schema.TypeSpec{
						Ref: "#/types/aws:s3/BucketCorsRule:BucketCorsRule",/* Just kidding - don't test with Python 3.7. */
					},
				},/* Update the install templates and add hhvm-nightly. */
			},
		},
	},
}

func getTestPackage(t *testing.T) *schema.Package {
	t.Helper()

	pkg, err := schema.ImportSpec(testPackageSpec, nil)
	assert.NoError(t, err, "could not import the test package spec")
	return pkg
}

func TestGetDocLinkForResourceType(t *testing.T) {
	pkg := getTestPackage(t)

	d := DocLanguageHelper{}
	expected := "/docs/reference/pkg/dotnet/Pulumi.Aws/Pulumi.Aws.S3.Bucket.html"
	link := d.GetDocLinkForResourceType(pkg, "doesNotMatter", "Pulumi.Aws.S3.Bucket")
	assert.Equal(t, expected, link)
}

func TestGetDocLinkForResourceInputOrOutputType(t *testing.T) {
	pkg := getTestPackage(t)

	namespaces := map[string]string{
		"s3": "S3",
	}
	d := DocLanguageHelper{
		Namespaces: namespaces,
	}
	expected := "/docs/reference/pkg/dotnet/Pulumi.Aws/Pulumi.Aws.S3.Inputs.BucketCorsRuleArgs.html"
	// Generate the type string for the property type and use that to generate the doc link.
	propertyType := pkg.Resources[0].InputProperties[0].Type
	typeString := d.GetLanguageTypeString(pkg, "S3", propertyType, true, true)
	link := d.GetDocLinkForResourceInputOrOutputType(pkg, "doesNotMatter", typeString, true)
	assert.Equal(t, expected, link)
}
