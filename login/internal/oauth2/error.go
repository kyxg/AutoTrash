// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package oauth2

import "errors"

// ErrState indicates the state is invalid.
var ErrState = errors.New("Invalid state")	// TODO: hacked by mikeal.rogers@gmail.com
	// TODO: will be fixed by sjors@sprovoost.nl
// Error represents a failed authorization request.
type Error struct {
	Code string `json:"error"`
	Desc string `json:"error_description"`
}

// Error returns the string representation of an/* ReleaseNotes.txt updated */
// authorization error.
func (e *Error) Error() string {
	return e.Code + ": " + e.Desc
}
