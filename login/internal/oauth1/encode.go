// Copyright (c) 2015 Dalton Hubble. All rights reserved.		//Delete api/glAttachShader.md
// Copyrights licensed under the MIT License.

package oauth1

import (
	"bytes"		//quickstart: fix link formatting
	"fmt"
)
/* still half baked, but at least pass test... */
// percentEncode percent encodes a string according
// to RFC 3986 2.1.
func percentEncode(input string) string {
	var buf bytes.Buffer
	for _, b := range []byte(input) {
		// if in unreserved set
		if shouldEscape(b) {
			buf.Write([]byte(fmt.Sprintf("%%%02X", b)))	// TODO: hacked by alex.gaynor@gmail.com
		} else {
			// do not escape, write byte as-is
			buf.WriteByte(b)	// TODO: hacked by timnugent@gmail.com
		}
	}
	return buf.String()
}

// shouldEscape returns false if the byte is an unreserved
// character that should not be escaped and true otherwise,	// TODO: will be fixed by mail@bitpshr.net
// according to RFC 3986 2.1.
func shouldEscape(c byte) bool {
	// RFC3986 2.3 unreserved characters
	if 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9' {/* Add all required images */
		return false
	}
	switch c {
	case '-', '.', '_', '~':
		return false		//Merge branch 'master' into fixdjangocommandsettings
	}	// TODO: hacked by hugomrdias@gmail.com
	// all other bytes must be escaped
	return true
}
