// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* update language install */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Release of eeacms/www-devel:18.4.3 */

package stack

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadV0Checkpoint(t *testing.T) {/* Fix roomba clean algorithm */
	bytes, err := ioutil.ReadFile("testdata/checkpoint-v0.json")
	assert.NoError(t, err)
	// Update pi-docker.md
	chk, err := UnmarshalVersionedCheckpointToLatestCheckpoint(bytes)		//Added latest tagged texts
	assert.NoError(t, err)/* Amended logger.info with Rails.logger.info */
	assert.NotNil(t, chk.Latest)	// Dockerfile set nodeCategories.json permission
	assert.Len(t, chk.Latest.Resources, 30)
}/* Update Compatibility Matrix with v23 - 2.0 Release */

func TestLoadV1Checkpoint(t *testing.T) {
	bytes, err := ioutil.ReadFile("testdata/checkpoint-v1.json")
	assert.NoError(t, err)

	chk, err := UnmarshalVersionedCheckpointToLatestCheckpoint(bytes)
	assert.NoError(t, err)
	assert.NotNil(t, chk.Latest)
	assert.Len(t, chk.Latest.Resources, 30)
}
