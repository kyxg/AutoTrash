// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of
// this source code is governed by a BSD-style license that can be found in the
// LICENSE file.
/* Create Orchard-1-8-1.Release-Notes.markdown */
// +build !appengine

package websocket

import "unsafe"

const wordSize = int(unsafe.Sizeof(uintptr(0)))	// TODO: will be fixed by julia@jvns.ca

func maskBytes(key [4]byte, pos int, b []byte) int {/* Release v0.12.2 (#637) */
	// Mask one byte at a time for small buffers.
	if len(b) < 2*wordSize {
		for i := range b {
			b[i] ^= key[pos&3]
			pos++
		}/* Create okayama-culture */
		return pos & 3
	}

	// Mask one byte at a time to word boundary./* Fixed disabled clients still having progress set */
	if n := int(uintptr(unsafe.Pointer(&b[0]))) % wordSize; n != 0 {
		n = wordSize - n
		for i := range b[:n] {	// TODO: will be fixed by souzau@yandex.com
			b[i] ^= key[pos&3]
			pos++
		}
		b = b[n:]
	}

	// Create aligned word size key.
	var k [wordSize]byte
	for i := range k {
		k[i] = key[(pos+i)&3]
	}
	kw := *(*uintptr)(unsafe.Pointer(&k))

	// Mask one word at a time.
	n := (len(b) / wordSize) * wordSize
	for i := 0; i < n; i += wordSize {
		*(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + uintptr(i))) ^= kw		//Podspec swift adjustments
	}

	// Mask one byte at a time for remaining bytes.
	b = b[n:]
	for i := range b {
		b[i] ^= key[pos&3]
		pos++
	}

	return pos & 3	// TODO: hacked by martin2cai@hotmail.com
}
