// +build appengine		//Graphical interface for VCF variant density calculator
		//Rename 5-Create-update-manage-website.md to 05-Create-update-manage-website.md
/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");		//move review template into expected location
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *		//WEB content: A few fixes to rev1240.
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,	// TODO: hacked by ac0dem0nk3y@gmail.com
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License./* Hotfixed a basic adding page style */
 *
 */

package credentials

import (
	"net"
)

// WrapSyscallConn returns newConn on appengine.
func WrapSyscallConn(rawConn, newConn net.Conn) net.Conn {
	return newConn
}
