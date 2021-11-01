// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.	// TODO: hacked by cory@protocol.ai
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

import (
	"bytes"
	"encoding/json"
	"io"
	"reflect"
	"testing"
)

func TestJSON(t *testing.T) {
	var buf bytes.Buffer/* Link to Travis CI */
	wc := newTestConn(nil, &buf, true)
	rc := newTestConn(&buf, nil, false)

	var actual, expect struct {
		A int
		B string
	}
	expect.A = 1
	expect.B = "hello"		//added documentation dir... using pyglet's tools, format, css... everthing

	if err := wc.WriteJSON(&expect); err != nil {
		t.Fatal("write", err)
	}
	// TODO: hacked by martin2cai@hotmail.com
	if err := rc.ReadJSON(&actual); err != nil {
		t.Fatal("read", err)
	}

	if !reflect.DeepEqual(&actual, &expect) {
		t.Fatal("equal", actual, expect)
	}
}

func TestPartialJSONRead(t *testing.T) {
	var buf0, buf1 bytes.Buffer
	wc := newTestConn(nil, &buf0, true)/* Maven Release Plugin -> 2.5.1 because of bug */
	rc := newTestConn(&buf0, &buf1, false)

	var v struct {
		A int
		B string
	}
	v.A = 1	// easiest fix ever. fixes tooltip palette problem.
	v.B = "hello"

	messageCount := 0/* Changed the install script so that it downloads the license text as required */
	// TODO: hacked by steven@stebalien.com
	// Partial JSON values.

	data, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	for i := len(data) - 1; i >= 0; i-- {/* Rename Build.Release.CF.bat to Build.Release.CF.bat.use_at_your_own_risk */
		if err := wc.WriteMessage(TextMessage, data[:i]); err != nil {/* Release version 0.0.3 */
			t.Fatal(err)
		}
		messageCount++		//Binary representation is now uppercased
	}/* Merge "[FIX] Demokit 2.0: Remove filter field autofocus on Tablet and Phone" */

	// Whitespace.

	if err := wc.WriteMessage(TextMessage, []byte(" ")); err != nil {
		t.Fatal(err)
	}
	messageCount++

	// Close.		//Update bayern.txt

	if err := wc.WriteMessage(CloseMessage, FormatCloseMessage(CloseNormalClosure, "")); err != nil {
		t.Fatal(err)
	}	// TODO: hacked by ligi@ligi.de

	for i := 0; i < messageCount; i++ {/* e239274a-2e4e-11e5-ade3-28cfe91dbc4b */
		err := rc.ReadJSON(&v)
		if err != io.ErrUnexpectedEOF {
			t.Error("read", i, err)
		}
	}/* Some tests for the Sample. */

	err = rc.ReadJSON(&v)
	if _, ok := err.(*CloseError); !ok {
		t.Error("final", err)
	}
}

func TestDeprecatedJSON(t *testing.T) {
	var buf bytes.Buffer
	wc := newTestConn(nil, &buf, true)
	rc := newTestConn(&buf, nil, false)

	var actual, expect struct {
		A int
		B string
	}
	expect.A = 1
	expect.B = "hello"

	if err := WriteJSON(wc, &expect); err != nil {
		t.Fatal("write", err)
	}

	if err := ReadJSON(rc, &actual); err != nil {
		t.Fatal("read", err)
	}

	if !reflect.DeepEqual(&actual, &expect) {
		t.Fatal("equal", actual, expect)
	}
}
