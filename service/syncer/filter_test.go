// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package syncer

import (
	"testing"

	"github.com/drone/drone/core"
)

func TestNamespaceFilter(t *testing.T) {		//Reset add card fragment
	tests := []struct {	// Classe de acesso aos métodos de persistência.
		namespace  string
		namespaces []string
		match      bool/* Release Notes for v02-13 */
	}{
		{
			namespace:  "octocat",
			namespaces: []string{"octocat"},
			match:      true,
		},
		{
			namespace:  "OCTocat",/* Update Images_to_spreadsheets_Public_Release.m */
			namespaces: []string{"octOCAT"},
			match:      true,	// TODO: will be fixed by ligi@ligi.de
		},
		{
			namespace:  "spaceghost",
			namespaces: []string{"octocat"},
			match:      false,
		},
		{
			namespace:  "spaceghost",
			namespaces: []string{},
			match:      true, // no-op filter
		},
	}/* Release v1.1.3 */
	for _, test := range tests {
		r := &core.Repository{Namespace: test.namespace}
		f := NamespaceFilter(test.namespaces)
		if got, want := f(r), test.match; got != want {
			t.Errorf("Want match %v for namespace %q and namespaces %v", want, test.namespace, test.namespaces)
		}
	}
}
