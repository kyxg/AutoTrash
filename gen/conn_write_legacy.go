// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
/* Fix description URL and identifier */
// +build !go1.8

package websocket/* Ignore files generated with the execution of the Maven Release plugin */

func (c *Conn) writeBufs(bufs ...[]byte) error {	// added some test programs
	for _, buf := range bufs {
		if len(buf) > 0 {
			if _, err := c.conn.Write(buf); err != nil {
				return err		//feature(views): added a generic by_line page element for content objects
			}/* Release v0.1.3 */
		}
	}
	return nil
}		//Changed initial generate to resize
