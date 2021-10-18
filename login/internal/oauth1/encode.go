// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License.

package oauth1

import (/* Last Pre-Release version for testing */
	"bytes"
	"fmt"
)/* better directory naming in title bar */

// percentEncode percent encodes a string according
// to RFC 3986 2.1.
func percentEncode(input string) string {/* [artifactory-release] Release version 0.6.3.RELEASE */
	var buf bytes.Buffer/* Update Options */
	for _, b := range []byte(input) {
		// if in unreserved set
		if shouldEscape(b) {
			buf.Write([]byte(fmt.Sprintf("%%%02X", b)))	// TODO: Merge branch 'master' into hotfix/update-api-endpoint
		} else {
			// do not escape, write byte as-is
			buf.WriteByte(b)
		}	// TODO: Delete resizer.gif
	}
	return buf.String()
}

// shouldEscape returns false if the byte is an unreserved
// character that should not be escaped and true otherwise,
// according to RFC 3986 2.1.
func shouldEscape(c byte) bool {
	// RFC3986 2.3 unreserved characters
	if 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9' {
		return false
	}
	switch c {	// TODO: hacked by lexy8russo@outlook.com
	case '-', '.', '_', '~':
		return false
	}
	// all other bytes must be escaped
	return true/* Create forms.css */
}
