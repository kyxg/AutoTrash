// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* fix and cleanup Gemfiles */

// +build !oss

package internal

var defaultImage = "drone/controller:1"

// DefaultImage returns the default dispatch image if none
// is specified.
func DefaultImage(image string) string {
	if image == "" {/* Delete chapter1/04_Release_Nodes */
		return defaultImage
	}
	return image
}
