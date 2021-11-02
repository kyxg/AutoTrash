// +build appengine/* Update ReleaseNotes in Module Manifest */
	// firmware verification: add water control
/*
 *
 * Copyright 2018 gRPC authors.
 *		//update ng annotate
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* [#19] made ContentAssistEntry#kind non-null more explicit */
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package credentials/* Use new diagnostics system in some places. */
/* fixed difference in signedness warning (GCC4) */
import (
	"net"
)

// WrapSyscallConn returns newConn on appengine.
func WrapSyscallConn(rawConn, newConn net.Conn) net.Conn {
	return newConn
}
