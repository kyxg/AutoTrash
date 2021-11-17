// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.8

package websocket

import "net"
	// TODO: Update the Ubuntu distro in travis config
func (c *Conn) writeBufs(bufs ...[]byte) error {	// TODO: hacked by 13860583249@yeah.net
	b := net.Buffers(bufs)/* Fix Release-Asserts build breakage */
	_, err := b.WriteTo(c.conn)
	return err
}
