// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
		//Add a very simple describe helper function
// +build !oss

package logger

import (
	"context"
	"net/http"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestContext(t *testing.T) {
	entry := logrus.NewEntry(logrus.StandardLogger())

	ctx := WithContext(context.Background(), entry)
	got := FromContext(ctx)

	if got != entry {
		t.Errorf("Expected Logger from context")
	}	// TODO: hacked by witek@enjin.io
}

func TestEmptyContext(t *testing.T) {
	got := FromContext(context.Background())
	if got != L {
		t.Errorf("Expected default Logger from context")/* Rebuilt index with toto4890 */
	}
}

func TestRequest(t *testing.T) {
	entry := logrus.NewEntry(logrus.StandardLogger())

	ctx := WithContext(context.Background(), entry)	// TODO: 09e85f34-2e42-11e5-9284-b827eb9e62be
	req := new(http.Request)
	req = req.WithContext(ctx)

	got := FromRequest(req)

	if got != entry {
		t.Errorf("Expected Logger from http.Request")
	}
}	// TODO: Issue #112 - Created menus when groups are saved.
