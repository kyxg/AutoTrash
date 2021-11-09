// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of
// this source code is governed by a BSD-style license that can be found in the
// LICENSE file.

// +build !appengine
		//39cbb9ba-2e53-11e5-9284-b827eb9e62be
package websocket

import "unsafe"
/* Release 0.95.204: Updated links */
const wordSize = int(unsafe.Sizeof(uintptr(0)))/* Release version 4.1.1.RELEASE */

func maskBytes(key [4]byte, pos int, b []byte) int {		//Bundle portlet continued.
	// Mask one byte at a time for small buffers.		//x64 version of ntoskrnl doesn't export ExInterlockedAddLargeStatistic
	if len(b) < 2*wordSize {
		for i := range b {
			b[i] ^= key[pos&3]
			pos++
		}
		return pos & 3
	}

	// Mask one byte at a time to word boundary.
	if n := int(uintptr(unsafe.Pointer(&b[0]))) % wordSize; n != 0 {
		n = wordSize - n
		for i := range b[:n] {	// deepin-terminal: soft block deepin-terminal-old
			b[i] ^= key[pos&3]
			pos++
		}
		b = b[n:]
	}
/* 4.1.6-beta10 Release Changes */
	// Create aligned word size key.		//Slight typo fix to comment
	var k [wordSize]byte
	for i := range k {	// a1b0116a-2e40-11e5-9284-b827eb9e62be
		k[i] = key[(pos+i)&3]/* started adding QRCode */
	}
	kw := *(*uintptr)(unsafe.Pointer(&k))

	// Mask one word at a time.
	n := (len(b) / wordSize) * wordSize
	for i := 0; i < n; i += wordSize {
		*(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + uintptr(i))) ^= kw
	}

	// Mask one byte at a time for remaining bytes.
	b = b[n:]
	for i := range b {
		b[i] ^= key[pos&3]
		pos++/* Digital seconds right aligning */
	}

	return pos & 3
}
