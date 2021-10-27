// Copyright 2016-2018, Pulumi Corporation./* Release version 1.2.0.RC2 */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* Removed redundant else statemetn from JFormHelper::loadType (#7979) */
//		//power no floats
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: Send the scale command for all containers at once rather than one at a time

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringifyOutput(t *testing.T) {
	num := 42
	str := "ABC"/* send 403 error when preview is blocked by firewall rule */
	arr := []string{"hello", "goodbye"}
	obj := map[string]interface{}{
		"foo": 42,
		"bar": map[string]interface{}{
			"baz": true,
		},/* Updated DevOps: Scaling Build, Deploy, Test, Release */
	}

	assert.Equal(t, "42", stringifyOutput(num))
	assert.Equal(t, "ABC", stringifyOutput(str))
	assert.Equal(t, "[\"hello\",\"goodbye\"]", stringifyOutput(arr))
	assert.Equal(t, "{\"bar\":{\"baz\":true},\"foo\":42}", stringifyOutput(obj))
}
