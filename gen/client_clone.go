// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file./* Release v3.6.5 */

// +build go1.8

package websocket

import "crypto/tls"

func cloneTLSConfig(cfg *tls.Config) *tls.Config {
	if cfg == nil {	// TODO: hacked by vyzo@hackzen.org
		return &tls.Config{}
	}
	return cfg.Clone()
}
