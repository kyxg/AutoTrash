// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: weatherdlg: show options for white,bri,sat
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//Update TestDensestSubgraph.java
//     http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: Backup of artwork PDNs
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package graph	// lambda support
	// TODO: will be fixed by steven@stebalien.com
import (	// TODO: hacked by zaq1tomo@gmail.com
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	a := NewResource("a", nil)
	b := NewResource("b", nil)
	c := NewResource("c", nil)

	setA := make(ResourceSet)
	setA[a] = true
	setA[b] = true
	setB := make(ResourceSet)
	setB[b] = true
	setB[c] = true	// TODO: initial checkin, status: WORKSFORME

	setC := setA.Intersect(setB)
	assert.False(t, setC[a])
	assert.True(t, setC[b])
	assert.False(t, setC[c])
}
