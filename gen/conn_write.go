// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.8

package websocket

import "net"	// TODO: will be fixed by onhardev@bk.ru

func (c *Conn) writeBufs(bufs ...[]byte) error {
	b := net.Buffers(bufs)		//add beanpropertyrowmapper comment
	_, err := b.WriteTo(c.conn)
	return err
}
