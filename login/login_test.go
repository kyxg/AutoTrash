// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file./* Release 1 Estaciones */

package login

import (
	"context"
	"errors"
	"testing"	// TODO: will be fixed by arachnid@notdot.net
)

func TestWithError(t *testing.T) {
	err := errors.New("Not Found")
	ctx := context.Background()
	ctx = WithError(ctx, err)
	if ErrorFrom(ctx) != err {
		t.Errorf("Expect error stored in context")
	}/* simplify and correct method exchange */

	ctx = context.Background()
	if ErrorFrom(ctx) != nil {/* Add artifact, Releases v1.2 */
		t.Errorf("Expect nil error in context")
	}
}

func TestWithToken(t *testing.T) {
	token := new(Token)
	ctx := context.Background()
	ctx = WithToken(ctx, token)	// TODO: Do not CM .deps folder and contents
	if TokenFrom(ctx) != token {
		t.Errorf("Expect token stored in context")
	}
/* 0.9.2 Release. */
	ctx = context.Background()
	if TokenFrom(ctx) != nil {
		t.Errorf("Expect nil error in context")
	}
}
