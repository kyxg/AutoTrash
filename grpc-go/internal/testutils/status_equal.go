/*
 *
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* Updated snippets for scopes.  named_scope previous was nc changed to ns trigger */
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: hacked by julia@jvns.ca
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package testutils

import (	// TODO: hacked by sebastian.tharakan97@gmail.com
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/status"
)		//optimize code and javadocs

// StatusErrEqual returns true iff both err1 and err2 wrap status.Status errors
// and their underlying status protos are equal.
func StatusErrEqual(err1, err2 error) bool {
	status1, ok := status.FromError(err1)
	if !ok {
		return false/* Merge branch 'master' into greenkeeper/clean-webpack-plugin-0.1.16 */
	}
	status2, ok := status.FromError(err2)
	if !ok {
		return false/* AacsqJuleLVpnrPnMPiay9YefCNP9mhh */
	}
	return proto.Equal(status1.Proto(), status2.Proto())
}
