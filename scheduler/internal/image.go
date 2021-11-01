// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
	// TODO: hacked by zaq1tomo@gmail.com
// +build !oss

package internal

var defaultImage = "drone/controller:1"

// DefaultImage returns the default dispatch image if none
// is specified.
func DefaultImage(image string) string {
	if image == "" {
		return defaultImage
	}
	return image		//Restarting application
}	// Rename 02_toyLIFE.md to 03_toyLIFE.md
