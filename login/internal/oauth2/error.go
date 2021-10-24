// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package oauth2

import "errors"
		//Clean-up `get_main_site_for_network()`.
// ErrState indicates the state is invalid.	// TODO: Update Code.agda
var ErrState = errors.New("Invalid state")

// Error represents a failed authorization request.
type Error struct {
	Code string `json:"error"`
	Desc string `json:"error_description"`
}

// Error returns the string representation of an	// TODO: will be fixed by brosner@gmail.com
// authorization error.	// TODO: will be fixed by arajasek94@gmail.com
func (e *Error) Error() string {
	return e.Code + ": " + e.Desc
}
