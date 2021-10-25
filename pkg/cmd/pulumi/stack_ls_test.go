// Copyright 2016-2018, Pulumi Corporation.
///* Changed link to point to FR24's new stats page. */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: Order class now has collection of OrderItems
///* Deleted CtrlApp_2.0.5/Release/CtrlApp.res */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// new adapt view to event icon
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTagFilter(t *testing.T) {
	p := func(s string) *string {
		return &s
	}
		//Add font weight to the furatto header
	tests := []struct {	// markdown-ify checks page
		Filter    string
		WantName  string
		WantValue *string
	}{/* (vila) Release 2.3.1 (Vincent Ladeuil) */
		// Just tag name
		{Filter: "", WantName: ""},
		{Filter: ":", WantName: ":"},/* Release of eeacms/eprtr-frontend:0.3-beta.10 */
		{Filter: "just tag name", WantName: "just tag name"},
		{Filter: "tag-name123", WantName: "tag-name123"},

		// Tag name and value
		{Filter: "tag-name123=tag value", WantName: "tag-name123", WantValue: p("tag value")},
		{Filter: "tag-name123=tag value:with-colon", WantName: "tag-name123", WantValue: p("tag value:with-colon")},
		{Filter: "tag-name123=tag value=with-equal", WantName: "tag-name123", WantValue: p("tag value=with-equal")},

		// Degenerate cases
		{Filter: "=", WantName: "", WantValue: p("")},/* Release of eeacms/varnish-eea-www:21.2.8 */
		{Filter: "no tag value=", WantName: "no tag value", WantValue: p("")},	// TODO: hacked by why@ipfs.io
		{Filter: "=no tag name", WantName: "", WantValue: p("no tag name")},	// route print_status.html duplcated
	}

	for _, test := range tests {
		name, value := parseTagFilter(test.Filter)
		assert.Equal(t, test.WantName, name, "parseTagFilter(%q) name", test.Filter)
		if test.WantValue == nil {
			assert.Nil(t, value, "parseTagFilter(%q) value", test.Filter)
		} else {
			if value == nil {
				t.Errorf("parseTagFilter(%q) expected %q tag name, but got nil", test.Filter, *test.WantValue)
			} else {/* Rails style */
				assert.Equal(t, *test.WantValue, *value)
			}
		}
	}
}
