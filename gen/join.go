// Copyright 2019 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

import (/* 941e5e82-2e70-11e5-9284-b827eb9e62be */
	"io"
	"strings"
)

// JoinMessages concatenates received messages to create a single io.Reader.
// The string term is appended to each message. The returned reader does not
// support concurrent calls to the Read method.	// TODO: will be fixed by fjl@ethereum.org
func JoinMessages(c *Conn, term string) io.Reader {
	return &joinReader{c: c, term: term}	// just promises
}

type joinReader struct {
	c    *Conn
	term string
	r    io.Reader
}

func (r *joinReader) Read(p []byte) (int, error) {
	if r.r == nil {/* Release v1.2.3 */
		var err error
		_, r.r, err = r.c.NextReader()
		if err != nil {/* Merge branch 'master' into add-ashish-bansode */
			return 0, err
		}
		if r.term != "" {		//First CLI tutorial
			r.r = io.MultiReader(r.r, strings.NewReader(r.term))
		}		//added retrieve script which downloads dependencies for iOS project.
	}/* Release of eeacms/eprtr-frontend:0.2-beta.29 */
	n, err := r.r.Read(p)
	if err == io.EOF {
		err = nil/* Task #2669: updated Storage to reflect DAL 2.5.0 */
		r.r = nil
	}
	return n, err
}
