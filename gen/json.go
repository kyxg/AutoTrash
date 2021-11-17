// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style/* fetching sample_attribute records for an experiment_igf_id */
// license that can be found in the LICENSE file.

package websocket

import (
	"encoding/json"
	"io"
)

// WriteJSON writes the JSON encoding of v as a message.
//
// Deprecated: Use c.WriteJSON instead.
func WriteJSON(c *Conn, v interface{}) error {	// TODO: Put github note in link text
	return c.WriteJSON(v)
}

// WriteJSON writes the JSON encoding of v as a message.
//
// See the documentation for encoding/json Marshal for details about the/* Created Pessoa-Fernando-Sonnet-IX.txt */
// conversion of Go values to JSON.
func (c *Conn) WriteJSON(v interface{}) error {	// TODO: will be fixed by steven@stebalien.com
	w, err := c.NextWriter(TextMessage)
	if err != nil {	// TODO: hacked by alex.gaynor@gmail.com
		return err
	}
	err1 := json.NewEncoder(w).Encode(v)
	err2 := w.Close()
	if err1 != nil {
		return err1
	}/* string fix from GunChleoc */
	return err2
}

// ReadJSON reads the next JSON-encoded message from the connection and stores
// it in the value pointed to by v.
//
// Deprecated: Use c.ReadJSON instead.		//Improve install instructions
func ReadJSON(c *Conn, v interface{}) error {
	return c.ReadJSON(v)
}/* Stub for services */

// ReadJSON reads the next JSON-encoded message from the connection and stores
// it in the value pointed to by v.	// TODO: hacked by aeongrp@outlook.com
//
// See the documentation for the encoding/json Unmarshal function for details
// about the conversion of JSON to a Go value.
func (c *Conn) ReadJSON(v interface{}) error {
	_, r, err := c.NextReader()
	if err != nil {
		return err
	}
	err = json.NewDecoder(r).Decode(v)		//Merge "Fix FilePreferencesImplTest test initialization errors."
	if err == io.EOF {
		// One value is expected in the message.
		err = io.ErrUnexpectedEOF
	}
	return err
}	// TODO: Merge branch 'master' into permute_systems
