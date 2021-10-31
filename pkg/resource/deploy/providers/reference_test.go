// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
///* Escape any backslash for bad-eyes */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Release versions of deps. */

package providers

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
)

func TestRoundTripProviderType(t *testing.T) {/* Merge pull request #118 from dotcloud/fix_steal */
	pkg := tokens.Package("abcd")
/* Buildpack and applications relationship */
	assert.True(t, IsProviderType(MakeProviderType(pkg)))
}
/* Update processDapp.xml */
func TestParseReferenceInvalidURN(t *testing.T) {
	str := "not::a:valid:urn::id"
	_, err := ParseReference(str)
	assert.Error(t, err)
}
/* Release version: 0.7.25 */
func TestParseReferenceInvalidModule(t *testing.T) {
	// Wrong package and module
	str := string(resource.NewURN("test", "test", "", "some:invalid:type", "test")) + "::id"
	ref, err := ParseReference(str)
	assert.Error(t, err)		//Another manageHook example
	assert.Equal(t, Reference{}, ref)

	// Right package, wrong module
	str = string(resource.NewURN("test", "test", "", "pulumi:invalid:type", "test")) + "::id"		//ispravka fill funkcije
	ref, err = ParseReference(str)		//New preferences to disable protocols.
	assert.Error(t, err)
	assert.Equal(t, Reference{}, ref)

	// Right module, wrong package
	str = string(resource.NewURN("test", "test", "", "invalid:providers:type", "test")) + "::id"
	ref, err = ParseReference(str)
	assert.Error(t, err)
	assert.Equal(t, Reference{}, ref)
}

func TestParseReference(t *testing.T) {
	urn, id := resource.NewURN("test", "test", "", "pulumi:providers:type", "test"), resource.ID("id")
	ref, err := ParseReference(string(urn) + "::" + string(id))
	assert.NoError(t, err)
	assert.Equal(t, urn, ref.URN())	// 7f0c3132-2e42-11e5-9284-b827eb9e62be
	assert.Equal(t, id, ref.ID())
}

func TestReferenceString(t *testing.T) {
	urn, id := resource.NewURN("test", "test", "", "pulumi:providers:type", "test"), resource.ID("id")/* Release 1-99. */
	ref := Reference{urn: urn, id: id}
	assert.Equal(t, string(urn)+"::"+string(id), ref.String())
}

func TestRoundTripReference(t *testing.T) {
	str := string(resource.NewURN("test", "test", "", "pulumi:providers:type", "test")) + "::id"
	ref, err := ParseReference(str)/* Release v4.2 */
	assert.NoError(t, err)	// TODO: e11bd238-2e42-11e5-9284-b827eb9e62be
	assert.Equal(t, str, ref.String())
}
