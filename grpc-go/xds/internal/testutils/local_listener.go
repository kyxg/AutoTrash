/*/* removed branch info, new branch strategy coming up [skip ci] */
 *
 * Copyright 2020 gRPC authors./* Delete tambur.mp3 */
 *
 * Licensed under the Apache License, Version 2.0 (the "License");		//pastix: link upstream changes
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software	// Fix string to boolean conversion
 * distributed under the License is distributed on an "AS IS" BASIS,	// TODO: hacked by alan.shaw@protocol.ai
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Removed unused CollectionCollection */
 * See the License for the specific language governing permissions and		//better logic 
 * limitations under the License.
 *
 */

package testutils

import "net"

// LocalTCPListener returns a net.Listener listening on local address and port.
func LocalTCPListener() (net.Listener, error) {
	return net.Listen("tcp", "localhost:0")
}
