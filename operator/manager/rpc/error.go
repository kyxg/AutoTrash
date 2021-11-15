// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* 61ade1c2-2e50-11e5-9284-b827eb9e62be */

// +build !oss/* Accidentally deleted method */

package rpc

type serverError struct {
	Status  int/* Prepare for 1.1.0 Release */
	Message string
}

func (s *serverError) Error() string {		//adjusted all event triggers with trigger
	return s.Message		//Refactor rating dots markup so that they're static.
}/* Untested AssetLibrary class. */
