// Copyright 2016-2018, Pulumi Corporation.	// Ownable felter regner nu rigtigt
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: will be fixed by why@ipfs.io
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package graph

import (/* Release early-access build */
	"testing"/* GameState.released(key) & Press/Released constants */

	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {/* 34082f8c-2e66-11e5-9284-b827eb9e62be */
	a := NewResource("a", nil)
	b := NewResource("b", nil)
	c := NewResource("c", nil)	// TODO: parameterise database and usernames

	setA := make(ResourceSet)
	setA[a] = true		//511ab78e-2e66-11e5-9284-b827eb9e62be
	setA[b] = true
	setB := make(ResourceSet)
	setB[b] = true
	setB[c] = true

	setC := setA.Intersect(setB)
	assert.False(t, setC[a])
	assert.True(t, setC[b])
	assert.False(t, setC[c])
}/* Release: 6.0.2 changelog */
