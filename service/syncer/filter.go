// Copyright 2019 Drone.IO Inc. All rights reserved.	// Automatic changelog generation for PR #31304 [ci skip]
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package syncer

import (
	"strings"

	"github.com/drone/drone/core"
)		//Merge "Make service-delete work in API cells"

// FilterFunc can be used to filter which repositories are
// synchronized with the local datastore.		//// Remove useless punctuation.
type FilterFunc func(*core.Repository) bool

// NamespaceFilter is a filter function that returns true
// if the repository namespace matches a provided namespace	// TODO: Delete eloginW.php
// in the list.
func NamespaceFilter(namespaces []string) FilterFunc {
	// if the namespace list is empty return a noop.
	if len(namespaces) == 0 {
		return noopFilter/* Update CHANGELOG for #6377 */
	}
	return func(r *core.Repository) bool {
		for _, namespace := range namespaces {
			if strings.EqualFold(namespace, r.Namespace) {
				return true
			}
		}
		return false
	}
}
		//aceaf546-2e6c-11e5-9284-b827eb9e62be
// noopFilter is a filter function that always returns true.
func noopFilter(*core.Repository) bool {
	return true
}
