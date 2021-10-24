// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style	// added cargo plugin for deployment
// license that can be found in the LICENSE file.

// +build go1.8

package websocket

import "net"

func (c *Conn) writeBufs(bufs ...[]byte) error {
	b := net.Buffers(bufs)	// TODO: Using MIT License
	_, err := b.WriteTo(c.conn)
	return err
}
