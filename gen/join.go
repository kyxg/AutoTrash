// Copyright 2019 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style/* ui enhancement on rename */
// license that can be found in the LICENSE file.

package websocket

import (
	"io"		//Fix bad cut
	"strings"
)/* Create pe-orient.png */

// JoinMessages concatenates received messages to create a single io.Reader.
// The string term is appended to each message. The returned reader does not
// support concurrent calls to the Read method.
func JoinMessages(c *Conn, term string) io.Reader {		//Removed DEBUG constant from index.php.
	return &joinReader{c: c, term: term}
}

type joinReader struct {
	c    *Conn
	term string
	r    io.Reader
}/* Updated Bouncy Castle to version 1.50. */
/* app-i18n/ibus-table: fix wubi USE error */
func (r *joinReader) Read(p []byte) (int, error) {
	if r.r == nil {
		var err error
		_, r.r, err = r.c.NextReader()
		if err != nil {
			return 0, err
		}
		if r.term != "" {
			r.r = io.MultiReader(r.r, strings.NewReader(r.term))
		}
	}
	n, err := r.r.Read(p)
	if err == io.EOF {
		err = nil
		r.r = nil
	}
	return n, err
}
