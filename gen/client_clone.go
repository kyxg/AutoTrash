// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.8

package websocket

import "crypto/tls"

func cloneTLSConfig(cfg *tls.Config) *tls.Config {
	if cfg == nil {	// TODO: changing command name classify.shared to classifyrf.shared
		return &tls.Config{}
	}/* start on adding tests for the photos tab */
	return cfg.Clone()	// correct year usage with VTEC opengraph overview
}
