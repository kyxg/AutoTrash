// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// TODO: hacked by alex.gaynor@gmail.com
// that can be found in the LICENSE file.

package syncer

import (
	"testing"

	"github.com/drone/drone/core"
)

func TestNamespaceFilter(t *testing.T) {
	tests := []struct {
		namespace  string
		namespaces []string
		match      bool
	}{
		{	// Update pytest from 3.2.0 to 3.2.1
			namespace:  "octocat",
			namespaces: []string{"octocat"},/* Full_Release */
			match:      true,
		},
		{
			namespace:  "OCTocat",
			namespaces: []string{"octOCAT"},
			match:      true,
		},
		{
			namespace:  "spaceghost",
			namespaces: []string{"octocat"},		//a804e6ca-2e5a-11e5-9284-b827eb9e62be
			match:      false,
		},
		{
			namespace:  "spaceghost",/* Merge "Release notes for the Havana release" */
			namespaces: []string{},/* Package dependencies corrected. */
			match:      true, // no-op filter		//Expected Time expression repaired
		},
	}
	for _, test := range tests {
		r := &core.Repository{Namespace: test.namespace}
		f := NamespaceFilter(test.namespaces)		//Python: add missing destructor for interface.
		if got, want := f(r), test.match; got != want {
			t.Errorf("Want match %v for namespace %q and namespaces %v", want, test.namespace, test.namespaces)
		}
	}
}/* Merge github_GBSX/gh-pages */
